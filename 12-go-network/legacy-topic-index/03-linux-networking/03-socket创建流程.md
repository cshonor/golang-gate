# socket 创建流程

> **03-linux-networking · Linux 网络底层**  
> 上篇：[02-file-descriptor-fd](./02-file-descriptor-fd.md)

---

## 一、核心背景

一个 **TCP 服务端** 在经典 Berkeley API 下遵循：

**`socket` → `bind` → `listen` → `accept`（循环）**

**TCP 客户端** 常见路径：

**`socket` →（可选 `bind` 本地地址）→ `connect`**

Go 的 **`net.Listen` / `net.Dial`** 是对上述路径的封装（外加解析、双栈、`Dialer` 等逻辑）。

---

## 二、TCP 服务端（POSIX 心智）

### 1. `socket()`

```c
int socket(int domain, int type, int protocol);
```

- **`domain`**：如 **`AF_INET` / `AF_INET6`**。  
- **`type`**：TCP 用 **`SOCK_STREAM`**。  
- **`protocol`**：常传 **`0`** 表示由系统选择默认协议（IPv4 下为 TCP）。  
- **返回**：新 **fd**；此时通常尚未绑定端口、未进入监听。

### 2. `bind()`

将 socket 与 **本地 IP:端口** 关联。服务端必须让客户端能 **定位到端口**。

- 常见：`INADDR_ANY` + 端口 `8080` → 在所有地址上监听 `8080`。  
- 失败典型：**`EADDRINUSE`**（端口已被占用或未 `SO_REUSEADDR` 策略冲突）。

### 3. `listen()`

```c
int listen(int sockfd, int backlog);
```

- 将 socket 置于 **被动监听**，开始处理到达的 **SYN**（三次握手前半段）。  
- **`backlog`**：与 **全连接队列（accept queue）** 长度相关，内核还会结合 **`somaxconn`** 等做上限裁剪（详见 [07](./07-半连接队列与全连接队列.md)）。

### 4. `accept()`

```c
int accept(int sockfd, struct sockaddr *addr, socklen_t *addrlen);
```

- 从 **已完成三次握手** 的全连接队列取一项，返回 **代表该连接的新 fd**。  
- **监听 fd** 与 **连接 fd 分离**：监听 fd 继续 `accept`，连接 fd 负责与该客户端 `read`/`write`。

---

## 三、TCP 客户端

### 1. `socket()`

同服务端，创建流式 socket。

### 2. `connect()`

向对端 **三元组（IP,端口,协议）** 发起连接，驱动 **三次握手**。

- 客户端 **源端口** 常由内核 **临时分配**（`ephemeral port`），也可先 `bind` 固定源端口（少见）。  
- 失败典型：**`ECONNREFUSED`**（对端无监听）、**`ETIMEDOUT`**（路由/防火墙丢弃等）。

---

## 四、Go 中的对应关系（摘要）

### 1. 服务端 `net.Listen("tcp", ":8080")`

- 典型底层序列：**`socket` → `setsockopt`（若干）→ `bind` → `listen`**。  
- 返回 **`net.Listener`**，内部持有 **监听 fd**。  
- **`listener.Accept()`** → 底层 **`accept`**，得到 **`net.Conn`**（封装 **连接 fd**）。

### 2. 客户端 `net.Dial("tcp", "host:8080")`

- 典型底层序列：**`socket` →（可选源地址绑定）→ `connect`**。  
- 返回 **`net.Conn`**，封装 **已连接 fd**。  
- 名称解析、Happy Eyeballs、**`Dialer` 超时/KeepAlive** 等由 `net` 包在 **`connect`** 前后组合完成（读源码以 **`src/net`** 为准）。

### 3. 常见错误与排查

| 错误 | 常见根因 |
|------|----------|
| `address already in use` | 端口占用、`TIME_WAIT` 复用策略 |
| `connection refused` | 对端未监听、防火墙、错误 IP/端口 |
| `i/o timeout` | `Dial`/`Read` 超时时限过短或网络差 |

---

## 五、与前后章节

- 握手与队列内核行为：[04-三次握手内核做了什么](./04-三次握手内核做了什么.md)、[07](./07-半连接队列与全连接队列.md)。  
- 用户态协议笔记：[04-tcp · 三次握手](../04-tcp/03-三次握手.md)。

---

## 六、极简总结

- 服务端：**`socket` → `bind` → `listen` → `accept`**。  
- 客户端：**`socket` → `connect`**。  
- **`accept` 返回新 fd** 是并发服务「一线程/一协程一连接」的物理基础之一。

---

## 导航

- 上一篇：[02-file-descriptor-fd](./02-file-descriptor-fd.md)  
- 下一篇：[04-三次握手内核做了什么](./04-三次握手内核做了什么.md)
