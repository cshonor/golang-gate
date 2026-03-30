# Listener 接口

> **06-go-net-internals · Go net 包源码级理解**

## 内容大纲

- Accept 返回 Conn；Addr 报告监听地址
- TCPListener：File、`SetDeadline`、`Close` 行为
- UnixListener、PacketConn 变体（扩展）
- 优雅关闭：`Close` 后 Accept 返回错误类型
- 与 `http.Server`：`Shutdown` 如何关 Listener

## 正文

（待补充）
