# Channel 阻塞协程：原理与现象（合并篇）

> 本篇由原 **`10-Channel阻塞协程的原理与现象`** 与 **`10-Channel阻塞协程的原理与现象`** 合并，现象与实现对照阅读；合并版速记仍见 [summary/channel阻塞_现象到原理.md](./summary/channel阻塞_现象到原理.md)。

---

## 第一部分：现象（用户态看到什么）

### 1. 现象描述

- **无缓冲 channel**：发送方在对方未接收前 **阻塞**；接收方在对方未发送前 **阻塞**。  
- **有缓冲 channel**：缓冲区 **满** 时发送阻塞；**空** 时接收阻塞。

### 2. 从运行时角度理解（直觉）

- 阻塞的 goroutine **不会一直占满 CPU**：会被 **挂起**，进入等待队列，**让出 P**。  
- 代码停在 `ch <- v` 或 `<-ch`，本质是 **调度器不再运行该 G，直到被唤醒**。

### 3. 与 GMP 的粗关系

- **G**：被阻塞的 goroutine。  
- **M / P**：线程与逻辑处理器继续跑其他 G；被阻塞的 G 与 channel 的 **`recvq` / `sendq`**（`waitq`）关联，节点为 **`sudog`**（详见 [08-sudog核心数据结构与作用详解.md](./08-sudog核心数据结构与作用详解.md)）。

### 4. 小实验思路

- `main` 里无缓冲 send 且无 receive → 死锁（单 G 无法配对）。  
- 缓冲为 1，连续两次 send 无 receive → 第二次 send 阻塞。

---

## 第二部分：原理（`hchan` 与 park/ready）

### 1. `hchan` 里大致有什么

channel 运行时对应 **`hchan`**（`src/runtime/chan.go`），典型包含：环形缓冲（若有）、等待队列、互斥锁等。

```go
type hchan struct {
	qcount   uint
	dataqsiz uint
	buf      unsafe.Pointer
	elemsize uint16
	closed   uint32
	elemtype *_type
	sendx    uint
	recvx    uint
	recvq    waitq
	sendq    waitq
	lock     mutex
}

type waitq struct {
	first *sudog
	last  *sudog
}
```

更完整的字段表与延伸阅读：[11-Channel核心数据结构hchan详解.md](./11-Channel核心数据结构hchan详解.md)。

### 2. 缓冲区（有缓冲时为环形，FIFO）

- 由 `buf` 指向连续内存，`sendx` / `recvx` 维护下标，**FIFO**。  
- **`dataqsiz == 0`**：无独立环形区，走「直接对接」的同步交换路径。  
- **`dataqsiz > 0`**：异步缓冲，发送不必立刻碰到接收。

环形算法细节：[12-Channel环形队列实现原理.md](./12-Channel环形队列实现原理.md)。

### 3. `sendq` / `recvq` 与 `lock`

- **`sendq`**：缓冲满等空间时发送方排队。  
- **`recvq`**：缓冲空等数据时接收方排队。  
- 节点为 **`sudog`**，携带对应 **G** 等。  
- **`lock`**：保护整个 `hchan`，发送/接收/关闭路径先拿锁（优化路径以版本为准）。

### 4. 无缓冲：同步交换

- 收发**对上眼**：持锁路径下常可 **直接拷贝**，双方继续或一方先跑。  
- **对不上**：当前 G 以 **sudog** 入 `sendq` 或 `recvq`，**park**，**让出 M**。

### 5. 有缓冲：满 / 空

| 情况 | 行为（直觉） |
|------|----------------|
| 缓冲未满 | `send` 写入环形槽；可能唤醒 `recvq` 里等待的接收方。 |
| 缓冲满 | `send` 进 `sendq`，G 阻塞。 |
| 缓冲非空 | `recv` 从槽取；可能唤醒 `sendq` 里等待的发送方。 |
| 缓冲空 | `recv` 进 `recvq`，G 阻塞。 |

### 6. 现象 ↔ 实现

用户态「卡在 `<-ch` / `ch <- v`」对应 runtime：**G 挂在 `sendq`/`recvq` 上并 park**；条件满足时 **ready**，重新进入可运行队列。与 [07-Goroutine阻塞后的调度行为.md](./07-Goroutine阻塞后的调度行为.md) 对照。

---

## 延伸阅读

- 源码：`src/runtime/chan.go`（`chansend`、`chanrecv`、`closechan`）。  
- 关闭语义：[13-closed_channel核心特性与坑.md](./13-closed_channel核心特性与坑.md)。  
- GMP 与 chan 总览：[14-学GMP必学Channel总览.md](./14-学GMP必学Channel总览.md)。

---

## 复习速记

| 记什么 | 记一句 |
|--------|--------|
| 等谁排队 | `sendq` / `recvq`，节点是 **sudog** |
| 缓冲 | 有缓冲时 `buf` 为环形 FIFO |
| 阻塞 | 条件不满足 → sudog 入队 → G 等待 |

---

## 自检

- 无缓冲 channel 同步的是「**时间点**」还是「**数据**」？  
- 缓冲 channel 的「满/空」与 len/cap 的关系？  
- 无缓冲 channel 是否还需要环形 `buf`？（结合 `dataqsiz == 0`。）  
- 为什么 channel 内部往往用一把**覆盖整个 `hchan`** 的锁？
