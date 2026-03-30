# Go 的 Read/Write 为什么看起来阻塞

> **07-go-netpoll · Go netpoll 高并发核心**

## 内容大纲

- netFD.Read：EAGAIN 时调用 pollWait
- gopark 与 netpoll goroutine 唤醒路径
- 与 SetDeadline：timer 与 poll 取消
- 阻塞的用户感知与线程状态
- 实验：strace 或 dlv 小 demo 观察（可选）

## 正文

（待补充）
