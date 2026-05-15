# 《Network Programming with Go Language, Second Edition》章节目录

> **作者**：Jan Newmarch、Ronald Petty  
> **说明**：第二版基于 **Go 1.18** 全面更新；下列为**中英对照章节目录**（便于与本仓库 `12-go-network` 笔记对照学习）。  
> **注意**：正式书名、小节标题以出版社 PDF/纸质书为准；此处为学习用提纲整理。

全书从底层 **Socket** 到应用层 **HTTP / REST / WebSocket**，并覆盖安全、测试与 Go 1.18 相关附录（模糊测试、泛型）等主题。

---

## 第 1 章：架构分层（Architectural Layers）

- **协议层（Protocol Layers）**：OSI 模型与 TCP/IP 协议栈  
- **网络基础**：网关（Gateways）、主机级网络与数据包封装  
- **连接模型**：面向连接（Connection Oriented）与无连接（Connectionless）  
- **通信与分布式计算模型**：消息传递（Message Passing）、远程过程调用（RPC）、C/S 与 P2P  
- **分布式系统的八大谬误（Eight Fallacies of Distributed Computing）**：可靠性、延迟、带宽与安全等常见误区  

## 第 2 章：Go 语言概述（Overview of the Go Language）

- **类型系统（Types）**：数组与切片、Maps、指针、函数、结构体与方法  
- **并发机制**：多线程（Multithreading）与 Goroutines  
- **项目与生态**：Packages、Modules 与 GOPATH  
- **错误处理（Error Values）**  

## 第 3 章：Socket 级编程（Socket-Level Programming）

- **TCP/IP 栈**：IP 数据报、UDP 与 TCP  
- **互联网地址（Internet Addresses）**：IPv4/IPv6、`IP` 类型与子网掩码（IPMask）、基本路由实现  
- **DNS 与主机解析**：CNAME 与 IP 查找  
- **TCP Sockets**：客户端、多线程服务端、超时与 KeepAlive  
- **UDP 数据报**与**多 Socket 监听**  
- **原始套接字（Raw Sockets）与 IPConn**：自定义底层协议（如 Ping/ICMP）  

## 第 4 章：数据序列化（Data Serialization）

- **结构化数据与互通协议（Mutual Agreement）**  
- **ASN.1**：基础语法、DER 与 Go 中的读写  
- **JSON**：`encoding/json` 与 C/S 通信  
- **Gob**：Go 原生二进制序列化  
- **Protocol Buffers**：安装、`protoc` 与生成 Go 代码  

## 第 5 章：应用层协议（Application-Level Protocols）

- **协议设计与版本控制**：HTTP 0.9/1.0/1.1/2/3 演进  
- **消息格式**：基于字节与基于字符的协议  
- **状态信息（State Information）**：系统状态、状态转换图与 C/S 模型  

## 第 6 章：管理字符集与编码（Managing Character Sets and Encodings）

- **定义**：字符、字符集、字符编码与传输编码（Transport Encoding）  
- **ASCII 与 UTF-8**：Go 中的处理与常见问题  
- **UTF-16 与 Go**：大端/小端字节序  

## 第 7 章：安全（Security）

- **ISO 安全架构**：完整性与认证  
- **加密算法**：对称密钥与公钥加密  
- **X.509 与 TLS**：TLS 服务器、自签名证书与安全传输  

## 第 8 章：HTTP

- **URLs、资源与 HTTP 版本**：HTTP/1.1、HTTP/2、HTTP/3（QUIC）  
- **简单 User Agents**：`Response` 与 GET/HEAD  
- **配置 HTTP 请求**：`Client` 与代理（Proxy）  
- **服务端**：文件服务、Handler、自定义 Multiplexer  
- **HTTPS** 服务器配置  

## 第 9 章：模板（Templates）

- **插入对象值**与**管道（Pipelines）**  
- **自定义模板函数**与**条件**  
- **`html/template`**：上下文感知与 XSS 防护  

## 第 10 章：构建完整的 Web 服务器（A Complete Web Server）

- 综合运用静态资源、模板与表单，完成「闪卡（Flashcard）」多页面项目  

## 第 11 章：HTML

- **Tokenizing HTML**：`golang.org/x/net/html`、HTML5 解析与标签树  

## 第 12 章：XML

- **XML 解析**：结构体 Marshal/Unmarshal  
- **XML 树遍历**：`StartElement`、`EndElement`、`CharData` 等  

## 第 13 章：远程过程调用（Remote Procedure Call）

- **Go 原生 RPC**  
- **HTTP RPC 与 TCP RPC**  
- **JSON RPC** 与跨语言通信  

## 第 14 章：REST

- **REST 风格**：URI 与资源、表现形式、动词约束  
- **无状态（Stateless）与 HATEOAS**  
- **REST 事务与 Richardson 成熟度模型**  
- **实战**：`ServeMux`、内容协商（Content Negotiation）、闪卡应用 REST 化  

## 第 15 章：WebSockets

- **协议原理**：从 HTTP 握手升级到全双工  
- **`golang.org/x/net/websocket`**：Message、JSON、自定义 Codec  
- **WSS** 与浏览器 JavaScript  
- **`github.com/gorilla/websocket`**  

## 第 16 章：Gorilla 工具包

- **中间件（Middleware）**  
- **`gorilla/mux`**：按主机、前缀、Header、Query 等路由  
- **Gorilla Handlers**：日志、内容协商等  
- **扩展**：`gorilla/rpc`、`gorilla/schema`、`gorilla/securecookie`  

## 第 17 章：测试（Testing）

- **网络测试的困境与反模式**  
- **`httptest`**：`NewRecorder`、`NewServer`  
- **`net.Pipe()`**：内存双工连接做隔离测试  
- **标准库测试用例阅读**  

## 附录 A：模糊测试（Fuzzing）

- Go 1.18 **Fuzzing** 简介  
- 在网络相关代码上的应用思路  

## 附录 B：泛型（Generics）

- Go 1.18 **泛型**  
- **重构示例**：从普通代码到泛型、自定义约束（Constraints）  
- **集合与并发中的泛型**及**不宜使用泛型**的场景  
