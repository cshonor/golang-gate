// 示例：Go 语言的类型转换
// 演示数值类型、字符串、布尔值之间的转换和类型断言

package main

import (
	"fmt"
	"strconv"
)

func main() {
	// ============================================
	// 1. Go 语言的强类型特性
	// ============================================
	fmt.Println("=== Go 语言的强类型特性 ===")

	// Go 是一门强类型语言，编译器不允许在不同类型之间进行隐式的混合操作
	var num int = 10
	var str string = "20"

	// 错误：不能直接混合使用不同类型
	// result := num + str  // ❌ 编译错误：invalid operation: num + str (mismatched types int and string)

	// 正确：必须先转换为相同类型
	strNum, _ := strconv.Atoi(str)
	result := num + strNum
	fmt.Printf("num (%d) + strNum (%d) = %d\n", num, strNum, result)

	fmt.Println()

	// ============================================
	// 2. 数值类型之间的转换
	// ============================================
	fmt.Println("=== 数值类型之间的转换 ===")

	// int 转 float64
	var i int = 42
	var f float64 = float64(i)
	fmt.Printf("int(%d) -> float64(%f)\n", i, f)

	// float64 转 int（会截断小数部分）
	var f2 float64 = 3.14
	var i2 int = int(f2)
	fmt.Printf("float64(%f) -> int(%d) (注意：小数部分被截断)\n", f2, i2)

	// int32 转 int64
	var i32 int32 = 100
	var i64 int64 = int64(i32)
	fmt.Printf("int32(%d) -> int64(%d)\n", i32, i64)

	// int64 转 int32（可能溢出）
	var i64_2 int64 = 2147483647 // int32 的最大值
	var i32_2 int32 = int32(i64_2)
	fmt.Printf("int64(%d) -> int32(%d)\n", i64_2, i32_2)

	// byte 转 rune
	var b byte = 'A'
	var r rune = rune(b)
	fmt.Printf("byte('%c') -> rune('%c')\n", b, r)

	// rune 转 byte（可能丢失信息，因为 rune 范围更大）
	var r2 rune = '中'
	// var b2 byte = byte(r2)  // 可能溢出，不推荐
	fmt.Printf("rune('%c', U+%04X) 不能安全转换为 byte\n", r2, r2)

	fmt.Println()

	// ============================================
	// 3. 数值与字符串的转换（strconv 包）
	// ============================================
	fmt.Println("=== 数值与字符串的转换 ===")

	// 字符串转整数：strconv.Atoi（ASCII to Integer）
	str1 := "123"
	num1, err := strconv.Atoi(str1)
	if err != nil {
		fmt.Printf("转换失败: %v\n", err)
	} else {
		fmt.Printf("strconv.Atoi(%q) = %d\n", str1, num1)
	}

	// 字符串转整数：strconv.ParseInt（更灵活，可指定进制）
	str2 := "1010"
	num2, err := strconv.ParseInt(str2, 2, 64) // 二进制转十进制
	if err != nil {
		fmt.Printf("转换失败: %v\n", err)
	} else {
		fmt.Printf("strconv.ParseInt(%q, 2, 64) = %d (二进制转十进制)\n", str2, num2)
	}

	// 字符串转浮点数：strconv.ParseFloat
	str3 := "3.14159"
	f3, err := strconv.ParseFloat(str3, 64)
	if err != nil {
		fmt.Printf("转换失败: %v\n", err)
	} else {
		fmt.Printf("strconv.ParseFloat(%q, 64) = %f\n", str3, f3)
	}

	// 整数转字符串：strconv.Itoa（Integer to ASCII）
	num3 := 456
	str4 := strconv.Itoa(num3)
	fmt.Printf("strconv.Itoa(%d) = %q\n", num3, str4)

	// 整数转字符串：strconv.FormatInt（可指定进制）
	num4 := int64(255)
	str5 := strconv.FormatInt(num4, 16) // 转十六进制字符串
	fmt.Printf("strconv.FormatInt(%d, 16) = %q (十进制转十六进制)\n", num4, str5)

	// 浮点数转字符串：strconv.FormatFloat
	f4 := 3.14159
	str6 := strconv.FormatFloat(f4, 'f', 2, 64) // 'f'=固定格式, 2=保留2位小数, 64=float64
	fmt.Printf("strconv.FormatFloat(%f, 'f', 2, 64) = %q\n", f4, str6)

	// 浮点数转字符串：科学计数法格式
	str7 := strconv.FormatFloat(f4, 'e', 2, 64) // 'e'=科学计数法
	fmt.Printf("strconv.FormatFloat(%f, 'e', 2, 64) = %q\n", f4, str7)

	fmt.Println()

	// ============================================
	// 4. 布尔值与其他类型的转换
	// ============================================
	fmt.Println("=== 布尔值与其他类型的转换 ===")

	// Go 不支持布尔值与数值/字符串的隐式转换，必须显式处理

	// 布尔值转字符串：使用 strconv.FormatBool
	b1 := true
	str8 := strconv.FormatBool(b1)
	fmt.Printf("strconv.FormatBool(%t) = %q\n", b1, str8)

	// 字符串转布尔值：使用 strconv.ParseBool
	str9 := "true"
	b2, err := strconv.ParseBool(str9)
	if err != nil {
		fmt.Printf("转换失败: %v\n", err)
	} else {
		fmt.Printf("strconv.ParseBool(%q) = %t\n", str9, b2)
	}

	// 布尔值不能直接转数值，需要手动处理
	// var num5 int = int(b1)  // ❌ 编译错误：cannot convert bool to int
	// 需要手动转换
	var num5 int
	if b1 {
		num5 = 1
	} else {
		num5 = 0
	}
	fmt.Printf("布尔值 %t 手动转数值: %d\n", b1, num5)

	// 数值转布尔值：需要手动判断
	var num6 int = 0
	var b3 bool = num6 != 0
	fmt.Printf("数值 %d 转布尔值: %t\n", num6, b3)

	var num7 int = 42
	var b4 bool = num7 != 0
	fmt.Printf("数值 %d 转布尔值: %t\n", num7, b4)

	fmt.Println()

	// ============================================
	// 5. 类型断言（Type Assertion）
	// ============================================
	fmt.Println("=== 类型断言 ===")

	// 类型断言用于在接口类型中提取具体类型的值
	var i_interface interface{} = 42

	// 方式1：安全断言（返回两个值：转换后的值和是否成功）
	value, ok := i_interface.(int)
	if ok {
		fmt.Printf("类型断言成功: %d (类型: %T)\n", value, value)
	} else {
		fmt.Println("类型断言失败")
	}

	// 方式2：直接断言（如果失败会 panic）
	var i_interface2 interface{} = "hello"
	// value2 := i_interface2.(int)  // ❌ 会 panic，因为实际类型是 string
	value2 := i_interface2.(string) // ✅ 正确
	fmt.Printf("直接类型断言: %q (类型: %T)\n", value2, value2)

	// 类型断言用于接口类型检查
	var i_interface3 interface{} = 3.14
	switch v := i_interface3.(type) {
	case int:
		fmt.Printf("是 int 类型: %d\n", v)
	case float64:
		fmt.Printf("是 float64 类型: %f\n", v)
	case string:
		fmt.Printf("是 string 类型: %q\n", v)
	default:
		fmt.Printf("未知类型: %v (类型: %T)\n", v, v)
	}

	fmt.Println()

	// ============================================
	// 6. 常见转换场景示例
	// ============================================
	fmt.Println("=== 常见转换场景示例 ===")

	// 场景1：从用户输入读取字符串，转换为数值进行计算
	userInput := "100"
	price, _ := strconv.ParseFloat(userInput, 64)
	tax := price * 0.1
	total := price + tax
	fmt.Printf("价格: %.2f, 税费: %.2f, 总计: %.2f\n", price, tax, total)

	// 场景2：将计算结果转换为字符串显示
	resultValue := 123.456
	resultStr := fmt.Sprintf("结果: %.2f", resultValue) // 使用 fmt.Sprintf 格式化
	fmt.Println(resultStr)

	// 场景3：不同进制之间的转换
	decimal := 255
	binaryStr := strconv.FormatInt(int64(decimal), 2)
	hexStr := strconv.FormatInt(int64(decimal), 16)
	fmt.Printf("十进制 %d = 二进制 %s = 十六进制 %s\n", decimal, binaryStr, hexStr)

	// 场景4：字符串切片转数值切片
	strSlice := []string{"1", "2", "3", "4", "5"}
	intSlice := make([]int, len(strSlice))
	for i, s := range strSlice {
		intSlice[i], _ = strconv.Atoi(s)
	}
	fmt.Printf("字符串切片: %v -> 整数切片: %v\n", strSlice, intSlice)

	fmt.Println()

	// ============================================
	// 7. 转换时的注意事项
	// ============================================
	fmt.Println("=== 转换时的注意事项 ===")

	// 注意1：浮点数转整数会截断小数部分
	f5 := 3.9
	i5 := int(f5)
	fmt.Printf("float64(%f) -> int(%d) (小数部分被截断，不是四舍五入)\n", f5, i5)

	// 注意2：大类型转小类型可能溢出
	var bigInt int64 = 300
	var smallInt int8 = int8(bigInt) // 300 超过 int8 范围（-128~127），会溢出
	fmt.Printf("int64(%d) -> int8(%d) (溢出回绕)\n", bigInt, smallInt)

	// 注意3：字符串转数值时，格式错误会返回错误
	invalidStr := "abc"
	_, err = strconv.Atoi(invalidStr)
	if err != nil {
		fmt.Printf("字符串 %q 转整数失败: %v\n", invalidStr, err)
	}

	// 注意4：类型断言失败会 panic（如果不用安全断言）
	var i_interface4 interface{} = 42
	// var str10 string = i_interface4.(string)  // ❌ 会 panic
	// 应该使用安全断言
	str10, ok := i_interface4.(string)
	if !ok {
		fmt.Printf("类型断言失败: 无法将 %T 转换为 string\n", i_interface4)
	} else {
		fmt.Printf("类型断言成功: %q\n", str10)
	}

	fmt.Println()

	// ============================================
	// 8. 类型转换速查表
	// ============================================
	fmt.Println("=== 类型转换速查表 ===")
	fmt.Println("数值类型之间:")
	fmt.Println("  int -> float64: float64(i)")
	fmt.Println("  float64 -> int: int(f) (截断小数)")
	fmt.Println("  int32 -> int64: int64(i32)")
	fmt.Println("  int64 -> int32: int32(i64) (可能溢出)")
	fmt.Println()
	fmt.Println("数值 <-> 字符串:")
	fmt.Println("  字符串 -> int: strconv.Atoi(s)")
	fmt.Println("  字符串 -> int64: strconv.ParseInt(s, 10, 64)")
	fmt.Println("  字符串 -> float64: strconv.ParseFloat(s, 64)")
	fmt.Println("  int -> 字符串: strconv.Itoa(i)")
	fmt.Println("  int64 -> 字符串: strconv.FormatInt(i, 10)")
	fmt.Println("  float64 -> 字符串: strconv.FormatFloat(f, 'f', 2, 64)")
	fmt.Println()
	fmt.Println("布尔值 <-> 字符串:")
	fmt.Println("  布尔值 -> 字符串: strconv.FormatBool(b)")
	fmt.Println("  字符串 -> 布尔值: strconv.ParseBool(s)")
	fmt.Println()
	fmt.Println("类型断言:")
	fmt.Println("  安全断言: value, ok := i.(Type)")
	fmt.Println("  直接断言: value := i.(Type) (失败会 panic)")
	fmt.Println("  类型开关: switch v := i.(type) { case Type: ... }")
}

