# 接口与反射（面试向笔记）

本目录整理 **接口语义、空接口 `eface`、类型断言、`reflect.TypeOf`/`ValueOf`、反射改值与性能、面试口述** 等高频考点；根目录笔记文件名带 **`01-`～`11-`** 前缀。各篇采用统一 **`#` / `##`** 层级、规范代码围栏，文末常附 **复习速记** 与 **延伸阅读**。

**GMP / Channel** 见 **`../GMP and channel/`**；**map / slice** 见 **`../datastruct/`**；**GC** 见 **`../GC and memory/`**。

> **说明**：`05-reflect.TypeOf …` 与 `03-_type …` 文件名中含 **不间断空格（U+00A0）**（历史遗留）；下方链接已用 `%C2%A0` 编码，便于在 GitHub 上跳转。

## 主题 ↔ 文件索引

| 序号 | 主题 | 笔记文件 |
|:----:|------|----------|
| 01 | Go `interface` 语义鸟瞰 | [01-Go 接口（Interface）.md](./01-Go%20接口（Interface）.md) |
| 02 | 空接口底层、`eface`、接口 nil 坑 | [02-空接口eface与nil.md](./02-空接口eface与nil.md) |
| 03 | `_type` 与类型元数据直觉 | [03-_type …](./03-_type%C2%A0%20到底是什么.md) |
| 04 | 类型断言：`.(T)` 与 `ok` | [04-类型断言.md](./04-类型断言.md) |
| 05 | `reflect.TypeOf` / `reflect.ValueOf` | [05-reflect.TypeOf …](./05-reflect.TypeOf%C2%A0%20and%20reflect.ValueOf.md) |
| 06 | Go 语言反射（提纲） | [06-Go 语言反射（Reflection）.md](./06-Go%20语言反射（Reflection）.md) |
| 07 | `reflect.Value` 补充 | [07-value.md](./07-value.md) |
| 08 | `CanSet`、`Elem`、改值 | [08-反射CanSet与修改值.md](./08-反射CanSet与修改值.md) |
| 09 | 反射为什么慢、何时用 | [09-反射性能.md](./09-反射性能.md) |
| 10 | 一段话口述 | [10-面试口述版.md](./10-面试口述版.md) |
| 11 | 练习：真题 | [11-面试真题.md](./11-面试真题.md) |

## 学习顺序建议

1. **01 → 04**：接口与 `eface` → `_type` 直觉 → **类型断言**（语义基础）  
2. **05 → 07**：`TypeOf`/`ValueOf` → 反射提纲 → `Value` 细节  
3. **08 → 09**：`CanSet` 与改值 → **性能**（会用且知道代价）  
4. **10 → 11**：口述 → 真题  

## 进阶阅读（源码）

- `runtime/iface.go`（`eface` / `iface`、`convT` 等）
- `reflect/type.go`（`reflect.Type`、`rtype`）
- `reflect/value.go`（`Value`、`CanSet`、`Elem`）
