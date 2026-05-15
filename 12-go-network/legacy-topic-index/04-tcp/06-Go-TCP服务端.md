# Go TCP 服务端

> **04-tcp · TCP 协议与编程**  
> 上篇：[05-TCP重传、滑动窗口、拥塞控制](./05-TCP重传、滑动窗口、拥塞控制.md)；socket 与队列见 [03-linux](../03-linux-networking/03-socket创建流程.md)。

---

## 一、最小可运行骨架

```go
ln, err := net.Listen("tcp", ":8080")
if err != nil { /* log fatal */ }
defer ln.Close()

for {
    conn, err := ln.Accept()
    if err != nil { /* log continue */ }
    go handle(conn) // 或投递 worker pool
}
```

- **`Listen`**：对应 **`socket`+`bind`+`listen`**。  
- **`Accept`**：从 **全连接队列** 取 fd（内核已完成握手，见 linux 篇）。

---

## 二、并发模型取舍

| 模型 | 优点 | 风险 |
|------|------|------|
| **每连接一 goroutine** | 代码简单 | 海量连接时 **调度与内存** 压力 |
| **固定 worker + channel** | 背压可控 | 需设计 **任务丢弃/排队上限** |
| **多 `Listen`（`SO_REUSEPORT`）** | 多进程扩容 | 平台与内核版本差异 |

---

## 三、工程必备

1. **`conn.SetDeadline` / `ReadDeadline`**：防慢客户端占满 goroutine。  
2. **`context` 取消**：与 **关停** 联动，注意 **不会自动关 conn**。  
3. **错误处理**：`errors.Is(err, net.ErrClosed)`、`syscall.ECONNRESET` 等（见 `06-go-net-internals/09`）。  
4. **优雅退出**：`Close` listener → `Accept` 返回 **`use of closed network connection`**，跳出循环；在途 conn 用 **`WaitGroup`** 或 **shutdown 信号** 收尾。

---

## 四、与内核调优衔接

- **`Accept` 慢** → **全连接队列满** → 客户端超时。  
- **`SetReadBuffer`** 等 → [03-linux 08](../03-linux-networking/08-socket缓冲区SO_RCVBUF与SO_SNDBUF.md)。

---

## 五、极简总结

- **`Listen` + `Accept` 循环** 是服务端心脏。  
- **并发模型** 决定你能不能扛 **C10K/C10M**。  
- **超时 + 优雅退出** 和 **协议正确性** 一样重要。

---

## 导航

- 上一篇：[05-TCP重传、滑动窗口、拥塞控制](./05-TCP重传、滑动窗口、拥塞控制.md)  
- 下一篇：[07-Go-TCP客户端](./07-Go-TCP客户端.md)
