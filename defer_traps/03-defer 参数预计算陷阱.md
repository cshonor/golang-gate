# 03 - defer 参数预计算陷阱

## 1. 核心陷阱

`defer` 的参数在**注册 defer 当下**就求值，不是执行时求值。

```go
x := 1
defer fmt.Println(x) // 这里就记住了 1
x = 2
// 输出 1
```

## 2. 如果要读“最终值”

用闭包：

```go
x := 1
defer func() { fmt.Println(x) }()
x = 2
// 输出 2
```

