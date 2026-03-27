# Error Handling（错误处理）

本目录聚焦 **Go 错误处理** 的核心写法与面试高频点：返回错误、包装、判等、类型判断、分层与日志边界。

补充：`panic/recover` 可以理解为**错误处理的最后兜底（latterly）**，但工程上应控制在边界使用（已并入本目录）。

## 文件索引

| 主题 | 笔记文件 |
|------|----------|
| 错误处理核心套路（可背） | [error_handling.md](./error_handling.md) |
| 最后兜底：panic / recover | [panic_recover.md](./panic_recover.md) |

## 学习顺序建议

1. 先把 `error_handling.md` 里“**三件事：返回/包装/判断**”背熟  
2. 再用同一套套路去看你项目里：文件 IO、网络请求、数据库调用的错误链路

