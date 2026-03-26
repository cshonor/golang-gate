# channel 阻塞：现象到原理（实现向提纲）

> 由 `channel阻塞协程现象.md` 与 `channel阻塞原理.md` 合并整理。

## 1. 先看现象（面试开场用）

- **无缓冲 channel**：发送方在对方未接收前阻塞；接收方在对方未发送前阻塞（同步交换）。
- **有缓冲 channel**：缓冲区**满**时发送阻塞；缓冲区**空**时接收阻塞。

运行时直觉：

- 阻塞的 goroutine 不会一直占 CPU：会被 **挂起（park）**，进入等待队列，让出执行机会；直到匹配操作发生再 **就绪（ready）**。

## 2. channel 内部大致结构（抓住 3 件事）

1. **缓冲区**：有缓冲时是环形队列（FIFO）
2. **sendq / recvq**：等待发送 / 等待接收的 G 队列（底层是 `sudog` 链表）
3. **锁**：保护 channel 内部状态（channel 不是无锁队列）

（提纲版结构）

```go
type hchan struct {
    qcount   uint           // 缓冲区中元素数量
    dataqsiz uint           // 缓冲区大小
    buf      unsafe.Pointer // 指向环形队列
    elemsize uint16
    closed   uint32
    elemtype *_type

    sendx    uint           // 发送索引（写）
    recvx    uint           // 接收索引（读）

    recvq    waitq          // 等待接收队列
    sendq    waitq          // 等待发送队列

    lock     mutex          // 互斥锁
}

type waitq struct {
    first *sudog
    last  *sudog
}
```

## 2.1 sudog 在 channel 里到底干嘛用？（把“等”说清楚）

当 G 因为 send/recv 需要等待时，runtime 不会把整个 `g` 结构体塞进队列，而是创建/复用一个 **sudog** 节点：

- **挂到哪里**：`hchan.sendq` 或 `hchan.recvq`
- **携带什么**：等待的 G、数据地址（elem）、关联的 channel 等
- **用来做什么**：未来配对成功时，直接拿 sudog 上的信息完成拷贝/唤醒

这就是你在代码里看到的“卡住”，在 runtime 里对应的“入队 + park”。

## 3. 无缓冲：同步交换（对不上就阻塞）

- 发送方找到正在等待的接收方（或反之）→ **直接拷贝数据** → 双方就绪
- 对不上 → 当前 G 入队（sendq/recvq）并阻塞，调度器切走

把它写成“分支流程”更像源码（记逻辑即可）：

```text
send(unbuffered):
  lock(hchan)
  if recvq not empty:
     dequeue receiver sudog
     copy(sender -> receiver)
     ready(receiver.g)
     unlock; return
  else:
     enqueue sender sudog into sendq
     park current g
     unlock; return (after wakeup)
```

## 4. 有缓冲：队列满/空（环形队列 + 等待队列）

- **缓冲区未满**：send 写入环形槽；必要时唤醒 recvq 等待的 G
- **缓冲区满**：send 进入 sendq，G 阻塞
- **缓冲区空**：recv 进入 recvq，G 阻塞

对应到两个核心指针：

- **`sendx`**：下一次写入位置（写指针）
- **`recvx`**：下一次读取位置（读指针）

所以你可以把 buffered channel 记成：**“环形队列 + 两个等待队列 + 一把锁”**。

## 5. 把“现象”映射到“原理”

- 用户看到卡在 `ch <- v` 或 `<-ch`  
  = runtime 把 G（sudog）挂到 `hchan.sendq/recvq` + `park`  
  = 匹配操作到来再 `ready`

## 6. close/closed channel 的阻塞语义（面试很爱问）

1. **close 之后再接收**：不会阻塞；会立刻返回零值 + `ok=false`（语义层面）
2. **close 之后再发送**：panic（语义层面）
3. **close 时在等待队列里的 G**：会被唤醒（senders/receivers 的处理细节见 `closechan`）

> 更完整的语义清单见同目录 `closed_channel相关特性.md`。

## 7. 为什么 channel 内部通常是一把“粗锁”？（回答模板）

- channel 操作需要同时维护：缓冲区计数、读写索引、等待队列、关闭标记等  
- 这些状态之间强相关，拆成多把锁容易引入复杂死锁/一致性问题  
- Go runtime 选择用一把锁保护整体，换来更易维护的正确性；性能依赖于“快路径”（配对/写入缓冲）尽量短

## 8. 延伸阅读（按函数名对照源码）

- `src/runtime/chan.go`：`chansend`、`chanrecv`、`closechan` 等

还可以顺手关注：

- `gopark` / `goready`：G 的 park/ready（挂起与就绪）
- `sudog`：等待节点结构（chan/mutex 等都会用到）

## 9. 自检

- 无缓冲 channel 是否需要 buffer？为什么？
- 为什么 channel 内部通常是一把锁保护整体状态？
 - `sendq/recvq` 里存的是 `g` 还是 `sudog`？为什么？

