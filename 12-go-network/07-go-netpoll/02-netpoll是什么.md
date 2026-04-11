# netpoll 是什么

> **07-go-netpoll · Go netpoll 高并发核心**

## 内容大纲

- runtime 内部对可读可写事件的统一抽象
- 每个网络 fd 关联 pollDesc；g 在 fd 上 park
- 就绪时 netpoll 把 g 标记 runnable 推回队列
- 与 epoll_wait、kevent、IOCP 的桥接
- 查找入口：runtime 目录下 netpoll 相关源码

## 扩写索引（闭环）

- **06 桥梁（`internal/poll`）**：[07-pollDesc核心结构与原理](../06-go-net-internals/07-pollDesc核心结构与原理.md)  
- **G 如何睡/醒 + 调度**：[08-netpoll与GMP调度深度联动](./08-netpoll与GMP调度深度联动.md)、[09-pollDesc等待队列与唤醒原理](./09-pollDesc等待队列与唤醒原理.md)  
- **线上坑**：[10-netpoll常见坑与优化](./10-netpoll常见坑与优化.md)

## 正文

（待补充）
