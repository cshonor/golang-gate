# GC 与内存管理（面试向笔记）

本目录整理 **堆栈与分配器、逃逸分析、三色并发标记、写屏障、STW、调参与排查** 等主线，与 runtime 实现对照阅读效果更佳。

**GMP / Channel** 见 **`../GMP and channel/`**；**map / slice** 见 **`../datastruct/`**；**锁 / WaitGroup** 见 **`../atomic and lock/`**。

## 主题 ↔ 文件索引

| 主题 | 笔记文件 |
|------|----------|
| 内存分配：堆 vs 栈，`mcache` / `mcentral` / `mheap` | [内存分配堆与栈.md](./内存分配堆与栈.md) |
| 逃逸分析与 `-gcflags=-m` | [逃逸分析.md](./逃逸分析.md) |
| 三色标记、并发漏标、GC 版本演进 | [三色标记清除与GC演进.md](./三色标记清除与GC演进.md) |
| 插入/删除/混合写屏障 | [写屏障.md](./写屏障.md) |
| STW 与异步抢占等 | [STW优化.md](./STW优化.md) |
| 触发条件、`GOGC`、`GOMEMLIMIT`、pprof | [GC触发与性能优化.md](./GC触发与性能优化.md) |
| 面试简答 | [GC面试高频真题.md](./GC面试高频真题.md) |

## 学习顺序建议

1. **内存分配堆与栈** → **逃逸分析**（先建立「为何堆分配才要 GC」）  
2. **三色标记清除与GC演进** → **写屏障**（并发标记正确性）  
3. **STW优化** → **GC触发与性能优化**（观测与调参）  
4. 最后过一遍 **GC面试高频真题**  

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
