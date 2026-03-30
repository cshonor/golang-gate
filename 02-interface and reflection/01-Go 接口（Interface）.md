# Go 接口（Interface）

新手可以先把接口想成一张 **「技能清单 / 契约」**：

- 清单里只写 **要会做什么**（方法名 + 参数返回值），**不写具体怎么做**。
- 你的类型只要 **真的实现了清单上的全部方法**，就 **自动算** 实现了这个接口——**不用**像 Java 那样写 `implements`（这是 Go 很「省心」的地方）。

下面按「是什么 → 有啥用 → 和空接口/断言啥关系」的顺序看，**不用一次全背**。

## 1. 先建立直觉：多态从「一个函数」开始

没有接口时，往往要给每种动物写一个函数，重复多；有了接口，可以把「会叫」抽成 **同一个参数类型**：

```go
type Speaker interface {
	Speak() string
}

type Dog struct{ Name string }
func (d Dog) Speak() string { return "汪汪" }

type Cat struct{ Name string }
func (c Cat) Speak() string { return "喵喵" }

func LetSpeak(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	LetSpeak(Dog{Name: "阿黄"})
	LetSpeak(Cat{Name: "咪咪"})
}
```

**一句话**：同一个 `LetSpeak`，传进去的具体类型不同，执行的 `Speak` 不同——这就是 **多态**。

## 2. Go 接口的几个「大白话」特点

**非侵入**：不用在 `type Dog` 旁边声明「我实现了 Speaker」；方法集对上了，就能赋给 `Speaker`。

**空接口 `any` / `interface{}`**：没有任何方法的接口 → **所有类型都算**满足它，所以能装任意值（后面再靠类型断言等「拆包」）。

```go
var a any
a = 100
a = "hello"
a = Dog{}
```

**和底层怎么对上号（先记个名字）**：

- 空接口：运行时往往是 **`eface`**（类型 + 数据两根指针）。
- 带方法的接口：往往是 **`iface`**（多一个方法表相关的东西）。

想细一点就看 [02-空接口eface与nil.md](./02-空接口eface与nil.md)。

## 3. 接口和「类型断言」是啥关系？

把具体类型塞进接口变量以后，**静态类型**是接口，**运行时**才带着「真实类型」信息。要**还原成 Dog、Cat** 或试试能不能转成别的接口，就用 **类型断言**（详见 [04-类型断言.md](./04-类型断言.md)）。

```go
var s Speaker = Dog{Name: "旺财"}
var a any = s

dog, ok := a.(Dog)
if ok {
	fmt.Println(dog.Speak())
}
```

## 4. 平时在哪用？

- **标准库**：比如 `error` 就是 `Error() string` 的小接口；`fmt.Stringer` 之类。
- **你自己的代码**：参数写成接口，方便换实现、写测试（mock）。
- **泛型之前**：用 `any` + 断言/反射做通用工具（现在有泛型，很多场景可以更简单）。

## 5. 纠正一个常见说法

- 不是「接口实现了方法」，而是 **具体类型（Dog、Cat）实现了方法**，从而 **满足** 接口。
- 少实现一个方法，都不能赋给那个接口类型。

---

## 复习速记

| 考点 | 一句话 |
|------|--------|
| 接口是什么 | 方法集合（契约） |
| Go 特色 | 隐式实现、非侵入 |
| 空接口 | `any`，底层常对应 `eface` |
| 拆类型 | 类型断言 `.()` |

## 延伸阅读

- `eface` / nil：[02-空接口eface与nil.md](./02-空接口eface与nil.md)
- `_type`：[03-_type  到底是什么.md](./03-_type%C2%A0%20到底是什么.md)
