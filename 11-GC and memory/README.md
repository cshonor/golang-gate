# GC 与内存管理（面试向笔记）

本目录整理 **堆栈与分配器（mcache / mcentral / mheap）、逃逸分析、三色并发标记、写屏障、STW、调参与排查** 等主线，与 runtime 实现对照阅读效果更佳。

文件名统一为 **`01-`～`09-` 序号**；**01～03** 偏 **分配器与堆布局**，**04** 起进入 **逃逸与 GC 算法**。

**GMP / Channel** 见 **`../07-GMP and channel/`**；**map / slice** 见 **`../01-datastruct/`**；**锁 / WaitGroup** 见 **`../08-atomic and lock/`**；**接口与反射** 见 **`../02-interface-and-reflection/`**。

## 主题 ↔ 文件索引

| 序号 | 主题 | 笔记 |
|:----:|------|------|
| 01 | 内存分配：堆 vs 栈，分配路径直觉 | [01-内存分配堆与栈.md](./01-内存分配堆与栈.md) |
| 02 | `mcache` / `mcentral` / `mspan` | [02-mcache mcentral mspan.md](./02-mcache%20mcentral%20mspan.md) |
| 03 | `mheap` 全局堆 | [03-mheap.md](./03-mheap.md) |
| 04 | 逃逸分析与 `-gcflags=-m` | [04-逃逸分析.md](./04-逃逸分析.md) |
| 05 | 三色标记、并发漏标、GC 演进 | [05-三色标记清除与GC演进.md](./05-三色标记清除与GC演进.md) |
| 06 | 插入/删除/混合写屏障 | [06-写屏障.md](./06-写屏障.md) |
| 07 | STW 与异步抢占等 | [07-STW优化.md](./07-STW优化.md) |
| 08 | 触发条件、`GOGC`、`GOMEMLIMIT`、pprof | [08-GC触发与性能优化.md](./08-GC触发与性能优化.md) |
| 09 | 面试简答 | [09-GC面试高频真题.md](./09-GC面试高频真题.md) |

## 学习顺序建议

1. **01 → 03**：堆栈与分配器（`mcache` / `mcentral` / `mheap`）  
2. **04**：逃逸分析（建立「堆分配才主要走 GC」）  
3. **05 → 06**：三色标记 → 写屏障（并发标记正确性）  
4. **07 → 08**：STW → 触发与调优（观测与 pprof）  
5. **09**：面试高频真题  

## 进阶阅读（源码）

- `src/runtime/mgc.go`（GC 主流程）  
- `src/runtime/malloc.go`（分配器）  
- `src/runtime/mbarrier.go`（写屏障）

## 速记（背诵用）

- **栈**：per-goroutine，随调用结束回收，堆压力小  
- **堆**：逃逸/共享/动态大小等，走分配器 + GC  
- **逃逸**：`-gcflags="-m -l"` 看 `moved to heap`  
- **三色**：白/灰/黑，并发标记靠 **写屏障** 保正确性  
- **调参**：`GOGC`、`GOMEMLIMIT`；**诊断**：`pprof`、`runtime/trace`  
