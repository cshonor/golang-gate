# Chapter 03 — Socket-Level Programming（Socket 级编程）

原书第 3 章：TCP/UDP、地址、DNS、超时与 KeepAlive、Raw socket 等。  
本仓库中 **绝大部分「硬核」笔记**落在 **`legacy-topic-index/`** 下列目录，请按组阅读。

## Linux / fd / 握手挥手

- [01-socket是什么](../../legacy-topic-index/03-linux-networking/01-socket是什么.md)
- [02-file-descriptor-fd](../../legacy-topic-index/03-linux-networking/02-file-descriptor-fd.md)
- [03-socket创建流程](../../legacy-topic-index/03-linux-networking/03-socket创建流程.md)
- [04-三次握手内核做了什么](../../legacy-topic-index/03-linux-networking/04-三次握手内核做了什么.md)
- [05-四次挥手内核做了什么](../../legacy-topic-index/03-linux-networking/05-四次挥手内核做了什么.md)
- [06-TCP连接状态](../../legacy-topic-index/03-linux-networking/06-TCP连接状态.md)
- [07-半连接队列与全连接队列](../../legacy-topic-index/03-linux-networking/07-半连接队列与全连接队列.md)
- [08-socket缓冲区SO_RCVBUF与SO_SNDBUF](../../legacy-topic-index/03-linux-networking/08-socket缓冲区SO_RCVBUF与SO_SNDBUF.md)

## TCP 协议与 Go 编程

- [01-TCP特点](../../legacy-topic-index/04-tcp/01-TCP特点.md) … [09-TCP长连接与短连接](../../legacy-topic-index/04-tcp/09-TCP长连接与短连接.md)（见该目录内全部文件）

## UDP / 组播

- [01-UDP特点](../../legacy-topic-index/05-udp-multicast/01-UDP特点.md) … [06-UDP可靠传输设计](../../legacy-topic-index/05-udp-multicast/06-UDP可靠传输设计.md)

## Go `net` 与 netpoll（高并发关键）

- [06-go-net-internals](../../legacy-topic-index/06-go-net-internals/) 下各篇
- [07-go-netpoll](../../legacy-topic-index/07-go-netpoll/) 下各篇

## IO 模型（理解「阻塞 / 多路复用」）

- [02-io-models](../../legacy-topic-index/02-io-models/) 下各篇
- [01-io-fundamentals](../../legacy-topic-index/01-io-fundamentals/) 下各篇

## 调优与线上问题（与连接治理相关）

- [12-tuning-and-issues](../../legacy-topic-index/12-tuning-and-issues/) 下各篇

## 可选实战

- [13-projects-optional](../../legacy-topic-index/13-projects-optional/)（Echo、聊天室、代理等）
