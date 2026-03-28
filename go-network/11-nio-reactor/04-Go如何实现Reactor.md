# Go 如何实现 Reactor

> **11-nio-reactor · 非阻塞 IO 与 Reactor 模型**

## 内容大纲

- netpoll 即内核事件源；M 执行 goroutine 即 worker
- 手写最小 echo：是否必要用 epoll（通常不必要）
- 与 cgo 或 syscall.EpollWait 自建 loop 的边界
- 性能：不要与 runtime 打架
- 结论：理解 Reactor 有助于读 Go 源码，不必业务里再造轮子

## 正文

（待补充）
