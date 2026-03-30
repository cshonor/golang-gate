# `_type` 到底是什么（runtime 里的「类型身份证」）

如果你刚看完 `eface`，一定会问：**`eface` 里那个 `_type` 指针，到底指向啥？**

可以把它想成：**Go 在运行时给每种类型准备的一份「说明书」**——大小、布局、能不能做相等比较、GC 要扫哪里……都在这类元数据里。  
下面用**能看懂**为主，不追求把每个字段背下来。

## 1. 一句话

`_type`（以及围绕它扩展出来的一族结构）是 **runtime 里描述类型的根元数据**；空接口 `eface._type` 就指向「当前这份值」对应的说明书。

## 2. 长什么样（简化版，够用就行）

不同 Go 版本字段略有出入，**理解用**可以看个简化版：

```go
type _type struct {
	size       uintptr
	ptrdata    uintptr
	hash       uint32
	tflag      tflag
	align      uint8
	fieldAlign uint8
	kind       uint8
	equal      func(unsafe.Pointer, unsafe.Pointer) bool
	// 还有名字、指针类型偏移等，先不展开
}
```

| 字段（直觉） | 人话 |
|--------------|------|
| `size` | 这个类型占多少字节 |
| `ptrdata` | 里面有多少「带指针」的区域，GC 要扫 |
| `kind` | 大类：int、struct、ptr 等（和 `reflect.Kind` 对应） |
| `equal` 等 | map 键、比较运算等会用到 |

## 3. 为啥需要它？（和反射、GC 的关系）

1. **空接口知道「我装的是谁」**：靠 `_type` 指向的那份类型元数据。
2. **反射 `TypeOf`**：读的就是这类信息（再经 `reflect` 包包装成好用的 API）。
3. **GC**：要知道一块内存里哪些字段是指针，离不开类型布局。
4. **map 键、比较**：也要靠类型里挂的函数指针等。

## 4. 和 `eface` / `iface` 放一起记

```go
type eface struct {
	_type *_type
	data  unsafe.Pointer
}

type iface struct {
	tab  *itab
	data unsafe.Pointer
}
```

非空接口那边用 `itab` 把「动态类型 + 方法表」绑在一起，**里面也会关联到类型元数据**——你先把 **「类型信息总离不开 _type 这一套」** 记住就行。

## 5. 再和「nil 坑」串一下

```go
var a any
fmt.Println(a == nil) // true：相当于「没类型没数据」

var p *int = nil
var b any = p
fmt.Println(b == nil) // false：类型信息已经说明是 *int 了，只是 data 可以是 nil
```

**人话**：第二种情况，**接口这个盒子已经不是「空盒子」了**，只是盒子里装了一个「nil 指针」。

---

## 复习速记

| 考点 | 一句话 |
|------|--------|
| `_type` | 类型的运行时元数据（身份证） |
| 和 `eface` | `eface._type` 指向它 |
| 别死背字段 | 知道「干啥用的」更重要 |

## 延伸阅读

- 空接口：[02-空接口eface与nil.md](./02-空接口eface与nil.md)
- 反射入口：[05-reflect.TypeOf  and reflect.ValueOf.md](./05-reflect.TypeOf%C2%A0%20and%20reflect.ValueOf.md)
