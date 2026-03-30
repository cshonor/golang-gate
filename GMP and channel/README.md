# 并发：GMP、Channel（与课表对应笔记）

本目录笔记对应 **「后端进阶」** 中与 **协程 / GMP / channel** 相关的条目。根目录下 **14 篇** 笔记文件名带 **`01-`～`14-`** 前缀；**第 15 项**为外链至 **`atomic and lock`** 的 `WaitGroup` 专题。

**map / slice** 见 **`../datastruct/`**；**WaitGroup** 见 **`../atomic and lock/14-WaitGroup特性及原理.md`**；**GC** 见 **`../GC and memory/`**。

## 课表 ↔ 文件索引

| 序号 | 课表标题 | 笔记文件 |
|:----:|----------|----------|
| 01 | 协程和线程的关系 | [01-协程和线程的关系.md](./01-协程和线程的关系.md) |
| 02 | GMP 关系（调度模型鸟瞰） | [02-GMP关系.md](./02-GMP关系.md) |
| 03 | G、M、P、sudog 四个结构体的关系 | [03-G、M、P、sudog 四个结构体的关系.md](./03-G、M、P、sudog%20四个结构体的关系.md) |
| 04 | Goroutine 的数据结构（runtime.g） | [04-Goroutine 的数据结构（runtime.g）.md](./04-Goroutine%20的%20数据结构（runtime.g）.md) |
| 05 | Go GMP 里 P 的核心数据结构（关键属性） | [05-Go GMP 里 P 的核心数据结构 关键属性.md](./05-Go%20GMP%20里%20P%20的核心数据结构%20关键属性.md) |
| 06 | GMP 里一共就 3 个队列 | [06-GMP 里一共就 3 个队列.md](./06-GMP%20里一共就%203%20个队列.md) |
| 07 | G 阻塞后队列归属与 M、P 行为 | [07-G 阻塞后队列归属与 M、P 行为.md](./07-G%20阻塞后队列归属与%20M、P%20行为.md) |
| 08 | sudog 详细介绍 | [08-sudog详细介绍.md](./08-sudog详细介绍.md) |
| 09 | channel 阻塞协程现象 | [09-channel阻塞协程现象.md](./09-channel阻塞协程现象.md) |
| 10 | channel 阻塞原理 | [10-channel阻塞原理.md](./10-channel阻塞原理.md) |
| 11 | closed channel 相关特性 | [11-closed_channel相关特性.md](./11-closed_channel相关特性.md) |
| 12 | 学 GMP 必学 chan | [12-学 GMP 必学 chan.md](./12-学%20GMP%20必学%20chan.md) |
| 13 | channel 读取优化 | [13-channel读取优化.md](./13-channel读取优化.md) |
| 14 | 环形队列 | [14-环形队列.md](./14-环形队列.md) |
| 15 | 并发 WaitGroup 特性及原理 | [14-WaitGroup特性及原理.md](../atomic%20and%20lock/14-WaitGroup特性及原理.md)（`atomic and lock` 目录） |

## 合并版总结（可选）

更少文件、一条主线快速复习，见 `summary/`（原始分篇已全部保留并带 **01–14** 编号）：

- [summary/GMP总览_GMP关系与sudog.md](./summary/GMP总览_GMP关系与sudog.md)
- [summary/G与P的关键数据结构.md](./summary/G与P的关键数据结构.md)
- [summary/GMP队列与阻塞流转.md](./summary/GMP队列与阻塞流转.md)
- [summary/channel阻塞_现象到原理.md](./summary/channel阻塞_现象到原理.md)

## 学习顺序建议（可按课表顺序）

与上表 **序号 01–15** 一致；同一步里多篇可连读。

- **01** [01-协程和线程的关系.md](./01-协程和线程的关系.md)（心智模型）  
- **02–03** [02-GMP关系.md](./02-GMP关系.md) → [03-G、M、P、sudog 四个结构体的关系.md](./03-G、M、P、sudog%20四个结构体的关系.md)  
- **04–05** [04-Goroutine 的数据结构（runtime.g）.md](./04-Goroutine%20的%20数据结构（runtime.g）.md) → [05-Go GMP 里 P 的核心数据结构 关键属性.md](./05-Go%20GMP%20里%20P%20的核心数据结构%20关键属性.md)  
- **06–08** [06-GMP 里一共就 3 个队列.md](./06-GMP%20里一共就%203%20个队列.md) → [07-G 阻塞后队列归属与 M、P 行为.md](./07-G%20阻塞后队列归属与%20M、P%20行为.md) → [08-sudog详细介绍.md](./08-sudog详细介绍.md)  
- **09–11** [09-channel阻塞协程现象.md](./09-channel阻塞协程现象.md) → [10-channel阻塞原理.md](./10-channel阻塞原理.md) → [11-closed_channel相关特性.md](./11-closed_channel相关特性.md)  
- **12** [12-学 GMP 必学 chan.md](./12-学%20GMP%20必学%20chan.md)（把调度模型映射到 channel 路径）  
- **13** [13-channel读取优化.md](./13-channel读取优化.md)（实现/性能视角）  
- **14** [14-环形队列.md](./14-环形队列.md)（与有缓冲 channel / FIFO 对照）  
- **15** [14-WaitGroup特性及原理.md](../atomic%20and%20lock/14-WaitGroup特性及原理.md)（同步收尾）

视频里手写 **「目标：数据结构 / FIFO」**：FIFO 是有缓冲 channel 环形队列的常见实现特征，可与 [10-channel阻塞原理.md](./10-channel阻塞原理.md) 对照。
