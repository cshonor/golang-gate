# IO 多路复用：poll

> **02-io-models · IO 模型全解**

## 内容大纲

- `pollfd` 结构：fd、events、revents；无固定 1024 上限但仍 O(n) 扫描
- 水平触发语义：与 select 类似，易写不易踩 ET 坑
- 性能瓶颈：大量连接时每次传入全量数组、内核扫描全表
- 与 select 选型：简单服务、连接中等时够用
- 过渡到 epoll：何时 poll 成为明显短板（量化直觉）

## 正文

（待补充）
