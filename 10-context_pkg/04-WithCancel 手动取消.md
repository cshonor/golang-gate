# 04 - WithCancel：手动取消

## 1. 什么时候用 WithCancel

当你启动后台 goroutine 做工作时，需要一个“停止开关”：

- 请求被取消
- 主流程结束
- 发生错误，需要提前停掉下游工作

## 2. 标准模板（必须记住 defer cancel）

```go
ctx, cancel := context.WithCancel(parent)
defer cancel()

go func() {
    select {
    case <-ctx.Done():
        return
    case job := <-jobs:
        _ = job
    }
}()
```

## 3. 常见坑

- 忘记 `cancel()`：可能导致 goroutine 泄漏、timer/资源释放延迟
- 只 cancel 但下游不监听 `ctx.Done()`：等于没取消

## 4. 面试一句话

> `WithCancel` 用于“手动取消”，让一组 goroutine 能统一退出；取消要靠 `Done()` 协作实现。

