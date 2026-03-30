# 异步 IO（AIO）

> **02-io-models · IO 模型全解**

## 内容大纲

- POSIX AIO 在 Linux 上的真实程度：部分场景仍线程池模拟
- io_uring：真异步与批量提交（可作为扩展）
- 网络 AIO：Windows IOCP 模型对照
- 与 Proactor：`WSARecv` + completion port
- Go 选型：为何多数网络代码仍走 netpoll 而非原生 AIO

## 正文

（待补充）
