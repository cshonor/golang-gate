// 示例：Go 语言的整数类型
// 演示整数类型的分类、溢出回绕、类型转换和 fmt.Printf 占位符

package main

import (
	"fmt"
	"math/big"
	"math/bits"
)

func main() {
	// ============================================
	// 1. 整数类型的分类和取值范围
	// ============================================
	fmt.Println("=== 整数类型分类 ===")

	// 有符号整数：能表示正、负、0
	var a int8 = 127   // 1 字节，-128~127
	var b int16 = 32767 // 2 字节，-32768~32767
	var c int32 = 2147483647 // 4 字节，-2^31~2^31-1
	var d int64 = 9223372036854775807 // 8 字节，-2^63~2^63-1

	fmt.Printf("int8:  %T, 值: %d\n", a, a)
	fmt.Printf("int16: %T, 值: %d\n", b, b)
	fmt.Printf("int32: %T, 值: %d\n", c, c)
	fmt.Printf("int64: %T, 值: %d\n", d, d)

	// 无符号整数：只能表示正、0
	var e uint8 = 255   // 1 字节，0~255
	var f uint16 = 65535 // 2 字节，0~65535
	var g uint32 = 4294967295 // 4 字节，0~2^32-1
	var h uint64 = 18446744073709551615 // 8 字节，0~2^64-1

	fmt.Printf("uint8:  %T, 值: %d\n", e, e)
	fmt.Printf("uint16: %T, 值: %d\n", f, f)
	fmt.Printf("uint32: %T, 值: %d\n", g, g)
	fmt.Printf("uint64: %T, 值: %d\n", h, h)

	// int 和 uint：位数跟操作系统一致（32位系统32位，64位系统64位）
	var i int = 100   // 不确定位数，根据系统而定
	var j uint = 200  // 不确定位数，根据系统而定
	fmt.Printf("int:  %T, 值: %d\n", i, i)
	fmt.Printf("uint: %T, 值: %d\n", j, j)

	fmt.Println()

	// ============================================
	// 2. 整数溢出和回绕现象
	// ============================================
	fmt.Println("=== 整数溢出和回绕 ===")

	// 有符号整数溢出：会回绕到最小值
	var a1 int8 = 127
	fmt.Printf("int8(127) + 1 = %d (溢出回绕到最小值)\n", a1+1) // 输出 -128

	var a2 int8 = -128
	fmt.Printf("int8(-128) - 1 = %d (溢出回绕到最大值)\n", a2-1) // 输出 127

	// 无符号整数溢出：会回绕到最小值（0）
	var b1 uint8 = 255
	fmt.Printf("uint8(255) + 1 = %d (溢出回绕到最小值0)\n", b1+1) // 输出 0

	var b2 uint8 = 0
	fmt.Printf("uint8(0) - 1 = %d (溢出回绕到最大值255)\n", b2-1) // 输出 255

	// int64 溢出示例（假设是64位系统，int就是int64）
	var c1 int = 9223372036854775807 // int64最大值
	fmt.Printf("int最大值: %d\n", c1)
	fmt.Printf("int最大值 + 1 = %d (溢出回绕到最小值)\n", c1+1)

	fmt.Println()

	// ============================================
	// 3. 类型转换
	// ============================================
	fmt.Println("=== 类型转换 ===")

	// Go 中不同类型不能直接运算，必须先转换
	var x int32 = 100
	var y int64 = 200
	// var z = x + y  // 错误：不能直接运算
	var z = int64(x) + y // 正确：先转换再运算
	fmt.Printf("int32(100) + int64(200) = %d\n", z)

	// 无符号和有符号之间的转换
	var u uint8 = 255
	var s int8 = int8(u) // 转换时要注意溢出
	fmt.Printf("uint8(255) 转 int8 = %d (注意：可能溢出)\n", s)

	fmt.Println()

	// ============================================
	// 4. 使用场景示例
	// ============================================
	fmt.Println("=== 使用场景示例 ===")

	// 普通计数/循环：用 int
	for i := 0; i < 5; i++ {
		fmt.Printf("循环计数: %d\n", i)
	}

	// 表示字节/二进制数据：用 uint8
	var pixel uint8 = 255 // RGB 像素值
	fmt.Printf("像素值: %d\n", pixel)

	// 时间戳/大数值：用 int64
	var timestamp int64 = 1609459200 // Unix 时间戳
	fmt.Printf("时间戳: %d\n", timestamp)

	// 位运算/状态标记：用无符号整数
	var status uint32 = 0b1010 // 用不同位表示不同状态
	fmt.Printf("状态标记: %b (二进制)\n", status)

	fmt.Println()

	// ============================================
	// 5. 溢出检测示例
	// ============================================
	fmt.Println("=== 溢出检测示例 ===")

	// 注意：Go 的整数运算不会自动报错，溢出会直接回绕
	// 实际开发中，如果可能出现大数值，应该：
	// 1. 提前选对类型（比如用 int64 代替 int）
	// 2. 手动检查数值范围
	// 3. 超大数用 big.Int

	maxInt64 := int64(9223372036854775807)
	fmt.Printf("int64 最大值: %d\n", maxInt64)
	fmt.Printf("int64 最大值 + 1 = %d (溢出回绕)\n", maxInt64+1)

	// 使用 math/bits 包可以检测无符号整数的溢出
	u1 := uint64(18446744073709551615) // uint64最大值
	u2 := uint64(1)
	sum, carryOut := bits.Add64(u1, u2, 0)
	if carryOut > 0 {
		fmt.Printf("检测到无符号整数溢出！%d + %d 会溢出\n", u1, u2)
		fmt.Printf("结果（回绕后）: %d\n", sum)
	}

	fmt.Println()

	// ============================================
	// 6. 超大整数：big.Int（对应 Java 的 BigInteger）
	// ============================================
	fmt.Println("=== 超大整数 big.Int ===")

	// 超过 int64 范围的大整数
	bigA := big.NewInt(9223372036854775807) // int64最大值
	bigB := big.NewInt(1)
	bigC := new(big.Int).Add(bigA, bigB) // 大整数相加
	fmt.Printf("big.Int 相加: %s (不会溢出)\n", bigC.String())

	fmt.Println()

	// ============================================
	// 7. fmt.Printf 占位符：%T（类型占位符）
	// ============================================
	fmt.Println("=== fmt.Printf 占位符：%T ===")

	var num1 int = 10
	var num2 uint8 = 255
	var num3 float64 = 3.14
	var num4 = []int{1, 2, 3} // 切片
	type MyType string
	var num5 MyType = "test"

	fmt.Printf("num1 的类型: %T\n", num1)   // 输出: int
	fmt.Printf("num2 的类型: %T\n", num2)   // 输出: uint8
	fmt.Printf("num3 的类型: %T\n", num3)   // 输出: float64
	fmt.Printf("num4 的类型: %T\n", num4)   // 输出: []int
	fmt.Printf("num5 的类型: %T\n", num5)   // 输出: main.MyType

	fmt.Println()

	// ============================================
	// 8. fmt.Printf 占位符：%c（字符占位符）
	// ============================================
	fmt.Println("=== fmt.Printf 占位符：%c ===")

	var char1 rune = 'A'
	var char2 rune = '中'
	var ascii int = 97 // 小写 a 的 ASCII 码

	fmt.Printf("%c\n", char1)  // 输出: A
	fmt.Printf("%c\n", char2)  // 输出: 中
	fmt.Printf("%c\n", ascii)  // 输出: a（ASCII码转字符）

	fmt.Println()

	// ============================================
	// 9. fmt.Printf 占位符：%v（通用占位符）
	// ============================================
	fmt.Println("=== fmt.Printf 占位符：%v ===")

	var v1 int = 123
	var v2 string = "hello"
	var v3 = []int{1, 2, 3}
	type Person struct {
		Name string
		Age  int
	}
	var v4 = Person{"Tom", 20}

	fmt.Printf("%v\n", v1)  // 输出: 123
	fmt.Printf("%v\n", v2)  // 输出: hello
	fmt.Printf("%v\n", v3)  // 输出: [1 2 3]
	fmt.Printf("%v\n", v4)  // 输出: {Tom 20}

	// %+v：输出更详细的信息（结构体带字段名）
	fmt.Printf("%+v\n", v4) // 输出: {Name:Tom Age:20}

	// %#v：输出带引号和类型信息
	fmt.Printf("%#v\n", v2) // 输出: "hello"
	fmt.Printf("%#v\n", v4) // 输出: main.Person{Name:"Tom", Age:20}

	fmt.Println()

	// ============================================
	// 10. fmt.Printf 占位符：%[n]v（位置标记）
	// ============================================
	fmt.Println("=== fmt.Printf 占位符：%[n]v ===")

	// %[1]v 表示使用第1个变量，%[2]v 表示使用第2个变量
	fmt.Printf("%[1]v %[1]v\n", 10)              // 输出: 10 10（两个占位符都用第1个变量）
	fmt.Printf("%[2]v - %[1]v\n", "Go", "语言")    // 输出: 语言 - Go（第2个变量在前）

	// 混合使用位置标记和顺序占位符
	fmt.Printf("%[1]T %[1]v %v\n", 123, "hello") // 输出: int 123 hello
	// 第1个占位符 %[1]T 对应第1个变量（显示类型）
	// 第2个占位符 %[1]v 对应第1个变量（显示值）
	// 第3个占位符 %v 对应第2个变量（显示值）

	fmt.Println()

	// ============================================
	// 11. 占位符顺序对应关系
	// ============================================
	fmt.Println("=== 占位符顺序对应关系 ===")

	// 占位符和变量按顺序一一对应
	a_var := 123       // 第1个变量：int类型
	b_var := "hello"   // 第2个变量：string类型
	c_var := 3.14      // 第3个变量：float64类型

	// 三个占位符按顺序对应三个变量
	fmt.Printf("%T %v %v\n", a_var, b_var, c_var)
	// 输出: int hello 3.14
	// 第1个 %T 对应 a_var（显示类型）
	// 第1个 %v 对应 b_var（显示值）
	// 第2个 %v 对应 c_var（显示值）

	fmt.Println()

	// ============================================
	// 12. 无符号整数的陷阱
	// ============================================
	fmt.Println("=== 无符号整数的陷阱 ===")

	// 注意：如果用 uint 做倒序循环，会无限循环
	// for i := uint(0); i >= 0; i-- { ... }  // 错误：uint 不会小于 0，会无限循环

	// 正确的倒序循环方式
	fmt.Println("正确的倒序循环（用 int）：")
	for i := 5; i >= 0; i-- {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// 如果要用 uint 倒序，需要特殊处理
	fmt.Println("用 uint 倒序（需要特殊处理）：")
	for i := uint(5); i > 0; i-- {
		fmt.Printf("%d ", i)
	}
	fmt.Printf("0\n") // 手动输出 0

	fmt.Println()

	// ============================================
	// 13. 总结：Go vs Java 类型对应关系
	// ============================================
	fmt.Println("=== Go vs Java 类型对应 ===")
	fmt.Println("Java long  → Go int64")
	fmt.Println("Java int   → Go int32")
	fmt.Println("Java short → Go int16")
	fmt.Println("Java byte  → Go int8")
	fmt.Println("Java BigInteger → Go math/big.Int")
}

