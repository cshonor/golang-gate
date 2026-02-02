// 示例：为函数声明新类型
// 演示如何为函数类型创建语义化的类型别名，提高代码可读性

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ============================================
// 1. 定义温度单位类型
// ============================================

type kelvin float64

// ============================================
// 2. 为函数类型声明新的语义化类型
// ============================================

// sensor 为「无参、返回 kelvin 的函数」声明一个新类型
// 这不是定义函数，而是为函数类型创建别名，赋予语义化含义
type sensor func() kelvin

// ============================================
// 3. 实现传感器函数（符合 sensor 类型的签名）
// ============================================

// fakeSensor 模拟传感器：生成150K~300K的随机温度
func fakeSensor() kelvin {
	rand.Seed(time.Now().UnixNano())
	return kelvin(rand.Intn(151) + 150) // 150~300K
}

// realSensor 真实传感器：这里用固定值模拟，实际项目中会连接硬件
func realSensor() kelvin {
	return 290.0 // 模拟真实传感器读数
}

// ============================================
// 4. 改写后的函数，参数为 sensor 类型
// ============================================

// measureTemperature 测量温度（使用语义化的 sensor 类型）
func measureTemperature(samples int, s sensor) {
	for i := 0; i < samples; i++ {
		temp := s()
		fmt.Printf("采样 %d: %.2f K\n", i+1, temp)
	}
}

// ============================================
// 5. 对比：改写前 vs 改写后
// ============================================

// measureTemperatureOld 旧版本：使用模糊的函数类型
func measureTemperatureOld(samples int, s func() kelvin) {
	for i := 0; i < samples; i++ {
		temp := s()
		fmt.Printf("采样 %d: %.2f K\n", i+1, temp)
	}
}

func main() {
	// ============================================
	// 6. 使用新类型的好处
	// ============================================
	fmt.Println("=== 为函数声明新类型的好处 ===")
	fmt.Println()
	fmt.Println("改写前: func measureTemperature(samples int, s func() kelvin)")
	fmt.Println("  问题: s func() kelvin 是模糊的函数类型，可读性差")
	fmt.Println()
	fmt.Println("改写后: func measureTemperature(samples int, s sensor)")
	fmt.Println("  优势: s sensor 语义清晰，一看就知道是传感器函数")
	fmt.Println()

	// ============================================
	// 7. 使用示例
	// ============================================
	fmt.Println("=== 使用示例 ===")

	// 传入 fakeSensor，因为它符合 sensor 类型
	fmt.Println("模拟传感器采样：")
	measureTemperature(3, fakeSensor)

	// 传入 realSensor，同样符合 sensor 类型
	fmt.Println("\n真实传感器采样：")
	measureTemperature(2, realSensor)

	// 也可以传入匿名函数（只要签名匹配）
	fmt.Println("\n匿名函数传感器采样：")
	measureTemperature(2, func() kelvin {
		return kelvin(250.0)
	})

	fmt.Println()

	// ============================================
	// 8. 对比：旧版本 vs 新版本
	// ============================================
	fmt.Println("=== 对比：旧版本 vs 新版本 ===")
	fmt.Println("旧版本调用:")
	measureTemperatureOld(2, fakeSensor)
	fmt.Println("\n新版本调用（更清晰）:")
	measureTemperature(2, fakeSensor)
	fmt.Println()

	// ============================================
	// 9. 带错误处理的扩展版本
	// ============================================
	fmt.Println("=== 带错误处理的扩展版本 ===")

	// 定义带错误处理的传感器类型
	type sensorWithError func() (kelvin, error)

	// 实现带错误处理的传感器
	fakeSensorWithError := func() (kelvin, error) {
		rand.Seed(time.Now().UnixNano())
		temp := kelvin(rand.Intn(151) + 150)
		if temp < 200 {
			return temp, fmt.Errorf("温度过低警告: %.2f K", temp)
		}
		return temp, nil
	}

	realSensorWithError := func() (kelvin, error) {
		return 290.0, nil
	}

	// 带错误处理的测量函数
	measureTemperatureWithError := func(samples int, s sensorWithError) {
		for i := 0; i < samples; i++ {
			temp, err := s()
			if err != nil {
				fmt.Printf("采样 %d: 错误 - %v\n", i+1, err)
			} else {
				fmt.Printf("采样 %d: %.2f K\n", i+1, temp)
			}
		}
	}

	fmt.Println("带错误处理的模拟传感器:")
	measureTemperatureWithError(3, fakeSensorWithError)

	fmt.Println("\n带错误处理的真实传感器:")
	measureTemperatureWithError(2, realSensorWithError)

	fmt.Println()

	// ============================================
	// 10. 更多函数类型别名示例
	// ============================================
	fmt.Println("=== 更多函数类型别名示例 ===")

	// 定义不同的函数类型别名
	type Transformer func(float64) float64
	type Validator func(string) bool
	type Processor func([]int) []int

	// 使用这些类型别名
	celsiusToFahrenheit := func(c float64) float64 {
		return c*1.8 + 32
	}

	isEmail := func(s string) bool {
		return len(s) > 0 && s[len(s)-4:] == ".com"
	}

	doubleNumbers := func(nums []int) []int {
		result := make([]int, len(nums))
		for i, n := range nums {
			result[i] = n * 2
		}
		return result
	}

	// 使用类型别名的函数
	applyTransform := func(value float64, t Transformer) float64 {
		return t(value)
	}

	validateInput := func(input string, v Validator) bool {
		return v(input)
	}

	processData := func(data []int, p Processor) []int {
		return p(data)
	}

	fmt.Printf("温度转换: %.2f °C = %.2f °F\n", 25.0, applyTransform(25.0, celsiusToFahrenheit))
	fmt.Printf("验证邮箱: %t\n", validateInput("test@example.com", isEmail))
	fmt.Printf("处理数据: %v\n", processData([]int{1, 2, 3, 4, 5}, doubleNumbers))

	fmt.Println()

	// ============================================
	// 11. 核心总结
	// ============================================
	fmt.Println("=== 核心总结 ===")
	fmt.Println("1. type sensor func() kelvin 不是定义函数，而是为函数类型创建别名")
	fmt.Println("2. 为函数类型创建别名可以赋予语义化含义，提高代码可读性")
	fmt.Println("3. 改写后函数参数从 s func() kelvin 变成 s sensor，更简洁明确")
	fmt.Println("4. 只要函数签名和类型别名一致，就能作为参数传入")
	fmt.Println("5. 可以创建带错误处理的函数类型别名")
	fmt.Println("6. 函数类型别名可以用于各种场景：转换器、验证器、处理器等")
}

