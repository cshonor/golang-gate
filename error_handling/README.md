# Error Handling（错误处理）

本目录聚焦 **Go 错误处理** 的核心写法与面试高频点：返回错误、包装、判等、类型判断、分层与日志边界。

补充：`panic/recover` 可以理解为**错误处理的最后兜底（latterly）**，但工程上应控制在边界使用（已并入本目录）。

## 文件索引

| 主题 | 笔记文件 |
|------|----------|
| 01 - 错误介绍 | [01 - 错误介绍.md](./01 - 错误介绍.md) |
| 02 - 错误基本语法 | [02 - 错误基本语法.md](./02 - 错误基本语法.md) |
| 03 - 自定义错误类型 | [03 - 自定义错误类型.md](./03 - 自定义错误类型.md) |
| 04 - 错误包装与错误链 | [04 - 错误包装与错误链.md](./04 - 错误包装与错误链.md) |
| 05 - errors.Is 与 errors.As 使用 | [05-errors.Is 与 errors.As 使用.md](./05-errors.Is 与 errors.As 使用.md) |
| 06 - panic 与 recover 捕获异常 | [06-panic 与 recover 捕获异常.md](./06-panic 与 recover 捕获异常.md) |
| 进阶总览（可背版） | [error_handling.md](./error_handling.md) |
| 兜底专题（可背版） | [panic_recover.md](./panic_recover.md) |

## 学习顺序建议

1. 先学 01-03（概念、语法、类型）  
2. 再学 04-05（错误链、Is/As 判定）  
3. 最后学 06（panic/recover 兜底）  
4. 复习时看 `error_handling.md` 和 `panic_recover.md`

