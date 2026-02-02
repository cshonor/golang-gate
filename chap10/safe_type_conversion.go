// 示例：Go 语言安全类型转换和防止溢出
// 演示类型转换中的常见错误、防止溢出的方法和安全转换模板

package main

import (
	"fmt"
	"math"
)

func main() {
	// ============================================
	// 1. 类型转换中的常见错误
	// ============================================
	fmt.Println("=== 类型转换中的常见错误 ===")

	// 错误1：整数转字符串时的错误
	fmt.Println("错误1：整数转字符串")
	var num int = 65
	// 错误写法：string(num) 会把 ASCII 码 65 对应的字符 'A' 转成字符串
	wrongStr := string(num)
	fmt.Printf("  错误写法: string(%d) = %q (期望: \"65\")\n", num, wrongStr)

	// 正确写法：使用 strconv.Itoa
	// import "strconv"
	// correctStr := strconv.Itoa(num) // "65"
	fmt.Println("  正确写法: strconv.Itoa(num) = \"65\"")

	// 错误2：浮点数转整数时的精度丢失
	fmt.Println("\n错误2：浮点数转整数")
	var f float64 = 3.9
	wrongInt := int(f) // 直接截断小数部分
	fmt.Printf("  错误写法: int(%f) = %d (期望: 4，实际是截断不是四舍五入)\n", f, wrongInt)

	// 正确写法：需要四舍五入时使用 math.Round
	correctInt := int(math.Round(f))
	fmt.Printf("  正确写法: int(math.Round(%f)) = %d\n", f, correctInt)

	// 错误3：大数值溢出转换
	fmt.Println("\n错误3：大数值溢出转换")
	var num32 int32 = 300
	wrongByte := byte(num32) // 300 超出了 byte（0-255）的范围
	fmt.Printf("  错误写法: byte(%d) = %d (发生溢出，300 - 256 = 44)\n", num32, wrongByte)

	// 正确写法：先检查范围
	if num32 >= 0 && num32 <= math.MaxUint8 {
		correctByte := byte(num32)
		fmt.Printf("  正确写法: 先检查范围，byte(%d) = %d\n", num32, correctByte)
	} else {
		fmt.Printf("  正确写法: %d 超出 byte 范围（0~255），不能转换\n", num32)
	}

	fmt.Println()

	// ============================================
	// 2. 防止溢出的安全转换方法
	// ============================================
	fmt.Println("=== 防止溢出的安全转换方法 ===")

	// 示例：安全转换为 uint8
	fmt.Println("安全转换为 uint8:")
	testValues := []int{123, 300, -10, 255, 0}
	for _, v := range testValues {
		if res, err := SafeUint8(v); err != nil {
			fmt.Printf("  SafeUint8(%d): 转换失败 - %v\n", v, err)
		} else {
			fmt.Printf("  SafeUint8(%d): %d\n", v, res)
		}
	}

	fmt.Println()

	// ============================================
	// 3. 各种整数类型的安全转换模板
	// ============================================
	fmt.Println("=== 各种整数类型的安全转换模板 ===")

	// 测试不同整数类型的安全转换
	testInt := 32767
	fmt.Printf("测试值: %d\n", testInt)

	// int8
	if res, err := SafeInt8(testInt); err != nil {
		fmt.Printf("  SafeInt8: %v\n", err)
	} else {
		fmt.Printf("  SafeInt8: %d\n", res)
	}

	// uint8
	if res, err := SafeUint8(testInt); err != nil {
		fmt.Printf("  SafeUint8: %v\n", err)
	} else {
		fmt.Printf("  SafeUint8: %d\n", res)
	}

	// int16
	if res, err := SafeInt16(testInt); err != nil {
		fmt.Printf("  SafeInt16: %v\n", err)
	} else {
		fmt.Printf("  SafeInt16: %d\n", res)
	}

	// uint16
	if res, err := SafeUint16(testInt); err != nil {
		fmt.Printf("  SafeUint16: %v\n", err)
	} else {
		fmt.Printf("  SafeUint16: %d\n", res)
	}

	fmt.Println()

	// ============================================
	// 4. 判断值是否在合法范围内
	// ============================================
	fmt.Println("=== 判断值是否在合法范围内 ===")

	// 方式1：直接数值比较
	v1 := 128
	if v1 >= 0 && v1 <= 255 {
		fmt.Printf("方式1: %d 处于无符号8位整数的合法范围（0~255）\n", v1)
	} else {
		fmt.Printf("方式1: %d 超出无符号8位整数的合法范围\n", v1)
	}

	// 方式2：使用标准库常量（更规范）
	v2 := 300
	if v2 >= 0 && v2 <= math.MaxUint8 {
		fmt.Printf("方式2: %d 处于无符号8位整数的合法范围（0~%d）\n", v2, math.MaxUint8)
	} else {
		fmt.Printf("方式2: %d 超出无符号8位整数的合法范围（0~%d）\n", v2, math.MaxUint8)
	}

	fmt.Println()

	// ============================================
	// 5. math 包提供的整数类型范围常量
	// ============================================
	fmt.Println("=== math 包提供的整数类型范围常量 ===")

	fmt.Printf("int8  范围: %d ~ %d\n", math.MinInt8, math.MaxInt8)
	fmt.Printf("int16 范围: %d ~ %d\n", math.MinInt16, math.MaxInt16)
	fmt.Printf("int32 范围: %d ~ %d\n", math.MinInt32, math.MaxInt32)
	fmt.Printf("int64 范围: %d ~ %d\n", math.MinInt64, math.MaxInt64)
	fmt.Printf("uint8  范围: 0 ~ %d\n", math.MaxUint8)
	fmt.Printf("uint16 范围: 0 ~ %d\n", math.MaxUint16)
	fmt.Printf("uint32 范围: 0 ~ %d\n", math.MaxUint32)
	fmt.Printf("uint64 范围: 0 ~ %v\n", uint64(math.MaxUint64))

	fmt.Println()

	// ============================================
	// 6. 安全转换函数使用示例
	// ============================================
	fmt.Println("=== 安全转换函数使用示例 ===")

	testCases := []struct {
		name string
		val  int
	}{
		{"正常值", 100},
		{"边界值（最大值）", 255},
		{"超出范围", 300},
		{"负数", -10},
		{"零值", 0},
	}

	for _, tc := range testCases {
		fmt.Printf("\n测试: %s = %d\n", tc.name, tc.val)
		if res, err := SafeUint8(tc.val); err != nil {
			fmt.Printf("  结果: 转换失败 - %v\n", err)
		} else {
			fmt.Printf("  结果: 成功转换为 %d\n", res)
		}
	}

	fmt.Println()

	// ============================================
	// 7. 类型转换速查表
	// ============================================
	fmt.Println("=== 类型转换速查表 ===")
	fmt.Println("转换方向\t核心方法/函数\t注意事项")
	fmt.Println("整数 ↔ 浮点数\tT(v)\t浮点数转整数会截断小数，不是四舍五入")
	fmt.Println("整数 ↔ 字符串\tstrconv.Itoa/Atoi\tAtoi 需处理错误")
	fmt.Println("浮点数 ↔ 字符串\tstrconv.FormatFloat/ParseFloat\tFormatFloat 需指定格式和精度")
	fmt.Println("布尔值 ↔ 字符串\tstrconv.FormatBool/ParseBool\tParseBool 仅识别 true/false")
	fmt.Println("接口 → 具体类型\tv.(T) 或 v, ok := v.(T)\t直接断言失败会 panic，安全断言更推荐")
	fmt.Println()
	fmt.Println("防止溢出:")
	fmt.Println("  1. 使用 math 包的常量检查范围")
	fmt.Println("  2. 转换前先判断值是否在目标类型范围内")
	fmt.Println("  3. 使用安全转换函数，返回错误而不是静默溢出")
}

// ============================================
// 安全转换函数模板
// ============================================

// SafeUint8 安全地将任意整数转换为 uint8
func SafeUint8(v int) (uint8, error) {
	if v < 0 || v > math.MaxUint8 {
		return 0, fmt.Errorf("值 %d 超出 uint8 范围（0~%d）", v, math.MaxUint8)
	}
	return uint8(v), nil
}

// SafeInt8 安全地将任意整数转换为 int8
func SafeInt8(v int) (int8, error) {
	if v < math.MinInt8 || v > math.MaxInt8 {
		return 0, fmt.Errorf("值 %d 超出 int8 范围（%d~%d）", v, math.MinInt8, math.MaxInt8)
	}
	return int8(v), nil
}

// SafeUint16 安全地将任意整数转换为 uint16
func SafeUint16(v int) (uint16, error) {
	if v < 0 || v > math.MaxUint16 {
		return 0, fmt.Errorf("值 %d 超出 uint16 范围（0~%d）", v, math.MaxUint16)
	}
	return uint16(v), nil
}

// SafeInt16 安全地将任意整数转换为 int16
func SafeInt16(v int) (int16, error) {
	if v < math.MinInt16 || v > math.MaxInt16 {
		return 0, fmt.Errorf("值 %d 超出 int16 范围（%d~%d）", v, math.MinInt16, math.MaxInt16)
	}
	return int16(v), nil
}

// SafeUint32 安全地将任意整数转换为 uint32
func SafeUint32(v int64) (uint32, error) {
	if v < 0 || v > math.MaxUint32 {
		return 0, fmt.Errorf("值 %d 超出 uint32 范围（0~%d）", v, math.MaxUint32)
	}
	return uint32(v), nil
}

// SafeInt32 安全地将任意整数转换为 int32
func SafeInt32(v int64) (int32, error) {
	if v < math.MinInt32 || v > math.MaxInt32 {
		return 0, fmt.Errorf("值 %d 超出 int32 范围（%d~%d）", v, math.MinInt32, math.MaxInt32)
	}
	return int32(v), nil
}

