# Go 语言反射（Reflection）

**反射**听起来玄，其实可以先用一句话：**程序跑起来之后，还能「看」变量是什么类型、「动」它的值**——编译时你甚至可以完全不知道具体类型。

当然，能力强，代价也大（慢、容易踩坑），所以：**业务热路径少用，框架里很常见**（JSON 之类）。

## 1. 两个入口：先记死名字

| 入口 | 干啥 |
|------|------|
| `reflect.TypeOf(x)` | 看类型：`Kind`、字段、方法…… |
| `reflect.ValueOf(x)` | 看/改值、调方法（受 `CanSet`、可导出等限制） |

## 2. 最小例子：一套代码「瞟」多种类型

```go
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func inspect(x any) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	fmt.Printf("类型：%v，Kind：%v\n", t, t.Kind())
	if t.Kind() == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			fmt.Printf("  %s: %v\n", t.Field(i).Name, v.Field(i).Interface())
		}
	}
}

func main() {
	inspect(100)
	inspect(User{Name: "张三", Age: 20})
}
```

**人话**：同一段 `inspect`，`int` 也能进，`struct` 也能进——这就是反射「通用」的地方。

## 3. 反射一般能干啥？

1. **查类型**：`Kind`、`NumField`、`Field`、`Method`……
2. **读写值**：`SetInt`、`SetString`、`Field`……（注意指针、`Elem`、可导出）
3. **动态调用**：`Method(i).Call(...)`
4. **造值**：`reflect.New`、`MakeSlice`、`MakeMap` 等

## 4. 底层别慌：还是 `eface` 那一套

```go
type eface struct {
	_type *_type
	data  unsafe.Pointer
}
```

`TypeOf` / `ValueOf` 都是先把 `x` 变成 `eface`，再分别从 **`_type`**、**`data`** 做文章——所以前面 **`_type`、空接口** 学顺了，反射就不玄学。

## 5. 为啥慢？（知道方向就行）

1. **运行时**才查表、解析名字，路走得长。
2. 编译器 **不好帮你内联、专项优化**。
3. `Value`、`Type` 等对象 **分配多**，GC 压力大。
4. **指针跳来跳去**，cache 不友好。

**量级**：简单读写慢一个数量级以上很常见，具体以你自己的 `benchmark` 为准（详见 [05-反射性能.md](./05-反射性能.md)）。

## 6. 使用原则（新手版）

- **能写具体类型，就别反射**；能 **泛型** 解决的，优先泛型。
- **JSON / ORM / 配置映射** 这类库，反射很常见，属于「框架替你付代价」。
- **坑**：改值要指针 + `Elem`；小写字段包外改不了；类型不对会 `panic`，要先 `Kind` / `ok`。

## 7. 和类型断言、泛型放一起看

| 特性 | 反射 | 类型断言 | 泛型 |
|------|------|----------|------|
| 时机 | 运行时 | 运行时 | 编译期 |
| 性能 | 一般较差 | 好 | 好 |
| 类型安全 | 弱（多运行时检查） | 中 | 强 |

---

## 复习速记

| 记一句 | 内容 |
|--------|------|
| 入口 | `TypeOf` / `ValueOf` |
| 底层 | 先把 `x` 变 `eface` |
| 原则 | 框架多用，业务热点慎用 |

## 延伸阅读

- `Value`：[03-value.md](./03-value.md)
- 改值：[04-反射CanSet与修改值.md](./04-反射CanSet与修改值.md)
- 性能：[05-反射性能.md](./05-反射性能.md)
