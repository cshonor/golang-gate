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
	fmt.Println("\n有符号整数溢出示例:")
	var a1 int8 = 127
	fmt.Printf("int8(127) + 1 = %d (溢出回绕到最小值 -128)\n", a1+1) // 输出 -128

	var a2 int8 = -128
	fmt.Printf("int8(-128) - 1 = %d (溢出回绕到最大值 127)\n", a2-1) // 输出 127

	// int16 溢出示例
	var a3 int16 = 32767
	fmt.Printf("int16(32767) + 1 = %d (溢出回绕到最小值 -32768)\n", a3+1)

	var a4 int16 = -32768
	fmt.Printf("int16(-32768) - 1 = %d (溢出回绕到最大值 32767)\n", a4-1)

	// int32 溢出示例
	var a5 int32 = 2147483647
	fmt.Printf("int32(2147483647) + 1 = %d (溢出回绕到最小值)\n", a5+1)

	// int64 溢出示例
	var a6 int64 = 9223372036854775807
	fmt.Printf("int64(最大值) + 1 = %d (溢出回绕到最小值)\n", a6+1)

	// 无符号整数溢出：会回绕到最小值（0）
	fmt.Println("\n无符号整数溢出示例:")
	var b1 uint8 = 255
	fmt.Printf("uint8(255) + 1 = %d (溢出回绕到最小值 0)\n", b1+1) // 输出 0

	var b2 uint8 = 0
	fmt.Printf("uint8(0) - 1 = %d (溢出回绕到最大值 255)\n", b2-1) // 输出 255

	// uint16 溢出示例
	var b3 uint16 = 65535
	fmt.Printf("uint16(65535) + 1 = %d (溢出回绕到最小值 0)\n", b3+1)

	var b4 uint16 = 0
	fmt.Printf("uint16(0) - 1 = %d (溢出回绕到最大值 65535)\n", b4-1)

	// uint32 溢出示例
	var b5 uint32 = 4294967295
	fmt.Printf("uint32(最大值) + 1 = %d (溢出回绕到最小值 0)\n", b5+1)

	// uint64 溢出示例
	var b6 uint64 = 18446744073709551615
	fmt.Printf("uint64(最大值) + 1 = %d (溢出回绕到最小值 0)\n", b6+1)

	// int 溢出示例（假设是64位系统，int就是int64）
	fmt.Println("\nint 类型溢出示例（64位系统）:")
	var c1 int = 9223372036854775807 // int64最大值
	fmt.Printf("int最大值: %d\n", c1)
	fmt.Printf("int最大值 + 1 = %d (溢出回绕到最小值)\n", c1+1)

	fmt.Println()

	// ============================================
	// 3. 类型转换
	// ============================================
	fmt.Println("=== 类型转换 ===")

	// Go 中不同类型不能直接运算，必须先转换
	fmt.Println("\n3.1 不同大小类型之间的转换:")
	var x int32 = 100
	var y int64 = 200
	// var z = x + y  // 错误：不能直接运算
	var z = int64(x) + y // 正确：先转换再运算
	fmt.Printf("int32(100) + int64(200) = %d\n", z)

	// 大类型转小类型：会截断高位
	var x1 int16 = 1000
	var x2 int8 = int8(x1) // 1000 超出 int8 范围，会截断
	fmt.Printf("int16(1000) 转 int8 = %d (注意：发生截断)\n", x2)

	// 小类型转大类型：安全，不会丢失数据
	var x3 int8 = 100
	var x4 int16 = int16(x3)
	fmt.Printf("int8(100) 转 int16 = %d (安全转换)\n", x4)

	// 无符号和有符号之间的转换
	fmt.Println("\n3.2 有符号和无符号之间的转换:")
	var u uint8 = 255
	var s int8 = int8(u) // 转换时要注意溢出
	fmt.Printf("uint8(255) 转 int8 = %d (注意：可能溢出)\n", s)

	var u2 uint8 = 127
	var s2 int8 = int8(u2) // 127 在 int8 范围内，安全
	fmt.Printf("uint8(127) 转 int8 = %d (安全转换)\n", s2)

	var s3 int8 = -1
	var u3 uint8 = uint8(s3) // -1 转无符号，会变成 255
	fmt.Printf("int8(-1) 转 uint8 = %d (注意：负数会变成大正数)\n", u3)

	// 同大小类型之间的转换
	fmt.Println("\n3.3 同大小类型之间的转换:")
	var i32 int32 = -100
	var u32 uint32 = uint32(i32) // 负数转无符号
	fmt.Printf("int32(-100) 转 uint32 = %d (注意：负数会变成大正数)\n", u32)

	var u32_2 uint32 = 300
	var i32_2 int32 = int32(u32_2) // 300 在 int32 范围内，安全
	fmt.Printf("uint32(300) 转 int32 = %d (安全转换)\n", i32_2)

	// 类型转换时的截断示例
	fmt.Println("\n3.4 类型转换时的截断示例:")
	var large int32 = 0x12345678 // 大数值
	var small int8 = int8(large) // 只保留低8位
	fmt.Printf("int32(0x12345678) 转 int8 = %d (0x%02x，只保留低8位)\n", small, small)

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

	fmt.Println("\n5.1 溢出回绕示例:")
	maxInt64 := int64(9223372036854775807)
	fmt.Printf("int64 最大值: %d\n", maxInt64)
	fmt.Printf("int64 最大值 + 1 = %d (溢出回绕)\n", maxInt64+1)

	maxInt32 := int32(2147483647)
	fmt.Printf("int32 最大值: %d\n", maxInt32)
	fmt.Printf("int32 最大值 + 1 = %d (溢出回绕)\n", maxInt32+1)

	// 使用 math/bits 包可以检测无符号整数的溢出
	fmt.Println("\n5.2 使用 math/bits 检测无符号整数溢出:")
	bitsU1 := uint64(18446744073709551615) // uint64最大值
	bitsU2 := uint64(1)
	sum, carryOut := bits.Add64(bitsU1, bitsU2, 0)
	if carryOut > 0 {
		fmt.Printf("检测到无符号整数溢出！%d + %d 会溢出\n", bitsU1, bitsU2)
		fmt.Printf("结果（回绕后）: %d\n", sum)
	}

	// 检测 uint32 溢出
	bitsU3 := uint32(4294967295) // uint32最大值
	bitsU4 := uint32(1)
	sum32, carryOut32 := bits.Add32(bitsU3, bitsU4, 0)
	if carryOut32 > 0 {
		fmt.Printf("检测到 uint32 溢出！%d + %d 会溢出\n", bitsU3, bitsU4)
		fmt.Printf("结果（回绕后）: %d\n", sum32)
	}

	// 检测 uint16 溢出（使用 Add32，因为 Go 标准库没有 Add16）
	bitsU5 := uint32(65535) // uint16最大值，用 uint32 表示
	bitsU6 := uint32(1)
	sum16, carryOut16 := bits.Add32(bitsU5, bitsU6, 0)
	if carryOut16 > 0 {
		fmt.Printf("检测到 uint16 溢出！%d + %d 会溢出\n", bitsU5, bitsU6)
		fmt.Printf("结果（回绕后）: %d\n", sum16)
	} else {
		// 手动检查是否超出 uint16 范围
		if sum16 > 65535 {
			fmt.Printf("uint16(65535) + uint16(1) = %d (超出 uint16 范围，回绕后为 %d)\n", sum16, uint16(sum16))
		}
	}

	// 手动检查有符号整数溢出
	fmt.Println("\n5.3 手动检查有符号整数溢出:")
	checkInt8Overflow := func(a, b int8) (int8, bool) {
		if a > 0 && b > 0 {
			if a > 127-b {
				return 0, true // 会溢出
			}
		}
		if a < 0 && b < 0 {
			if a < -128-b {
				return 0, true // 会溢出
			}
		}
		return a + b, false
	}

	result, overflow := checkInt8Overflow(100, 50)
	if overflow {
		fmt.Printf("int8(100) + int8(50) 会溢出\n")
	} else {
		fmt.Printf("int8(100) + int8(50) = %d (安全)\n", result)
	}

	result2, overflow2 := checkInt8Overflow(100, 50)
	if overflow2 {
		fmt.Printf("int8(100) + int8(50) 会溢出\n")
	} else {
		fmt.Printf("int8(100) + int8(50) = %d (安全)\n", result2)
	}

	result3, overflow3 := checkInt8Overflow(127, 1)
	if overflow3 {
		fmt.Printf("int8(127) + int8(1) 会溢出！\n")
	} else {
		fmt.Printf("int8(127) + int8(1) = %d\n", result3)
	}

	fmt.Println()

	// ============================================
	// 6. 超大整数：big.Int（对应 Java 的 BigInteger）
	// ============================================
	fmt.Println("=== 超大整数 big.Int ===")

	// 超过 int64 范围的大整数
	fmt.Println("\n6.1 基本运算:")
	bigA := big.NewInt(9223372036854775807) // int64最大值
	bigB := big.NewInt(1)
	bigC := new(big.Int).Add(bigA, bigB) // 大整数相加
	fmt.Printf("big.Int 相加: %s + %s = %s (不会溢出)\n", bigA.String(), bigB.String(), bigC.String())

	// 减法
	bigD := big.NewInt(100)
	bigE := big.NewInt(50)
	bigF := new(big.Int).Sub(bigD, bigE)
	fmt.Printf("big.Int 相减: %s - %s = %s\n", bigD.String(), bigE.String(), bigF.String())

	// 乘法
	bigG := big.NewInt(1000000)
	bigH := big.NewInt(2000000)
	bigI := new(big.Int).Mul(bigG, bigH)
	fmt.Printf("big.Int 相乘: %s * %s = %s\n", bigG.String(), bigH.String(), bigI.String())

	// 除法
	bigJ := big.NewInt(100)
	bigK := big.NewInt(3)
	bigL := new(big.Int).Div(bigJ, bigK)
	fmt.Printf("big.Int 相除: %s / %s = %s (整数除法)\n", bigJ.String(), bigK.String(), bigL.String())

	// 从字符串创建大整数
	fmt.Println("\n6.2 从字符串创建大整数:")
	bigStr := "123456789012345678901234567890"
	bigM, _ := new(big.Int).SetString(bigStr, 10)
	fmt.Printf("从字符串创建: %s\n", bigM.String())

	// 超大数运算
	bigN := big.NewInt(1)
	for i := 0; i < 100; i++ {
		bigN.Mul(bigN, big.NewInt(2)) // 计算 2^100
	}
	fmt.Printf("2^100 = %s\n", bigN.String())

	// 比较
	fmt.Println("\n6.3 大整数比较:")
	bigO := big.NewInt(100)
	bigP := big.NewInt(200)
	fmt.Printf("big.Int 比较: %s 和 %s\n", bigO.String(), bigP.String())
	fmt.Printf("  %s < %s: %t\n", bigO.String(), bigP.String(), bigO.Cmp(bigP) < 0)
	fmt.Printf("  %s > %s: %t\n", bigO.String(), bigP.String(), bigO.Cmp(bigP) > 0)
	fmt.Printf("  %s == %s: %t\n", bigO.String(), bigP.String(), bigO.Cmp(bigP) == 0)

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

	// 陷阱1：倒序循环会无限循环
	fmt.Println("\n陷阱1：倒序循环会无限循环")
	fmt.Println("错误示例（已注释）:")
	fmt.Println("  for i := uint(0); i >= 0; i-- { ... }")
	fmt.Println("  // 错误：uint 不会小于 0，会无限循环")

	// 正确的倒序循环方式
	fmt.Println("\n正确的倒序循环（用 int）：")
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

	// 陷阱2：无符号整数减法可能产生大数
	fmt.Println("\n陷阱2：无符号整数减法可能产生大数")
	var trapU1 uint8 = 5
	var trapU2 uint8 = 10
	trapResult := trapU1 - trapU2 // 5 - 10 = 251 (溢出回绕)
	fmt.Printf("uint8(5) - uint8(10) = %d (注意：不是 -5！)\n", trapResult)
	fmt.Println("  说明：无符号整数不能表示负数，减法溢出会回绕")

	// 陷阱3：无符号整数比较
	fmt.Println("\n陷阱3：无符号整数比较")
	var trapU3 uint8 = 0
	var trapU4 int8 = -1
	fmt.Printf("uint8(0) > int8(-1)? 不能直接比较，类型不同\n")
	fmt.Printf("但 uint8(0) 转 int8 = %d\n", int8(trapU3))
	fmt.Printf("int8(-1) 转 uint8 = %d (负数变成大正数)\n", uint8(trapU4))

	// 陷阱4：无符号整数在条件判断中
	fmt.Println("\n陷阱4：无符号整数在条件判断中")
	var trapU5 uint8 = 0
	if trapU5-1 > 0 { // 0 - 1 = 255，255 > 0 为 true
		fmt.Printf("uint8(0) - 1 = %d > 0 为 true (注意：这不是预期的行为)\n", trapU5-1)
	}

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

/*
运行结果：

=== 整数类型分类 ===
int8:  int8, 值: 127
int16: int16, 值: 32767
int32: int32, 值: 2147483647
int64: int64, 值: 9223372036854775807
uint8:  uint8, 值: 255
uint16: uint16, 值: 65535
uint32: uint32, 值: 4294967295
uint64: uint64, 值: 18446744073709551615
int:  int, 值: 100
uint: uint, 值: 200

=== 整数溢出和回绕 ===

有符号整数溢出示例:
int8(127) + 1 = -128 (溢出回绕到最小值 -128)
int8(-128) - 1 = 127 (溢出回绕到最大值 127)
int16(32767) + 1 = -32768 (溢出回绕到最小值 -32768)
int16(-32768) - 1 = 32767 (溢出回绕到最大值 32767)
int32(2147483647) + 1 = -2147483648 (溢出回绕到最小值)
int64(最大值) + 1 = -9223372036854775808 (溢出回绕到最小值)

无符号整数溢出示例:
uint8(255) + 1 = 0 (溢出回绕到最小值 0)
uint8(0) - 1 = 255 (溢出回绕到最大值 255)
uint16(65535) + 1 = 0 (溢出回绕到最小值 0)
uint16(0) - 1 = 65535 (溢出回绕到最大值 65535)
uint32(最大值) + 1 = 0 (溢出回绕到最小值 0)
uint64(最大值) + 1 = 0 (溢出回绕到最小值 0)

int 类型溢出示例（64位系统）:
int最大值: 9223372036854775807
int最大值 + 1 = -9223372036854775808 (溢出回绕到最小值)

=== 类型转换 ===

3.1 不同大小类型之间的转换:
int32(100) + int64(200) = 300
int16(1000) 转 int8 = -24 (注意：发生截断)
int8(100) 转 int16 = 100 (安全转换)

3.2 有符号和无符号之间的转换:
uint8(255) 转 int8 = -1 (注意：可能溢出)
uint8(127) 转 int8 = 127 (安全转换)
int8(-1) 转 uint8 = 255 (注意：负数会变成大正数)

3.3 同大小类型之间的转换:
int32(-100) 转 uint32 = 4294967196 (注意：负数会变成大正数)
uint32(300) 转 int32 = 300 (安全转换)

3.4 类型转换时的截断示例:
int32(0x12345678) 转 int8 = 120 (0x78，只保留低8位)

=== 使用场景示例 ===
循环计数: 0
循环计数: 1
循环计数: 2
循环计数: 3
循环计数: 4
像素值: 255
时间戳: 1609459200
状态标记: 1010 (二进制)

=== 溢出检测示例 ===

5.1 溢出回绕示例:
int64 最大值: 9223372036854775807
int64 最大值 + 1 = -9223372036854775808 (溢出回绕)
int32 最大值: 2147483647
int32 最大值 + 1 = -2147483648 (溢出回绕)

5.2 使用 math/bits 检测无符号整数溢出:
检测到无符号整数溢出！18446744073709551615 + 1 会溢出
结果（回绕后）: 0
检测到 uint32 溢出！4294967295 + 1 会溢出
结果（回绕后）: 0
uint16(65535) + uint16(1) = 65536 (超出 uint16 范围，回绕后为 0)

5.3 手动检查有符号整数溢出:
int8(100) + int8(50) 会溢出
int8(100) + int8(50) 会溢出
int8(127) + int8(1) 会溢出！

=== 超大整数 big.Int ===

6.1 基本运算:
big.Int 相加: 9223372036854775807 + 1 = 9223372036854775808 (不会溢出)
big.Int 相减: 100 - 50 = 50
big.Int 相乘: 1000000 * 2000000 = 2000000000000
big.Int 相除: 100 / 3 = 33 (整数除法)

6.2 从字符串创建大整数:
从字符串创建: 123456789012345678901234567890
2^100 = 1267650600228229401496703205376

6.3 大整数比较:
big.Int 比较: 100 和 200
  100 < 200: true
  100 > 200: false
  100 == 200: false

=== fmt.Printf 占位符：%T ===
num1 的类型: int
num2 的类型: uint8
num3 的类型: float64
num4 的类型: []int
num5 的类型: main.MyType

=== fmt.Printf 占位符：%c ===
A
中
a

=== fmt.Printf 占位符：%v ===
123
hello
[1 2 3]
{Tom 20}
{Name:Tom Age:20}
"hello"
main.Person{Name:"Tom", Age:20}

=== fmt.Printf 占位符：%[n]v ===
10 10
语言 - Go
int 123 hello

=== 占位符顺序对应关系 ===
int hello 3.14

=== 无符号整数的陷阱 ===

陷阱1：倒序循环会无限循环
错误示例（已注释）:
  for i := uint(0); i >= 0; i-- { ... }
  // 错误：uint 不会小于 0，会无限循环

正确的倒序循环（用 int）：
5 4 3 2 1 0 
用 uint 倒序（需要特殊处理）：
5 4 3 2 1 0

陷阱2：无符号整数减法可能产生大数
uint8(5) - uint8(10) = 251 (注意：不是 -5！)
  说明：无符号整数不能表示负数，减法溢出会回绕

陷阱3：无符号整数比较
uint8(0) > int8(-1)? 不能直接比较，类型不同
但 uint8(0) 转 int8 = 0
int8(-1) 转 uint8 = 255 (负数变成大正数)

陷阱4：无符号整数在条件判断中
uint8(0) - 1 = 255 > 0 为 true (注意：这不是预期的行为)

=== Go vs Java 类型对应 ===
Java long  → Go int64
Java int   → Go int32
Java short → Go int16
Java byte  → Go int8
Java BigInteger → Go math/big.Int

*/

