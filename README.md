# Books

Go 语言入门学习项目。

## 项目描述

这是 **Go 语言入门 32 讲** 课程的配套练习项目。

### 配套教材

- **英文原版**：《Get Programming with Go》
- **中文版**：《Go 语言趣学指南》

本项目用于练习和实践课程中的 Go 语言知识点。

## 学习目标

通过本课程和项目，您将学习到：

- Go 语言的基础语法和特性
- 变量、函数、结构体等核心概念
- 并发编程（goroutine 和 channel）
- 包管理和模块系统
- 错误处理最佳实践
- 接口和面向对象编程
- 标准库的使用

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

## 快速开始

### 1. 安装 Go 语言

如果还没有安装 Go，请访问 [Go 官网](https://golang.org/dl/) 下载并安装。

### 2. 验证安装

安装完成后，在终端运行以下命令验证：

```bash
go version
```

如果显示版本信息（如 `go version go1.24.1 windows/amd64`），说明安装成功。

### 3. 获取项目

克隆或下载此项目到本地：

```bash
git clone <repository-url>
cd goland
```

或者直接下载项目文件。

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

## 学习建议

1. **按照课程顺序学习**：建议按照《Go 语言趣学指南》的章节顺序进行学习
2. **动手实践**：每学完一个知识点，尝试编写代码并运行
3. **多做练习**：完成教材中的练习题和示例代码
4. **查阅文档**：遇到问题时，参考 [Go 官方文档](https://golang.org/doc/)
5. **社区交流**：加入 Go 语言学习社区，与其他学习者交流

## 参考资料

- **课程教材**：
  - 《Get Programming with Go》（英文原版）
  - 《Go 语言趣学指南》（中文版）
- **官方资源**：
  - [Go 官方网站](https://golang.org/)
  - [Go 官方文档](https://golang.org/doc/)
  - [Go 语言之旅](https://go.dev/tour/)
- **在线资源**：
  - [Go by Example](https://gobyexample.com/)
  - [Go 语言中文网](https://studygolang.com/)

## 许可证

MIT

