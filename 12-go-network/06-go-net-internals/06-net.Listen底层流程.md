# `net.Listen` 底层流程（服务端监听）

> **06-go-net-internals · Go net 包源码级理解**  
> 上篇：[05-net.Dial底层流程](./05-net.Dial底层流程.md)；内核队列见 [03-linux 07](../03-linux-networking/07-半连接队列与全连接队列.md)。

---

## 一、总览（TCP）

`net.Listen("tcp", ":8080")` 概念序列：

1. **解析地址**：`":8080"` → 绑定 **`0.0.0.0:8080`**（IPv4 心智；IPv6 另论）。  
2. **`socket`**：创建监听 fd。  
3. **`setsockopt`**：`SO_REUSEADDR` 等（平台/版本差异）。  
4. **`bind`**：占端口。  
5. **`listen(backlog)`**：进入 **`LISTEN`**，启用 **SYN/accept 队列**（长度受 **`somaxconn`** 等裁剪）。  
6. **封装 `TCPListener`**：内部仍持有 **`netFD` + poll**，`Accept` 返回的新连接再各自 **`netFD`**。

---

## 二、`backlog` 与 Go

- Go 传入 `listen` 的 backlog **有默认与上限**；调大 **`somaxconn`** 而不加快 **`Accept` 处理**，队列仍会满。  
- **高并发**：`Accept` 循环必须 **轻**，把重活丢给 worker。

---

## 三、与 `Dial` 的差异

| 项 | Listen（TCP） | Dial（TCP） |
|----|----------------|-------------|
|  syscall 主链 | `socket`→`bind`→`listen` | `socket`→`connect` |
| 返回对象 | `Listener` | `Conn` |
| 阻塞热点 | **`Accept` 等连接** | **握手完成** |

---

## 四、极简总结

- **`Listen`** = **占坑 + 排队接客**（内核队列 + 用户态 `Accept`）。  
- **`Accept` 出来的每个 `Conn`** 都要 **独立注册 netpoll**。

---

## 导航

- 上一篇：[05-net.Dial底层流程](./05-net.Dial底层流程.md)  
- 下一篇：[07-pollDesc核心结构与原理](./07-pollDesc核心结构与原理.md)
