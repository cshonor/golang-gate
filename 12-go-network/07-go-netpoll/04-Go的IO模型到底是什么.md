# Go 的 IO 模型到底是什么

> **07-go-netpoll · Go netpoll 高并发核心**

## 内容大纲

- 对外：同步阻塞式 API；对内：非阻塞 fd 加多路复用
- 归类：Reactor 加线程池（M）的口述清晰版
- 与 Node.js、Netty、Java NIO 的异同表
- 文件 IO：为何不少路径仍阻塞 OS 线程（版本演进简述）
- 面试：一句话加一张图

## Reactor 定性（扩写）

口述：**多路复用器（netpoll）只负责「就绪」**；**业务 `Read`/`Write` 仍在 goroutine 里同步执行**——整体接近 **Reactor + worker（G 即轻量 worker）**。与 **11-nio-reactor**、本目录 **[08-netpoll与GMP调度深度联动](./08-netpoll与GMP调度深度联动.md)** 对照。

## 正文

（待补充）
