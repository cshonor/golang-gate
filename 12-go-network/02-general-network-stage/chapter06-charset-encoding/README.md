# Chapter 06 — Managing Character Sets and Encodings（字符集与编码）

原书第 6 章：术语栈、ASCII、ISO 8859、Unicode、UTF-8/`rune`、流式 UTF-8、UTF-16 与字节序、正规化、遗留编码（`x/text`）。

- [note.md](./note.md) — 本章结构化笔记（精读版，**6.1～6.14**）

**目录速览**：6.1 定义 · 6.2 ASCII · 6.3 ISO 8859 · 6.4 Unicode · 6.5 UTF-8 与 rune · 6.6 流式 UTF-8 · 6.7 ASCII 假设风险 · 6.8 UTF-16 与标准库 · 6.9 字节序与 BOM · 6.10 `x/text` UTF-16 · 6.11 陷阱与正规化 · 6.12 `charmap` · 6.13 GBK 等 · 6.14 小结与清单。

## 与本仓库其它模块

- [00-basic-types/04-字符串与rune.md](../../../00-basic-types/04-字符串与rune.md) — `string`/`[]byte`/`[]rune` 基础

与 **第 4 章序列化**、**第 5 章文本协议** 对照阅读，见 `note.md` 末节前向链接。
