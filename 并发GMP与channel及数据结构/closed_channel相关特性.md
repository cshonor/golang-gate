# closed channel 相关特性

## 1. `close(ch)` 之后的行为

| 操作 | 结果 |
|------|------|
| **再 `close`** | panic：`close of closed channel` |
| **发送** `ch <- v` | panic：`send on closed channel` |
| **接收** `<-ch` | 读完缓冲内剩余数据后，**零值 + false**（`v, ok := <-ch` 中 `ok==false`） |
| **带 `ok` 的接收** | 缓冲空且已关闭 → 零值，`ok=false` |

## 2. 设计意图

- **只应由发送方关闭**（或唯一拥有者）：接收方通常用 `range` 或 `ok` 判断结束。  
- 多发送方场景不要随意 `close`，易竞态；可用 **额外信号 channel** 或 `WaitGroup` 协调。

## 3. `for range ch`

- 等价于不断接收，直到 channel **关闭且缓冲读完** 才退出循环。

## 4. 与 `select`

- 对已关闭且无数据的 channel，`case <-ch` 会 **立刻就绪**，读到零值；需用 `ok` 或配合其他逻辑区分「真零值」与「已关闭」。

## 5. 易错点

- 把 **关闭** 当成「通知消费者」可以，但不要 **多 goroutine 同时 close**。  
- **nil channel** 上 `select` 永久阻塞，与 closed 不同。

## 6. 自检

- 为什么对已关闭 channel 发送会 panic，而接收不会？  
- 如何区分「业务上合法的零值」和「channel 已关」？
