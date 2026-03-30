# 09 - 中间件实战：用 context 透传 trace-id

本篇补齐生产常用的“链路标识”：把 `trace-id` 放进 `context`，日志与下游调用统一取。

---

## 1. 约定 key 与存取函数

```go
type ctxKey string

const traceIDKey ctxKey = "trace_id"

func WithTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceIDKey, traceID)
}

func TraceID(ctx context.Context) (string, bool) {
	v := ctx.Value(traceIDKey)
	s, ok := v.(string)
	return s, ok
}
```

---

## 2. 中间件：从请求头读 trace-id

伪代码（HTTP）：

- 若请求头已有 `X-Trace-Id`：沿用（便于跨服务传递）
- 否则生成一个（UUID/雪花等）

```go
func TraceMiddleware(next Handler) Handler {
	return func(ctx context.Context, req *Request) error {
		id := req.Header.Get("X-Trace-Id")
		if id == "" {
			id = newTraceID()
		}
		ctx = WithTraceID(ctx, id)
		return next(ctx, req)
	}
}
```

---

## 3. 业务层与日志：统一从 ctx 取

```go
id, _ := TraceID(ctx)
logger.With("trace_id", id).Info("handle request")
```

---

## 4. 下游调用：把 trace-id 透传出去

- 调用下游 HTTP/RPC 时，从 ctx 取 trace-id 并写入 header/metadata。

---

## 延伸阅读

- [06-WithValue 数据透传](./06-WithValue%20数据透传.md)
- [08-context常见陷阱与反模式](./08-context常见陷阱与反模式.md)
- [07-context在中间件中的实战](./07-context在中间件中的实战.md)
