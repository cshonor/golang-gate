# socket 缓冲区：SO_RCVBUF / SO_SNDBUF

> **03-linux-networking · Linux 网络底层**

## 内容大纲

- 内核自动调优 vs `setsockopt` 固定：各自适用
- 与 BDP（带宽时延积）：高带宽长肥管下的窗口关系
- 应用 `read` 慢导致接收缓冲区满：对端 `write` 阻塞与零窗口
- 内存估算：并发连接 × 每连接缓冲
- 与 `TCP_NOTSENT_LOWAT`、写节流（扩展）

## 正文

（待补充）
