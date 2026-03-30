# netpoll 是什么

> **07-go-netpoll · Go netpoll 高并发核心**

## 内容大纲

- runtime 内部对可读可写事件的统一抽象
- 每个网络 fd 关联 pollDesc；g 在 fd 上 park
- 就绪时 netpoll 把 g 标记 runnable 推回队列
- 与 epoll_wait、kevent、IOCP 的桥接
- 查找入口：runtime 目录下 netpoll 相关源码

## 正文

（待补充）
