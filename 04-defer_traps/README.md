# Defer Traps（defer 易错点）

本目录整理 **defer 的常见坑**：资源泄漏、锁没释放、返回值误解、循环 defer、`recover` 位置等；并补充与 **`error` 体系**、工程规范相关的协同章节。

**与 `03-error_handling` 的关系**：规则与陷阱见本目录；**`Close` / 事务 / `%w` / `errors.Join`** 等见 [07](./07-defer%20与错误处理协同.md) 及 `../03-error_handling/` 对应篇。

## 文件索引

| 主题 | 笔记文件 |
|------|----------|
| 01 - defer 基本用法 | [01-defer 基本用法.md](./01-defer%20基本用法.md) |
| 02 - defer 执行顺序规则 | [02-defer 执行顺序规则.md](./02-defer%20执行顺序规则.md) |
| 03 - defer 参数预计算陷阱 | [03-defer 参数预计算陷阱.md](./03-defer%20参数预计算陷阱.md) |
| 04 - defer 与 return 执行流程 | [04-defer 与 return 执行流程.md](./04-defer%20与%20return%20执行流程.md) |
| 05 - defer 在循环中的风险 | [05-defer 在循环中的风险.md](./05-defer%20在循环中的风险.md) |
| 06 - defer 底层实现原理 | [06-defer 底层实现原理.md](./06-defer%20底层实现原理.md) |
| 07 - defer 与错误处理协同 | [07-defer 与错误处理协同.md](./07-defer%20与错误处理协同.md) |
| 08 - defer 最佳实践与反模式 | [08-defer 最佳实践与反模式.md](./08-defer%20最佳实践与反模式.md) |
| 总览（可背版） | [defer_traps.md](./defer_traps.md) |

## 学习顺序建议

**主线**：`01 → 02 → 03 → 04 → 05 → 06 → 07 → 08`

1. **01–03**：会用、顺序、参数预计算陷阱  
2. **04–05**：return 与 defer 时序、循环 defer  
3. **06**：实现与性能直觉  
4. **07**：与 **`error` / 事务 / recover** 协同（工程刚需）  
5. **08**：规范与反模式总表（复习、面试）  
6. 速背：`defer_traps.md`  

**关联模块**：[错误处理（03-error_handling）](../03-error_handling/README.md)

