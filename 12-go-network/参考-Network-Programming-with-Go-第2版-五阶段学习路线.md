# 《Network Programming with Go Language, 2nd Ed.》五阶段学习路线

> 与 [章节目录](./参考-Network-Programming-with-Go-第2版-章节目录.md) 配套：按**打底 → 通用能力 → Web → 进阶协议 → 工程化**重组精读提纲（章号与书中一致）。  
> **关于第 5 章**：原书 **1～17 章 + 附录 A/B 一章不少**；第 5 章在书中本就存在，这里**刻意**放在**第一阶段末尾**（紧接 Socket），强调「应用层长什么样、如何设计协议」这一层**思想铺垫**，与第 8/13/14/15 等**落地实现**章呼应——并非遗漏，而是**阶段归类**。**书本对应的五阶段英文目录**已并入本模块顶层（[`01-foundation-stage`](./01-foundation-stage/README.md) … [`05-engineering-stage`](./05-engineering-stage/README.md)），见 [模块 README](./README.md)。  
> **阅读顺序提示**：原书纸面顺序是第 4 章在 5、6 之前；本路线**第一阶段读 1→2→3→5**，**第二阶段再从第 4 章序列化**接 6、7，逻辑上先「能发包」再「包里装什么、怎么编码、怎么加密」。

---

## 第一阶段：打底必学（地基阶段）

**核心目标**：掌握网络分层与 TCP/IP 直觉、建立 Go 并发/类型/错误模型与后续 Socket 编程的衔接，避免「只会调库、不懂原理」。

### 第 1 章 架构分层（Architectural Layers）

**核心知识点**

- **协议层**：以 **TCP/IP 协议栈** 为主（替代 OSI 的教学栈），理解各层职责（网络接口层 → 网络层 → 传输层 → 应用层）及数据在各层的封装/解封装。
- **网络基础**：网关（连接不同网络、转发数据包）、主机级网络、从应用数据到帧的封装与接收端反向解析。
- **连接模型**：**面向连接（TCP）** vs **无连接（UDP）**——握手/可靠 vs 无连接/尽力交付；与后续 Socket 选型一致。
- **通信与分布式模型**：消息传递、RPC、C/S 与 P2P。
- **分布式系统八大谬误**（强记，避坑）：
  1. 网络可靠  
  2. 延迟为零  
  3. 带宽无限  
  4. 网络安全  
  5. 拓扑不变  
  6. 只有一个管理员  
  7. 传输成本为零  
  8. 网络同构  

**关键总结**：建立**分层思维** + 记住八大谬误；后续 Socket、HTTP、RPC 都建立在这套认知上。

结构化笔记：[`01-foundation-stage/chapter01-architectural-layers/note.md`](./01-foundation-stage/chapter01-architectural-layers/note.md)。

### 第 2 章 Go 语言概述（Overview of the Go Language）

**核心知识点（面向网络编程）**

- **类型系统**：切片（缓冲、报文拼接）、`map`（配置、表驱动）、指针（少拷贝、共享状态）、结构体与方法（封装连接/会话对象）。
- **并发**：Goroutine 处理多连接；与 **GMP** 调度直觉衔接（不必在这一章抠完调度细节）。
- **生态**：`package`、`go mod` 依赖管理（替代纯 GOPATH 心智）。
- **错误处理**：`error` 值而非异常；连接/读写/超时/取消等错误路径是网络代码的主线。

**关键总结**：抓 **Goroutine + 切片/map + error**，够支撑后续高并发与健壮 I/O。

结构化笔记：[`01-foundation-stage/chapter02-go-language-overview/note.md`](./01-foundation-stage/chapter02-go-language-overview/note.md)。

### 第 3 章 Socket 级编程（Socket-Level Programming）

**核心知识点**

- **TCP/IP 细节**：IP 数据报；UDP（无连接、数据报）；TCP（连接、可靠、字节流）。
- **互联网地址**：IPv4/IPv6、`IP`、`IPMask`、路由表直觉。
- **DNS**：域名 → IP、`CNAME`、`net.LookupIP` 等。
- **TCP Socket（实战重心）**：
  - 客户端：`Dial`、读写、关闭；
  - 服务端：`Listen`、`Accept`、每连接 **goroutine**（或后续模型演进）；
  - **超时**、**KeepAlive** 等连接治理。
- **UDP**：`WriteTo` / `ReadFrom`、无连接语义（丢包/乱序风险）、多 socket 监听。
- **Raw / `IPConn`**：绕过传输层、自定义 IP 层行为（如 ICMP/Ping 类示例）；注意权限与平台差异。

**关键总结**：必须手写最小 **TCP 客户端/服务端** 与 **UDP 收发**；理解 TCP/UDP 差异与超时/KeepAlive，再学 HTTP/WebSocket 才有锚点。

结构化笔记：[`01-foundation-stage/chapter03-socket-programming/note.md`](./01-foundation-stage/chapter03-socket-programming/note.md)。

### 第 5 章 应用层协议（Application-Level Protocols）

**定位**：**思想层 / 设计层**——版本、报文形态、状态如何建模；**几乎没有独立「语法 API」**，但决定你后面读 HTTP、RPC、REST、WebSocket 时「为什么长这样」。

**核心知识点**

- **协议设计与版本控制**：兼容旧客户端、能力协商、渐进升级（与 HTTP/2/3 演进章衔接）。
- **消息格式**：**面向字节**（长度前缀、定长、 TLV）vs **面向字符/文本**（行协议、HTTP 头）；和「粘包/拆包、边界」问题挂钩。
- **状态信息**：无状态 vs 有状态会话；**状态机**直觉（请求生命周期、连接阶段）；在 C/S 模型里状态放服务端、客户端还是 token/cookie。
- **与后续章节的关系**：本章是「**应用层协议共性**」；第 8 章 HTTP、第 13 章 RPC、第 14 章 REST、第 15 章 WebSocket 都是**具体实例化**。

**关键总结**：把第 5 章当 **checklist**：每学一个具体协议，问自己——**帧边界怎么做？状态放哪？如何升级版本？失败怎么表示？**

结构化笔记：[`01-foundation-stage/chapter05-application-protocols/note.md`](./01-foundation-stage/chapter05-application-protocols/note.md)。

---
## 第二阶段：通用网络必备（跨服务通信）

**核心目标**：序列化、字符编码、安全传输——解决「数据如何高效、可读、安全地跨进程/跨语言」。

### 第 4 章 数据序列化（Data Serialization）

**核心知识点**

- 结构化数据与**双方约定格式**（互操作性）。
- **ASN.1 / DER**：`encoding/asn1`；偏证书与传统协议语境。
- **JSON**：`encoding/json`、`Marshal`/`Unmarshal`、结构体标签、嵌套/空值/类型差异；接口与爬虫 JSON 主力。
- **Gob**：Go 之间高效二进制、常见于 Go 内部 RPC。
- **Protobuf**：`protoc`、`.proto`、生成 Go 代码；微服务与跨语言高频。

**关键总结**：**JSON + Protobuf** 优先熟练；**Gob** 服务内；**ASN.1** 了解即可；按场景选型。

结构化笔记：[`02-general-network-stage/chapter04-data-serialization/note.md`](./02-general-network-stage/chapter04-data-serialization/note.md)。

### 第 6 章 管理字符集与编码（Managing Character Sets and Encodings）

**核心知识点**

- 字符 / 字符集 / 字符编码 / **传输编码**（如 Base64）区分清楚。
- **ASCII 与 UTF-8**：Go 源码与 `string` 默认 UTF-8；`len` 为**字节数**；中文常见多字节；截取与乱码排查与 `[]byte`/`[]rune` 的配合（可与仓库 `00-basic-types/04-字符串与rune.md` 对照）。
- **UTF-16**：大小端；`unicode/utf16` 与 `golang.org/x/text/encoding/unicode`；Windows/部分协议场景。

**关键总结**：日常以 **UTF-8** 为主；UTF-16 掌握字节序与何时出场即可。

结构化笔记：[`02-general-network-stage/chapter06-charset-encoding/note.md`](./02-general-network-stage/chapter06-charset-encoding/note.md)。

### 第 7 章 安全（Security）

**核心知识点**

- ISO 安全架构中的**完整性**与**认证**直觉。
- **对称加密**（如 AES，大量数据）vs **公钥加密**（密钥交换、签名，如 RSA）。
- **X.509 + TLS**：自签名（联调）vs CA（生产）；`crypto/tls`、**HTTPS / WSS** 配置。

**关键总结**：能配 **TLS/HTTPS**；分清对称/非对称适用场景；证书链与主机名校验要有基本概念。

结构化笔记：[`02-general-network-stage/chapter07-security/note.md`](./02-general-network-stage/chapter07-security/note.md)。

---

## 第三阶段：Web 服务核心（求职高频）

**核心目标**：`net/http`、路由与中间件、模板与小型完整站点；与本书闪卡案例对齐。

### 第 8 章 HTTP

**核心知识点**

- URL、资源、**HTTP 版本演进**（0.9 → 1.0 → 1.1 → 2 → 3/QUIC）与各自解决的问题。
- **Client**：GET/HEAD、`Response`、自定义 `Client`（超时、代理、Header）。
- **Server**：`FileServer`、自定义 `Handler`、`ServeMux` 与路由扩展。
- **HTTPS**：结合 TLS 配置证书与密钥。

**关键总结**：动手做最小 HTTP/HTTPS 服务 + 自定义 Handler + Client；**HTTP/1.1 与 HTTP/2** 特性面试常问。

结构化笔记：[`03-web-core-stage/chapter08-http/note.md`](./03-web-core-stage/chapter08-http/note.md)。

### 第 16 章 Gorilla 工具包

**核心知识点**

- **中间件**：嵌套 `Handler`、横切关注点（日志、鉴权、CORS）。
- **`gorilla/mux`**：主机名、前缀、Header、Query、路径参数、分组路由。
- **Gorilla Handlers**：日志、内容协商等现成中间件。
- **扩展**：`gorilla/rpc`、`gorilla/schema`（表单绑定）、`gorilla/securecookie`。

**关键总结**：企业里 **`mux` + 中间件** 很常见；`schema` / `securecookie` 按需深入。

结构化笔记：[`03-web-core-stage/chapter16-gorilla-toolkit/note.md`](./03-web-core-stage/chapter16-gorilla-toolkit/note.md)。

### 第 9 章 模板（Templates）与第 10 章 完整 Web 服务器

**核心知识点**

- 插值、**管道**、自定义函数、条件分支。
- **`html/template`**：上下文感知，防 **XSS**（Web 场景优先于 `text/template`）。
- **第 10 章**：静态资源 + 模板 + 表单，**闪卡**多页面串联。

**关键总结**：模板以 **安全默认** 为先；第 10 章建议完整敲一遍，形成项目级肌肉记忆。

结构化笔记：[`03-web-core-stage/chapter09-templates/note.md`](./03-web-core-stage/chapter09-templates/note.md) · [`03-web-core-stage/chapter10-complete-web-server/note.md`](./03-web-core-stage/chapter10-complete-web-server/note.md)。

---

## 第四阶段：进阶协议与架构

**核心目标**：RPC、REST、WebSocket；再补 HTML/XML 解析类技能（爬虫、协议适配）。

### 第 13 章 远程过程调用（RPC）

**核心知识点**

- `net/rpc` 方法签名约定、注册与调用。
- **TCP RPC** vs **HTTP RPC**（`rpc.HandleHTTP` 等）。
- **JSON-RPC**：跨语言、与 `encoding/json` 衔接。

**关键总结**：理解「远程像本地」的边界（网络错误、版本、幂等）；能搭最小 RPC 服务端/客户端。

结构化笔记：[`04-advanced-protocol-stage/chapter13-rpc/note.md`](./04-advanced-protocol-stage/chapter13-rpc/note.md)。

### 第 14 章 REST

**核心知识点**

- URI 与资源、**Representation**、HTTP 动词语义、**Stateless**、**HATEOAS**（进阶）。
- **Richardson 成熟度模型**、状态码与错误表达。
- **实战**：闪卡 REST 化、`ServeMux` 或 `gorilla/mux`、**Content Negotiation**（JSON/XML 等）。

**关键总结**：能设计可读 URI + 正确使用动词/状态码；能做一次从「页面思维」到「资源思维」的重构。

结构化笔记：[`04-advanced-protocol-stage/chapter14-rest/note.md`](./04-advanced-protocol-stage/chapter14-rest/note.md)。

### 第 15 章 WebSockets

**核心知识点**

- HTTP **Upgrade** 握手 → 全双工长连接。
- `golang.org/x/net/websocket`：Message、JSON、`Codec`（注意：生态上更多新项目偏向 **`gorilla/websocket`**；以书为准练习、以工程选型为准落地）。
- **WSS**、浏览器 `WebSocket` API。
- **`gorilla/websocket`**：消息类型、ping/pong、缓冲区与关闭语义。

**关键总结**：掌握握手、心跳、断线重连直觉；生产用 **WSS**。

结构化笔记：[`04-advanced-protocol-stage/chapter15-websocket/note.md`](./04-advanced-protocol-stage/chapter15-websocket/note.md)。

### 第 11 章 HTML 与第 12 章 XML

**核心知识点**

- **`golang.org/x/net/html`**：token、DOM 式遍历、抽取节点。
- **`encoding/xml`**：Marshal/Unmarshal；`Decoder` 与 `StartElement`/`EndElement`/`CharData` 流式解析。

**关键总结**：能解析页面/报文即可，不必实现解析器。

结构化笔记（第 11 章）：[`04-advanced-protocol-stage/chapter11-html-parse/note.md`](./04-advanced-protocol-stage/chapter11-html-parse/note.md)。  
结构化笔记（第 12 章）：[`04-advanced-protocol-stage/chapter12-xml-parse/note.md`](./04-advanced-protocol-stage/chapter12-xml-parse/note.md)。

---

## 第五阶段：工程化与新特性

**核心目标**：可测、可模糊测、可泛型抽象——对齐企业规范与 Go 1.18+ 工具链。

### 第 17 章 测试（Testing）

**核心知识点**

- 网络测试的**不稳定**来源与反模式（强依赖外网/真实端口）。
- **`net/http/httptest`**：`ResponseRecorder`、`NewServer`。
- **`net.Pipe`**：内存双工、隔离 Socket 级逻辑。
- 阅读标准库 **`net` / `net/http`** 测试作为范本。

**关键总结**：**httptest** 必会；能写 Handler 级单测与客户端假后端。

结构化笔记：[`05-engineering-stage/chapter17-testing/note.md`](./05-engineering-stage/chapter17-testing/note.md)。

### 附录 A 模糊测试（Fuzzing）与附录 B 泛型（Generics）

**核心知识点**

- **Fuzzing**：随机输入轰炸边界；适合协议解析、反序列化、复杂字符串处理。
- **泛型**：约束、集合/通道类工具代码复用；避免为泛型而泛型。

**关键总结**：Fuzzing 抓「解析与输入边界」；泛型提升**可复用工具**质量，网络工具函数可逐步泛型化。

---

## 整体学习总结

路线覆盖：**底层理论 → Go 与 Socket → 应用层协议思想（第 5 章）→ 序列化/编码/安全 → HTTP 与 Web 栈 → RPC/REST/WebSocket → 解析与测试 → Fuzzing/泛型**。
务必配合**每章动手代码**（尤其 TCP/UDP、HTTP、RPC、WebSocket），并与本仓库 [12-go-network](./README.md) 中 IO、TCP、`net/http` 等笔记交叉复盘，形成自己的「原理 + 标准库 + 工程选型」闭环。
