# 08 - panic 与 recover 捕获异常

## 1. 定位

在错误处理体系里，`panic/recover` 是**最后兜底**，不是常规业务分支。`defer` 里如何合并 `Close` 等错误见 [09 - defer 与错误处理协同](./09%20-%20defer与错误处理协同.md)。

## 2. 原则

- 可预期错误：返回 `error`
- 不可恢复错误：`panic`
- 边界兜底：`defer + recover`

## 3. 典型模板

```go
defer func() {
    if r := recover(); r != nil {
        // 记录日志/报警
    }
}()
```

