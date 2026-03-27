# 04 - WithTimeout：超时控制

## 1. 什么时候用 WithTimeout

任何可能卡住的外部依赖都要有超时：

- DB 查询
- HTTP 请求
- RPC 调用
- 读写文件/队列（视业务）

## 2. 标准模板（必须 `defer cancel()`）

```go
ctx, cancel := context.WithTimeout(parent, 200*time.Millisecond)
defer cancel()

err := call(ctx)
if err != nil {
    return err
}
```

## 3. WithTimeout vs time.After

工程上优先 `WithTimeout`：

- 更容易“跨层传递”到下游（DB/RPC）
- 与取消统一（`ctx.Done()`）

## 4. 如何识别超时错误

```go
if errors.Is(err, context.DeadlineExceeded) {
    // 超时
}
```

