# channel 阻塞原理（实现向提纲）

## 1. channel 内部大致有什么

- **缓冲区**：有缓冲时为环形队列（**FIFO** 出队顺序，与课里「FIFO」目标一致）。  
- **sendq / recvq**：等待发送、等待接收的 **G 队列**（sudog 链表等，概念上叫 waitq）。  
- **锁**：保护 channel 内部状态（channel 是带锁的结构，不是无锁队列）。

type hchan struct {
    qcount   uint           // 缓冲区中元素数量
    dataqsiz uint           // 缓冲区大小（环形队列长度）
    buf      unsafe.Pointer // 指向环形队列的指针
    elemsize uint16         // 单个元素大小
    closed   uint32         // 通道是否关闭
    elemtype *_type         // 元素类型
    sendx    uint           // 发送索引（环形队列写入位置）
    recvx    uint           // 接收索引（环形队列读取位置）
    recvq    waitq          // 等待接收的 Goroutine 队列（recvg）
    sendq    waitq          // 等待发送的 Goroutine 队列（sendg）
    lock     mutex          // 互斥锁，保护所有内部状态
}

type waitq struct {
    first *sudog
    last  *sudog
}

## 2. 无缓冲：同步交换

- 发送方找到正在等待的接收方（或反之）→ **直接拷贝数据** 到对方栈或寄存器路径，双方就绪。  
- 对不上则当前 G **入队并阻塞**，调度器切换去跑别的 G。

## 3. 有缓冲：队列满/空

- **缓冲区未满**：send 写入环形槽，`recv` 从槽取；可能唤醒 recvq 里等待的 G。  
- **缓冲区满**：send 进 sendq，G 阻塞。  
- **缓冲区空**：recv 进 recvq，G 阻塞。

## 4. 与「阻塞协程现象」的对应

- 用户看到的「卡在 `<-` / `ch<-`」= runtime 把 G 挂在 channel 的队列上 + **park**；另一方操作到来时再 **ready**。

## 5. 延伸阅读

- `src/runtime/chan.go`：`chansend`、`chanrecv` 等。

## 6. 自检

- 无缓冲 channel 是否还需要环形 buffer？  
- 为什么 channel 内部通常要有一把大锁（粗粒度）？
