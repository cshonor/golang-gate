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
| 06 | WithValue：数据透传 | [06-WithValue 数据透传.md](./06-WithValue%20数据透传.md) |
| 07 | 中间件实战：层层包装 ctx | [07-context在中间件中的实战.md](./07-context在中间件中的实战.md) |

## 学习顺序建议

1. **01**（Goroutine 链 + 取消方向；可与 02 穿插）  
2. **02 / 03**（接口与上下文树）  
3. **04 / 05**（取消与超时，工程里用得最多）  
4. **06 / 07**（数据透传与中间件写法）
