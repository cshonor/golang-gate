# Channel 核心数据结构：`runtime.hchan` 详解

> **定位**：本篇补全「**Channel 自身在运行时里长什么样**」——`hchan` 的字段、环形缓冲、锁与等待队列；与 **[10-Channel阻塞](./10-Channel阻塞协程的原理与现象.md)**、**[14-学GMP必学Channel](./14-学GMP必学Channel总览.md)** 等讲的「**GMP 视角下为何阻塞/如何唤醒**」是**同一故事的两层**：先有容器与队列，再谈调度把 G 挂上去、摘下来。

权威定义以本机 Go 安装树下的 **`src/runtime/chan.go`** 为准；下文结构与字段名随版本可能微调，**背思想不背行号**。

---

## 1. 本章目录里各篇分工（和本篇的关系）

| 文档 | 侧重 | 是否覆盖 `hchan` 全图 |
|------|------|------------------------|
| [10-Channel阻塞协程的原理与现象](./10-Channel阻塞协程的原理与现象.md) | 阻塞现象 + 原理（合并篇） | 偏调度与实现 |
| [13-closed_channel核心特性与坑](./13-closed_channel核心特性与坑.md) | 关闭语义、`panic` 边界 | 偏语义 |
| [14-学GMP必学Channel总览](./14-学GMP必学Channel总览.md) | GMP 与 chan 路径总览 | 总览，未必拆字段 |
| [15-Channel性能优化详解](./15-Channel性能优化详解.md) | 性能与实现路径 | 会碰到 `hchan`，非系统拆解 |
| [12-Channel环形队列实现原理](./12-Channel环形队列实现原理.md) | `buf` / `sendx` / `recvx` / FIFO | **缓冲区的核心算法侧** |
| **本篇 11** | **`hchan` 全结构 + 与 sudog / 调度衔接** | **是** |

---

## 2. `hchan`：channel 在运行时的「盒子」

有缓冲 channel 在底层对应一个 **`hchan`**（`chan` 在运行时的具体类型）。概念上包含：**一块定长环形存储**、**两把等待队列**、**一把大锁**、**元数据**。

```go
// 示意：与 runtime 源码对齐读 chan.go；勿当作可 import 的类型
type hchan struct {
	qcount   uint           // 当前环形区里已有元素个数
	dataqsiz uint           // 环形区容量，即 make(chan T, N) 的 N；0 表示无缓冲
	buf      unsafe.Pointer // 指向 dataqsiz 个元素的数组首地址；无缓冲时常为 nil
	elemsize uint16         // 单个元素字节数
	closed   uint32         // 是否已关闭（原子访问语义由 runtime 保证）
	elemtype *_type         // 元素类型描述（拷贝、对齐、指针写屏障等）
	sendx    uint           // 下一次发送写入下标
	recvx    uint           // 下一次接收读出下标
	recvq    waitq          // 因「暂无数据可读」而阻塞在此 chan 上的接收方（<-ch）
	sendq    waitq          // 因「暂无空位可写」而阻塞在此 chan 上的发送方（ch<-）
	lock     mutex          // 保护上述字段与两个队列的一致性
}

type waitq struct {
	first *sudog
	last  *sudog
}
```

- **`recvq` / `sendq`**：挂的是 **`sudog`** 链表节点；每个 `sudog` 背后关联一个等待中的 **G**。详见 [08-sudog核心数据结构与作用详解](./08-sudog核心数据结构与作用详解.md)、[09-GMP与sudog四者联动关系](./09-GMP与sudog四者联动关系.md)。  
- **调度层**：G 进队后不再占用 P 的本地可运行队列，直到被 channel 逻辑唤醒——这条线在 [07-Goroutine阻塞后的调度行为](./07-Goroutine阻塞后的调度行为.md)、[10-Channel阻塞协程的原理与现象](./10-Channel阻塞协程的原理与现象.md) 里展开。

---

## 3. 字段速查表（面试够用）

| 字段 | 一句话 |
|------|--------|
| `qcount` | 当前缓冲区内元素个数；**判空/判满**常与 `dataqsiz` 配合 |
| `dataqsiz` | 容量；**0** 表示无缓冲：发送/接收更易直接和对方 **G** 握手，不经过环形区 |
| `buf` | 环形数组本体；元素布局连续，`unsafe.Pointer` + `elemsize` 寻址 |
| `sendx` / `recvx` | 环形写指针 / 读指针；与 [12-Channel环形队列实现原理](./12-Channel环形队列实现原理.md) 一致 |
| `elemtype` / `elemsize` | 类型与大小：决定**内存拷贝**与**是否指针类型**（影响 GC 写屏障路径） |
| `closed` | 关闭标记；与 [13-closed_channel核心特性与坑](./13-closed_channel核心特性与坑.md) 的语义对应 |
| `recvq` / `sendq` | **因 channel 条件不满足**而睡眠的 G，经 `sudog` 串起来 |
| `lock` | **同一 `hchan` 上 send/recv/close 互斥**；高竞争下会成为热点（见 [15-Channel性能优化详解](./15-Channel性能优化详解.md)） |

---

## 4. 无缓冲 vs 有缓冲（结构层差异）

| | 无缓冲 `make(chan T)` | 有缓冲 `make(chan T, N)` |
|--|------------------------|----------------------------|
| `dataqsiz` | **0** | **N > 0** |
| `buf` | 通常无独立环形区 | 指向 **N** 个 `T` 大小的槽位 |
| 典型路径 | 直接 **sendq ↔ recvq** 与对方 G 同步，或一方入队等待 | 先尝试 **写满/读空** 再在 `buf` 与队列间切换 |

---

## 5. 和 GMP 怎么「接上线」（一页纸）

1. **发送/接收/关闭**进入 `runtime.chansend` / `chanrecv` / `closechan` 等函数，先拿 **`hchan.lock`**。  
2. 若**能立刻完成**（对端已在等、或缓冲区有余地）：在锁内拷贝数据、唤醒对端 `sudog`，G 重新变为可运行。  
3. 若**不能立刻完成**：当前 G 打包成 **`sudog`**，入 **`sendq` 或 `recvq`**，G **挂起**并让出 M（细节回 **[10](./10-Channel阻塞协程的原理与现象.md)** / **[07](./07-Goroutine阻塞后的调度行为.md)**）。  
4. **环形区**只解决「**同一 goroutine 之间异步暂存**」；**阻塞与唤醒**仍由 **`waitq` + 调度器** 完成。

---

## 6. 推荐阅读顺序（把「结构」和「调度」拼成一块）

1. [12-Channel环形队列实现原理](./12-Channel环形队列实现原理.md)（`buf` / 下标 / FIFO）  
2. **本篇**（`hchan` 全貌 + `recvq`/`sendq`）  
3. [08-sudog核心数据结构与作用详解](./08-sudog核心数据结构与作用详解.md)（队列节点长什么样）  
4. [10-Channel阻塞协程的原理与现象](./10-Channel阻塞协程的原理与现象.md)（为何阻塞、谁唤醒）  
5. [14-学GMP必学Channel总览](./14-学GMP必学Channel总览.md)（把上述串进 GMP 心智模型）

---

## 7. 延伸

- 源码：`$GOROOT/src/runtime/chan.go`（`hchan`、`chansend`、`chanrecv`、`closechan`）。  
- 同步原语收尾（与本章课表 **16** 对应）：[WaitGroup 特性及原理](../08-atomic%20and%20lock/14-WaitGroup特性及原理.md)。
