# Go 的 IO 模型到底是什么？

> **07-go-netpoll · Go netpoll 高并发核心**  
> 上篇：[03-epoll-kqueue-IOCP支持](./03-epoll-kqueue-IOCP支持.md)；同步/异步定义见 [01-io-fundamentals/03](../01-io-fundamentals/03-同步与异步.md)。

---

## 一、对外：同步阻塞式 API

```go
n, err := conn.Read(buf) // 看起来「卡在这里直到有数据」
```

对业务：**同步**——这次调用返回时，**数据已在 `buf`**（或明确错误/`EOF`）。

---

## 二、对内：同步 IO + 非阻塞 fd + 多路复用（Reactor 色彩）

1. **fd 非阻塞**：无数据 → **`EAGAIN`**。  
2. **`internal/poll`**：把 **等可读** 转成 **`runtime_pollWait`**。  
3. **netpoll**：**`epoll_wait` 等** 批量阻塞在 **多 fd** 上。  
4. **就绪**：唤醒 **G**，回到用户态 **`Read`** 里继续 **syscall 拷贝数据**。

因此：**内核通知「就绪」**，用户态 **`Read` 完成拷贝**——教材常称 **Reactor** 范式；**不是**「内核把数据直接异步拷完再通知」的 **Proactor / POSIX AIO** 那种 **`io_uring` 全链路心智**（扩展见 [02-io-models/09](../02-io-models/09-异步IO%20AIO.md)）。

---

## 三、对比表（面试）

| 框架/语言 | 典型心智 |
|-----------|----------|
| **Go `net`** | **G 阻塞 + epoll/kqueue + 非阻塞 fd** |
| **Java NIO + Netty** | **Reactor + 线程池**（概念相近） |
| **Node** | **单线程事件循环 + 非阻塞**（形态不同） |

---

## 四、极简总结

- **Go 网络默认**：**同步 API + 异步就绪通知 + G 调度**。  
- **别称**：**「同步非阻塞 + 多路复用」** 口语版。  
- 深入：[08-netpoll与GMP调度深度联动](./08-netpoll与GMP调度深度联动.md)、[11-nio-reactor](../11-nio-reactor/02-Reactor模式.md)。

---

## 导航

- 上一篇：[03-epoll-kqueue-IOCP支持](./03-epoll-kqueue-IOCP支持.md)  
- 下一篇：[05-Go的Read与Write为什么看起来阻塞](./05-Go的Read与Write为什么看起来阻塞.md)
