# `reflect.Value` 核心逻辑

`reflect.Value` 是 **`reflect` 包中的结构体**（值语义返回），表示对某个值的**访问通道**：内部持有类型信息与数据指针（及 `flag`），**不拷贝用户数据本身**。

## 1. 结构示意

```go
type Value struct {
	typ  *rtype
	ptr  unsafe.Pointer
	flag uintptr
}
```

| 字段 | 含义（直觉） |
|------|----------------|
| `typ` | 与 `TypeOf` 同源的类型元数据 |
| `ptr` | 指向数据的指针（是否可改依赖是否可寻址等） |
| `flag` | 可寻址、种类、是否 nil 等位域，影响 `CanSet` / `Set` |

## 2. 传值 vs 传指针

```go
type User struct {
	Name string
	Age  int
}

func main() {
	u := User{Name: "张三", Age: 20}

	v1 := reflect.ValueOf(u)              // 副本路径，改字段常失败或需可导出+可设置条件
	v2 := reflect.ValueOf(&u).Elem()      // 指向原 u

	v2.FieldByName("Name").SetString("李四")
}
```

- **`ValueOf(值)`**：`ptr` 往往指向**临时副本**上的数据，修改不一定回写外层变量。
- **`ValueOf(&x).Elem()`**：得到代表 `x` 本身的可寻址 `Value`，字段 `Set*` 在条件满足时可改 `x`。

## 3. 与 `Type`、`eface` 的关系

```go
type eface struct {
	_type *_type
	data  unsafe.Pointer
}
```

`ValueOf(x)` / `TypeOf(x)` 都把 `x` 先变为 `eface`：  
`Type` 主要来自 `_type`；`Value.ptr` 主要来自 `data`（再结合 `flag`）。

## 4. 一句话

**`Value` = 运行时对「某块内存 + 类型」的操作句柄；想改原变量，优先走指针 + `Elem()`，并检查 `CanSet()`。**

---

## 复习速记

| 记什么 | 记一句 |
|--------|--------|
| 是什么 | `reflect` 里的值句柄，不是 `runtime` 类型名 |
| 改原值 | `ValueOf(&x).Elem()` + 可导出字段 |
| 与 Type | 同源 `typ`，多一个数据指针 |

## 延伸阅读

- `CanSet`：[08-反射CanSet与修改值.md](./08-反射CanSet与修改值.md)
- 入口对比：[05-reflect.TypeOf  and reflect.ValueOf.md](./05-reflect.TypeOf%C2%A0%20and%20reflect.ValueOf.md)
