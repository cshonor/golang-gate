# 06 - sync/atomic 原子操作原理与实战

本篇偏“怎么用”；实现与内存屏障详见 `08-atomic and lock/`。

---

## 1. 什么时候用 atomic

- 极短临界区、简单数值/指针状态（计数、标志位）。
- 不要用 atomic 去拼“复杂事务逻辑”，可读性与正确性风险高。

---

## 2. 常见 API

- `atomic.Add*`：加减计数
- `atomic.Load*` / `atomic.Store*`：读写
- `atomic.CompareAndSwap*`：CAS

---

## 3. 常见坑

- atomic 只保证单变量原子性，不保证多变量一致性。

---

## 延伸阅读

- `08-atomic and lock/04-atomic 包正确用法：Add 与 CAS 区分.md`
