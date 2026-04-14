# 数据结构（Go：`array` / `slice` / `map`）

本目录按 **专题子文件夹** 组织：**`00-array/`**、**`01-slice/`**、**`02-map/`**、**`03-common-optimization/`**。  
**Channel / GMP** 见 **`../07-GMP and channel/`**；**WaitGroup** 见 **`../08-atomic and lock/14-WaitGroup特性及原理.md`**。

---

## 目录结构

```
01-datastruct/
├── README.md                 # 本文件：学习路径与索引
├── 00-array/                 # Array 专题
│   ├── 01-array数据结构.md
│   ├── 02-array内存布局.md
│   ├── 03-array与slice的区别.md
│   ├── 04-array常见坑与最佳实践.md
│   ├── 05-array底层汇编分析.md       # 可选
│   └── README.md
├── 01-slice/                 # Slice 专题
│   ├── 01-slice数据结构.md
│   ├── 02-slice扩容.md
│   ├── 03-slice常见坑与最佳实践.md
│   └── 05-slice底层汇编分析.md      # 可选
├── 02-map/                   # Map 专题
│   ├── 01-map-hmap与bmap-直觉对照.md
│   ├── 02-map数据结构旧版.md
│   ├── 03-map渐进式扩容旧版本.md
│   ├── 04-map遍历顺序不固定.md
│   ├── 05-map乱序原理.md
│   ├── 06-map并发不安全.md
│   ├── 07-map并发安全方案.md
│   ├── 08-map内存模型与GC.md
│   └── 09-map新版实现（1.18+）.md   # 可选
└── 03-common-optimization/
    └── 01-数据结构性能优化.md
```

---

## 文件索引（扁平表）

### Array（`00-array/`）

| 编号 | 标题 | 文件 |
|:----:|------|------|
| 01 | array 数据结构 | [00-array/01-array数据结构.md](./00-array/01-array数据结构.md) |
| 02 | array 内存布局 | [00-array/02-array内存布局.md](./00-array/02-array内存布局.md) |
| 03 | array 与 slice 的区别 | [00-array/03-array与slice的区别.md](./00-array/03-array与slice的区别.md) |
| 04 | array 常见坑与最佳实践 | [00-array/04-array常见坑与最佳实践.md](./00-array/04-array常见坑与最佳实践.md) |
| 05 | array 底层汇编分析（可选） | [00-array/05-array底层汇编分析.md](./00-array/05-array底层汇编分析.md) |

### Slice（`01-slice/`）

| 编号 | 标题 | 文件 |
|:----:|------|------|
| 01 | slice 数据结构 | [01-slice/01-slice数据结构.md](./01-slice/01-slice数据结构.md) |
| 02 | slice 扩容 | [01-slice/02-slice扩容.md](./01-slice/02-slice扩容.md) |
| 03 | slice 常见坑与最佳实践 | [01-slice/03-slice常见坑与最佳实践.md](./01-slice/03-slice常见坑与最佳实践.md) |
| 04 | slice 底层汇编分析（可选） | [01-slice/05-slice底层汇编分析.md](./01-slice/05-slice底层汇编分析.md) |

### Map（`02-map/`）

| 编号 | 标题 | 文件 |
|:----:|------|------|
| 01 | hmap/bmap 直觉对照 | [02-map/01-map-hmap与bmap-直觉对照.md](./02-map/01-map-hmap与bmap-直觉对照.md) |
| 02 | map 数据结构（旧版） | [02-map/02-map数据结构旧版.md](./02-map/02-map数据结构旧版.md) |
| 03 | map 渐进式扩容（旧版本） | [02-map/03-map渐进式扩容旧版本.md](./02-map/03-map渐进式扩容旧版本.md) |
| 04 | map 遍历顺序不固定 | [02-map/04-map遍历顺序不固定.md](./02-map/04-map遍历顺序不固定.md) |
| 05 | map 乱序原理 | [02-map/05-map乱序原理.md](./02-map/05-map乱序原理.md) |
| 06 | map 并发不安全 | [02-map/06-map并发不安全.md](./02-map/06-map并发不安全.md) |
| 07 | map 并发安全方案 | [02-map/07-map并发安全方案.md](./02-map/07-map并发安全方案.md) |
| 08 | map 内存模型与 GC | [02-map/08-map内存模型与GC.md](./02-map/08-map内存模型与GC.md) |
| 09 | map 新版实现（1.18+，可选） | [02-map/09-map新版实现（1.18+）.md](./02-map/09-map新版实现（1.18+）.md) |

### 通用优化（`03-common-optimization/`）

| 编号 | 标题 | 文件 |
|:----:|------|------|
| 01 | 数据结构性能优化 | [03-common-optimization/01-数据结构性能优化.md](./03-common-optimization/01-数据结构性能优化.md) |

---

## 学习顺序建议

1. **Array → Slice → Map**：线性结构先打地基（数组）再看切片与扩容，然后进入哈希表。  
2. **Array**：`00-array` 内按 **01 → 04** 为主，**05** 可选。  
3. **Slice**：`01-slice` 内按 **01 → 03** 为主，**05** 可选；读完 **02** 后可看 **03 坑**。  
4. **Map**：`02-map` 内按 **01 → 09** 顺序；**01 直觉对照**可先于 **02** 速览。  
5. **并发**：map 并发按 **06 → 07**（问题 → 方案）。  
6. **优化与 GC**：需要时再读 map 的 **08**、以及 **`03-common-optimization/01`**。  
7. **版本细节**：map 的 **09** 强调对照当前 `runtime/map.go`，勿背死实现细节。

> 注：笔记中「旧版」指课程/教材常用简化模型；**抠实现与面试细节以当前 `src/runtime/map.go` 为准**。

---

## 面试与工程侧重点（摘要）

- **必会**：slice 头与共享数组、`append` 与扩容；map 的 `hmap`/桶/并发规则、`for range` 无序。  
- **高频坑**：子切片泄漏、并发写 map、`errors`/`map` 误用（错误处理见 `../03-error_handling/`）。  
- **生产**：日志与可观测、锁与 `sync.Map` 选型、预分配与 pprof。
