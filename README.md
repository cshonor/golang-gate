# Books

Go 语言入门学习项目。

## 项目描述

这是 **Go 语言入门 32 讲** 课程的配套练习项目。

### 配套教材

- **英文原版**：《Get Programming with Go》
- **中文版**：《Go 语言趣学指南》

本项目用于练习和实践课程中的 Go 语言知识点。

### 进阶笔记（后端进阶课表对照）

目录名带 **`01-`～`13-`** 序号，与下方学习顺序一致。

- [Go 基础数据类型](./00-basic-types/README.md)
- [数据结构：map / slice](./01-datastruct/README.md)
- [接口与反射](./02-interface-and-reflection/README.md)
- [错误处理（Error Handling）](./03-error_handling/README.md)
- [defer 易错点（Defer Traps）](./04-defer_traps/README.md)
- [泛型（Go Generics）](./05-Go%20Generics/README.md)
- [Go 1.21+ 新特性](./06-go_new_features/README.md)
- [并发：GMP / Channel](./07-GMP%20and%20channel/README.md)
- [锁实现原理](./08-atomic%20and%20lock/README.md)（含 `WaitGroup`）
- [并发模式与同步原语](./09-concurrency_patterns/README.md)
- [Context（中间件的灵魂）](./10-context_pkg/README.md)
- [GC 与内存管理](./11-GC%20and%20memory/README.md)
- [Go 网络编程 + IO 模型](./12-go-network/README.md)
- [unsafe（unsafe 包）](./13-unsafe_pkg/README.md)

### 进阶笔记建议学习顺序

按 **依赖关系** 排了一条主线（括号内是可并行或穿插的模块）：

1. [Go 基础数据类型](./00-basic-types/README.md) — 类型系统、零值、转换、字符串与指针，为后续所有模块打地基。  
2. [数据结构：map / slice](./01-datastruct/README.md) — 容器与复杂度，后面读 runtime、网络缓冲都更顺。  
3. [接口与反射](./02-interface-and-reflection/README.md) — `io.Reader`/`error` 等惯用法与类型系统。  
4. [错误处理（Error Handling）](./03-error_handling/README.md) — 日常编码与 API 设计基础。  
5. [defer 易错点（Defer Traps）](./04-defer_traps/README.md) — 与错误处理、资源关闭强相关，宜紧接其后。  
6. [泛型（Go Generics）](./05-Go%20Generics/README.md) — 现代标准库与第三方 API 常见写法（目录内从 `01-Introduce.md` 起按 README 顺序读）。  
7. [Go 1.21+ 新特性](./06-go_new_features/README.md) — 工具链与语言小特性（**可与 1～6 穿插**，不必死磕顺序）。  
8. [并发：GMP / Channel](./07-GMP%20and%20channel/README.md) — 调度与 channel 心智模型，**并发主线入口**。  
9. [锁实现原理](./08-atomic%20and%20lock/README.md) — Mutex、原子、WaitGroup；与 GMP「阻塞与让出」对照读效果更好。  
10. [并发模式与同步原语](./09-concurrency_patterns/README.md) — 在 GMP/锁之上收束常见写法与反模式。  
11. [Context（中间件的灵魂）](./10-context_pkg/README.md) — 超时、取消、请求级数据；**写 HTTP/RPC 前建议完成**。  
12. [GC 与内存管理](./11-GC%20and%20memory/README.md) — 逃逸、分配、调优；有并发与 IO 直觉后读更落地。  
13. [Go 网络编程 + IO 模型](./12-go-network/README.md) — 综合 OS、TCP/UDP、`net`/`netpoll`；**依赖 8～11 的并发与 Context 基础**。  
14. [unsafe（unsafe 包）](./13-unsafe_pkg/README.md) — 放最后：在已熟悉类型系统、内存与并发边界后再碰。

**简记**：语言与类型（1～7）→ 并发栈（8～11）→ 内存与系统（12～13）→ unsafe（14）。

## 系统要求

- Go 1.21 或更高版本

## 检查 Go 版本

在终端（PowerShell、CMD 或 Git Bash）中运行以下命令检查 Go 版本：

```bash
go version
```

输出示例：`go version go1.24.1 windows/amd64`

### 其他有用的 Go 环境命令

```bash
go env          # 查看所有 Go 环境变量
go env GOROOT   # 查看 Go 安装路径
go env GOPATH   # 查看 Go 工作空间路径
go env GOVERSION # 查看 Go 版本（另一种方式）
```

## 安装

1. 确保已安装 Go 语言环境：
   ```bash
   go version
   ```

2. 克隆或下载此项目到本地

## 运行

```bash
go run main.go
```

或者编译后运行：

```bash
go build
./books.exe  # Windows
# 或
./books      # Linux/Mac
```

## 项目结构

```
.
├── main.go      # 主程序文件
├── go.mod       # Go 模块配置文件
└── README.md    # 项目说明文档
```

## 许可证

MIT

