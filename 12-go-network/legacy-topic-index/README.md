# 按技术主题的深度笔记（Legacy 索引）

本目录为 **《Network Programming with Go》五阶段书本结构** 引入之前的笔记编排方式：按 **IO 模型 / Linux 网络 / TCP / netpoll / HTTP** 等技术线拆文件夹，正文路径未改，便于历史链接与 `tools/apply_outlines.py` 继续工作。

**主入口**：回到 [`../README.md`](../README.md)，按 **书本章节（五阶段）** 阅读；各章 `README` 内会链到本目录下对应篇目。

| 目录 | 主题 |
|------|------|
| `01-io-fundamentals` | IO 核心理论 |
| `02-io-models` | IO 模型全解 |
| `03-linux-networking` | Linux 网络底层 |
| `04-tcp` | TCP 协议与编程 |
| `05-udp-multicast` | UDP 与组播 |
| `06-go-net-internals` | Go `net` 包源码级理解 |
| `07-go-netpoll` | Go netpoll 高并发核心 |
| `08-framing-protocols` | 粘包拆包与协议设计 |
| `09-http-internals` | HTTP 标准库底层 |
| `10-server-architecture` | 高并发服务器架构 |
| `11-nio-reactor` | 非阻塞 IO 与 Reactor |
| `12-tuning-and-issues` | 性能调优与线上问题 |
| `13-projects-optional` | 实战项目（可选） |

**逐篇链接（完整）**：[FILES.md](./FILES.md)
