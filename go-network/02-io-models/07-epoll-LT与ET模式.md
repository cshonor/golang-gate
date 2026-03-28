# epoll：LT 与 ET 模式

> **02-io-models · IO 模型全解**

## 内容大纲

- LT（水平触发）：只要条件满足就持续通知；写起来像阻塞 IO 心智
- ET（边缘触发）：仅在状态变化边沿通知；必须读写到 `EAGAIN`
- ET 下 `read`/`write` 典型 bug：漏读导致「饿死」、写事件未再触发
- `EPOLLOUT` 注册策略：是否需要一直监听？
- 与 nginx：`edge`/`kqueue` 配置对照（若扩展阅读）

## 正文

（待补充）
