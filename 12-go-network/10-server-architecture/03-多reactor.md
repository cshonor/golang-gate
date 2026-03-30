# 多 Reactor

> **10-server-architecture · 高并发服务器架构**

## 内容大纲

- 主从：主 accept，子 reactor 处理已连接 fd
- 水平扩展：CPU 核数、SO_REUSEPORT（视平台）
- 负载均衡：连接分配、惊群历史与现状
- 与 Nginx、Netty 线程模型对照
- 在 Go 中模拟的思路（多 Listen、多进程等）

## 正文

（待补充）
