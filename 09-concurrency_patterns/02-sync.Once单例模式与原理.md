# 02 - sync.Once：单例模式与原理

## 1. 解决什么问题

`sync.Once` 保证某段初始化逻辑 **在并发场景只执行一次**。

典型：全局单例、配置加载、昂贵资源初始化（连接池、编译正则等）。

## 2. 标准用法

```go
var (
    once sync.Once
    cli  *Client
)

func GetClient() *Client {
    once.Do(func() {
        cli = NewClient()
    })
    return cli
}
```

## 3. 两个坑

### 坑 1：Do 里的函数 panic

如果 `Do` 的函数里 panic，Once 的状态可能会让后续调用永远认为“已经执行过/不再执行”，导致系统处于半初始化状态。工程上要避免：

- Do 函数尽量不 panic
- 或把可能失败的逻辑改成显式返回 error，并在外部控制重试策略

### 坑 2：Do 里不要递归调用 Do

会造成死锁或逻辑混乱（初始化依赖要设计清晰）。

## 4. 面试一句话

> `sync.Once` 解决并发下“只初始化一次”的需求；用它比自己写 double-check 锁更安全。

