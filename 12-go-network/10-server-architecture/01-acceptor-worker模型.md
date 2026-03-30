# acceptor-worker 模型

> **10-server-architecture · 高并发服务器架构**

## 内容大纲

- 单线程或少量线程 accept，分发到 worker 池
- 与半同步半异步：边界在哪条队列
- 背压：全连接队列、任务队列长度、拒绝策略
- 与 Go：Accept goroutine 加 chan 任务（经典写法）
- 对比：线程池与 goroutine 池

## 正文

（待补充）
