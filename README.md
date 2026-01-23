# Books

Go 语言入门学习项目。

## 项目描述

这是 **Go 语言入门 32 讲** 课程的配套练习项目。

### 配套教材

- **英文原版**：《Get Programming with Go》
- **中文版**：《Go 语言趣学指南》

本项目用于练习和实践课程中的 Go 语言知识点。

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

