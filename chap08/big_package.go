// 示例：Go 语言的 big 包
// 演示 big.Int、big.Float、big.Rat 的使用，解决原生数值类型范围限制的问题

package main

import (
	"fmt"
	"math/big"
)

func main() {
	// ============================================
	// 1. big 包的作用和使用场景
	// ============================================
	fmt.Println("=== big 包的作用 ===")

	// big 包解决的核心问题：Go 原生数值类型的范围限制
	// 比如 int64 最大只能表示 9223372036854775807
	maxInt64 := int64(9223372036854775807)
	fmt.Printf("int64 最大值: %d\n", maxInt64)

	// 如果需要计算更大的数（比如 1000 的阶乘、加密算法里的大质数）
	// 用 int64 肯定会溢出，但用 big.Int 就不会有这个问题
	// big.Int 能根据数值大小动态分配内存，理论上可以表示无限大的整数

	fmt.Println()

	// ============================================
	// 2. big 包的三种类型
	// ============================================
	fmt.Println("=== big 包的三种类型 ===")

	// big.Int：大整数，可以表示任意大小的整数
	// big.Float：大浮点数，可以设置任意精度
	// big.Rat：大有理数（分数），比如 1/2、3/4

	fmt.Println("big.Int   - 大整数")
	fmt.Println("big.Float - 大浮点数")
	fmt.Println("big.Rat   - 大有理数（分数）")

	fmt.Println()

	// ============================================
	// 3. 创建 big.Int 的两种方式
	// ============================================
	fmt.Println("=== 创建 big.Int 的两种方式 ===")

	// 方式1：使用 big.NewInt() 直接初始化（最常用）
	// big.NewInt() 的参数必须是 int64 类型
	secPerDay1 := big.NewInt(86400)
	fmt.Printf("方式1 - big.NewInt(86400): %s\n", secPerDay1.String())

	// 方式2：使用 new(big.Int) + SetString/SetInt64
	// new 是 Go 语言内置的关键字，不是 big 包专属的
	// new(Type) 会创建一个 Type 类型的指针，初始值为零值
	secPerDay2 := new(big.Int)
	secPerDay2.SetString("86400", 10) // SetString(字符串, 进制)
	fmt.Printf("方式2 - new(big.Int) + SetString: %s\n", secPerDay2.String())

	// 方式2的变种：用 SetInt64
	secPerDay3 := new(big.Int)
	secPerDay3.SetInt64(86400)
	fmt.Printf("方式2变种 - new(big.Int) + SetInt64: %s\n", secPerDay3.String())

	// 方式3：声明 big.Int 值类型（不是指针）
	var secPerDay4 big.Int
	secPerDay4.SetInt64(86400)
	fmt.Printf("方式3 - var big.Int + SetInt64: %s\n", secPerDay4.String())

	fmt.Println()

	// ============================================
	// 4. new 关键字是 Go 语言内置的
	// ============================================
	fmt.Println("=== new 关键字是 Go 语言内置的 ===")

	// new 不是 big 包专属的，它是 Go 语言内置的关键字
	// 可以用来创建任意类型的指针

	// 创建 int 类型的指针
	intPtr := new(int)
	*intPtr = 100
	fmt.Printf("new(int): %d\n", *intPtr)

	// 创建 string 类型的指针
	strPtr := new(string)
	*strPtr = "hello"
	fmt.Printf("new(string): %s\n", *strPtr)

	// 创建 big.Int 类型的指针
	bigIntPtr := new(big.Int)
	bigIntPtr.SetInt64(12345)
	fmt.Printf("new(big.Int): %s\n", bigIntPtr.String())

	fmt.Println()

	// ============================================
	// 5. SetString 和 SetInt64 方法详解
	// ============================================
	fmt.Println("=== SetString 和 SetInt64 方法详解 ===")

	// ============================================
	// 5.1 SetInt64 方法
	// ============================================
	fmt.Println("\n5.1 SetInt64 方法:")
	fmt.Println("作用：将 int64 类型的值设置到 big.Int 中")
	fmt.Println("参数：int64 类型的数值")
	fmt.Println("返回值：*big.Int（返回自身，支持链式调用）")
	fmt.Println("使用场景：当数值在 int64 范围内时使用")
	fmt.Println("限制：只能设置 int64 范围内的值（-9223372036854775808 到 9223372036854775807）")

	bigNum1 := new(big.Int)
	bigNum1.SetInt64(12345)
	fmt.Printf("示例：SetInt64(12345) = %s\n", bigNum1.String())

	bigNum2 := new(big.Int)
	bigNum2.SetInt64(-100)
	fmt.Printf("示例：SetInt64(-100) = %s（支持负数）\n", bigNum2.String())

	// SetInt64 的限制：不能设置超过 int64 范围的值
	fmt.Println("\nSetInt64 的限制：")
	fmt.Println("  - 最大值：9223372036854775807")
	fmt.Println("  - 最小值：-9223372036854775808")
	fmt.Println("  - 超过这个范围的值必须使用 SetString")

	// ============================================
	// 5.2 SetString 方法
	// ============================================
	fmt.Println("\n5.2 SetString 方法:")
	fmt.Println("作用：将字符串形式的数字转换成 big.Int")
	fmt.Println("参数1：数字字符串（必须是纯数字，不能有字母、符号）")
	fmt.Println("参数2：进制（10=十进制，2=二进制，16=十六进制，0=自动识别）")
	fmt.Println("返回值：(*big.Int, bool) - 大数指针和是否成功")
	fmt.Println("使用场景：")
	fmt.Println("  - 处理超过 int64 范围的超大数")
	fmt.Println("  - 从字符串（用户输入、文件读取）创建大数")
	fmt.Println("  - 支持不同进制（二进制、十六进制等）")

	bigNum := new(big.Int)

	// 十进制
	bigNum.SetString("86400", 10)
	fmt.Printf("SetString(\"86400\", 10): %s\n", bigNum.String())

	// 二进制 "1010" = 十进制的 10
	bigNum.SetString("1010", 2)
	fmt.Printf("SetString(\"1010\", 2): %s\n", bigNum.String())

	// 十六进制 "FF" = 十进制的 255
	bigNum.SetString("FF", 16)
	fmt.Printf("SetString(\"FF\", 16): %s\n", bigNum.String())

	// 自动识别进制（传 0）
	bigNum.SetString("0xFF", 0) // 0x 开头，自动识别为十六进制
	fmt.Printf("SetString(\"0xFF\", 0): %s\n", bigNum.String())

	bigNum.SetString("123", 0) // 没有前缀，自动识别为十进制
	fmt.Printf("SetString(\"123\", 0): %s\n", bigNum.String())

	// SetString 返回两个值：转换后的大数指针 和 是否成功（bool）
	bigNum2Str := new(big.Int)
	success, ok := bigNum2Str.SetString("86400a", 10) // 字符串里有字母，转换会失败
	if !ok {
		fmt.Printf("SetString(\"86400a\", 10) 失败: 字符串格式错误\n")
	} else {
		fmt.Printf("SetString 成功: %s\n", success.String())
	}

	// 处理超大数：超过 int64 范围的数字只能用 SetString
	veryBigNum := new(big.Int)
	veryBigNum.SetString("1234567890123456789012345678901234567890", 10)
	fmt.Printf("超大数: %s\n", veryBigNum.String())

	// ============================================
	// 5.3 SetString 和 SetInt64 的对比
	// ============================================
	fmt.Println("\n5.3 SetString 和 SetInt64 的对比:")
	fmt.Println("┌─────────────┬──────────────┬──────────────┐")
	fmt.Println("│   特性      │   SetInt64   │   SetString  │")
	fmt.Println("├─────────────┼──────────────┼──────────────┤")
	fmt.Println("│ 参数类型    │    int64     │    string    │")
	fmt.Println("│ 数值范围    │ 受 int64 限制│    无限制    │")
	fmt.Println("│ 进制支持    │    不支持    │     支持     │")
	fmt.Println("│ 返回值      │   *big.Int   │ (*big.Int,bool)│")
	fmt.Println("│ 使用场景    │ 小到中等数值 │ 任意大小数值 │")
	fmt.Println("└─────────────┴──────────────┴──────────────┘")

	// 实际应用示例对比
	fmt.Println("\n实际应用示例对比:")
	
	// 场景1：小数值（两种方法都可以）
	fmt.Println("\n场景1：小数值（两种方法都可以）")
	small1 := new(big.Int)
	small1.SetInt64(100)
	fmt.Printf("  SetInt64(100): %s\n", small1.String())
	
	small2 := new(big.Int)
	small2.SetString("100", 10)
	fmt.Printf("  SetString(\"100\", 10): %s\n", small2.String())
	
	// 场景2：超大数（只能用 SetString）
	fmt.Println("\n场景2：超大数（只能用 SetString）")
	// bigNum3.SetInt64(123456789012345678901234567890)  // ❌ 编译错误：超出 int64 范围
	bigNum3 := new(big.Int)
	bigNum3.SetString("123456789012345678901234567890", 10)
	fmt.Printf("  SetString(\"123456789012345678901234567890\", 10): %s\n", bigNum3.String())
	
	// 场景3：不同进制（只能用 SetString）
	fmt.Println("\n场景3：不同进制（只能用 SetString）")
	hexNum := new(big.Int)
	hexNum.SetString("FF", 16)
	fmt.Printf("  SetString(\"FF\", 16): %s（十六进制转十进制）\n", hexNum.String())
	
	binNum := new(big.Int)
	binNum.SetString("1010", 2)
	fmt.Printf("  SetString(\"1010\", 2): %s（二进制转十进制）\n", binNum.String())

	fmt.Println()

	// ============================================
	// 6. big.Int 的运算
	// ============================================
	fmt.Println("=== big.Int 的运算 ===")

	// big.Int 的运算需要用方法，不能用 + - * / 运算符
	a := big.NewInt(100)
	b := big.NewInt(200)

	// 加法：Add
	result := new(big.Int)
	result.Add(a, b)
	fmt.Printf("%s + %s = %s\n", a.String(), b.String(), result.String())

	// 乘法：Mul
	result.Mul(a, b)
	fmt.Printf("%s × %s = %s\n", a.String(), b.String(), result.String())

	// 减法：Sub
	result.Sub(b, a)
	fmt.Printf("%s - %s = %s\n", b.String(), a.String(), result.String())

	// 除法：Div
	result.Div(b, a)
	fmt.Printf("%s ÷ %s = %s\n", b.String(), a.String(), result.String())

	fmt.Println()

	// ============================================
	// 7. big.Float 的使用
	// ============================================
	fmt.Println("=== big.Float 的使用 ===")

	// 创建 big.Float
	// 方式1：big.NewFloat()
	float1 := big.NewFloat(3.141592653589793)
	fmt.Printf("big.NewFloat(3.14159...): %s\n", float1.String())

	// 方式2：new(big.Float) + SetFloat64
	float2 := new(big.Float)
	float2.SetFloat64(3.141592653589793)
	fmt.Printf("new(big.Float) + SetFloat64: %s\n", float2.String())

	// 设置精度：SetPrec（设置有效位数）
	float3 := new(big.Float)
	float3.SetPrec(100) // 设置 100 位有效数字
	float3.SetFloat64(1.0 / 3.0)
	fmt.Printf("高精度 1/3 (100位): %s\n", float3.Text('f', 50)) // 显示50位小数

	fmt.Println()

	// ============================================
	// 8. big.Rat 的使用（有理数/分数）
	// ============================================
	fmt.Println("=== big.Rat 的使用（有理数/分数）===")

	// 创建 big.Rat（分数）
	// 方式1：big.NewRat(分子, 分母)
	rat1 := big.NewRat(1, 2) // 1/2
	fmt.Printf("big.NewRat(1, 2): %s\n", rat1.String())

	rat2 := big.NewRat(3, 4) // 3/4
	fmt.Printf("big.NewRat(3, 4): %s\n", rat2.String())

	// 方式2：new(big.Rat) + SetString
	rat3 := new(big.Rat)
	rat3.SetString("1/3") // 字符串形式 "分子/分母"
	fmt.Printf("new(big.Rat) + SetString(\"1/3\"): %s\n", rat3.String())

	// 分数转浮点数：Float64()
	floatValue, _ := rat1.Float64()
	fmt.Printf("1/2 转浮点数: %f\n", floatValue)

	fmt.Println()

	// ============================================
	// 9. 常量的特性（无类型常量，不会溢出）
	// ============================================
	fmt.Println("=== 常量的特性（无类型常量，不会溢出）===")

	// Go 的常量在编译期处理，默认是"无类型常量"
	// 常量不会溢出，因为编译器会根据常量的大小，自动用足够大的临时类型去容纳它
	// 注意：虽然常量本身不会溢出，但直接用于 fmt.Printf 等函数时，如果超过 int 范围会报错
	// 所以这里用一个稍小但仍然很大的常量来演示
	const bigConst = "12345678901234567890123456789" // 用字符串形式表示超大数

	fmt.Printf("超大常量（字符串形式）: %s\n", bigConst)

	// 常量赋值给 big.Int（不会溢出）
	bigFromConst := big.NewInt(0)
	bigFromConst.SetString(bigConst, 10)
	fmt.Printf("常量转 big.Int: %s\n", bigFromConst.String())

	// 常量赋值给 int64（会编译报错，这里用注释说明）
	// const numConst = 12345678901234567890123456789
	// var num2 int64 = numConst // 编译器会提示：constant overflows int64

	fmt.Println()

	// ============================================
	// 10. 实际应用示例：计算阶乘
	// ============================================
	fmt.Println("=== 实际应用示例：计算阶乘 ===")

	// 计算 20 的阶乘（20!），这个值远超 int64 范围
	factorial := big.NewInt(1)
	for i := int64(1); i <= 20; i++ {
		factorial.Mul(factorial, big.NewInt(i))
	}
	fmt.Printf("20! = %s\n", factorial.String())

	// 计算 100 的阶乘（100!），只能用 big.Int
	factorial100 := big.NewInt(1)
	for i := int64(1); i <= 100; i++ {
		factorial100.Mul(factorial100, big.NewInt(i))
	}
	fmt.Printf("100! 的位数: %d 位\n", len(factorial100.String()))
	fmt.Printf("100! 的前50位: %s...\n", factorial100.String()[:50])

	fmt.Println()

	// ============================================
	// 11. 实际应用示例：高精度金融计算
	// ============================================
	fmt.Println("=== 实际应用示例：高精度金融计算 ===")

	// 金融计算需要高精度，避免浮点数精度损失
	// 使用 big.Float 设置高精度
	price := new(big.Float)
	price.SetPrec(100) // 设置 100 位精度
	price.SetFloat64(19.99)

	quantity := new(big.Float)
	quantity.SetPrec(100)
	quantity.SetFloat64(1000)

	resultFloat := new(big.Float)
	resultFloat.SetPrec(100)
	resultFloat.Mul(price, quantity)

	// big.Float 的 Text 方法需要两个参数：格式字符和精度
	// 'f' 表示固定小数点格式，第二个参数表示小数位数
	fmt.Printf("单价: %s\n", price.Text('f', 2))
	fmt.Printf("数量: %s\n", quantity.Text('f', 0))
	fmt.Printf("总价: %s\n", resultFloat.Text('f', 2))
	
	// 也可以先用 Float64() 转成 float64，再用 fmt.Printf 格式化
	priceVal, _ := price.Float64()
	quantityVal, _ := quantity.Float64()
	resultVal, _ := resultFloat.Float64()
	fmt.Printf("单价（保留2位）: %.2f\n", priceVal)
	fmt.Printf("数量: %.0f\n", quantityVal)
	fmt.Printf("总价（保留2位）: %.2f\n", resultVal)

	fmt.Println()

	// ============================================
	// 12. 总结
	// ============================================
	fmt.Println("=== 总结 ===")
	fmt.Println("1. big 包解决原生数值类型范围限制的问题")
	fmt.Println("2. big.Int：大整数，可以表示任意大小的整数")
	fmt.Println("3. big.Float：大浮点数，可以设置任意精度")
	fmt.Println("4. big.Rat：大有理数（分数）")
	fmt.Println("5. 创建方式：big.NewXXX() 或 new(big.XXX) + SetXXX()")
	fmt.Println("6. new 是 Go 语言内置关键字，不是 big 包专属的")
	fmt.Println("7. SetString(字符串, 进制) 用于从字符串创建大数")
	fmt.Println("8. 常量是无类型的，不会溢出，赋值时才检查类型")
	fmt.Println("9. big.Int 运算用方法（Add、Mul、Sub、Div），不用运算符")
}

