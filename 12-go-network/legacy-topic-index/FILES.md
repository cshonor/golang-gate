# 按技术主题：逐篇笔记索引（路径相对本目录）

> 与 [`README.md`](./README.md) 同级的**完整文件名列表**；主学习顺序请优先使用仓库 [`../README.md`](../README.md) 的 **五阶段书本目录**。

## 01-io-fundamentals

**建议顺序**：`01` → `02`/`03` → `04` → `05` → `06`，再进 **02-io-models**。

- [01-什么是IO](./01-io-fundamentals/01-什么是IO.md)
- [02-阻塞与非阻塞](./01-io-fundamentals/02-阻塞与非阻塞.md)
- [03-同步与异步](./01-io-fundamentals/03-同步与异步.md)
- [04-用户态与内核态](./01-io-fundamentals/04-用户态与内核态.md)
- [05-缓冲区、页缓存、零拷贝](./01-io-fundamentals/05-缓冲区、页缓存、零拷贝.md)
- [06-文件IO vs 网络IO](./01-io-fundamentals/06-文件IO%20vs%20网络IO.md)

## 02-io-models

**建议顺序**：`01` 总览 → `02`～`03`（BIO/NIO）→ `04`～`07`（多路复用与 epoll）→ `08`～`09` → `10` 收束。

- [01-五大IO模型总览](./02-io-models/01-五大IO模型总览.md)
- [02-阻塞IO BIO](./02-io-models/02-阻塞IO%20BIO.md)
- [03-非阻塞IO NIO](./02-io-models/03-非阻塞IO%20NIO.md)
- [04-IO多路复用-select](./02-io-models/04-IO多路复用-select.md)
- [05-IO多路复用-poll](./02-io-models/05-IO多路复用-poll.md)
- [06-IO多路复用-epoll原理](./02-io-models/06-IO多路复用-epoll原理.md)
- [07-epoll-LT与ET模式](./02-io-models/07-epoll-LT与ET模式.md)
- [08-信号驱动IO](./02-io-models/08-信号驱动IO.md)
- [09-异步IO AIO](./02-io-models/09-异步IO%20AIO.md)
- [10-五大IO模型对比总结](./02-io-models/10-五大IO模型对比总结.md)

## 03-linux-networking

**建议顺序**：`01`～`03` → `04`～`05` → `06`～`07` → `08`。

- [01-socket是什么](./03-linux-networking/01-socket是什么.md)
- [02-file-descriptor-fd](./03-linux-networking/02-file-descriptor-fd.md)
- [03-socket创建流程](./03-linux-networking/03-socket创建流程.md)
- [04-三次握手内核做了什么](./03-linux-networking/04-三次握手内核做了什么.md)
- [05-四次挥手内核做了什么](./03-linux-networking/05-四次挥手内核做了什么.md)
- [06-TCP连接状态](./03-linux-networking/06-TCP连接状态.md)
- [07-半连接队列与全连接队列](./03-linux-networking/07-半连接队列与全连接队列.md)
- [08-socket缓冲区SO_RCVBUF与SO_SNDBUF](./03-linux-networking/08-socket缓冲区SO_RCVBUF与SO_SNDBUF.md)

## 04-tcp

**建议顺序**：`01`～`05` → `06`～`07` → `08`～`09`。

- [01-TCP特点](./04-tcp/01-TCP特点.md)
- [02-TCP报文结构](./04-tcp/02-TCP报文结构.md)
- [03-三次握手](./04-tcp/03-三次握手.md)
- [04-四次挥手](./04-tcp/04-四次挥手.md)
- [05-TCP重传、滑动窗口、拥塞控制](./04-tcp/05-TCP重传、滑动窗口、拥塞控制.md)
- [06-Go-TCP服务端](./04-tcp/06-Go-TCP服务端.md)
- [07-Go-TCP客户端](./04-tcp/07-Go-TCP客户端.md)
- [08-TCP心跳保活](./04-tcp/08-TCP心跳保活.md)
- [09-TCP长连接与短连接](./04-tcp/09-TCP长连接与短连接.md)

## 05-udp-multicast

- [01-UDP特点](./05-udp-multicast/01-UDP特点.md)
- [02-UDP报文结构](./05-udp-multicast/02-UDP报文结构.md)
- [03-Go-UDP服务端](./05-udp-multicast/03-Go-UDP服务端.md)
- [04-Go-UDP客户端](./05-udp-multicast/04-Go-UDP客户端.md)
- [05-组播与广播](./05-udp-multicast/05-组播与广播.md)
- [06-UDP可靠传输设计](./05-udp-multicast/06-UDP可靠传输设计.md)

## 06-go-net-internals

- [01-net.Conn接口](./06-go-net-internals/01-net.Conn接口.md)
- [02-Listener接口](./06-go-net-internals/02-Listener接口.md)
- [03-TCPConn结构](./06-go-net-internals/03-TCPConn结构.md)
- [04-UDPConn结构](./06-go-net-internals/04-UDPConn结构.md)
- [05-net.Dial底层流程](./06-go-net-internals/05-net.Dial底层流程.md)
- [06-net.Listen底层流程](./06-go-net-internals/06-net.Listen底层流程.md)
- [07-pollDesc核心结构与原理](./06-go-net-internals/07-pollDesc核心结构与原理.md)
- [08-网络超时与Deadline底层实现](./06-go-net-internals/08-网络超时与Deadline底层实现.md)
- [09-网络错误分类与处理](./06-go-net-internals/09-网络错误分类与处理.md)
- [10-连接关闭与资源泄漏排查](./06-go-net-internals/10-连接关闭与资源泄漏排查.md)

**建议顺序**：`01`～`06` 与 **`07`** 穿插读，再 **`08`～`10`**。

## 07-go-netpoll

- [01-Go为什么高并发](./07-go-netpoll/01-Go为什么高并发.md)
- [02-netpoll是什么](./07-go-netpoll/02-netpoll是什么.md)
- [03-epoll-kqueue-IOCP支持](./07-go-netpoll/03-epoll-kqueue-IOCP支持.md)
- [04-Go的IO模型到底是什么](./07-go-netpoll/04-Go的IO模型到底是什么.md)
- [05-Go的Read与Write为什么看起来阻塞](./07-go-netpoll/05-Go的Read与Write为什么看起来阻塞.md)
- [06-Goroutine与netpoll调度](./07-go-netpoll/06-Goroutine与netpoll调度.md)
- [07-netpoll源码核心流程](./07-go-netpoll/07-netpoll源码核心流程.md)
- [08-netpoll与GMP调度深度联动](./07-go-netpoll/08-netpoll与GMP调度深度联动.md)
- [09-pollDesc等待队列与唤醒原理](./07-go-netpoll/09-pollDesc等待队列与唤醒原理.md)
- [10-netpoll常见坑与优化](./07-go-netpoll/10-netpoll常见坑与优化.md)

**建议顺序**：`01` → `02`→`03`→`04`→`05` → `06`→`07` → **`08`～`10`** 与 **06 之 `07-pollDesc`** 交叉闭环。

## 08-framing-protocols

- [01-什么是粘包拆包](./08-framing-protocols/01-什么是粘包拆包.md)
- [02-为什么TCP会粘包](./08-framing-protocols/02-为什么TCP会粘包.md)
- [03-固定长度协议](./08-framing-protocols/03-固定长度协议.md)
- [04-分隔符协议](./08-framing-protocols/04-分隔符协议.md)
- [05-长度+报文协议（最常用）](./08-framing-protocols/05-长度+报文协议（最常用）.md)
- [06-Protobuf协议](./08-framing-protocols/06-Protobuf协议.md)
- [07-自定义私有协议](./08-framing-protocols/07-自定义私有协议.md)

## 09-http-internals

- [01-HTTP1.1](./09-http-internals/01-HTTP1.1.md)
- [02-HTTP2](./09-http-internals/02-HTTP2.md)
- [03-HTTP3](./09-http-internals/03-HTTP3.md)
- [04-Go-net-http底层](./09-http-internals/04-Go-net-http底层.md)
- [05-路由、ServeMux](./09-http-internals/05-路由、ServeMux.md)
- [06-http.Server结构体](./09-http-internals/06-http.Server结构体.md)
- [07-长连接、流水线](./09-http-internals/07-长连接、流水线.md)

## 10-server-architecture

- [01-acceptor-worker模型](./10-server-architecture/01-acceptor-worker模型.md)
- [02-reactor模型](./10-server-architecture/02-reactor模型.md)
- [03-多reactor](./10-server-architecture/03-多reactor.md)
- [04-goroutine-per-conn](./10-server-architecture/04-goroutine-per-conn.md)
- [05-协程池](./10-server-architecture/05-协程池.md)
- [06-连接池](./10-server-architecture/06-连接池.md)
- [07-高并发最佳实践](./10-server-architecture/07-高并发最佳实践.md)

## 11-nio-reactor

- [01-非阻塞socket](./11-nio-reactor/01-非阻塞socket.md)
- [02-Reactor模式](./11-nio-reactor/02-Reactor模式.md)
- [03-主从Reactor](./11-nio-reactor/03-主从Reactor.md)
- [04-Go如何实现Reactor](./11-nio-reactor/04-Go如何实现Reactor.md)

## 12-tuning-and-issues

- [01-TCP参数调优](./12-tuning-and-issues/01-TCP参数调优.md)
- [02-端口范围、backlog](./12-tuning-and-issues/02-端口范围、backlog.md)
- [03-大量TIME_WAIT问题](./12-tuning-and-issues/03-大量TIME_WAIT问题.md)
- [04-CLOSE_WAIT问题](./12-tuning-and-issues/04-CLOSE_WAIT问题.md)
- [05-大量连接占内存问题](./12-tuning-and-issues/05-大量连接占内存问题.md)
- [06-pprof网络调优](./12-tuning-and-issues/06-pprof网络调优.md)
- [07-高并发压测指标](./12-tuning-and-issues/07-高并发压测指标.md)

## 13-projects-optional

- [01-Echo服务器](./13-projects-optional/01-Echo服务器.md)
- [02-聊天室](./13-projects-optional/02-聊天室.md)
- [03-文件传输](./13-projects-optional/03-文件传输.md)
- [04-代理服务器](./13-projects-optional/04-代理服务器.md)
- [05-高性能网关](./13-projects-optional/05-高性能网关.md)

## 建议阅读顺序（Legacy 技术线）

按 **01 → 13** 子文件夹序号；与 **五阶段书本目录** 并行：书本章节的 `README` 会把你带回此处对应篇目。
