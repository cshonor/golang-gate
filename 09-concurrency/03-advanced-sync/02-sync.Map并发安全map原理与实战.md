# 02 - sync.Map 并发安全 map 原理与实战

前置：`01-datastruct/02-map/06-map并发不安全.md` 与 `07-map并发安全方案.md`。

---

## 1. 适用场景

- 读多写少、key 集合相对稳定。

---

## 2. 选型

- 语义清晰优先：`map + RWMutex`。
- 真瓶颈再 profile 后考虑 `sync.Map`。

---

## 延伸阅读

- `01-datastruct/02-map/07-map并发安全方案.md`
