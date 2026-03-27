# Context（中间件的灵魂）

本目录把 `context.Context` 按“**中间件/后端工程实战**”的方式讲透：取消、超时、数据透传、上下文树、以及在 middleware 里的标准用法。

## 文件索引

| 主题 | 笔记 |
|------|------|
| 01 - context 是什么：作用与核心接口 | [01-context是什么、作用、核心接口.md](./01-context是什么、作用、核心接口.md) |
| 02 - Background/TODO 与上下文树 | [02-Background()、TODO()、上下文树.md](./02-Background()、TODO()、上下文树.md) |
| 03 - WithCancel：手动取消 | [03-WithCancel 手动取消.md](./03-WithCancel 手动取消.md) |
| 04 - WithTimeout：超时控制 | [04-WithTimeout 超时控制.md](./04-WithTimeout 超时控制.md) |
| 05 - WithValue：数据透传 | [05-WithValue 数据透传.md](./05-WithValue 数据透传.md) |
| 06 - 中间件实战：层层包装 ctx | [06-context在中间件中的实战.md](./06-context在中间件中的实战.md) |

## 学习顺序建议

1. 先学 01/02（接口与上下文树）  
2. 再学 03/04（取消与超时，工程里用得最多）  
3. 最后学 05/06（数据透传与中间件写法）

