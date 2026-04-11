# G 与 P 的关键数据结构（runtime.g / runtime.p）

> 由 `03-Goroutine核心数据结构runtime.g详解.md` 与 `04-Processor核心数据结构runtime.p详解.md`（上级目录）合并整理。

## 1. G（runtime.g）：协程的“身份证 + 运行档案”

Goroutine 在 runtime 底层对应结构体 **`g`**（小写），我们常叫 **G**。

### 关键字段（只记这些就够用）

1. **stack**：协程自己的栈（初始很小，动态扩容）
2. **pc / sp（在 sched 里）**：切换现场核心
3. **atomicstatus**：协程状态
   - `_Gidle` / `_Grunnable` / `_Grunning` / `_Gwaiting` / `_Gsyscall` / `_Gdead`
4. **m**：当前绑定的线程 M（正在谁上面跑）
5. **sched**：调度上下文（保存被暂停时现场：pc、sp、寄存器等）
6. **waitreason**：阻塞原因（chan、锁、sleep、GC 等）
7. **preempt**：抢占标记（运行太久会被调度器切走）

### 再补 4 个“面试常追问但容易漏”的点

1. **G 的栈是“可增长的连续栈”**  
   - 初始很小（常见 2KB 起步），不够会扩容并搬迁  
   - 这是 Go 能“起很多 goroutine 但不爆内存”的基础之一

2. **sched 不是“调度器”，而是“现场保存区”**  
   - 保存切换点的 `pc/sp` 等寄存器信息  
   - Goroutine 切走时保存，切回来恢复

3. **g0 / m0 / 主 goroutine**（只要能说清关系）  
   - 每个 M 有自己的 g0（runtime 调度用栈）  
   - `main` goroutine 也是普通 G，只是由 runtime 启动并执行 `main.main`

4. **G 的状态流转是理解阻塞/唤醒的关键**  
   - `_Grunnable`：可运行（在 runq/sched.runq）  
   - `_Grunning`：正在某个 M 上跑  
   - `_Gwaiting`：阻塞（挂到等待队列）  
   - `_Gsyscall`：在 syscall/cgo（可能导致 M 被内核阻塞）

一句话总结：

**`g` 保存了协程的栈、运行状态、切换现场、绑定关系与阻塞原因，是协程切换与阻塞唤醒的核心载体。**

## 2. P（runtime.p）：调度单元 + 执行上下文

P = Processor（逻辑处理器）：

- 承载可运行 G
- 是 M 执行 Go 代码必须绑定的“执行许可”
- 全局 P 数量由 `GOMAXPROCS` 控制（≈ 并行度）

### P 的核心数据结构（抓住最关键字段）

```go
type p struct {
    status      uint32
    id          int32
    m           muintptr

    // 本地可运行 G 队列（非常重要）
    runqhead    uint32
    runqtail    uint32
    runq        [256]guintptr

    preempt     bool
    // ... 还有 mcache、栈缓存、定时器等
}
```

### 必背：P 最重要的属性

1. **runq 本地队列（长度 256）**
   - 数组 + `runqhead/runqtail` 形成环形队列
   - 目标：尽量本地取 G，减少全局锁竞争
2. **m（绑定的 M）**
   - 同一时间：通常 **1P ↔ 1M** 绑定执行
3. **status**
   - `_Pidle` / `_Prunning` / `_Pgcstop` / `_Pdead`
4. **preempt**
   - 抢占相关标记（Go 1.14+ 异步抢占等与之相关）
5. **mcache**
   - 每个 P 一份小对象分配缓存（无锁快分配，和 GC/分配器主线相关）

### P 上还常出现的“面试加分点”（知道名字就够）

- **runnext**：下一次优先运行的 G（用于减少延迟/提高局部性，具体策略以版本为准）
- **timers**：定时器相关结构（`time.Sleep`、`time.After` 等会走到这里或相关路径）
- **gc 相关缓存/统计**：协助 GC 与写屏障的运行（细节随版本变化）

> 你不需要背全字段；面试要表达的是：**P 不只是 runq，它是“运行 Go 代码的上下文容器”。**

## 3. 结构体与文件对照（源码导航）

- `src/runtime/runtime2.go`：`type g` / `type m` / `type p`（可能拆分到其它文件，随版本变化）
- `src/runtime/proc.go`：调度、runq、work stealing、绑定/解绑
- `src/runtime/stack.go`：栈增长/搬迁（理解“连续栈”）


一句话总结：

**P 是 Go 调度的核心单元：本地 runq 无锁/低锁 + 绑定执行上下文，决定了并行度与调度效率。**

