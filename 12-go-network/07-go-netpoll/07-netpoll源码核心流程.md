# netpoll 源码核心流程

> **07-go-netpoll · Go netpoll 高并发核心**

## 内容大纲

- pollDesc：open、close、read、write、park
- netpollinit、netpollopen、netpollblock、netpollunblock
- epollctl 增删改事件集合
- 从 internal/poll 到 runtime 的边界
- 带问题阅读：一次 Accept 从 syscall 返回到用户 handler 的完整栈

## 扩写衔接

把 **「事件循环」** 与 **「G 调度」** 拼起来读：[08-netpoll与GMP调度深度联动](./08-netpoll与GMP调度深度联动.md)、[09-pollDesc等待队列与唤醒原理](./09-pollDesc等待队列与唤醒原理.md)。

## 正文

（待补充）
