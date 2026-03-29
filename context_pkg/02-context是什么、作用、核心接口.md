# 02 - context 是什么、作用、核心接口

## 1. 一句话结论（背）

> **Context = 请求的生命周期控制器**：用来做取消、超时、以及请求级数据透传；中间件就是一层层包装并传递 ctx。

## 2. context 解决的 3 个核心问题

1. **取消**：上游取消后，下游 goroutine 要能尽快退出  
2. **超时/截止时间**：请求超过 SLA 直接终止，避免堆积  
3. **请求级元数据透传**：trace id、user id、token、request id（小且稳定的元信息）

## 3. 核心接口（只记这 4 个）

```go
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key any) any
}
```

### 3.1 `Done()` / `Err()`：取消协作的关键

- `Done()`：关闭表示“该 ctx 已取消/超时”
- `Err()`：返回取消原因（`context.Canceled` / `context.DeadlineExceeded`）

## 4. 工程规范（面试必背）

- **入口创建/获取 ctx**：HTTP/gRPC handler 里拿到 ctx  
- **逐层传递**：函数签名优先 `func(ctx context.Context, ...)`  
- **不长期持有 ctx**：不要塞进 struct 里当字段缓存（会让生命周期失控）

