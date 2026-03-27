# 04 - defer 与 return 执行流程

## 1. 执行顺序

函数返回时大致顺序：

1. 计算返回值
2. 执行 defer
3. 真正返回

## 2. 命名返回值示例

```go
func f() (ret int) {
    defer func() { ret++ }()
    return 1
}
// 返回 2
```

## 3. 结论

defer 可修改命名返回值，项目里应谨慎使用，避免可读性下降。

