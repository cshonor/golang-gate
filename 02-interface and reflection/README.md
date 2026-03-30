# 接口与反射（面试向笔记）

本目录按 **专题子文件夹** 组织：**接口**、**反射**、**二者关联**。统一使用 **`#` / `##`** 层级与代码围栏；文末常附 **复习速记** / **延伸阅读**。

**GMP / Channel** 见 **`../07-GMP and channel/`**；**map / slice** 见 **`../01-datastruct/`**；**GC** 见 **`../11-GC and memory/`**。

---

## 目录结构（总览）

```
02-interface and reflection/
├── README.md                        # 本文件
├── 01-接口（Interface）专题/
│   ├── 01-Go 接口（Interface）.md
│   ├── 02-空接口eface与nil.md
│   ├── 03-接口iface与eface底层结构.md
│   ├── 04-_type 到底是什么.md
│   ├── 05-类型断言.md
│   ├── 06-接口常见坑与最佳实践.md
│   ├── 07-接口与多态实战.md
│   └── 08-接口汇编分析.md           # 可选
├── 02-反射（Reflection）专题/
│   ├── 01-reflect.TypeOf and reflect.ValueOf.md
│   ├── 02-Go 语言反射（Reflection）.md
│   ├── 03-value.md
│   ├── 04-反射CanSet与修改值.md
│   ├── 05-反射性能.md
│   ├── 06-反射常见坑与最佳实践.md
│   ├── 07-反射实战案例.md
│   ├── 08-反射unsafe包结合使用.md   # 可选
│   ├── 09-泛型与反射的区别.md       # 可选
│   ├── 10-面试口述版.md
│   └── 11-面试真题.md
└── 03-接口与反射关联/
    └── 01-接口与反射的关系.md
```

---

## 01 - 接口（Interface）专题

| 序号 | 主题 | 文件 |
|:----:|------|------|
| 01 | Go `interface` 语义鸟瞰 | [01-Go 接口（Interface）.md](./01-接口（Interface）专题/01-Go%20接口（Interface）.md) |
| 02 | 空接口、`eface`、接口 nil 坑 | [02-空接口eface与nil.md](./01-接口（Interface）专题/02-空接口eface与nil.md) |
| 03 | `iface` / `eface` 底层结构 | [03-接口iface与eface底层结构.md](./01-接口（Interface）专题/03-接口iface与eface底层结构.md) |
| 04 | `_type` 与类型元数据直觉 | [04-_type 到底是什么.md](./01-接口（Interface）专题/04-_type%20到底是什么.md) |
| 05 | 类型断言：`.(T)` 与 `ok` | [05-类型断言.md](./01-接口（Interface）专题/05-类型断言.md) |
| 06 | 接口常见坑与最佳实践 | [06-接口常见坑与最佳实践.md](./01-接口（Interface）专题/06-接口常见坑与最佳实践.md) |
| 07 | 接口与多态实战 | [07-接口与多态实战.md](./01-接口（Interface）专题/07-接口与多态实战.md) |
| 08 | 接口汇编分析（可选） | [08-接口汇编分析.md](./01-接口（Interface）专题/08-接口汇编分析.md) |

---

## 02 - 反射（Reflection）专题

| 序号 | 主题 | 文件 |
|:----:|------|------|
| 01 | `reflect.TypeOf` / `reflect.ValueOf` | [01-reflect.TypeOf and reflect.ValueOf.md](./02-反射（Reflection）专题/01-reflect.TypeOf%20and%20reflect.ValueOf.md) |
| 02 | Go 语言反射（提纲） | [02-Go 语言反射（Reflection）.md](./02-反射（Reflection）专题/02-Go%20语言反射（Reflection）.md) |
| 03 | `reflect.Value` 补充 | [03-value.md](./02-反射（Reflection）专题/03-value.md) |
| 04 | `CanSet`、`Elem`、改值 | [04-反射CanSet与修改值.md](./02-反射（Reflection）专题/04-反射CanSet与修改值.md) |
| 05 | 反射性能 | [05-反射性能.md](./02-反射（Reflection）专题/05-反射性能.md) |
| 06 | 反射常见坑与最佳实践 | [06-反射常见坑与最佳实践.md](./02-反射（Reflection）专题/06-反射常见坑与最佳实践.md) |
| 07 | 反射实战案例 | [07-反射实战案例.md](./02-反射（Reflection）专题/07-反射实战案例.md) |
| 08 | 反射与 `unsafe`（可选） | [08-反射unsafe包结合使用.md](./02-反射（Reflection）专题/08-反射unsafe包结合使用.md) |
| 09 | 泛型与反射的区别（可选） | [09-泛型与反射的区别.md](./02-反射（Reflection）专题/09-泛型与反射的区别.md) |
| 10 | 面试口述版 | [10-面试口述版.md](./02-反射（Reflection）专题/10-面试口述版.md) |
| 11 | 面试真题 | [11-面试真题.md](./02-反射（Reflection）专题/11-面试真题.md) |

---

## 03 - 接口与反射关联

| 序号 | 主题 | 文件 |
|:----:|------|------|
| 01 | 接口与反射的关系 | [01-接口与反射的关系.md](./03-接口与反射关联/01-接口与反射的关系.md) |

---

## 学习顺序建议

1. **接口专题**：按 **01 → 08**（**08** 可选）；**03** 建议在 **02** 之后读。  
2. **反射专题**：按 **01 → 09**（**08、09** 可选）；**06、07** 在掌握 **04、05** 后读更顺。  
3. **关联**：读完两侧入口后读 **`03-接口与反射关联/01`**。  
4. **冲刺**：**10 口述 → 11 真题**。

---

## 面试与工程侧重点（摘要）

- **必会**：`eface`/`iface`、接口 nil 与类型断言；`TypeOf`/`ValueOf`、`CanSet`/`Elem`；反射代价。  
- **源码**：`runtime/iface.go`、`reflect/type.go`、`reflect/value.go`（见各篇「进阶阅读」）。

> 说明：原扁平文件名中的 **NBSP** 已在新目录中尽量改为**普通空格**文件名（如 `01-reflect.TypeOf and reflect.ValueOf.md`），避免链接编码困扰。
