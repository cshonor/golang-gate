# 06 - select：并发模式与实战

## 1. select 是什么

`select` 用于同时等待多个 channel 事件，是 Go 并发编排的“控制台”。

## 2. 3 个最常用套路

### 2.1 超时控制

```go
select {
case v := <-ch:
    _ = v
case <-time.After(200 * time.Millisecond):
    return errors.New("timeout")
}
```

生产环境常用 `context.WithTimeout` 替代 `time.After`（避免计时器滥用/更好组合）。

### 2.2 退出信号（优雅停止 goroutine）

```go
for {
    select {
    case <-ctx.Done():
        return
    case job := <-jobs:
        handle(job)
    }
}
```

### 2.3 fan-in（多路合并）

思路：多个输入 channel 汇聚到一个输出 channel，由一个 goroutine 统一 select。

## 3. default 分支的坑

带 `default` 的 select 可能变成“忙等”（空转占 CPU）。

只有在你明确需要“非阻塞尝试”时才用 `default`，否则不要写。

## 4. 面试一句话

> `select` 是并发编排工具：超时、取消、合并多路事件，几乎所有高并发服务都会用到。

