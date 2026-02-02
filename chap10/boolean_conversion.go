// 示例：Go 语言的布尔值转换
// 演示布尔值与字符串、数值之间的转换，以及静态类型特性

package main

import (
	"fmt"
	"strconv"
)

func main() {
	// ============================================
	// 1. Go 的静态类型特性
	// ============================================
	fmt.Println("=== Go 的静态类型特性 ===")

	// Go 是静态类型语言，变量声明时绑定类型，类型不可变
	var countdown int = 10
	fmt.Printf("countdown 类型: %T, 值: %d\n", countdown, countdown)

	// 错误：不能将不同类型的值赋给已声明的变量
	// countdown = 0.5  // ❌ 编译错误：cannot use 0.5 (untyped float constant) as int value
	// countdown = "5 seconds"  // ❌ 编译错误：cannot use "5 seconds" (untyped string constant) as int value

	// 正确：只能赋相同类型的值
	countdown = 5
	fmt.Printf("countdown 重新赋值: %d\n", countdown)

	// 对比：动态类型语言（如 JavaScript/Python）可以随时改变类型
	fmt.Println("\n对比动态类型语言:")
	fmt.Println("  JavaScript: var countdown = 10; countdown = 0.5; countdown = '5 seconds' (允许)")
	fmt.Println("  Go: var countdown int = 10; countdown = 0.5 (不允许，编译错误)")

	fmt.Println()

	// ============================================
	// 2. 布尔值转字符串
	// ============================================
	fmt.Println("=== 布尔值转字符串 ===")

	var launch bool = false
	fmt.Printf("原始布尔值: %t\n", launch)

	// 方法1：使用 fmt.Sprintf（通用方式）
	launchText1 := fmt.Sprintf("%v", launch)
	fmt.Printf("方法1 - fmt.Sprintf(\"%%v\", %t): %q\n", launch, launchText1)

	// 方法2：使用 strconv.FormatBool（性能更优，推荐）
	launchText2 := strconv.FormatBool(launch)
	fmt.Printf("方法2 - strconv.FormatBool(%t): %q\n", launch, launchText2)

	// 测试 true 值
	var launchTrue bool = true
	launchText3 := strconv.FormatBool(launchTrue)
	fmt.Printf("strconv.FormatBool(%t): %q\n", launchTrue, launchText3)

	// 性能对比说明
	fmt.Println("\n性能说明:")
	fmt.Println("  strconv.FormatBool 性能更优，适合大量转换场景")
	fmt.Println("  fmt.Sprintf 更通用，但性能略低")

	fmt.Println()

	// ============================================
	// 3. 布尔值转数值
	// ============================================
	fmt.Println("=== 布尔值转数值 ===")

	// Go 没有直接转换语法，需要用 if 语句显式处理
	var launchBool bool = true
	var num int
	if launchBool {
		num = 1
	} else {
		num = 0
	}
	fmt.Printf("布尔值 %t 转数值: %d\n", launchBool, num)

	// 封装成函数
	fmt.Println("\n使用函数转换:")
	fmt.Printf("BoolToInt(%t) = %d\n", true, BoolToInt(true))
	fmt.Printf("BoolToInt(%t) = %d\n", false, BoolToInt(false))

	// 转换为其他数值类型
	var numFloat float64
	if launchBool {
		numFloat = 1.0
	} else {
		numFloat = 0.0
	}
	fmt.Printf("布尔值 %t 转 float64: %.1f\n", launchBool, numFloat)

	fmt.Println()

	// ============================================
	// 4. 字符串转布尔值
	// ============================================
	fmt.Println("=== 字符串转布尔值 ===")

	// 使用 strconv.ParseBool，仅识别 "true"/"false"（大小写敏感）
	testCases := []string{"true", "false", "TRUE", "FALSE", "True", "False", "1", "0", "yes", "no"}

	for _, s := range testCases {
		b, err := strconv.ParseBool(s)
		if err != nil {
			fmt.Printf("strconv.ParseBool(%q): 转换失败 - %v\n", s, err)
		} else {
			fmt.Printf("strconv.ParseBool(%q): %t\n", s, b)
		}
	}

	fmt.Println("\n注意事项:")
	fmt.Println("  strconv.ParseBool 识别以下格式（大小写不敏感）:")
	fmt.Println("    - \"true\", \"TRUE\", \"True\" -> true")
	fmt.Println("    - \"false\", \"FALSE\", \"False\" -> false")
	fmt.Println("    - \"1\" -> true, \"0\" -> false")
	fmt.Println("    - 其他字符串（如 \"yes\", \"no\"）会返回错误")

	fmt.Println()

	// ============================================
	// 5. 数值转布尔值
	// ============================================
	fmt.Println("=== 数值转布尔值 ===")

	// Go 没有直接转换语法，需要手动判断
	testNumbers := []int{0, 1, -1, 42, -100}

	for _, n := range testNumbers {
		b := n != 0 // 非零值为 true，零值为 false
		fmt.Printf("数值 %d 转布尔值: %t\n", n, b)
	}

	// 封装成函数
	fmt.Println("\n使用函数转换:")
	fmt.Printf("IntToBool(%d) = %t\n", 0, IntToBool(0))
	fmt.Printf("IntToBool(%d) = %t\n", 1, IntToBool(1))
	fmt.Printf("IntToBool(%d) = %t\n", 42, IntToBool(42))

	fmt.Println()

	// ============================================
	// 6. 常见转换场景示例
	// ============================================
	fmt.Println("=== 常见转换场景示例 ===")

	// 场景1：配置文件读取（字符串 -> 布尔值）
	configValue := "true"
	enabled, err := strconv.ParseBool(configValue)
	if err != nil {
		fmt.Printf("配置解析失败: %v\n", err)
	} else {
		fmt.Printf("配置值 %q 解析为: %t\n", configValue, enabled)
	}

	// 场景2：状态显示（布尔值 -> 字符串）
	isOnline := true
	statusText := strconv.FormatBool(isOnline)
	fmt.Printf("在线状态 %t 显示为: %q\n", isOnline, statusText)

	// 场景3：数据库存储（布尔值 -> 数值）
	hasPermission := true
	permissionFlag := BoolToInt(hasPermission)
	fmt.Printf("权限 %t 存储为数值: %d\n", hasPermission, permissionFlag)

	// 场景4：从数据库读取（数值 -> 布尔值）
	dbValue := 1
	hasAccess := IntToBool(dbValue)
	fmt.Printf("数据库值 %d 转换为权限: %t\n", dbValue, hasAccess)

	fmt.Println()

	// ============================================
	// 7. 接口的灵活性
	// ============================================
	fmt.Println("=== 接口的灵活性 ===")

	// 虽然 Go 是静态类型，但通过接口可以接受不同类型的值
	fmt.Println("fmt.Println 可以接受不同类型的值:")
	fmt.Println(true)        // bool
	fmt.Println(42)          // int
	fmt.Println(3.14)        // float64
	fmt.Println("hello")      // string
	fmt.Println([]int{1, 2})  // slice

	// 接口类型可以存储任意类型的值
	var i interface{} = true
	fmt.Printf("interface{} 存储 bool: %v (类型: %T)\n", i, i)

	i = 42
	fmt.Printf("interface{} 存储 int: %v (类型: %T)\n", i, i)

	i = "hello"
	fmt.Printf("interface{} 存储 string: %v (类型: %T)\n", i, i)

	fmt.Println()

	// ============================================
	// 8. 静态类型 vs 动态类型对比
	// ============================================
	fmt.Println("=== 静态类型 vs 动态类型对比 ===")

	fmt.Println("特性对比:")
	fmt.Println("  特性\t\t\tGo（静态类型）\t\tJavaScript/Python（动态类型）")
	fmt.Println("  类型绑定时机\t\t变量声明时绑定，类型不可变\t变量可以随时改变类型")
	fmt.Println("  性能\t\t\t编译期类型检查，运行速度快\t运行期动态解析，性能略低")
	fmt.Println("  错误检测\t\t编译期就能发现类型错误\t部分错误要到运行时才暴露")
	fmt.Println()
	fmt.Println("示例:")
	fmt.Println("  Go: var countdown int = 10; countdown = 0.5 (编译错误)")
	fmt.Println("  JavaScript: var countdown = 10; countdown = 0.5 (允许)")

	fmt.Println()

	// ============================================
	// 9. 布尔值转换速查表
	// ============================================
	fmt.Println("=== 布尔值转换速查表 ===")
	fmt.Println("转换方向\t\t方法\t\t\t示例")
	fmt.Println("布尔值 -> 字符串\tstrconv.FormatBool\tstrconv.FormatBool(true) = \"true\"")
	fmt.Println("布尔值 -> 字符串\tfmt.Sprintf(\"%v\")\tfmt.Sprintf(\"%v\", true) = \"true\"")
	fmt.Println("布尔值 -> 数值\t手动 if 语句\t\tif b { num = 1 } else { num = 0 }")
	fmt.Println("字符串 -> 布尔值\tstrconv.ParseBool\tstrconv.ParseBool(\"true\") = true, nil")
	fmt.Println("数值 -> 布尔值\t手动判断\t\tb := n != 0")
	fmt.Println()
	fmt.Println("注意事项:")
	fmt.Println("  1. strconv.ParseBool 支持 \"true\"/\"false\"（大小写不敏感）和 \"1\"/\"0\"")
	fmt.Println("  2. 布尔值转数值需要手动处理，没有直接转换语法")
	fmt.Println("  3. 数值转布尔值：非零值为 true，零值为 false")
	fmt.Println("  4. strconv.FormatBool 性能优于 fmt.Sprintf")
}

// ============================================
// 辅助函数
// ============================================

// BoolToInt 将布尔值转换为整数（true -> 1, false -> 0）
func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

// IntToBool 将整数转换为布尔值（非零 -> true, 零 -> false）
func IntToBool(n int) bool {
	return n != 0
}

// BoolToFloat64 将布尔值转换为浮点数（true -> 1.0, false -> 0.0）
func BoolToFloat64(b bool) float64 {
	if b {
		return 1.0
	}
	return 0.0
}

// Float64ToBool 将浮点数转换为布尔值（非零 -> true, 零 -> false）
func Float64ToBool(f float64) bool {
	return f != 0.0
}

