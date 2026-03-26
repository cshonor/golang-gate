# 接口与反射（面试向笔记）

本目录整理 **空接口 `eface`、类型断言、`reflect.TypeOf/ValueOf`、反射性能、`CanSet` 与改值** 等高频考点，表述偏「能背、能讲」。

**GMP / Channel** 见 **`../GMP and channel/`**；**map / slice** 见 **`../datastruct/`**；**GC** 见 **`../GC and memory/`**。

## 主题 ↔ 文件索引

| 主题 | 笔记文件 |
|------|----------|
| 空接口底层、`eface`、接口 nil 坑 | [空接口eface与nil.md](./空接口eface与nil.md) |
| 类型断言：`.(T)` 与 `ok` | [类型断言.md](./类型断言.md) |
| `reflect.TypeOf` / `reflect.ValueOf`（入口与底层指向） | [reflect.TypeOf  and reflect.ValueOf.md](./reflect.TypeOf  and reflect.ValueOf.md) |
| 反射为什么慢、何时用 | [反射性能.md](./反射性能.md) |
| `CanSet`、`Elem`、改值 | [反射CanSet与修改值.md](./反射CanSet与修改值.md) |
| 一段话口述 | [面试口述版.md](./面试口述版.md) |
| 练习：2 道真题 | [面试真题.md](./面试真题.md) |

## 学习顺序建议

1. **空接口eface与nil** → **类型断言**（接口语义基础）  
2. **反射性能** → **反射CanSet与修改值**（会用 `reflect` 且不出 panic）  
3. 过 **面试口述版**，再做 **面试真题**

## 进阶阅读（源码）

- `runtime/iface.go`（`eface` / `iface`、`convT` 等）
- `reflect/type.go`（`reflect.Type`、`rtype`）
- `reflect/value.go`（`Value`、`CanSet`、`Elem`）
