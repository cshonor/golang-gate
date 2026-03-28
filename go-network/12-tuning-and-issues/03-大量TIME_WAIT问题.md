# 大量 TIME_WAIT 问题

> **12-tuning-and-issues · 性能调优与线上问题**

## 内容大纲

- 成因：主动关闭方、短连接、HTTP 客户端
- 危害：端口、socket 结构占用、路由表压力（视规模）
- 手段：Keep-Alive、连接池、调内核（理解副作用）
- Linger 选项风险
- 与负载均衡：NAT 表项

## 正文

（待补充）
