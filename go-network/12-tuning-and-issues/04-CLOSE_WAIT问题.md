# CLOSE_WAIT 问题

> **12-tuning-and-issues · 性能调优与线上问题**

## 内容大纲

- 语义：对端已 FIN，本端应用未 close
- 典型 bug：漏关 Conn、goroutine 阻塞在读、池化泄漏
- 排查：ss、lsof、应用堆栈
- 与反向代理超时
- 修复：defer Close、context、finally 模式

## 正文

（待补充）
