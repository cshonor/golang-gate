# Context（中间件的灵魂）

本目录把 `context.Context` 按“**中间件/后端工程实战**”的方式讲透：Goroutine 链与取消方向、接口与上下文树、取消/超时、数据透传、middleware 标准用法。

## 文件索引

| 序号 | 主题 | 笔记 |
|------|------|------|
| 01 | Goroutine 链与 Context 取消（总览） | [01-Goroutine链与Context取消.md](./01-Goroutine链与Context取消.md) |
| 02 | context 是什么：作用与核心接口 | [02-context是什么、作用、核心接口.md](./02-context是什么、作用、核心接口.md) |
| 03 | Background/TODO 与上下文树 | [03-Background()、TODO()、上下文树.md](./03-Background()、TODO()、上下文树.md) |
| 04 | WithCancel：手动取消 | [04-WithCancel 手动取消.md](./04-WithCancel%20手动取消.md) |
| 05 | WithTimeout：超时控制 | [05-WithTimeout 超时控制.md](./05-WithTimeout%20超时控制.md) |
| 06 | WithDeadline：绝对时刻取消 | [06-WithDeadline 精确到时间点的取消控制.md](./06-WithDeadline%20精确到时间点的取消控制.md) |
| 07 | WithValue：数据透传 | [07-WithValue 数据透传.md](./07-WithValue%20数据透传.md) |
| 08 | 中间件实战：层层包装 ctx | [08-context在中间件中的实战.md](./08-context在中间件中的实战.md) |
| 09 | 常见陷阱与反模式 | [09-context常见陷阱与反模式.md](./09-context常见陷阱与反模式.md) |
| 10 | 中间件实战：trace-id 透传 | [10-中间件实战-trace-id透传.md](./10-中间件实战-trace-id透传.md) |

## 学习顺序建议

1. **01**（Goroutine 链 + 取消方向；可与 02 穿插）  
2. **02 / 03**（接口与上下文树）  
3. **04～06**（取消、相对超时、绝对 deadline）  
4. **07～10**（数据透传、中间件、陷阱、trace-id）



