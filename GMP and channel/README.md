# 并发：GMP、Channel（与课表对应笔记）

本目录笔记对应 **「后端进阶」** 中与 **协程 / GMP / channel** 相关的条目。  
**map / slice** 见 **`../datastruct/`**；**WaitGroup** 见 **`../atomic and lock/WaitGroup特性及原理.md`**；**GC** 见 **`../GC and memory/`**。

## 课表 ↔ 文件索引

| 课表标题 | 笔记文件 |
|----------|----------|
| 协程和线程的关系 | [协程和线程的关系.md](./协程和线程的关系.md) |
| GMP 关系（调度模型鸟瞰） | [GMP关系.md](./GMP关系.md) |
| G、M、P、sudog 四个结构体的关系 | [G、M、P、sudog 四个结构体的关系.md](./G、M、P、sudog%20四个结构体的关系.md) |
| Goroutine 的数据结构（runtime.g） | [Goroutine 的数据结构（runtime.g）.md](./Goroutine%20的%20数据结构（runtime.g）.md) |
| Go GMP 里 P 的核心数据结构（关键属性） | [Go GMP 里 P 的核心数据结构 关键属性.md](./Go%20GMP%20里%20P%20的核心数据结构%20关键属性.md) |
| GMP 里一共就 3 个队列 | [GMP 里一共就 3 个队列.md](./GMP%20里一共就%203%20个队列.md) |
| G 阻塞后队列归属与 M、P 行为 | [G 阻塞后队列归属与 M、P 行为.md](./G%20阻塞后队列归属与%20M、P%20行为.md) |
| sudog 详细介绍 | [sudog详细介绍.md](./sudog详细介绍.md) |
| channel 阻塞协程现象 | [channel阻塞协程现象.md](./channel阻塞协程现象.md) |
| channel 阻塞原理 | [channel阻塞原理.md](./channel阻塞原理.md) |
| 学 GMP 必学 chan | [学 GMP 必学 chan.md](./学%20GMP%20必学%20chan.md) |
| channel 读取优化 | [channel读取优化.md](./channel读取优化.md) |
| 环形队列 | [环形队列.md](./环形队列.md) |
| closed channel 相关特性 | [closed_channel相关特性.md](./closed_channel相关特性.md) |
| 并发 WaitGroup 特性及原理 | [WaitGroup特性及原理.md](../atomic%20and%20lock/WaitGroup特性及原理.md) |

## 合并版总结（可选）

如果你想用“更少文件 + 一条主线”快速复习，可以看 `summary/` 下的合并版（原始笔记已全部保留）：  

- [summary/GMP总览_GMP关系与sudog.md](./summary/GMP总览_GMP关系与sudog.md)
- [summary/G与P的关键数据结构.md](./summary/G与P的关键数据结构.md)
- [summary/GMP队列与阻塞流转.md](./summary/GMP队列与阻塞流转.md)
- [summary/channel阻塞_现象到原理.md](./summary/channel阻塞_现象到原理.md)

## 学习顺序建议（可按课表顺序）

1. 协程和线程的关系（心智模型）  
2. [GMP关系.md](./GMP关系.md) → [G、M、P、sudog 四个结构体的关系.md](./G、M、P、sudog%20四个结构体的关系.md)  
3. [Goroutine 的数据结构（runtime.g）.md](./Goroutine%20的%20数据结构（runtime.g）.md) → [Go GMP 里 P 的核心数据结构 关键属性.md](./Go%20GMP%20里%20P%20的核心数据结构%20关键属性.md)  
4. [GMP 里一共就 3 个队列.md](./GMP%20里一共就%203%20个队列.md) → [G 阻塞后队列归属与 M、P 行为.md](./G%20阻塞后队列归属与%20M、P%20行为.md) → [sudog详细介绍.md](./sudog详细介绍.md)  
5. [channel阻塞协程现象.md](./channel阻塞协程现象.md) → [channel阻塞原理.md](./channel阻塞原理.md) → [closed_channel相关特性.md](./closed_channel相关特性.md)  
6. [学 GMP 必学 chan.md](./学%20GMP%20必学%20chan.md)（把调度模型映射到 channel 路径）  
7. [channel读取优化.md](./channel读取优化.md)（实现/性能视角）  
8. [环形队列.md](./环形队列.md)（与有缓冲 channel / FIFO 对照）  
9. [WaitGroup特性及原理.md](../atomic%20and%20lock/WaitGroup特性及原理.md)（同步收尾）

视频里手写 **「目标：数据结构 / FIFO」**：FIFO 是有缓冲 channel 环形队列的常见实现特征，可与 `channel阻塞原理.md` 对照。
