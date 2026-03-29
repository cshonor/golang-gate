# 数据结构（Go：`map` / `slice` 与课表对应笔记）

本目录对应 **「后端进阶」** 中与 **map、slice** 相关的条目。  
**Channel / GMP** 见 **`../GMP and channel/`**；**WaitGroup** 见 **`../atomic and lock/14-WaitGroup特性及原理.md`**。

## 课表 ↔ 文件索引

| 课表标题 | 笔记文件 |
|----------|----------|
| map 数据结构（旧版） | [map数据结构旧版.md](./map数据结构旧版.md) |
| map 渐进式扩容（旧版本） | [map渐进式扩容旧版本.md](./map渐进式扩容旧版本.md) |
| map 并发不安全 | [map并发不安全.md](./map并发不安全.md) |
| map 遍历顺序不固定 | [map遍历顺序不固定.md](./map遍历顺序不固定.md) |
| map 乱序原理 | [map乱序原理.md](./map乱序原理.md) |
| slice 数据结构 | [slice数据结构.md](./slice数据结构.md) |
| slice 扩容 | [slice扩容.md](./slice扩容.md) |

## 学习顺序建议

1. `slice` 结构 → 扩容（内存与共享底层数组）  
2. `map` 结构（桶/bucket）→ 扩容 → 遍历/乱序 → 并发安全

> 注：不同 Go 版本 `map` 实现有演进，笔记中「旧版」指课程/教材对照的常见简化模型，读源码请以当前 `src/runtime/map.go` 为准。
