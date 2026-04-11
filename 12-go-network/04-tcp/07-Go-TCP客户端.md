# Go TCP 客户端

> **04-tcp · TCP 协议与编程**  
> 上篇：[06-Go-TCP服务端](./06-Go-TCP服务端.md)。

---

## 一、最简 `Dial`

```go
conn, err := net.Dial("tcp", "example.com:443")
```

- 解析 **DNS**、选择 **源地址**、执行 **`connect`**（三次握手）。  
- 失败常见：**`connection refused`**、**`i/o timeout`**、**`no such host`**。

---

## 二、生产应使用 `net.Dialer`

```go
d := net.Dialer{
    Timeout:   5 * time.Second,
    KeepAlive: 30 * time.Second,
}
conn, err := d.DialContext(ctx, "tcp", addr)
```

- **`Timeout`**：建立连接阶段上限。  
- **`DialContext`**：与 **`ctx`** 绑定取消。  
- **`KeepAlive`**：启用 **TCP keepalive** 探测（与 [08](./08-TCP心跳保活.md) 衔接）。

---

## 三、连接复用

- **短连接**：每次请求 `Dial` 新连接 → **`TIME_WAIT` 多**、握手多。  
- **长连接 / 池**：`database/sql`、HTTP `Transport` 等内部维护 **`idle` 连接**；注意 **服务端空闲断开** 与 **半开连接**。

---

## 四、读写到 `conn`

- **`Read` 可能返回部分长度**：循环读满或按协议解析。  
- **`Write` 可能部分写**：需处理 **`n < len(buf)`** 或缓冲 **`bufio.Writer`**。  
- **TLS**：`tls.Client(conn, cfg)` 包装后再读写。

---

## 五、极简总结

- **`Dialer` + `Context`** 是客户端标配。  
- **超时** 分 **连接建立** 与 **读写** 两层。  
- **连接复用** 决定延迟与 **`TIME_WAIT`** 压力。

---

## 导航

- 上一篇：[06-Go-TCP服务端](./06-Go-TCP服务端.md)  
- 下一篇：[08-TCP心跳保活](./08-TCP心跳保活.md)
