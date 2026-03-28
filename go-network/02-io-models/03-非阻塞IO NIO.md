# 非阻塞 IO（NIO）

> **02-io-models · IO 模型全解**

## 内容大纲

- 配置：`O_NONBLOCK`；`accept`/`read`/`write`/`connect` 的 `EAGAIN` 语义
- 忙轮询的代价：CPU 100% 空转；必须与复用或让出结合
- 自旋 + `sched_yield` vs 直接阻塞：何时得不偿失
- 与边缘触发（ET）配合时的读尽/写尽策略
- Go 中「非阻塞 fd + netpoll」如何把 NIO 包装成「协程级阻塞」

## 正文

（待补充）
