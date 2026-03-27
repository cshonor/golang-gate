# Defer Traps（defer 易错点）

本目录整理 **defer 的常见坑**：资源泄漏、锁没释放、返回值误解、循环 defer、recover 位置等。

## 文件索引

| 主题 | 笔记文件 |
|------|----------|
| defer 必背坑点与正确写法 | [defer_traps.md](./defer_traps.md) |

## 学习顺序建议

1. 先把 `defer_traps.md` 的“5 大坑”过一遍  
2. 再回看你项目里所有 `defer`：文件/连接/锁/计时器/trace 是否都能正确释放

