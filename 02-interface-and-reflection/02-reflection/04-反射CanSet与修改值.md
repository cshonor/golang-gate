# 反射修改值的限制（`CanSet`）

## 1）核心结论

**默认不能随便改：`CanSet() == true` 才能 `Set*` / `Set`。**

## 2）为什么不能「直接改」？

`reflect.ValueOf(x)` 若 `x` 是值类型，得到的是**一份值的拷贝**，改它**不会**写回原变量；Go 通过 `CanSet` 避免你误以为在改原值。

## 3）什么情况 `CanSet == false`？（常见）

- 传入的是**值**而非**可修改实体的指针**（未通过指针拿到可寻址对象）。
- `Value` **不可寻址**（例如 map 里取出的值、接口里的值副本等，依具体场景而定）。
- 要改**未导出（小写）字段**：包外**不能**用反射改写（语言限制）。

## 4）正确修改可寻址变量的步骤（背）

1. 传**指针**：`reflect.ValueOf(&x)`
2. **解引用**：`v.Elem()` 得到指向的变量
3. 再 `SetInt` / `SetString` / `Set` 等

### 正确示例

```go
var x int = 10
v := reflect.ValueOf(&x)
v.Elem().SetInt(20)
fmt.Println(x) // 20
```

### 错误示例（panic）

```go
v := reflect.ValueOf(x)
v.SetInt(20) // panic: reflect: call of reflect.Value.SetInt on int Value
```

## 5）面试一句话

> 反射改值必须拿到**可寻址、可设置**的 `Value`：通常 `ValueOf(指针)` 再 `Elem()`；否则 `CanSet` 为 false 会 panic；包外不能改未导出字段。

---

## 复习速记

| 步骤 | 写法 |
|------|------|
| 改 `x` | `reflect.ValueOf(&x).Elem()` 再 `Set*` |
| 判断 | `CanSet()` 为 true 才能改 |

## 延伸阅读

- `Value` 细节：[03-value.md](./03-value.md)
- 性能：[05-反射性能.md](./05-反射性能.md)
