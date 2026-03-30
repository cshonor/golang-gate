# 二、mcache / mcentral / mspan

顺着 Go 内存分配的「三级缓存 + 最小单元」架构，从下到上、从分配到回收，把每个组件的本质、结构、工作逻辑、设计亮点串起来。

## 1. 总览：Go 内存分配的完整层级

Go runtime 堆内存管理是分层、分锁、分大小类的架构，核心链路如下：

```text
操作系统 (OS)
    ↓ mmap 申请大块内存
mheap（全局总管家，唯一全局锁）
    ↓ 按大小类划分 mspan，分配给 mcentral
mcentral（按大小类分锁，全局共享）
    ↓ 为每个 P 的 mcache 补充 mspan
mcache（每个 P 私有，无锁，分配最快）
    ↓ 直接分配给用户对象
mspan（所有层级的最小内存单元，承载实际内存）
```

- **`mcache`**：线程（P）本地缓存，无锁，小对象分配的主战场。
- **`mcentral`**：大小类全局管理器，分锁，是 `mcache` 的「补给站」。
- **`mspan`**：堆内存的最小管理单元，所有内存的「物理载体」。

## 2. mcache：P 私有、无锁的极速分配层

### 1. 核心定位

`mcache` 是每个逻辑处理器 P 私有的内存缓存，是 Go 内存分配中最靠近用户、速度最快的一层；绝大多数小对象分配在这里完成，**完全无锁**，避免多线程竞争。

### 2. 核心数据结构（源码简化）

```go
type mcache struct {
	// 每个大小类对应的空闲 mspan 链表（共 numSpanClasses 个 size class）
	alloc [numSpanClasses]*mspan

	// 扫描相关：GC 标记阶段用
	scanAlloc uintptr
	// 其他统计、辅助字段...
}
```

- **`alloc` 数组**：每个元素对应一个 **size class**，存该类的空闲 `mspan` 链表。
- **与 P 绑定**：每个 P 独占一个 `mcache`；P 调度到哪个 M，`mcache` 就跟到哪个 M，线程本地访问无锁。

### 3. 核心职责

1. **无锁分配小对象**：申请 ≤32KB 的小对象时，直接从当前 P 的 `mcache` 对应大小类的 `mspan` 里分配，全程无锁。
2. **缓存复用**：缓存 GC 回收的小对象内存，减少向 `mcentral` / `mheap` 的申请。
3. **不足时补给**：某大小类的 `mspan` 耗尽时，向对应 `mcentral` 申请新 `mspan` 填入。
4. **GC 本地扫描**：标记阶段遍历 `mcache` 中的 `mspan`，标记存活对象。

### 4. 关键设计亮点

- **无锁性能**：利用 P 的私有性分散分配压力，消除小对象分配的锁竞争。
- **大小类匹配**：每个大小类对应独立 `mspan`，控制内部碎片。
- **批量补给**：一次从 `mcentral` 拉整段 `mspan`（含大量 slot），后续分配走本地缓存，减少全局操作。

### 5. 分配流程示例（小对象）

用户申请 20B → 归到例如 32B 的 size class → 从 `mcache.alloc[对应索引]` 的 `mspan` 取空闲块 → 完成（无锁，近似 O(1)）。若该 `mspan` 无空闲块 → 向 `mcentral` 申请新 `mspan` → 填入 `mcache` 后再分配。

## 3. mcentral：按大小类分锁的全局补给站

### 1. 核心定位

`mcentral` 是全局共享、**按大小类分锁**的中间层，是 `mcache` 与 `mheap` 的桥梁。每个大小类对应一个 `mcentral`，管理该类的全部 `mspan`，并为 `mcache` 补货。

### 2. 核心数据结构（源码简化）

```go
type mcentral struct {
	lock      mutex     // 每个 mcentral 独立锁，仅操作当前大小类时加锁
	spanClass spanClass // 对应的大小类

	nonempty mSpanList // 有空闲对象的 mspan，可分给 mcache
	empty    mSpanList // 无空闲或正被 mcache 使用的 mspan

	// 统计字段...
}
```

- **`lock`**：每个 `mcentral` 一把锁，只锁当前大小类，避免全局大锁。
- **`nonempty`**：有空闲块的 `mspan`，可直接分给 `mcache`。
- **`empty`**：无空闲块或已交给 `mcache` 的 `mspan`；GC 回收有空闲后可再挂回 `nonempty`。

### 3. 核心职责

1. **按 size class 管理 `mspan`**：例如 32B、64B、128B 各由对应 `mcentral` 管理。
2. **给 `mcache` 补给**：`mcache` 某类耗尽时，向该类 `mcentral` 要 `mspan`；通常从 `nonempty` 取出，并可能移到 `empty`。
3. **回收与复用**：`mcache` 归还或 GC 后有空闲的 `mspan` 回到 `mcentral`，进入可复用链表。
4. **`nonempty` 为空时**：向 `mheap` 申请新 `mspan`，再挂入链表。

### 4. 关键设计亮点

- **分锁降竞争**：不同 size class 互不抢同一把锁。
- **`nonempty` / `empty` 分离**：快速区分可分配与已占满，便于复用与 GC 协作。
- **缓冲层**：夹在无锁 `mcache` 与全局 `mheap` 之间，降低 `mheap` 调用频率。

### 5. 工作流程示例

`mcache` 上某 32B 的 `mspan` 用尽 → 对应该类的 `mcentral` 加锁 → 从 `nonempty` 取 `mspan`（逻辑上迁到 `empty` 等，视实现阶段）→ 解锁 → 交给 `mcache`。若 `nonempty` 为空 → 向 `mheap` 申请新 `mspan` → 再放入链表并分配。

## 4. mspan：堆内存的最小管理单元

### 1. 核心定位

`mspan` 是 Go 堆的**最小管理单元**，是真实内存页的承载方。`mcache`、`mcentral`、`mheap` 管理对象都是 `mspan`。

### 2. 核心数据结构（源码简化）

```go
type mspan struct {
	startAddr uintptr
	npages    uintptr // 连续页数量（常见 1 页 = 8KB）
	next      *mspan
	prev      *mspan

	spanClass spanClass
	elemsize  uintptr // 每个对象槽大小
	nelems    uintptr // 槽数量
	freeindex uintptr
	allocCache uint64
	allocBits   *gcBits
	gcmarkBits  *gcBits

	list *mspanList
	// 其他 GC、统计字段...
}
```

- **页与槽**：一个 `mspan` 由若干连续页组成；再按 `elemsize` 切成多个 slot（例如 1 页 8KB、32B 一类约 256 槽，具体以 size class 为准）。
- **`allocBits`**：记录每个槽是否已分配。
- **`gcmarkBits`**：GC 标记阶段记录存活信息，与清扫配合回收。

### 3. 核心职责

1. **承载真实内存**：连续页 + 定长槽，供分配器使用。
2. **维护分配状态**：位图快速查询空闲槽。
3. **GC 的载体**：标记阶段扫槽；清扫阶段回收死亡槽，位图可复用。
4. **链表节点**：在 `mcache` / `mcentral` / `mheap` 之间流转。

### 4. 关键设计亮点

- **连续页**：利于减少外部碎片；配合 size class 控制内部碎片。
- **位图**：分配/回收可做到很快的槽级操作。
- **统一抽象**：小对象少页、大对象多页，同一套 `mspan` 描述。
- **双位图**：分配与标记分离，便于并发 GC。

### 5. 工作流程示例（以某小对象 size class 为例）

1. `mheap` 分配页给 `mspan`，按 `elemsize` 划槽；`allocBits` 初始全空，`freeindex` 从 0 起。
2. 分配：取 `freeindex` 对应槽，置分配位，`freeindex` 推进。
3. GC：根据 `gcmarkBits` 判断存活；死亡槽清回分配位图，可再次分配。

## 5. 四组件联动（申请与回收）

### 1. 内存申请（小对象 ≤32KB）

```text
用户申请 → 匹配 size class
  → 优先从当前 P 的 mcache 对应 mspan 分配（无锁）
  ↓ mcache 该 mspan 无空闲槽
向对应 mcentral 取 mspan（加该 mcentral 锁）
  ↓ mcentral.nonempty 为空
向 mheap 申请新 mspan（加 mheap 全局锁）
  ↓ 堆上无足够空闲页
mheap 通过 mmap 向 OS 申请虚拟内存
  → 划页、构造 mspan，经 mcentral 进入链表
  → 回填 mcache → 分配给用户
```

### 2. 回收与复用（简要）

- **标记 / 清扫**：在 `mspan` 的位图上完成；死亡对象槽释放后可再分配。
- **链表迁移**：`mspan` 在 `nonempty` / `empty` 以及 `mcache` 之间迁移，实现复用。
- **页级归还**：最终由 `mheap` 的页分配器合并空闲页，必要时 `munmap` 还给操作系统（细节见 `mheap` 笔记）。

---

## 复习速记

| 组件 | 一句话 |
|------|--------|
| `mcache` | 每 P 私有，无锁，小对象主路径 |
| `mcentral` | 按 size class 分锁，给 `mcache` 补 `mspan` |
| `mspan` | 连续页 + 定长槽，位图管理，最小管理单元 |
| 顺序记忆 | OS → `mheap` → `mcentral` → `mcache` → 对象；执行分配时从 `mcache` 往外「要」 |
