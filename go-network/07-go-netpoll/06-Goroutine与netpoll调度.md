# Goroutine 与 netpoll 调度

> **07-go-netpoll · Go netpoll 高并发核心**

## 内容大纲

- netpoll 在调度循环中的调用点：findrunnable 等（概念）
- 批量就绪：一次 epoll_wait 唤醒多个 g 的公平性
- 与 sysmon、抢占的关系（浅尝）
- 高负载下：poll 延迟、尾延迟分析思路
- 反模式：在热点路径滥用 runtime.LockOSThread

## 正文

（待补充）
