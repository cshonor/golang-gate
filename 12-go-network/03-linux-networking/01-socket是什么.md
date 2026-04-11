# socket 是什么

> **03-linux-networking · Linux 网络底层**  
> 承接：[01-什么是 IO](../01-io-fundamentals/01-什么是IO.md)、[06-文件IO vs 网络IO](../01-io-fundamentals/06-文件IO%20vs%20网络IO.md)；IO 多路复用见 [02-io-models](../02-io-models/01-五大IO模型总览.md)。

---

## 一、核心定义

**Socket（套接字）** 是操作系统提供给用户程序的 **网络通信端点抽象**：应用通过 **一组系统调用** 与内核 **TCP/IP 协议栈** 交互，而不必直接操作网卡驱动。

- **本质**：在 Unix/Linux 上，socket 通常表现为 **文件描述符（fd）**——「一切皆文件」抽象的一部分。  
- **作用**：用相对统一的 **`read`/`write`/`send`/`recv`/`close`** 语义收发字节；具体语义随 **协议族（domain）**、**类型（type）** 而变。  
- **定位**：**应用层代码 ↔ 内核网络协议栈** 之间的编程接口。

---

## 二、Socket 与 fd

- 创建 socket 后，内核返回一个 **非负整数 fd**，进程后续对该 socket 的 IO 都通过该 fd 路由。  
- 与普通文件 fd 的相同点：**进程 fd 表索引**、`read`/`write`/`close` 的统一入口。  
- 不同点：背后不是 inode 磁盘文件，而是 **`struct socket` / `struct sock`** 等内核网络对象（概念了解即可，不必背字段）。

详见：[02-file-descriptor-fd](./02-file-descriptor-fd.md)。

---

## 三、TCP 与四元组

一条 **TCP 连接** 在 IP 层可由 **四元组** 唯一标识：

**（源 IP，源端口，目的 IP，目的端口）**

- **监听 socket**：常绑定 **`*:端口`**（如 `0.0.0.0:8080`），表示在本机多个地址上接受该端口连接（具体以 `bind` 地址为准）。  
- **已连接 socket**：`accept` 返回的 **新 fd** 代表一条与某客户端四元组对应的 **已建立连接**；**监听 fd** 仍继续监听。

---

## 四、常见 socket 类型（`type`）

| 类型 | 协议心智 | 典型用途 |
|------|----------|----------|
| **`SOCK_STREAM`** | 面向连接、可靠、有序 **字节流**（TCP） | HTTP、RPC、数据库协议 |
| **`SOCK_DGRAM`** | 无连接 **报文**（UDP） | DNS、日志、实时音视频（常配合应用层可靠性） |
| **`SOCK_RAW`** | 更底层访问 IP/协议号 | ping、自定义协议、抓包类工具（需权限） |

另有 **Unix Domain Socket（`AF_UNIX`）**：本机进程间通信，不经网络栈（Go 中 `net.Dial("unix", path)` 等）。

---

## 五、Socket 在工程里的角色

1. **入口**：所有 TCP/UDP 客户端、服务端最终都落到 **socket + fd**。  
2. **屏蔽细节**：校验和、重传、排序、窗口等由 **内核 TCP** 完成（UDP 则弱得多）。  
3. **高并发调优抓手**：缓冲区、backlog、`epoll` 监听的就是 **socket fd**（见 [08](./08-socket缓冲区SO_RCVBUF与SO_SNDBUF.md)、`02-io-models`）。

---

## 六、与 Go 的对应关系

- **`net.Listen` / `net.Dial`**：在 Linux 上对应 **`socket` → `bind`/`listen`/`accept`** 或 **`socket` → `connect`** 等组合（具体路径随解析器、`Dialer`、双栈等略有差异）。  
- **`net.Conn`**：封装 **已连接** 的 fd，**`Read`/`Write`** 对应用户缓冲与内核 socket 缓冲之间的拷贝（见 [01-io-fundamentals/05](../01-io-fundamentals/05-缓冲区、页缓存、零拷贝.md)）。

---

## 七、极简总结

- **Socket** = 网络端点抽象；在 Linux 上 **首先是 fd**。  
- **TCP** = 流式套接字 + 四元组标识的连接。  
- **Go 网络编程** 的底层对象仍是 **socket fd**；调优要回到 **队列、缓冲、状态机**。

---

## 导航

- 下一篇：[02-file-descriptor-fd](./02-file-descriptor-fd.md)  
- 后续协议与编程：[04-tcp · TCP 特点](../04-tcp/01-TCP特点.md)
