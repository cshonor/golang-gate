# 03 - sync.Cond：条件变量与生产者消费者

## 1. Cond 是什么

`sync.Cond` 用于“**条件不满足就等待**，条件满足就唤醒”的协作。

它通常和某把锁绑定：

- 等待时：释放锁并睡眠
- 被唤醒后：重新抢锁再检查条件

## 2. 为什么要“循环检查条件”

Cond 经典写法是 `for !condition { cond.Wait() }`，而不是 `if`：

- 可能有“虚假唤醒”（被唤醒但条件仍不满足）
- 多个等待者竞争，条件被别的 goroutine 消费掉

## 3. 生产者消费者最小模板

```go
type Queue struct {
    mu    sync.Mutex
    cond  *sync.Cond
    items []int
}

func NewQueue() *Queue {
    q := &Queue{}
    q.cond = sync.NewCond(&q.mu)
    return q
}

func (q *Queue) Push(x int) {
    q.mu.Lock()
    q.items = append(q.items, x)
    q.mu.Unlock()
    q.cond.Signal() // 唤醒一个等待者
}

func (q *Queue) Pop() int {
    q.mu.Lock()
    for len(q.items) == 0 {
        q.cond.Wait()
    }
    x := q.items[0]
    q.items = q.items[1:]
    q.mu.Unlock()
    return x
}
```

## 4. Cond vs channel

工程上，很多场景 **channel 更简单**（自带阻塞与唤醒语义）。

Cond 常见于：

- runtime/标准库内部实现
- 需要“广播唤醒”、或需要和复杂共享状态强绑定

