# net.Listen 底层流程

> **06-go-net-internals · Go net 包源码级理解**

## 内容大纲

- 解析 `host:port`、wildcard、`SO_REUSEADDR`/`SO_REUSEPORT`（平台）
- bind/listen syscall 序列
- accept 与 netpoll：如何把新 fd 纳入 runtime
- 多监听：`ListenConfig`、`*TCPAddr`
- 常见坑：EADDRINUSE、IPv6 only、容器端口映射

## 正文

（待补充）
