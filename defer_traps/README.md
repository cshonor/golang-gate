# Defer Traps（defer 易错点）

本目录整理 **defer 的常见坑**：资源泄漏、锁没释放、返回值误解、循环 defer、recover 位置等。

## 文件索引

| 主题 | 笔记文件 |
|------|----------|
| 01 - defer 基本用法 | [01-defer 基本用法.md](./01-defer 基本用法.md) |
| 02 - defer 执行顺序规则 | [02-defer 执行顺序规则.md](./02-defer 执行顺序规则.md) |
| 03 - defer 参数预计算陷阱 | [03-defer 参数预计算陷阱.md](./03-defer 参数预计算陷阱.md) |
| 04 - defer 与 return 执行流程 | [04-defer 与 return 执行流程.md](./04-defer 与 return 执行流程.md) |
| 05 - defer 在循环中的风险 | [05-defer 在循环中的风险.md](./05-defer 在循环中的风险.md) |
| 06 - defer 底层实现原理 | [06-defer 底层实现原理.md](./06-defer 底层实现原理.md) |

## 学习顺序建议

1. 先学 01-03（语义、顺序、参数陷阱）  
2. 再学 04-05（return 流程、循环风险）  
3. 最后学 06（实现原理）

