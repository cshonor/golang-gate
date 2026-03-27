# 01 - defer 基本用法

## 1. defer 做什么

`defer` 用于在函数返回前执行“收尾逻辑”。

常见用途：

- `file.Close()`
- `mu.Unlock()`
- `tx.Rollback()`

## 2. 模板

```go
f, err := os.Open(path)
if err != nil { return err }
defer f.Close()
```

## 3. 记忆点

拿到资源后第一时间 `defer`，是最稳的资源管理习惯。

