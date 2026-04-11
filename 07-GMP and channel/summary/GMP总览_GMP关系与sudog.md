# GMP 总览：G / M / P + sudog（面试背诵版）

> 由 `02-GMP模型核心概念总览.md` 与 `09-GMP与sudog四者联动关系.md`（上级目录）合并整理。

## 1. 三个字母（先背这一段）

| 符号 | 含义 | 直觉 |
|------|------|------|
| **G** | Goroutine | 要执行的函数闭包 + 栈 + 状态 |
| **M** | Machine（OS 线程） | 真正在 CPU 上跑代码的载体 |
| **P** | Processor（逻辑处理器） | 本地可运行 G 队列 + 执行 Go 代码所需上下文；数量 ≈ `GOMAXPROCS` |

- **P**：更像「并行度闸门」+「调度上下文」
- **M**：工作线程（可以多、可以阻塞、可以替换）
- **G**：任务（数量可远大于 M/P）

层级直觉：**G（最多） >> M（可多于 P） > P（≈ CPU 并行度）**

## 2. 协作关系（最重要的调度句子）

- **M 必须绑定 P 才能执行 Go 代码**（常见路径）。  
- **P 从本地队列 + 全局队列 + 偷取其他 P** 取 G 给 M 执行（work stealing）。  
- **G 阻塞在 channel/锁/网络等** 时，G 会进入对应等待队列；M/P 的行为取决于阻塞类型（见 `GMP队列与阻塞流转.md`）。

把这三句话展开成“调度循环”更好理解：

```text
for {
  g := nextRunnableG(p)      // 本地 runq 优先；否则偷；否则全局
  if g == nil { parkM(); continue }
  execute(g)                 // g 运行一段时间片/直到阻塞/退出
}
```

> 真实实现还涉及 `runnext`、sysmon、netpoll、GC 等；面试先讲清“取 G → 运行 → 阻塞/让出 → 再取”的主线即可。

## 3. 四大结构体：G / M / P / sudog（文字版关系图）

一句话背会：

- **G（g）**：协程本体，存运行现场、栈、状态
- **M（m）**：OS 线程，真正执行代码
- **P（p）**：逻辑处理器，存运行队列，控制并行度
- **sudog**：G 的包装节点，用于阻塞等待队列（尤其是 chan/mutex 等）

核心链路：

- **G 要跑 → 必须绑 M → M 必须绑 P**
- **G 阻塞 → 变成 sudog → 进入等待队列**

## 4. 各自核心字段（只记关键）

### 4.1 G（runtime.g）—— 协程本体

```text
g {
    stack       // 协程栈
    sched       // 调度现场（pc、sp 等，切换用）
    status      // running / waiting / syscall / ...
    m           // 当前绑定的 M
    waitreason  // 等待原因（chan、mutex、sleep）
}
```

补充两个高频点：

- **每个 M 都有一个 g0**：g0 是调度用栈，负责运行 runtime 的调度逻辑；普通 goroutine 在用户栈上跑。
- **抢占与安全点**：Go 1.14+ 的异步抢占让“停不下来的 goroutine”更可控，减少 GC/调度的长尾（细节分散在 `src/runtime/proc.go`、`src/runtime/preempt.go` 等）。

### 4.2 M（runtime.m）—— 系统线程

```text
m {
    g0          // 调度用的栈 G
    curg        // 当前正在运行的 G
    p           // 当前绑定的 P
    nextp       // 即将绑定的 P
}
```

面试补充：

- **M 可以多于 P**：syscall/cgo 可能阻塞 OS 线程；为了不让 P 闲着，runtime 会创建/唤醒更多 M。
- **M 绑定/解绑 P** 是调度器“让 P 保持忙碌”的关键动作之一。

### 4.3 P（runtime.p）—— 逻辑处理器

```text
p {
    runq[]      // 本地 G 运行队列
    runqhead
    runqtail
    m           // 绑定的 M
    gFree       // G 缓存池
}
```

面试补充：

- **P 的本地队列 runq** 是“减少全局锁竞争”的关键设计：大多数情况下取 G 不需要碰全局 `sched`。
- **P 也携带运行 Go 代码所需的一堆上下文**：例如分配器 `mcache`、计时器、GC 相关缓存等（字段随版本演进）。

### 4.4 sudog —— 阻塞等待队列节点

```text
sudog {
    g       // 被阻塞的 G
    c       // 对应的 channel（或其他同步对象语义）
    elem    // 数据地址（收发用）
    next
    prev
}
```

面试补充：

- **sudog 是“把 G 挂到等待队列”的节点**：例如 channel 的 `sendq/recvq` 底层就是 sudog 链表。
- sudog 常携带 **数据地址（elem）**、关联对象（例如 `hchan`），用于“配对拷贝/唤醒”。

## 5. 常见现象对照（面试用）

| 现象 | 本质（GMP 视角） |
|------|-------------------|
| CPU 打满 | 多个 G 在少数 P 上时间片轮转 |
| 网络 IO | netpoller 等让 G 等待而不长期占 M（概念上异步就绪） |
| 阻塞 syscall | 可能占住 M；运行时可能需要更多 M 维持 P 可运行 |

## 6. 与 channel 的关系（串联一句话）

- channel 操作在 runtime 会改变 **G 的状态**（就绪/等待）并把 G（准确说是 sudog）挂到 `hchan` 的等待队列；实现细节见 `channel阻塞_现象到原理.md`。

## 7. 面试高频追问（背答案）

### Q1：为什么需要 P？没有 P 只用 G-M 不行吗？

- 没有 P 时，取 G 大概率依赖全局队列 → **锁竞争**严重  
- P 提供本地队列与调度上下文 → **把竞争分散到 per-P**  
- P 的数量控制并行度（`GOMAXPROCS`）→ **避免过多 OS 线程上下文切换**

### Q2：为什么 M 会多于 P？

- 网络 IO 多数走 netpoll，不一定长期占 M  
- 但 **syscall/cgo 会阻塞 OS 线程**  
- 为了让 P 继续跑，就需要更多可用的 M 来绑定空闲 P

### Q3：G 阻塞时到底发生了什么？

- G：running → waiting  
- 被封装/挂到对应对象的等待队列（常见是 sudog 链表）  
- M 解绑 P，P 继续调度其它 runnable G（详见 `GMP队列与阻塞流转.md`）

## 8. 延伸阅读（建议按需对照）

- `src/runtime/proc.go`：调度主线、P/M/G 绑定与偷取  
- `src/runtime/runtime2.go`：核心结构体定义（随版本变化）  
- `src/runtime/chan.go`：channel 实现（send/recv/close）  
- `src/runtime/netpoll.go`：网络轮询相关  

- `src/runtime/proc.go`、`runtime2.go`、`runtime/iface.go`（按主题拆着读）

