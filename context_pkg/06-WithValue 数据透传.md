# 06 - WithValue：数据透传（你说的“穿透”）

## 1. WithValue 做什么

把“请求级元信息”挂到 ctx 上，向下游透传：

- trace id / request id
- user id
- token（或解析后的信息）

## 2. 最重要的边界（背）

- **只放小且稳定的元信息**  
- 不要放大对象、不放业务可变参数、不把 ctx 当 map 用

原因：可维护性与内存占用会失控。

## 3. key 的最佳实践

避免 key 冲突：用自定义类型当 key。

```go
type ctxKey string

const traceIDKey ctxKey = "trace_id"

ctx = context.WithValue(ctx, traceIDKey, "abc123")
```

## 4. 取值模板

```go
v := ctx.Value(traceIDKey)
id, _ := v.(string)
```

## 5. 面试一句话

> WithValue 用于“跨层透传请求元信息”，但要克制使用，避免把业务状态塞进 ctx。

