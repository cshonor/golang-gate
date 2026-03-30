# 三、mheap 深度解析

`mheap` 是 Go runtime 里**整个堆的全局总管理者**，处于 `mcache → mcentral → mheap → OS` 链条的顶端：堆内存的申请、分配、回收、与 OS 的交互都由它统一调度。

## 1. 核心定位与本质

### 1. 本质

`mheap` 对应全局唯一的 `runtime.mheap`，进程启动时初始化，生命周期与进程一致。它是堆的「总仓库」：对接操作系统，并为 `mcentral`、`mcache` 提供 `mspan` 与页级资源。

### 2. 核心职责

| 职责 | 说明 |
|------|------|
| 全局内存统筹 | 管理堆地址空间，掌握已分配、空闲、待回收的块 |
| 对接操作系统 | `mmap` / `munmap` 等，控制进程堆占用 |
| 为 `mcentral` 供货 | 某 size class 的 `mspan` 不足时，由 `mheap` 分配新 `mspan` |
| 配合 GC | 标记、清扫、页回收与整理参与堆管理 |
| 碎片治理 | 合并相邻空闲页等，提高利用率 |

## 2. 核心数据结构（源码视角）

在 `runtime/mheap.go` 中，`mheap` 可简化为：

```go
type mheap struct {
	lock mutex // 操作 mheap 需加锁（全局共享）

	// 每个 size class 对应一个 mcentral（含 cache line 对齐填充）
	central [numSpanClasses]struct {
		mcentral mcentral
		pad      [cpu.CacheLinePadSize - unsafe.Sizeof(mcentral{})%cpu.CacheLinePadSize]byte
	}

	// 堆内存的总范围（虚拟地址空间）
	arenaStart uintptr // 堆起始地址
	arenaEnd   uintptr // 堆结束地址
	arenaUsed  uintptr // 已使用的堆内存大小
	arenaAlloc uintptr // 已分配给 mspan 的内存大小

	pages pageAlloc // 页分配器：管理页的分配与回收

	spans []*mspan // 按页号索引 mspan，便于查找与 GC
	// 其他 GC、统计字段...
}
```

### 关键字段拆解

#### 1. 全局锁 `lock`

`mheap` 全局共享，申请/释放内存、部分 GC 路径等会竞争这把锁。为减少触碰频率，才有了 **`mcache`（P 私有无锁）** 与 **`mcentral`（按类分锁）**；多数小对象分配不走到 `mheap`。

#### 2. `central` 数组

`mheap` 持有全部 `mcentral` 实例，每个对应一个 size class（版本不同，数量如 60+ / 67 等，以源码为准）。某 `mcentral` 的 `mspan` 用尽时，向 `mheap` 申请新 `mspan` 并挂到该链路上。

#### 3. `pages` 页分配器

`mheap` 以**页**（常见 8KB）为粒度管理内存。`pageAlloc` 负责：

- 记录空闲 / 已分配页；
- 为 `mspan` 分配连续 1～N 页；
- 回收 GC 释放的页，合并相邻空闲页以降低碎片。

#### 4. `spans` 索引

堆上每一页可映射到所属 `mspan`，便于 O(1) 量级按页反查，支撑标记与回收。

## 3. 完整工作流程

### 1. 内存申请（从 OS 到对象）

```text
程序申请内存 → 优先当前 P 的 mcache（无锁）
  ↓ 不足
向对应 size class 的 mcentral 申请 mspan
  ↓ 不足
向 mheap 申请新 mspan
  ↓ 堆上空闲页不足
mheap 通过 mmap 向 OS 申请大块虚拟内存
  → 划页、构造 mspan，经 mcentral 同步到 mcache
  → 最终分配给用户对象
```

### 2. 内存回收（从对象到 OS）

```text
GC 标记：标记存活对象，识别可回收内存
  ↓
GC 清扫：回收死亡对象占用的空间，页归还页分配器
  ↓
页分配器合并相邻空闲页
  ↓
空闲超过策略阈值时，可 munmap 将内存还给 OS（行为受版本与策略影响）
  ↓
进程 RSS 可能下降
```

## 4. 核心设计亮点

### 1. 分层锁，降低热点

- **`mcache`**：P 私有，无锁，小对象主路径。
- **`mcentral`**：按 size class 分锁。
- **`mheap`**：全局锁，仅在更上层不足时触发，频率相对较低。

### 2. 页级管理

- 以页为单元灵活组合 `mspan`，适配不同对象大小。
- **大对象**（如常见实现中大于 32KB）往往直接从 `mheap` 分配，绕过小对象缓存，避免大对象占满缓存。

### 3. 主动与 GC 协同

- 配合 GC 回收与页归还，避免长期占用无用堆。
- 具体释放回 OS 的时机与阈值依 runtime 版本与配置而定。

### 4. 碎片治理

- 页分配器合并空闲页，减轻外部碎片；
- size class 减轻小对象内部碎片。

## 5. 与其他组件的关系

| 组件 | 与 mheap 的关系 | 典型交互 |
|------|-----------------|----------|
| `mcache` | 下层（P 私有） | 经 `mcentral` 间接触发 `mheap` 申请 |
| `mcentral` | 由 `mheap.central` 持有 | 直接向 `mheap` 要新 `mspan` |
| `mspan` | 页由 `mheap` 划出 | `mheap` 统一管理页与索引 |
| 操作系统 | 上层 | `mmap` / `munmap` 等 |

## 6. 常见问题

### 1. 大对象（如大于 32KB）如何分配？

通常视为大对象，**直接从 `mheap` 走连续页分配**，不经 `mcache` 小对象路径：占用大、复用模式不同，避免浪费 per-P 缓存。

### 2. `mheap` 会不会成为瓶颈？

绝大多数小对象停留在 `mcache`；只有补给链路上顶不住时才进 `mheap`。runtime 也在持续优化锁与分配路径，但全局锁仍是需要理解的成本点。

### 3. `mheap` 与 GC

- **标记**：借助 `spans` 等结构遍历 `mspan`、追踪存活。
- **清扫**：回收死亡对象关联的页与 `mspan`，归还页分配器。
- **整理**：与版本相关的压缩/合并策略在演进，以源码与 release note 为准。

---

## 复习速记

| 要点 | 内容 |
|------|------|
| 定位 | 全局唯一堆管理器，对接 OS |
| 锁 | 全局 `lock`，应少触碰 |
| 页 | `pageAlloc` 管页，`mspan` 由页组成 |
| 索引 | `spans` 按页找 `mspan` |
| 大对象 | 常绕过 `mcache`，直走 `mheap` |

## 延伸阅读

1. `runtime/mheap.go`、`runtime/mcentral.go`、`runtime/mcache.go`
2. size class 与 `mspan` 划分
3. 三色标记与清扫如何消费 `mspan` / 位图
4. 调参：`GOGC`、`GOMEMLIMIT` 与堆占用观测（`pprof`、`runtime/trace`）
