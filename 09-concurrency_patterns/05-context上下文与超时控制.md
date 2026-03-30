# 05 - context：上下文与超时控制

## 1. context 是什么

`context.Context` 用来做 3 件事：

1. **取消**（cancellation）
2. **超时/截止时间**（timeout/deadline）
3. **跨层传递请求级元数据**（request scoped values）

> 你已经有一整套更完整的 context 专题笔记：见 `../10-context_pkg/`。

## 2. 最重要的工程规范（背）

- **入口创建**：HTTP/RPC 入口创建/拿到 ctx
- **逐层传递**：函数签名 `func (x) Do(ctx context.Context, ...)`
- **不要存到 struct 里长期持有**（会泄漏/难管理）

## 3. 超时模板

```go
ctx, cancel := context.WithTimeout(parent, 200*time.Millisecond)
defer cancel()

err := call(ctx)
```

## 4. select 配合取消

```go
select {
case <-ctx.Done():
    return ctx.Err()
case v := <-ch:
    return use(v)
}
```

## 5. ctx.Value 的使用边界

只放“请求级元数据”（trace id、user id 等），不要把大对象/业务参数塞进去。

