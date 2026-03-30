# `reflect.TypeOf` 与 `reflect.ValueOf`

学反射时最容易懵的一点：**这俩到底「拿到了啥」？**

可以这么记：

- **`TypeOf`**：偏 **「看类型」**（多大、啥 Kind、几个字段……）。
- **`ValueOf`**：偏 **「动值」**（读、写、调用方法……当然要满足可寻址、可导出等条件）。

底层套路是一样的：传进来的 `x` 会先变成 **空接口 `eface`**，再分别从 **`_type`** 和 **`data`** 包装出去。

## 1. 对照表（先背这张就够应付一半面试）

| 函数 | 返回 | 大白话 |
|------|------|--------|
| `reflect.TypeOf(x)` | `reflect.Type`（**接口**） | 类型的「只读说明书」，底层连着 `runtime._type` 那套 |
| `reflect.ValueOf(x)` | `reflect.Value`（**结构体值**） | 一个「句柄」，里面有类型指针 + 数据指针 + 标记位 |

## 2. `TypeOf`：类型身份证

`reflect.Type` 是接口，实现里常见 `*rtype`，你可以理解成 **对 `_type` 的封装**：

```go
type Type interface { /* Kind, Name, Size, ... */ }
type rtype _type // 与 runtime 布局对齐的别名
```

流程：`x` → 变 `eface` → 拿出 `_type` → 当作 `Type` 给你。

## 3. `ValueOf`：能摸内存的句柄（但有规矩）

```go
type Value struct {
	typ  *rtype
	ptr  unsafe.Pointer
	flag uintptr
}
```

注意：**`ValueOf` 返回的是 `Value` 结构体本身，不是指针**。  
想 **改原变量**，通常要 **`ValueOf(&x).Elem()`** 这一套（详见 [04-反射CanSet与修改值.md](./04-反射CanSet与修改值.md)）。

```go
x := 100
v1 := reflect.ValueOf(x)
// v1.SetInt(200) // 往往不行：不可设置

v2 := reflect.ValueOf(&x).Elem()
v2.SetInt(200) // x 变成 200
```

## 4. 内存流向（文字版流程图）

```text
x → 先变成 interface{}（eface：_type + data）
      ↓
TypeOf(x)  → 主要吃 _type 这一半
ValueOf(x) → _type + data 都要，再塞进 Value
```

## 5. 和 `runtime` 啥关系？

- **类型元数据**在 `runtime`（`_type` / `rtype`）。
- **`reflect.Value`** 是 **`reflect` 包自己定义的结构体**，不是 `runtime` 里叫 `Value` 的类型，但里面 **指着** runtime 的类型/数据指针。

---

## 复习速记

| 记什么 | 记一句 |
|--------|--------|
| TypeOf | 看类型 |
| ValueOf | 摸值（读写看条件） |
| 改原变量 | 常要 `ValueOf(&x).Elem()` |

## 延伸阅读

- `_type`：[04-_type 到底是什么.md](../01-接口（Interface）专题/04-_type%20到底是什么.md)
- `Value` 细节：[03-value.md](./03-value.md)
- 反射提纲：[02-Go 语言反射（Reflection）.md](./02-Go%20语言反射（Reflection）.md)
