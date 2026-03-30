# 03 - 接口 iface 与 eface 底层结构

前置：[01-Go 接口（Interface）](./01-Go%20接口（Interface）.md)、[02-空接口eface与nil](./02-空接口eface与nil.md)。本篇把 **iface / eface** 与 runtime 布局对齐到直觉层，**细节以当前 src/runtime/iface.go 为准**。

---

## 1. 两种接口表示

- **eface（empty interface）**：`interface{}` / `any`。动态类型 + 动态值（常见为 `_type` 与 `data`）。
- **iface（非空接口）**：如 `io.Reader`。除类型与值外，通过 **itab** 做方法分派。

---

## 2. 与反射的衔接

理解二者有助于读 `reflect.TypeOf` / `ValueOf` 与 `Elem()`。关联：[../03-接口与反射关联/01-接口与反射的关系.md](../03-接口与反射关联/01-接口与反射的关系.md)。

---

## 延伸阅读

- [04-_type 到底是什么](./04-_type%20到底是什么.md)
- [../README.md](../README.md)
