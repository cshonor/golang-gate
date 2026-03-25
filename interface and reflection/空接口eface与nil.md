# 1. 空接口 `interface{}` / `any` 的底层结构（面试必考）

> Go 1.18 起 `any` 是 `interface{}` 的别名，面试说「空接口 / any」均可。

## 1）空接口是什么

- 没有定义任何方法的接口：`interface{}`（或 `any`）
- 可以**接收任意类型的值**：`int`、`string`、`struct`、指针、`nil` 都能装（注意 **`nil` 无类型** 与 **有类型的 nil 指针** 的区别）

## 2）真正的底层结构

Go 源码里，**空接口**对应结构体 `eface`（`runtime2.go` 概念，实际定义在 `runtime` 包内部）：

```go
type eface struct {
    _type *_type            // 类型元信息：到底是什么类型
    data  unsafe.Pointer   // 数据：指向真实值的地址
}
```

一句话记：**空接口 = 类型指针 + 数据指针**。

- `_type`：这个值**是什么类型**（`int`？`string`？`*User`？）
- `data`：数据存放位置（通常是指针；小对象实现细节以实现为准）

## 3）带方法的接口 vs 空接口（简单区分）

| | 空接口 | 带方法的接口（如 `io.Reader`） |
|--|--------|--------------------------------|
| 底层 | `eface`：仅类型 + 数据 | `iface`：类型 + 数据 + **方法表（itab）** |
| 典型用途 | `any` 传任意值 | 多态、只关心行为 |

## 4）最经典坑：接口值不是 nil ≠ 里面「存了 nil 指针」

```go
var p *int = nil
var i interface{} = p

fmt.Println(i == nil) // false
```

原因：

- 此时 `i` 已经是一个**非 nil 的接口值**：`_type` 指向 `*int`，类型信息存在。
- **接口值为 nil** 需要：`_type == nil` 且 `data == nil`（概念上两者都空）。这里 `_type` 非空，故 `i != nil`。

面试表述：**判断接口是否为 nil，要看整个接口值是否为零值，不能只看「里面是不是 nil 指针」。**
