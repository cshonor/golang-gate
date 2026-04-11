# 08 - context 在中间件中的实战（层层包装 ctx）

## 1. 一句话结论（背）

> 中间件 = 一层层包装 ctx（加超时、加值、响应取消），再把 ctx 传到业务逻辑。

## 2. 典型中间件能力 → 对应 ctx 技术点

- 超时控制 → `context.WithTimeout`
- 主动取消 → `context.WithCancel` + `cancel()`
- TraceID / UserID 透传 → `context.WithValue`
- 停止 goroutine → 在循环里 `select { case <-ctx.Done(): return }`

## 3. 伪代码：中间件链路长什么样

```text
request
  -> middleware A: ctx = WithTimeout(ctx)
  -> middleware B: ctx = WithValue(ctx, traceID)
  -> middleware C: ctx = WithValue(ctx, userID)
  -> handler(ctx)
```

## 4. goroutine 泄漏的“中间件级”根因

如果你开 goroutine 做后台工作，但不监听 `ctx.Done()`，请求结束后 goroutine 仍在跑：

- 资源泄漏（连接/内存/CPU）
- 指标异常（goroutine 数持续上涨）

最小修复：把 `ctx.Done()` 放进 select 里。

## 5. 工程落地清单

- 所有跨网络/跨进程调用必须接收 ctx
- 所有后台 goroutine 必须监听 ctx.Done
- 超时不要到处乱写，统一由入口或上层中间件控制

