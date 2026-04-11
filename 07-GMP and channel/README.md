# 并发：GMP、Channel（与课表对应笔记）

本目录 **`01-`～`15-`** 为 **GMP + Channel** 主线笔记（已按「认知 → 对象 → 队列/阻塞 → Channel 深度 → 优化」**物理重命名**）；**第 16 项** 为外链 **`WaitGroup`**（`08-atomic and lock`）。

**大图导航**：[summary/学习主线与目录映射.md](./summary/学习主线与目录映射.md)（**推荐先读**）。

**map / slice** 见 **`../01-datastruct/`**；**GC** 见 **`../11-GC and memory/`**。

---

## 推荐学习主线（四阶段）

### 阶段一：认知与「对象长什么样」

| 顺序 | 读什么 |
|:----:|--------|
| 1 | [01-协程与线程的本质区别](./01-协程与线程的本质区别.md) |
| 2 | [02-GMP模型核心概念总览](./02-GMP模型核心概念总览.md) |
| 3 | [03-Goroutine核心数据结构runtime.g详解](./03-Goroutine核心数据结构runtime.g详解.md) |
| 4 | [04-Processor核心数据结构runtime.p详解](./04-Processor核心数据结构runtime.p详解.md) |
| 5 | [05-Machine核心数据结构runtime.m详解](./05-Machine核心数据结构runtime.m详解.md) |
| 6（回扣） | [09-GMP与sudog四者联动关系](./09-GMP与sudog四者联动关系.md) |

### 阶段二：队列、阻塞、sudog

| 顺序 | 读什么 |
|:----:|--------|
| 7 | [06-GMP中的核心队列详解](./06-GMP中的核心队列详解.md) |
| 8 | [07-Goroutine阻塞后的调度行为](./07-Goroutine阻塞后的调度行为.md) |
| 9 | [08-sudog核心数据结构与作用详解](./08-sudog核心数据结构与作用详解.md) |

### 阶段三：Channel（结构 → 环形 → 阻塞合并篇 → 关闭 → 总览 → 优化）

| 顺序 | 读什么 |
|:----:|--------|
| 10 | [11-Channel核心数据结构hchan详解](./11-Channel核心数据结构hchan详解.md) |
| 11 | [12-Channel环形队列实现原理](./12-Channel环形队列实现原理.md) |
| 12 | [10-Channel阻塞协程的原理与现象](./10-Channel阻塞协程的原理与现象.md)（**原 09+10 合并**；提纲版见 [summary/channel阻塞_现象到原理](./summary/channel阻塞_现象到原理.md)） |
| 13 | [13-closed_channel核心特性与坑](./13-closed_channel核心特性与坑.md) |
| 14 | [14-学GMP必学Channel总览](./14-学GMP必学Channel总览.md) |
| 15 | [15-Channel性能优化详解](./15-Channel性能优化详解.md) |

若习惯「先现象再结构」，把 **12** 挪到 **10** 之前即可。

### 阶段四：总结与外延

| 内容 | 链接 |
|------|------|
| 合并速记 | [summary/](./summary/) |
| WaitGroup | [14-WaitGroup特性及原理](../08-atomic%20and%20lock/14-WaitGroup特性及原理.md) |

---

## 课表 ↔ 文件索引（01～16）

| 序号 | 课表标题 | 笔记文件 |
|:----:|----------|----------|
| 01 | 协程与线程的本质区别 | [01-协程与线程的本质区别.md](./01-协程与线程的本质区别.md) |
| 02 | GMP 模型核心概念总览 | [02-GMP模型核心概念总览.md](./02-GMP模型核心概念总览.md) |
| 03 | Goroutine 核心数据结构（`runtime.g`） | [03-Goroutine核心数据结构runtime.g详解.md](./03-Goroutine核心数据结构runtime.g详解.md) |
| 04 | Processor 核心数据结构（`runtime.p`） | [04-Processor核心数据结构runtime.p详解.md](./04-Processor核心数据结构runtime.p详解.md) |
| 05 | Machine 核心数据结构（`runtime.m`）导读 | [05-Machine核心数据结构runtime.m详解.md](./05-Machine核心数据结构runtime.m详解.md) |
| 06 | GMP 中的核心队列 | [06-GMP中的核心队列详解.md](./06-GMP中的核心队列详解.md) |
| 07 | Goroutine 阻塞后的调度行为 | [07-Goroutine阻塞后的调度行为.md](./07-Goroutine阻塞后的调度行为.md) |
| 08 | sudog 核心数据结构与作用 | [08-sudog核心数据结构与作用详解.md](./08-sudog核心数据结构与作用详解.md) |
| 09 | GMP 与 sudog 四者联动 | [09-GMP与sudog四者联动关系.md](./09-GMP与sudog四者联动关系.md) |
| 10 | Channel 阻塞：原理与现象（合并） | [10-Channel阻塞协程的原理与现象.md](./10-Channel阻塞协程的原理与现象.md) |
| 11 | Channel 核心数据结构（`hchan`） | [11-Channel核心数据结构hchan详解.md](./11-Channel核心数据结构hchan详解.md) |
| 12 | Channel 环形队列实现 | [12-Channel环形队列实现原理.md](./12-Channel环形队列实现原理.md) |
| 13 | closed channel 特性与坑 | [13-closed_channel核心特性与坑.md](./13-closed_channel核心特性与坑.md) |
| 14 | 学 GMP 必学 Channel 总览 | [14-学GMP必学Channel总览.md](./14-学GMP必学Channel总览.md) |
| 15 | Channel 性能优化 | [15-Channel性能优化详解.md](./15-Channel性能优化详解.md) |
| 16 | 并发 WaitGroup 特性及原理 | [14-WaitGroup特性及原理.md](../08-atomic%20and%20lock/14-WaitGroup特性及原理.md)（`08-atomic and lock`） |

---

## 合并版总结（可选）

- [学习主线与目录映射.md](./summary/学习主线与目录映射.md)  
- [GMP总览_GMP关系与sudog.md](./summary/GMP总览_GMP关系与sudog.md)  
- [G与P的关键数据结构.md](./summary/G与P的关键数据结构.md)  
- [GMP队列与阻塞流转.md](./summary/GMP队列与阻塞流转.md)  
- [channel阻塞_现象到原理.md](./summary/channel阻塞_现象到原理.md)  

---

## 按课表编号 01–16 顺序（与「推荐主线」二选一）

- **01** [01-协程与线程的本质区别.md](./01-协程与线程的本质区别.md)  
- **02** [02-GMP模型核心概念总览.md](./02-GMP模型核心概念总览.md)  
- **03** [03-Goroutine核心数据结构runtime.g详解.md](./03-Goroutine核心数据结构runtime.g详解.md)  
- **04** [04-Processor核心数据结构runtime.p详解.md](./04-Processor核心数据结构runtime.p详解.md)  
- **05** [05-Machine核心数据结构runtime.m详解.md](./05-Machine核心数据结构runtime.m详解.md)  
- **06** [06-GMP中的核心队列详解.md](./06-GMP中的核心队列详解.md)  
- **07** [07-Goroutine阻塞后的调度行为.md](./07-Goroutine阻塞后的调度行为.md)  
- **08** [08-sudog核心数据结构与作用详解.md](./08-sudog核心数据结构与作用详解.md)  
- **09** [09-GMP与sudog四者联动关系.md](./09-GMP与sudog四者联动关系.md)  
- **10** [10-Channel阻塞协程的原理与现象.md](./10-Channel阻塞协程的原理与现象.md)  
- **11** [11-Channel核心数据结构hchan详解.md](./11-Channel核心数据结构hchan详解.md)  
- **12** [12-Channel环形队列实现原理.md](./12-Channel环形队列实现原理.md)  
- **13** [13-closed_channel核心特性与坑.md](./13-closed_channel核心特性与坑.md)  
- **14** [14-学GMP必学Channel总览.md](./14-学GMP必学Channel总览.md)  
- **15** [15-Channel性能优化详解.md](./15-Channel性能优化详解.md)  
- **16** [14-WaitGroup特性及原理.md](../08-atomic%20and%20lock/14-WaitGroup特性及原理.md)  

FIFO、`hchan`、阻塞合并篇对照：[10-Channel阻塞协程的原理与现象.md](./10-Channel阻塞协程的原理与现象.md)、[11-Channel核心数据结构hchan详解.md](./11-Channel核心数据结构hchan详解.md)。
