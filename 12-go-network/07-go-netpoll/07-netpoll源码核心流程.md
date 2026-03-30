# netpoll 源码核心流程

> **07-go-netpoll · Go netpoll 高并发核心**

## 内容大纲

- pollDesc：open、close、read、write、park
- netpollinit、netpollopen、netpollblock、netpollunblock
- epollctl 增删改事件集合
- 从 internal/poll 到 runtime 的边界
- 带问题阅读：一次 Accept 从 syscall 返回到用户 handler 的完整栈

## 正文

（待补充）
