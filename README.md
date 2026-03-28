# Books

Go 语言入门学习项目。

## 项目描述

这是 **Go 语言入门 32 讲** 课程的配套练习项目。

### 配套教材

- **英文原版**：《Get Programming with Go》
- **中文版**：《Go 语言趣学指南》

本项目用于练习和实践课程中的 Go 语言知识点。

### 进阶笔记（后端进阶课表对照）

- [锁实现原理](./atomic%20and%20lock/README.md)（含 `WaitGroup`）
- [并发：GMP / Channel](./GMP%20and%20channel/README.md)
- [并发模式与同步原语](./concurrency_patterns/README.md)
- [Context（中间件的灵魂）](./context_pkg/README.md)
- [数据结构：map / slice](./datastruct/README.md)
- [泛型（Go Generics）](./Go%20Generics/1.introduce.md)
- [Go 1.21+ 新特性](./go_new_features/README.md)
- [GC 与内存管理](./GC%20and%20memory/README.md)
- [接口与反射](./interface%20and%20reflection/README.md)
- [错误处理（Error Handling）](./error_handling/README.md)
- [defer 易错点（Defer Traps）](./defer_traps/README.md)
- [unsafe（unsafe 包）](./unsafe_pkg/README.md)
- [Go 网络编程 + IO 模型](./go-network/README.md)

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

