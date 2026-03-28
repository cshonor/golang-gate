# TCP 心跳保活

> **04-tcp · TCP 协议与编程**

## 内容大纲

- 内核 `SO_KEEPALIVE` 与默认周期（平台差异）
- 应用层心跳：协议设计、间隔与超时、与 NAT 会话表关系
- 与 HTTP/2 ping、gRPC keepalive 对照
- 误用：过密心跳打满 CPU/带宽
- Go：`SetKeepAlive`、`SetKeepAlivePeriod`（平台支持说明）

## 正文

（待补充）
