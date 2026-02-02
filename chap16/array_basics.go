// 示例：数组的声明和访问元素
// 演示数组的声明、零值初始化、索引访问和数组长度

package main

import "fmt"

func main() {
	// ============================================
	// 1. 声明数组
	// ============================================
	fmt.Println("=== 声明数组 ===")

	// 声明长度为8的字符串数组
	var planets [8]string
	fmt.Printf("planets 类型: %T, 长度: %d\n", planets, len(planets))
	fmt.Printf("planets 零值: %v\n", planets)
	fmt.Println()

	// ============================================
	// 2. 为数组元素赋值
	// ============================================
	fmt.Println("=== 为数组元素赋值 ===")

	// 为部分元素赋值
	planets[0] = "Mercury"
	planets[1] = "Venus"
	planets[2] = "Earth"
	planets[3] = "Mars"
	planets[4] = "Jupiter"
	planets[5] = "Saturn"
	planets[6] = "Uranus"
	planets[7] = "Neptune"

	fmt.Printf("赋值后的 planets: %v\n", planets)
	fmt.Println()

	// ============================================
	// 3. 访问数组元素（速查16-1解答）
	// ============================================
	fmt.Println("=== 访问数组元素（速查16-1解答）===")

	// 1. 访问 planets 数组的第一个元素
	// 答案：通过索引 0 访问，即 planets[0]
	firstPlanet := planets[0]
	fmt.Printf("第一个元素 planets[0]: %s\n", firstPlanet)

	// 访问第三个元素（索引2）
	earth := planets[2]
	fmt.Printf("第三个元素 planets[2]: %s\n", earth)

	// 访问最后一个元素
	lastIndex := len(planets) - 1
	lastPlanet := planets[lastIndex]
	fmt.Printf("最后一个元素 planets[%d]: %s\n", lastIndex, lastPlanet)

	fmt.Println()
	fmt.Println("关键点:")
	fmt.Println("  - Go 语言的数组索引从 0 开始")
	fmt.Println("  - 第一个元素对应索引 0")
	fmt.Println("  - 最后一个元素对应索引 len(planets)-1")
	fmt.Println()

	// ============================================
	// 4. 零值初始化（速查16-1解答）
	// ============================================
	fmt.Println("=== 零值初始化（速查16-1解答）===")

	// 2. 新创建的整数数组元素默认值
	// 答案：未被赋值的元素会被初始化为对应类型的零值，整数数组的零值是 0
	var numbers [5]int
	fmt.Printf("整数数组 numbers: %v\n", numbers)
	fmt.Printf("numbers[0] 的零值: %d\n", numbers[0])
	fmt.Printf("numbers[3] 的零值: %d\n", numbers[3])
	fmt.Println("  说明: 整数数组的零值是 0")

	// 字符串数组的零值
	var strings [3]string
	fmt.Printf("\n字符串数组 strings: %v\n", strings)
	fmt.Printf("strings[0] 的零值: %q\n", strings[0])
	fmt.Printf("strings[0] == \"\": %t\n", strings[0] == "")
	fmt.Println("  说明: 字符串数组的零值是空字符串 \"\"")

	// 布尔数组的零值
	var bools [3]bool
	fmt.Printf("\n布尔数组 bools: %v\n", bools)
	fmt.Printf("bools[0] 的零值: %t\n", bools[0])
	fmt.Println("  说明: 布尔数组的零值是 false")

	// 浮点数数组的零值
	var floats [3]float64
	fmt.Printf("\n浮点数数组 floats: %v\n", floats)
	fmt.Printf("floats[0] 的零值: %.2f\n", floats[0])
	fmt.Println("  说明: 浮点数数组的零值是 0.0")

	fmt.Println()
	fmt.Println("常见类型的零值:")
	fmt.Println("  整数类型: 0")
	fmt.Println("  字符串: \"\" (空字符串)")
	fmt.Println("  布尔值: false")
	fmt.Println("  指针: nil")
	fmt.Println("  浮点数: 0.0")
	fmt.Println()

	// ============================================
	// 5. 验证未赋值元素的零值
	// ============================================
	fmt.Println("=== 验证未赋值元素的零值 ===")

	// 只给前3个元素赋值
	var planets2 [8]string
	planets2[0] = "Mercury"
	planets2[1] = "Venus"
	planets2[2] = "Earth"

	fmt.Printf("planets2: %v\n", planets2)
	fmt.Printf("已赋值元素 planets2[2]: %q\n", planets2[2])
	fmt.Printf("未赋值元素 planets2[3]: %q\n", planets2[3])
	fmt.Printf("planets2[3] == \"\": %t\n", planets2[3] == "")
	fmt.Printf("未赋值元素 planets2[7]: %q\n", planets2[7])
	fmt.Println()

	// ============================================
	// 6. 数组长度
	// ============================================
	fmt.Println("=== 数组长度 ===")

	fmt.Printf("planets 数组长度: %d\n", len(planets))
	fmt.Printf("numbers 数组长度: %d\n", len(numbers))
	fmt.Println()
	fmt.Println("关键点:")
	fmt.Println("  - 数组长度通过 len() 函数获取")
	fmt.Println("  - 数组长度在声明后不可改变")
	fmt.Println("  - 长度是数组类型的一部分")
	fmt.Println()

	// ============================================
	// 7. 索引访问规则
	// ============================================
	fmt.Println("=== 索引访问规则 ===")

	fmt.Println("索引范围: 0 到 len(array)-1")
	fmt.Printf("planets 数组索引范围: 0 到 %d\n", len(planets)-1)

	// 访问所有元素
	fmt.Println("\n访问所有元素:")
	for i := 0; i < len(planets); i++ {
		fmt.Printf("  planets[%d] = %s\n", i, planets[i])
	}

	// 错误示例（注释掉，避免编译错误）
	fmt.Println("\n注意:")
	fmt.Println("  - 访问超出范围的索引会导致运行时错误（panic）")
	fmt.Println("  - 例如: planets[8] 会 panic（索引范围是 0-7）")
	fmt.Println()

	// ============================================
	// 8. 完整示例：array.go
	// ============================================
	fmt.Println("=== 完整示例：array.go ===")

	// 声明长度为8的字符串数组
	var planets3 [8]string

	// 为部分元素赋值
	planets3[0] = "Mercury"
	planets3[1] = "Venus"
	planets3[2] = "Earth"

	// 访问并打印第三个元素（索引2）
	earth2 := planets3[2]
	fmt.Println("第三个元素:", earth2) // 输出：Earth

	// 打印数组长度
	fmt.Println("数组长度:", len(planets3)) // 输出：8

	// 验证未赋值元素的零值
	fmt.Println("planets3[3] == \"\":", planets3[3] == "") // 输出：true
	fmt.Println()

	// ============================================
	// 9. 数组常用操作清单
	// ============================================
	fmt.Println("=== 数组常用操作清单 ===")
	fmt.Println()
	fmt.Println("1. 声明数组:")
	fmt.Println("   var arr [5]int              // 声明长度为5的整数数组")
	fmt.Println("   var arr [5]string           // 声明长度为5的字符串数组")
	fmt.Println()
	fmt.Println("2. 初始化:")
	fmt.Println("   var arr [3]int = [3]int{1, 2, 3}  // 完整初始化")
	fmt.Println("   arr := [3]int{1, 2, 3}            // 短变量声明")
	fmt.Println("   arr := [...]int{1, 2, 3}          // 编译器推断长度")
	fmt.Println()
	fmt.Println("3. 访问元素:")
	fmt.Println("   arr[0]                     // 第一个元素（索引0）")
	fmt.Println("   arr[len(arr)-1]            // 最后一个元素")
	fmt.Println()
	fmt.Println("4. 赋值元素:")
	fmt.Println("   arr[0] = 10                // 给第一个元素赋值")
	fmt.Println()
	fmt.Println("5. 获取长度:")
	fmt.Println("   len(arr)                   // 返回数组长度")
	fmt.Println()
	fmt.Println("6. 零值:")
	fmt.Println("   整数: 0")
	fmt.Println("   字符串: \"\"")
	fmt.Println("   布尔值: false")
	fmt.Println("   指针: nil")
	fmt.Println()
	fmt.Println("7. 迭代:")
	fmt.Println("   for i := 0; i < len(arr); i++ { ... }")
	fmt.Println("   for i, v := range arr { ... }")
	fmt.Println()

	// ============================================
	// 10. 总结
	// ============================================
	fmt.Println("=== 总结 ===")
	fmt.Println("1. 数组索引从 0 开始，第一个元素是 arr[0]")
	fmt.Println("2. 最后一个元素的索引是 len(arr)-1")
	fmt.Println("3. 未赋值的元素会自动初始化为对应类型的零值")
	fmt.Println("4. 整数数组的零值是 0，字符串数组的零值是空字符串")
	fmt.Println("5. 数组长度通过 len() 函数获取，且不可改变")
	fmt.Println("6. 访问超出范围的索引会导致运行时错误")
}

