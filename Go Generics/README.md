# Go Generics（泛型）

本目录整理 **Go 1.18+ 泛型**：类型参数、约束、泛型函数/类型/方法，以及与 `interface{}`、反射的取舍。示例中 **比较大小的约束**优先使用 **Go 1.21+** 标准库 **`cmp.Ordered`**；更早版本可用 `golang.org/x/exp/constraints` 或手写联合类型。

**进阶笔记总入口**见仓库根目录 [README.md](../README.md)；**数据结构**见 [`../datastruct/README.md`](../datastruct/README.md)。

## 文件索引

| 顺序 | 主题 | 笔记文件 |
|------|------|----------|
| 1 | 介绍与动机 | [1.introduce.md](./1.introduce.md) |
| 2 | 泛型语法（Go 版） | [2. 泛型语法（Go 版）.md](./2.%20泛型语法（Go%20版）.md) |
| 3 | 泛型函数 | [3. generic-function.md](./3.%20generic-function.md) |
| 4 | 泛型结构体 | [4. generic-struct.md](./4.%20generic-struct.md) |
| 5 | 约束（constraints） | [5. 泛型约束.md](./5.%20泛型约束.md) |
| 6 | 泛型方法 | [6. 泛型方法.md](./6.%20泛型方法.md) |

> 链接中的 `%20` 为空格编码，便于在 GitHub 上正确跳转。

## 学习顺序建议

1. **1 → 2**：为什么要有泛型、语法长什么样。  
2. **3 → 4**：从函数到自定义泛型类型。  
3. **5**：`comparable`、`constraints`、`~` 底层类型等（面试常问）。  
4. **6**：方法上的类型参数、与「类型自己的方法」区别。

## 速记

- **类型参数**：`func F[T any](x T)`、`type List[T any] struct { ... }`  
- **约束**：`any`、`comparable`、`cmp.Ordered`（1.21+）、自定义 `interface{ A | B }`  
- **何时不用泛型**：只在一处用到、或反射/插件化更合适时不必硬上

## 延伸阅读

- 官方：[Type Parameters Proposal](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md)（设计文档）  
- 源码习惯：`src/go/types`、`cmd/compile` 中与泛型相关的类型检查（按需）
