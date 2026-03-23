# 并发：GMP、Channel（与课表对应笔记）

本目录笔记对应 **「后端进阶」** 中与 **协程 / GMP / channel** 相关的条目。  
**map / slice** 见同级目录 **`数据结构/`**；**WaitGroup** 见 **`锁实现原理/WaitGroup特性及原理.md`**。

## 课表 ↔ 文件索引

| 课表标题 | 笔记文件 |
|----------|----------|
| channel 阻塞协程现象 | [channel阻塞协程现象.md](./channel阻塞协程现象.md) |
| 协程和线程的关系 | [协程和线程的关系.md](./协程和线程的关系.md) |
| GMP 关系 | [GMP关系.md](./GMP关系.md) |
| channel 阻塞原理 | [channel阻塞原理.md](./channel阻塞原理.md) |
| channel 读取优化 | [channel读取优化.md](./channel读取优化.md) |
| closed channel 相关特性 | [closed_channel相关特性.md](./closed_channel相关特性.md) |
| 并发 WaitGroup 特性及原理 | → `../锁实现原理/WaitGroup特性及原理.md` |

## 学习顺序建议（可按课表顺序）

1. 协程和线程的关系（心智模型）  
2. GMP 关系（谁在调度）  
3. channel 阻塞协程现象 → channel 阻塞原理 → closed channel  
4. channel 读取优化（实现/性能视角）  
5. WaitGroup（同步收尾）

视频里手写 **「目标：数据结构 / FIFO」**：FIFO 是有缓冲 channel 环形队列的常见实现特征，可与 `channel阻塞原理.md` 对照。
