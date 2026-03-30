🧠 Go 内存分配全组件详解：mcache / mcentral / mspan
 
我们顺着 Go 内存分配的「三级缓存 + 最小单元」架构，从下到上、从分配到回收，把每个组件的本质、结构、工作逻辑、设计亮点一次性讲透。
 
 
 
一、先总览：Go 内存分配的完整层级
 
Go runtime 堆内存管理是一个分层、分锁、分大小类的架构，核心链路如下：
 
plaintext  
操作系统 (OS)
    ↓ mmap 申请大块内存
mheap (全局总管家，唯一全局锁)
    ↓ 按大小类划分 mspan，分配给 mcentral
mcentral (按大小类分锁，全局共享)
    ↓ 为每个 P 的 mcache 补充 mspan
mcache (每个 P 私有，无锁，分配最快)
    ↓ 直接分配给用户对象
mspan (所有层级的最小内存单元，承载实际内存)
 
 
- mcache：线程（P）本地缓存，无锁，小对象分配的主战场
- mcentral：大小类全局管理器，分锁，mcache 的「补给站」
- mspan：堆内存的最小管理单元，所有内存的「物理载体」
 
 
 
二、mcache：P 私有、无锁的极速分配层
 
1. 核心定位
 
 mcache  是每个逻辑处理器 P 私有的内存缓存，是 Go 内存分配中最靠近用户、速度最快的一层，99% 的小对象分配都在这里完成，完全无锁，彻底避免多线程竞争。
 
2. 核心数据结构（Go 源码简化版）
 
go  
type mcache struct {
    // 每个大小类对应的空闲 mspan 链表（共67个size class）
    alloc [numSpanClasses]*mspan

    // 扫描相关：GC 标记阶段用
    scanAlloc uintptr
    // 其他统计、辅助字段...
}
 
 
-  alloc  数组：每个元素对应一个  size class （大小类），存储该大小类的空闲  mspan  链表。
- 每个 P 独占一个  mcache ，P 调度到哪个 M（线程）， mcache  就跟着哪个 M，线程本地访问无锁。
 
3. 核心职责
 
1. 无锁分配小对象：用户申请 ≤32KB 的小对象时，直接从当前 P 的  mcache  对应大小类的  mspan  中分配，全程无锁，纳秒级完成。
2. 缓存复用内存：缓存 GC 回收的小对象内存，避免频繁向  mcentral / mheap  申请。
3. 内存不足时申请补给：当某个大小类的  mspan  耗尽（没有空闲块），向对应  mcentral  申请新的  mspan  填充。
4. GC 本地扫描：GC 标记阶段，遍历  mcache  中的  mspan ，标记存活对象。
 
4. 关键设计亮点
 
- 无锁极致性能：利用 P 的私有性，把分配压力分散到每个 P，彻底消除小对象分配的锁竞争，这是 Go 内存分配高性能的核心。
- 大小类精准匹配：每个大小类对应独立的  mspan ，避免内部碎片（比如 10B 对象分配 16B 块，浪费仅 6B）。
- 批量补给，减少交互： mcache  一次从  mcentral  申请一整个  mspan （包含上百个小对象块），后续分配直接用缓存，大幅减少全局操作。
 
5. 分配流程示例（小对象）
 
用户申请 20B 内存 → 匹配到 32B 大小类 → 从  mcache.alloc[对应索引]  的  mspan  中取一个空闲块 → 分配完成（无锁，O(1)）。
如果  mspan  无空闲块 → 向  mcentral  申请新的  mspan  → 填充到  mcache  → 分配。
 
 
 
三、mcentral：按大小类分锁的全局补给站
 
1. 核心定位
 
 mcentral  是全局共享、按大小类分锁的中间层，是  mcache  和  mheap  之间的桥梁。每个大小类对应一个  mcentral ，负责管理该大小类的所有  mspan ，为  mcache  补充内存。
 
2. 核心数据结构（Go 源码简化版）
 
go  
type mcentral struct {
    lock mutex // 每个mcentral独立锁，仅操作当前大小类时加锁

    spanClass spanClass // 对应的大小类
    // 空闲mspan链表：可直接分配给mcache
    nonempty mSpanList
    // 已分配mspan链表：正在被mcache使用
    empty mSpanList

    // 统计字段...
}
 
 
-  lock ：每个  mcentral  一把独立锁，仅操作当前大小类时加锁，避免全局锁竞争。
-  nonempty ：存储有空闲块的  mspan ，可直接分配给  mcache 。
-  empty ：存储无空闲块（或已分配给  mcache ）的  mspan ，等待 GC 回收后重新放入  nonempty 。
 
3. 核心职责
 
1. 按大小类管理 mspan：每个  mcentral  只负责一个大小类，比如 32B、64B、128B 等，精准管理对应规格的  mspan 。
2. 为 mcache 补给 mspan：当  mcache  某个大小类的  mspan  耗尽，向对应  mcentral  申请， mcentral  从  nonempty  链表取出  mspan  分配给  mcache ，并移到  empty  链表。
3. 回收 mcache 归还的 mspan： mcache  用完的  mspan （或 GC 回收后有空闲块的  mspan ）归还到  mcentral ，重新放入  nonempty  链表，等待复用。
4. 向 mheap 申请新 mspan：当  nonempty  链表为空（无可用  mspan ），向  mheap  申请新的  mspan ，补充到  nonempty  链表。
 
4. 关键设计亮点
 
- 分锁设计，降低竞争：每个大小类一把锁，不同大小类的操作互不干扰，只有同大小类的  mcache  申请才会竞争，远低于全局锁的开销。
- 两级链表，高效复用： nonempty / empty  分离，快速区分可用/不可用  mspan ，GC 回收后直接复用，避免频繁向  mheap  申请。
- 承上启下，分层缓冲：在  mcache （无锁）和  mheap （全局锁）之间加一层缓冲，大幅减少  mheap  的调用频率，提升整体性能。
 
5. 工作流程示例
 
 mcache  32B 大小类  mspan  耗尽 → 加锁对应  mcentral  → 从  nonempty  取  mspan  → 移到  empty  → 解锁 → 分配给  mcache 。
如果  nonempty  为空 → 向  mheap  申请新  mspan  → 加入  nonempty  → 分配。
 
 
 
四、mspan：堆内存的最小管理单元
 
1. 核心定位
 
 mspan  是 Go 堆内存的最小管理单元，是所有内存分配的「物理载体」。不管是  mcache 、 mcentral  还是  mheap ，管理的都是  mspan ，它承载了实际的内存页和对象分配信息。
 
2. 核心数据结构（Go 源码简化版）
 
go  
type mspan struct {
    // 内存范围：起始地址、大小、页号
    startAddr uintptr // mspan 起始虚拟地址
    npages    uintptr // 包含的页数量（1页=8KB）
    next      *mspan  // 链表下一个节点
    prev      *mspan  // 链表上一个节点

    // 大小类与对象信息
    spanClass   spanClass // 所属大小类
    elemsize    uintptr   // 每个对象块的大小（如32B、64B）
    nelems      uintptr   // mspan 包含的对象块总数
    freeindex   uintptr   // 下一个可分配的对象块索引
    allocCache  uint64    // 分配位缓存（快速判断空闲块）
    allocBits  *gcBits   // 分配位图：记录每个块是否已分配
    gcmarkBits *gcBits   // GC标记位图：记录每个块是否存活

    // 所属层级：属于哪个mcentral/mcache
    list *mspanList
    // 其他GC、统计字段...
}
 
 
- 核心本质：一个  mspan  由 N 个连续的 8KB 页组成（N 由大小类决定，比如 32B 大小类对应 1 页=8KB，可划分 256 个 32B 块）。
-  allocBits  分配位图：用 bit 位记录每个对象块的分配状态（1=已分配，0=空闲），O(1) 快速查找空闲块。
-  gcmarkBits  GC 标记位图：GC 标记阶段用，记录对象是否存活，回收时根据位图释放内存。
 
3. 核心职责
 
1. 承载实际内存：管理一片连续的内存页，划分成固定大小的对象块，供分配使用。
2. 记录分配状态：通过分配位图，快速跟踪每个块的分配/空闲状态，实现高效分配和回收。
3. GC 回收的核心载体：GC 标记阶段，遍历  mspan  的对象块，标记存活对象；清扫阶段，根据标记位图回收死亡对象的块，重置分配位图，重新变为空闲块。
4. 链表节点，层级管理：作为链表节点，被  mcache / mcentral / mheap  管理，在不同层级间流转（ mcache  →  mcentral  →  mheap ）。
 
4. 关键设计亮点
 
- 页级连续，减少碎片：以连续页为单位，避免小对象分配导致的外部碎片，同时通过大小类减少内部碎片。
- 位图极速分配/回收：用 bit 位记录状态，分配时直接找 0 位，回收时直接置 0，O(1) 操作，性能极高。
- 统一管理大小对象：小对象用 1 页  mspan ，大对象用多页  mspan （比如 100KB 对象用 13 页=104KB），一套结构适配所有场景。
- GC 友好设计：分配位图和标记位图分离，GC 标记和内存分配并行不冲突，支持并发 GC。
 
5. 工作流程示例（32B 大小类 mspan）
 
1.  mheap  分配 1 页（8KB）内存给  mspan  → 划分成 8KB / 32B = 256 个对象块。
2.  allocBits  初始全 0（全空闲）， freeindex  指向 0 号块。
3. 分配时：取  freeindex  对应块，置  allocBits  对应位为 1， freeindex  后移 → 分配完成。
4. GC 回收时：遍历  gcmarkBits ，死亡对象对应位为 0 → 置  allocBits  对应位为 0 → 块变为空闲，可重新分配。
 
 
 
五、四组件完整联动流程（从申请到回收）
 
1. 内存申请流程（小对象 ≤32KB）
 
plaintext  
用户申请内存 → 匹配大小类 → 优先从当前P的mcache对应大小类的mspan分配（无锁，最快）
↓（mcache的mspan耗尽）
向对应大小类的mcentral申请mspan（加mcentral锁，分锁竞争小）
↓（mcentral的nonempty链表为空）
向mheap