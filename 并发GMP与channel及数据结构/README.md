# 并发：GMP、Channel 与数据结构

本目录用于整理 **Go 运行时并发模型（GMP）**、**channel 语义与实现要点**，以及 **并发场景下常用数据结构**。

## 建议收录主题

### GMP

- **G（Goroutine）**、**M（OS 线程）**、**P（逻辑处理器）** 的职责与协作
- **调度**：work stealing、抢占（协作式与异步抢占）、`sysmon` 的作用（概念层）
- **与系统线程的关系**：`GOMAXPROCS`、阻塞 syscall 时的行为（概念层）

### Channel

- **语义**：发送/接收阻塞条件、`close` 后的行为、`select` 多路复用
- **实现要点**：环形缓冲、锁与等待队列（waitq）、有缓冲 vs 无缓冲（概念对照源码时更易读）

### 并发数据结构

- **无锁/少锁**：`sync/atomic`、无锁队列思路（概念）
- **带锁容器**：`sync.Map` 适用场景、分片锁（sharded lock）思路
- **经典模式**：worker pool、fan-out / fan-in、pipeline

## 学习顺序建议

1. goroutine + channel 用法（与 `chap30` 等章节对照）
2. `select`、超时、取消（context）
3. GMP 鸟瞰（建立“谁在调度、为何卡住”的直觉）
4. channel 与锁在状态共享上的取舍（与 `chap31` 对照）

可将笔记、流程图和示例代码放在本目录下的子文件中（如 `gmp_notes.md`、`channel_notes.md`、`patterns.go`）。
