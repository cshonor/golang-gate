# channel 阻塞原理（实现向提纲）

## 1. `hchan` 里大致有什么

channel 运行时对应 **`hchan`**（见 `src/runtime/chan.go`），典型包含：环形缓冲（若有）、等待队列、互斥锁等。

```go
type hchan struct {
	qcount   uint           // 缓冲区内元素个数
	dataqsiz uint           // 缓冲容量（0 表示无缓冲）
	buf      unsafe.Pointer // 指向环形缓冲区
	elemsize uint16
	closed   uint32
	elemtype *_type
	sendx    uint // 发送下标（环形）
	recvx    uint // 接收下标（环形）
	recvq    waitq // 等待接收（sudog 队列）
	sendq    waitq // 等待发送（sudog 队列）
	lock     mutex
}

type waitq struct {
	first *sudog
	last  *sudog
}
```

### 1. 缓冲区（有缓冲时为环形，FIFO）

- 由 `buf` 指向连续内存，`sendx` / `recvx` 维护读写下标，元素出队顺序为 **FIFO**。
- **`dataqsiz == 0`**：无独立环形区，走「直接对接」的同步交换路径（仍可能有 `sendq`/`recvq`）。
- **`dataqsiz > 0`**：异步缓冲，发送方不必立刻碰到接收方。

### 2. `sendq` / `recvq`（`waitq`，挂 **sudog**）

- **`sendq`**：发送方在**缓冲满**等空间时阻塞排队。
- **`recvq`**：接收方在**缓冲空**等数据时阻塞排队。
- 底层是 **`sudog` 链表**，每个 sudog 携带对应的 **G** 与数据指针等（不是「裸 G 队列」三个字能概括的，但语义上等价于「等着的 G 通过 sudog 排队」）。

### 3. `lock`

- channel **带一把锁**保护 `hchan` 内部状态，发送/接收路径要先拿锁。
- 因此 channel **不是**通用意义上的无锁队列；粗粒度锁是常见实现特征（具体优化以版本为准）。

## 2. 无缓冲：同步交换

- 发送方与接收方**对上眼**时：可在持锁路径下**直接拷贝**数据到接收方或约定位置，双方继续或一方先跑（实现细节见 `chansend`/`chanrecv`）。
- **对不上**：当前 G 以 **sudog** 入 `sendq` 或 `recvq`，进入等待，**让出 M**。

## 3. 有缓冲：满 / 空

| 情况 | 行为（直觉） |
|------|----------------|
| 缓冲未满 | `send` 写入环形槽；可能唤醒 `recvq` 里等待的接收方。 |
| 缓冲满 | `send` 进 `sendq`，G 阻塞。 |
| 缓冲非空 | `recv` 从槽取；可能唤醒 `sendq` 里等待的发送方。 |
| 缓冲空 | `recv` 进 `recvq`，G 阻塞。 |

## 4. 与「阻塞协程现象」的对应

用户态看到「卡在 `<-ch` / `ch <- v`」对应 runtime：**G 挂在 channel 的 `sendq`/`recvq` 上并 park**；对端或缓冲条件满足时再 **ready**，重新进入可运行队列。

## 5. 延伸阅读

- `src/runtime/chan.go`：`chansend`、`chanrecv`、`closechan` 等。
- sudog 细节：[08-sudog详细介绍.md](./08-sudog详细介绍.md)

## 6. 自检

- 无缓冲 channel 是否还需要环形 `buf`？（结合 `dataqsiz == 0` 语义思考。）
- 为什么 channel 内部往往用一把**覆盖整个 `hchan`** 的锁？

---

## 复习速记

| 记什么 | 记一句 |
|--------|--------|
| 等谁排队 | `sendq` / `recvq`，节点是 **sudog** |
| 缓冲 | 有缓冲时 `buf` 为环形 FIFO |
| 阻塞 | 条件不满足 → sudog 入队 → G 等待 |
