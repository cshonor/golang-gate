# 并发（concurrency）

本目录是「并发工程化」主线：从基础 goroutine / select，到 `sync` 同步原语、`context` 生命周期，再到 channel 工程模式、race 与性能。

与其它目录关系：

- **调度与 channel 心智模型**：`07-GMP and channel/`
- **锁实现原理 / atomic / WaitGroup 深挖**：`08-atomic and lock/`

---

## 目录结构

```
09-concurrency/
├── README.md
├── 01-basics/
├── 02-core-sync-primitives/
├── 03-advanced-sync/
├── 04-context-and-control/
├── 05-channel-and-practice/
├── 06-optimization-and-practice/
└── 07-interview/
```

---

## 学习顺序建议

1. `01-basics/`：先把 goroutine / select 的坑补齐。
2. `02-core-sync-primitives/`：Mutex/RWMutex/WaitGroup/Once/Cond/atomic。
3. `04-context-and-control/`：把退出、超时、取消串起来。
4. `05-channel-and-practice/`：channel 工程模式与排查。
5. `06-optimization-and-practice/`：race + 性能。
6. `07-interview/`：最后冲刺口述与真题。

---

## 文件索引

### 01 - basics

- [01-同步原语总览](./01-basics/01-同步原语总览.md)
- [02-goroutine基础与常见坑](./01-basics/02-goroutine基础与常见坑.md)
- [03-select并发模式与实战](./01-basics/03-select并发模式与实战.md)

### 02 - core sync primitives

- [01-sync.Mutex互斥锁原理与实战](./02-core-sync-primitives/01-sync.Mutex互斥锁原理与实战.md)
- [02-sync.RWMutex读写锁原理与实战](./02-core-sync-primitives/02-sync.RWMutex读写锁原理与实战.md)
- [03-sync.WaitGroup原理与实战](./02-core-sync-primitives/03-sync.WaitGroup原理与实战.md)
- [04-sync.Once单例模式与原理](./02-core-sync-primitives/04-sync.Once单例模式与原理.md)
- [05-sync.Cond条件变量与生产者消费者](./02-core-sync-primitives/05-sync.Cond条件变量与生产者消费者.md)
- [06-sync.atomic原子操作原理与实战](./02-core-sync-primitives/06-sync.atomic原子操作原理与实战.md)

### 03 - advanced sync

- [01-sync.Pool对象复用与性能优化](./03-advanced-sync/01-sync.Pool对象复用与性能优化.md)
- [02-sync.Map并发安全map原理与实战](./03-advanced-sync/02-sync.Map并发安全map原理与实战.md)
- [03-sync.ErrGroup并发错误处理](./03-advanced-sync/03-sync.ErrGroup并发错误处理.md)

### 04 - context & control

- [01-context上下文与超时控制](./04-context-and-control/01-context上下文与超时控制.md)

### 05 - channel & practice

- [01-channel基础与常见坑](./05-channel-and-practice/01-channel基础与常见坑.md)
- [02-channel经典并发模式](./05-channel-and-practice/02-channel经典并发模式.md)

### 06 - optimization & practice

- [01-数据竞争与race检测](./06-optimization-and-practice/01-数据竞争与race检测.md)
- [02-并发性能优化](./06-optimization-and-practice/02-并发性能优化.md)

### 07 - interview

- [01-面试口述版](./07-interview/01-面试口述版.md)
- [02-面试真题](./07-interview/02-面试真题.md)
