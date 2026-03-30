# IO 多路复用：epoll 原理

> **02-io-models · IO 模型全解**

## 内容大纲

- `epoll_create` / `epoll_ctl` / `epoll_wait`：红黑树 + 就绪链表（概念级）
- 边缘触发 vs 水平触发预告；`EPOLLONESHOT` 等标志位
- 为什么 epoll 是 O(1) 就绪通知（相对 poll 的全量扫描）
- 与 `timerfd`/`eventfd`：统一事件源
- 常见面试：epoll 为什么高效？与 nginx/redis 事件循环关系

## 正文

（待补充）
