2 泛型语法（Go 版）.md
 
Go 泛型语法
 
一、Go 泛型概述
 
Go 1.18+ 正式支持泛型，核心是通过类型参数实现代码复用，编译时类型安全，无需手动类型断言。
 
核心三要素：
 
- 类型参数： [T Constraint]  占位符类型
- 类型约束：限制类型参数的范围
- 泛型函数 / 泛型类型：支持多种数据类型
 
 
 
二、基础语法：泛型函数
 
1. 语法格式
 
go  
// 定义
func 函数名[类型参数 约束](参数) 返回值 {
    函数体
}

// 调用（自动推导类型，推荐）
函数名(参数)

// 显式指定类型
函数名[具体类型](参数)
 
 
2. 简单示例
 
go  
package main

import "golang.org/x/exp/constraints"

// T 支持可比较大小的类型：int/float64/string
func Max[T constraints.Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}

func main() {
    println(Max(10, 20))      // int
    println(Max(3.14, 2.71))  // float64
    println(Max("go", "java"))// string
}
 
 
 
 
三、类型约束
 
1. any（任意类型）
 
go  
func Print[T any](value T) {
    fmt.Println(value)
}
 
 
2. comparable（可比较类型：== / !=）
 
map 键必须使用该约束
 
go  
// 切片去重（仅支持可比较类型）
func Deduplicate[T comparable](slice []T) []T {
    m := make(map[T]bool)
    res := make([]T, 0, len(slice))
    for _, v := range slice {
        if !m[v] {
            m[v] = true
            res = append(res, v)
        }
    }
    return res
}
 
 
3. 自定义联合类型
 
go  
// 只允许 int / int64 / float64
type Number interface {
    int | int64 | float64
}

func Sum[T Number](a, b T) T {
    return a + b
}
 
 
4. 方法接口约束
 
go  
type Stringer interface {
    String() string
}

// T 必须实现 String() 方法
func Show[T Stringer](t T) {
    fmt.Println(t.String())
}
 
 
 
 
四、泛型类型（结构体 / 切片 / Map）
 
1. 泛型结构体
 
go  
// 泛型容器
type Box[T any] struct {
    Content T
}

func main() {
    intBox := Box[int]{Content: 100}
    strBox := Box[string]{Content: "hello go"}
}
 
 
2. 泛型切片
 
go  
func PrintSlice[T any](s []T) {
    for _, v := range s {
        fmt.Println(v)
    }
}
 
 
3. 泛型 Map
 
go  
// K：map 键必须 comparable
// V：值任意类型
type MyMap[K comparable, V any] struct {
    data map[K]V
}
 
 
 
 
五、泛型方法
 
方法必须依附泛型类型，不能单独定义
 
go  
type Container[T any] struct {
    value T
}

// 泛型方法
func (c *Container[T]) Set(v T) {
    c.value = v
}

func (c Container[T]) Get() T {
    return c.value
}
 
 
 
 
六、常用标准库约束
 
 import "golang.org/x/exp/constraints" 
 
约束 说明 
any 任意类型 
comparable 可比较（== / !=） 
Ordered 可排序（int/float/string） 
Integer 所有整数 
Float 所有浮点数 
 
 
 
七、Go 泛型重要规则
 
1. Go 版本 ≥ 1.18
2. 仅编译期生效，运行时无泛型类型信息
3. map 键必须使用  comparable  约束
4. 泛型方法的类型参数必须与结构体一致
5. 不支持泛型数组创建： [5]T  非法
 
 
 
八、完整示例：通用栈结构
 
go  
package main

import "fmt"

// 泛型栈
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
    if len(s.items) == 0 {
        var zero T
        return zero, false
    }
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item, true
}

func main() {
    // int 栈
    intStack := Stack[int]{}
    intStack.Push(10)
    intStack.Push(20)

    // string 栈
    strStack := Stack[string]{}
    strStack.Push("go")
    strStack.Push("泛型")
}
 
 
 
 
总结
 
1. 语法： 函数/类型 [T 约束] 
2. 核心：类型参数 + 类型约束
3. 常用约束： any 、 comparable 、 constraints.Ordered 
4. 支持：泛型函数、泛型结构体、泛型方法、泛型集合