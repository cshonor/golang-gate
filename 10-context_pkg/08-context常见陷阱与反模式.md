# 08 - context 常见陷阱与反模式

本篇补齐 `context` 的工程避坑：`WithValue` 的断言风险、ctx 生命周期边界等。

---

## 1. WithValue 的类型断言风险

`Value` 返回 `any`，取值需要断言：

```go
v := ctx.Value(traceIDKey)
id, ok := v.(string)
if !ok {
	// key 不存在或类型不匹配
}
```

**坑**：

- key 选错 / key 冲突（不同包用同一个字符串）
- 写入类型与读取类型不一致（写了 `[]byte`，读按 `string` 断言）

**建议**：

- key 用自定义类型（见 [06-WithValue 数据透传](./06-WithValue%20数据透传.md)）
- 封装 `GetTraceID(ctx) (string, bool)` 统一取值与断言

---

## 2. 不要把 context 当成业务参数包

只放“小且稳定”的请求元信息（trace id / user id）。不要放：

- 大对象（大 slice/map/struct）
- 可变业务状态（会变、会被频繁更新）

---

## 3. 生命周期边界：不要把长生命周期 ctx 传给短生命周期函数（以及反过来）

核心是：**ctx 的取消/超时语义应与“这段工作”的生命周期一致**。

### 常见反模式

- 用 `context.Background()` 调用一切：下游永远无法取消。
- 把“全局 ctx”（永不 cancel）传给请求级工作：超时与取消失效。
- 把请求 ctx 存到全局结构体：请求结束后 ctx 被取消，后续异步任务拿到的是已取消 ctx。

### 建议

- 请求处理：用上游传入的 `ctx`（HTTP/RPC 框架提供）。
- 后台任务：用独立 ctx（可被自己的停止逻辑 cancel），不要复用请求 ctx。

---

## 4. 自检

- [ ] 每个 goroutine 都监听 `ctx.Done()`？
- [ ] `WithTimeout` 的 `cancel()` 是否 `defer` 调用？
- [ ] `WithValue` 的 key 是否自定义类型且取值封装？

---

## 延伸阅读

- [06-WithValue 数据透传](./06-WithValue%20数据透传.md)
- [07-context在中间件中的实战](./07-context在中间件中的实战.md)
