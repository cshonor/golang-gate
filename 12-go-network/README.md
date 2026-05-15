# Go 网络编程（按书本五阶段 + Legacy 深度笔记）

本模块 **Jan Newmarch / Ronald Petty《Network Programming with Go Language, 2nd Ed.》** 为纲：**顶层按原书重组为 5 个阶段目录**，每阶段下按 **chapterXX-英文slug** 对齐章节；原先按 **IO / TCP / netpoll / HTTP** 等技术线撰写的 **中文逐篇笔记** 整体迁入 [`legacy-topic-index/`](./legacy-topic-index/)，**文件名与正文不变**，由各章 `README` 与 [`legacy-topic-index/FILES.md`](./legacy-topic-index/FILES.md) 做入口。

---

## 五阶段（主入口 · 与书本一致）

| 阶段 | 目录 | 原书范围（章号） |
|------|------|------------------|
| 第一阶段 打底 | [01-foundation-stage](./01-foundation-stage/README.md) | Ch.1、2、3、5（建议读序 **1→2→3→5**） |
| 第二阶段 通用网络 | [02-general-network-stage](./02-general-network-stage/README.md) | Ch.4、6、7 |
| 第三阶段 Web 核心 | [03-web-core-stage](./03-web-core-stage/README.md) | Ch.8、9、10、16 |
| 第四阶段 进阶协议 | [04-advanced-protocol-stage](./04-advanced-protocol-stage/README.md) | Ch.11～15 |
| 第五阶段 工程化 | [05-engineering-stage](./05-engineering-stage/README.md) | Ch.17、附录 A、B |

**精读提纲**（阶段目标 + 分章要点）：[参考-Network-Programming-with-Go-第2版-五阶段学习路线.md](./参考-Network-Programming-with-Go-第2版-五阶段学习路线.md)  
**全书章节目录**：[参考-Network-Programming-with-Go-第2版-章节目录.md](./参考-Network-Programming-with-Go-第2版-章节目录.md)

---

## 深度笔记（Legacy 技术线）

- **索引说明**：[legacy-topic-index/README.md](./legacy-topic-index/README.md)
- **逐篇完整链接**：[legacy-topic-index/FILES.md](./legacy-topic-index/FILES.md)

`tools/apply_outlines.py` 写入路径已指向 **`legacy-topic-index/`**（与 `tools/outlines/*.json` 中 `rel` 一致）。

---

## 维护工具

在仓库根目录执行：

```bash
python 12-go-network/tools/apply_outlines.py
```

大纲数据在 `12-go-network/tools/outlines/*.json`。

---

## Windows 文件名说明（历史）

以下主题在 Windows 上不能用 `/` 作文件名，已用 **与** 或 **-** 代替（与正文无关）：

| 原计划 | 实际文件名 |
|--------|------------|
| 半连接队列/全连接队列 | `07-半连接队列与全连接队列.md` |
| SO_RCVBUF/SO_SNDBUF | `08-socket缓冲区SO_RCVBUF与SO_SNDBUF.md` |
| epoll/kqueue/IOCP | `03-epoll-kqueue-IOCP支持.md` |
| Read/Write 阻塞 | `05-Go的Read与Write为什么看起来阻塞.md` |

---

## 建议怎么读

1. **跟书**：从 **01-foundation-stage** 起，打开每章 **`README.md`**，按其中链接跳转到 `legacy-topic-index` 或仓库其它模块。  
2. **跟技术线深挖**：直接打开 [`legacy-topic-index/FILES.md`](./legacy-topic-index/FILES.md)，按 **01→13** 子目录顺序读（与重构前一致）。  
3. **第 3 章 + netpoll**：以 **chapter03-socket-programming** 的链接表为主干，穿插 **06-go-net-internals** 与 **07-go-netpoll**。
