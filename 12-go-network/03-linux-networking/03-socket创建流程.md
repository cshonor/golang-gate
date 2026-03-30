# socket 创建流程

> **03-linux-networking · Linux 网络底层**

## 内容大纲

- `socket(AF_INET, SOCK_STREAM, 0)` → `bind` → `listen` → `accept` 服务端路径
- 客户端：`socket` → `connect`（可选 `bind` 本地源端口）
- UDP：`sendto`/`recvfrom` 无连接语义
- 错误路径：`EADDRINUSE`、`ECONNREFUSED` 常见根因
- Go 映射：`net.ListenTCP`、`resolver`、`Dialer` 控制项

## 正文

（待补充）
