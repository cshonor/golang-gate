# 04 - errors.Join：错误链补充

## 1. 解决什么问题

并发执行多个任务时，你可能会拿到多个错误，过去要么只返回第一个，要么自己拼接字符串。

`errors.Join(err1, err2, ...)` 可以把多个错误合成一个 error。

## 2. 典型场景：并发汇总错误

```go
var errs []error
// ... append err
return errors.Join(errs...)
```

## 3. 和 errors.Is/As 的关系

Join 后的 error 仍然可以用 `errors.Is/As` 判断其内部是否包含某个错误（适合和 `03-error_handling/05` 配套）。

