# Chapter 03 — Socket-Level Programming（套接字层级编程）

> 对应原书 *Network Programming with Go Language, 2nd Ed.* 第 3 章。  
> **Socket** 是端到端通信在操作系统中的核心抽象：从 IoT 边缘到云原生高并发，最终都落在「**如何寻址、如何收发字节**」上。本章对齐 **TCP/IP 分层直觉** 与 **Go `net` 包** 的常用 API。

**Legacy 深度笔记**（Linux 队列、TCP 状态机、`netpoll`、调优等）：见 [本目录 README](./README.md) 中的 **`legacy-topic-index/`** 链接表。

---

## 3.1 TCP/IP 协议栈（The TCP/IP Stack）

TCP/IP 强调 **「实现先行、RFC 固化」**，与 OSI 的「先标准后落地」形成对照；工程上互联网以 **TCP/IP** 为事实标准。

| OSI（概念分层） | TCP/IP（常见归纳） | 职责 | 常见协议/组件 |
|----------------|-------------------|------|----------------|
| 5–7 应用/表示/会话 | **应用层** | 进程间协议语义 | HTTP、DNS、TFTP、FTP |
| 4 传输层 | **传输层** | 端到端端口、可靠/尽力 | TCP、UDP、SCTP、QUIC |
| 3 网络层 | **网际层** | 逻辑寻址、路由 | IPv4/IPv6、ICMP |
| 1–2 物理/链路 | **网络接口**（链路与物理常合并讲） | 帧与比特传输 | 以太网、Wi‑Fi、蜂窝 |

**校验与语义（IPv4 视角）**

- **IP 数据报**：**无连接、尽力交付**。**IPv4 首部**带首部校验和；**TTL 每经一跳递减**时，转发设备需**重算 IPv4 首部校验和**（IPv6 首部不再包含等价字段，错误检测更多依赖上层与链路层）。  
- **UDP**：在 IP 上加**端口**与**长度**等；**不保证**顺序与到达 → 实时音视频、遥测等可容忍丢包场景。  
- **TCP**：在不可靠 IP 上构造**可靠、有序的字节流**（序列号、ACK、重传、拥塞控制等）。

---

## 3.2 互联网地址（Internet Addresses）

**IPv4**：32 位，点分十进制；传统 **A/B/C 类**划分浪费地址 → **CIDR** 用**可变前缀长度**（如 `/22` 表示网络前缀 22 位，主机部分 `32-22=10` 位，需再扣广播/网络地址语义依上下文）。  

**IPv6**：128 位，冒分十六进制；地址空间巨大，**设计上更鼓励端到端**（生产环境仍可能出现翻译/过渡方案，不宜教条认为「绝对没有 NAT」）。

**常见结构（教学口径，非唯一分配策略）**：IPv6 全球单播地址常讨论 **/64 子网边界** + **64 位接口标识符**（IID）等模型；**真实前缀长度**由运营商与站点规划决定（如 `/48`、`/56` 等），**不要**把某一种「64+16+48」位切分当成全网固定真理。

---

## 3.3 `net.IP` 与地址解析

`net.IP` 底层为 **16 字节**表示（兼容嵌入 IPv4 的 IPv4-mapped IPv6），API 上可同时承载 **IPv4 / IPv6**。

```go
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	addr := net.ParseIP(os.Args[1])
	if addr == nil {
		fmt.Println("Invalid address format")
		return
	}
	fmt.Println("Normalized:", addr.String())
}
```

**易错点**：`net.ParseIP` **只接受规范形式**（点分十进制 IPv4、RFC5952 风格 IPv6 等）。**不要**依赖某些 Shell/工具把 `127.1` 展开成 `127.0.0.1` 的行为——标准库为严谨性通常**不**支持这类「暧昧缩写」。社区对解析边界有过长期讨论，工程上应以 **`net` 文档与测试**为准。

---

## 3.4 文档与标准库示例（Using Documentation）

- **`go doc net.IP`**：快速查看 `IsLoopback`、`To4`、`DefaultMask` 等。  
- **读 `$GOROOT/src/net/*.go` 与 `*_test.go`**：边缘用例与回归用例往往最诚实。  
- 在 `$(go env GOROOT)/src/net` 下可用 **`go test -list 'ParseIP'`**（或更宽的模式）列出与解析相关的测试名，再打开对应 `*_test.go` 阅读。

---

## 3.5 `IPMask` 与子网

```go
package main

import (
	"fmt"
	"net"
)

func main() {
	addr := net.ParseIP("103.232.159.187")
	mask := net.CIDRMask(24, 32)
	network := addr.Mask(mask)
	fmt.Printf("Mask: %s\n", mask.String())
	fmt.Printf("Network: %s\n", network.String())
}
```

**`DefaultMask`**：历史分类网络遗留，**主要对 IPv4 有意义**；CIDR/IPv6 场景应 **`CIDRMask` 显式**给出前缀长度。

---

## 3.6 基础路由（示意）

「最长前缀匹配」等完整逻辑在内核路由表；下面仅示意 **「目的 IP 属于哪条前缀 → 下一跳」** 的朴素比较。

```go
package main

import (
	"fmt"
	"net"
)

func main() {
	table := []struct {
		mask    net.IPMask
		network net.IP
		nextHop net.IP
	}{
		{net.CIDRMask(24, 32), net.IPv4(192, 17, 7, 0), net.IPv4(192, 12, 7, 251)},
		{net.CIDRMask(0, 32), net.IPv4(0, 0, 0, 0), net.IPv4(10, 10, 10, 10)}, // 默认路由示意
	}
	dest := net.ParseIP("192.17.7.20")
	for _, e := range table {
		if dest.Mask(e.mask).Equal(e.network) {
			fmt.Printf("Next hop for %s → %s\n", dest, e.nextHop)
			break
		}
	}
}
```

生产路由应使用**最长前缀匹配**、策略路由、多路径等；此处仅为读懂「Mask + Equal」的心智模型。

---

## 3.7 `IPAddr` 类型

**`net.IPAddr`** 封装 **IP** 与可选 **Zone**（IPv6 **链路本地** `fe80::/10` 等必须区分**网卡接口**时，Zone 填接口名如 `eth0`）。

**工程建议**：**`ResolveIPAddr`** 往往只体现「解析到的某个地址」；需要**多 A/AAAA 轮询、客户端负载均衡或故障转移**时，优先 **`LookupHost`** / **`LookupIP`** 或自定义 **`net.Resolver`**。

---

## 3.8 主机规范名称与地址查找

**`LookupCNAME`**、**`LookupHost`** 等见下例（错误处理生产代码中不应忽略）。

```go
package main

import (
	"fmt"
	"net"
)

func main() {
	const name = "go.dev"
	if cname, err := net.LookupCNAME(name); err == nil {
		fmt.Println("CNAME:", cname)
	}
	if addrs, err := net.LookupHost(name); err == nil {
		for _, a := range addrs {
			fmt.Println("A/AAAA:", a)
		}
	}
}
```

---

## 3.9 服务与端口（Services）

端口区分同一 IP 上的进程；**0–1023** 为知名端口（如 HTTP 80）。**`net.LookupPort("tcp", "http")`** 可把服务名映射到端口（依系统 `services` 文件，行为随环境略有差异）。

---

## 3.10 `TCPAddr`

`TCPAddr` 聚合 **IP、Port、Zone**。  
**实践**：在热路径上避免重复 DNS——可在启动阶段 **`ResolveTCPAddr`** 缓存结果，或对长连接复用已建立的 `net.Conn`（仍注意 DNS TTL 与扩缩容）。

---

## 3.11 TCP 套接字编程

TCP 提供**字节流**语义：无消息边界，应用要自己**组帧**（见第 4 章与 `08-framing-protocols`）。

### 客户端（示例：发 HTTP HEAD）

```go
package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	const service = "127.0.0.1:80" // 示例：本机有 HTTP 服务时
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	if err != nil {
		panic(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	if _, err := conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n")); err != nil {
		panic(err)
	}
	result, err := io.ReadAll(conn)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(result))
}
```

说明：历史上常用 **`ioutil.ReadAll`**，现代 Go 请用 **`io.ReadAll`**（`ioutil` 已弃用）。

### 顺序服务器（单连接处理完再接下一个）

吞吐受「单次处理延迟」严重限制；仅作教学。

```go
package main

import (
	"net"
	"time"
)

func main() {
	listener, _ := net.Listen("tcp", ":1200")
	for {
		conn, _ := listener.Accept()
		_, _ = conn.Write([]byte(time.Now().String()))
		_ = conn.Close()
	}
}
```

### 并发服务器（每连接一 Goroutine）

```go
package main

import (
	"net"
	"time"
)

func main() {
	listener, _ := net.Listen("tcp", ":1200")
	for {
		conn, _ := listener.Accept()
		go func(c net.Conn) {
			defer c.Close()
			_, _ = c.Write([]byte(time.Now().String()))
		}(conn)
	}
}
```

**评估**：Goroutine 比「一线程一连接」轻，但上限仍受 **fd 上限、内存、带宽、应用逻辑**约束；高阶模型见 **`legacy-topic-index/07-go-netpoll`** 与 **`10-server-architecture`**。

---

## 3.12 控制 TCP 连接

- **`SetDeadline` / `SetReadDeadline` / `SetWriteDeadline`**：避免读写**永久阻塞**；对端死锁、分区、半开连接时仍能回收。  
- **KeepAlive**：由 TCP 保活探测对端是否仍存活（参数与 OS 默认值需结合运维调优）。

---

## 3.13 UDP 数据报

无连接；**`ReadFrom` / `WriteTo`**（或 UDPConn 的 `ReadFromUDP` 等）必须带上对端地址。应用层需自管**丢包、乱序、重复、分片与 MTU**。

```go
// 概念：ReadFrom → 处理 → WriteTo 回同一 addr
// 详见 legacy-topic-index/05-udp-multicast
```

---

## 3.14 多协议监听

C 时代常在 **`select`/`epoll`** 里多路复用多个 fd；Go 里常见做法是 **多个 `Accept` 循环各自跑在 Goroutine** 中，由 **`net` + netpoll** 与调度器对接底层多路复用（见 `07-go-netpoll`）。

---

## 3.15 `Conn`、`PacketConn`、`Listener`

- **`net.Conn`**：`Read`/`Write`/`Close`… — **流式**（TCP、`TLS.Conn` 等）。  
- **`net.PacketConn`**：**数据报**语义（`ReadFrom`/`WriteTo`）。  
- **`net.Listener`**：`Accept` 抽象。

接口驱动便于写 **`Handle(c net.Conn)`** 一类通用逻辑（TCP 与 TLS 上均可复用心智）。

---

## 3.16 原始套接字与 `IPConn`

可绕过 TCP/UDP，直接构造 **IP 层**报文（如 **ICMP Echo**）。  

**警告**：通常需要 **管理员/root** 权限；错误报文可能干扰本机协议栈；**ICMP 校验和与跨平台差异**必须按 OS 文档实现。下面仅为**结构占位**，非完整可 ping 实现：

```go
// 概念示意：net.DialIP("ip4:icmp", laddr, raddr) 后自行组 ICMP 并计算校验和
// 生产请使用成熟库或系统 ping，并遵守权限与安全策略
```

---

## 3.17 本章小结

本章解决 **「如何寻址」** 与 **「如何传输字节流/数据报」**：IP/掩码/DNS、TCP/UDP 与 `net` 抽象、连接治理与并发模型。  
**仅传字节还不够**——下一章 **数据序列化** 讨论如何把内存对象变成**跨进程、可演进**的线路格式（见 [`chapter04` 笔记](../../02-general-network-stage/chapter04-data-serialization/note.md)）。
