# net.Listen 底层流程

> **06-go-net-internals · Go net 包源码级理解**

## 内容大纲

- 解析 `host:port`、wildcard、`SO_REUSEADDR`/`SO_REUSEPORT`（平台）
- bind/listen syscall 序列
- accept 与 netpoll：如何把新 fd 纳入 runtime
- 多监听：`ListenConfig`、`*TCPAddr`
- 常见坑：EADDRINUSE、IPv6 only、容器端口映射

## 与 netpoll 的桥梁（扩写）

`Accept` 得到的 **`netFD`** 需 **注册到 netpoll** 才能融入 Go 的网络调度模型。见 **[07-pollDesc核心结构与原理](./07-pollDesc核心结构与原理.md)** 与 **07** 章 [09-pollDesc等待队列与唤醒原理](../07-go-netpoll/09-pollDesc等待队列与唤醒原理.md)。

## 正文

（待补充）
