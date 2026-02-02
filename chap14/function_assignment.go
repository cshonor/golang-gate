// 示例：将函数赋值给变量
// 演示"赋值函数本身"和"赋值函数返回值"的区别

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ============================================
// 温度类型定义
// ============================================

type kelvin float64

// ============================================
// 传感器函数
// ============================================

// fakeSensor 模拟传感器：生成150K~300K的随机温度
func fakeSensor() kelvin {
	rand.Seed(time.Now().UnixNano())
	return kelvin(rand.Intn(151) + 150)
}

// realSensor 真实传感器：这里用固定值模拟，实际项目中会连接硬件
func realSensor() kelvin {
	// 实际项目中这里会是硬件读取逻辑
	return 290.0
}

func main() {
	// ============================================
	// 1. 核心区别：赋值函数本身 vs 赋值函数返回值
	// ============================================
	fmt.Println("=== 核心区别：赋值函数本身 vs 赋值函数返回值 ===")
	fmt.Println()

	// 方式1：赋值函数本身（不带括号）
	fmt.Println("方式1：赋值函数本身（不带括号）")
	fmt.Println("  代码: sensor := fakeSensor")
	sensor := fakeSensor
	fmt.Printf("  sensor 的类型: %T\n", sensor)
	fmt.Printf("  sensor 的值: %v (函数地址)\n", sensor)
	fmt.Printf("  调用 sensor(): %.2f K\n", sensor())
	fmt.Println("  说明: sensor 是一个函数变量，可以多次调用")
	fmt.Printf("  再次调用 sensor(): %.2f K (可能不同，因为是随机值)\n", sensor())
	fmt.Println()

	// 方式2：赋值函数返回值（带括号）
	fmt.Println("方式2：赋值函数返回值（带括号）")
	fmt.Println("  代码: temperature := fakeSensor()")
	temperature := fakeSensor()
	fmt.Printf("  temperature 的类型: %T\n", temperature)
	fmt.Printf("  temperature 的值: %.2f K\n", temperature)
	fmt.Println("  说明: temperature 是一个 kelvin 类型的值，不是函数")
	fmt.Println("  注意: temperature 不能调用，因为它不是函数")
	fmt.Println()

	// ============================================
	// 2. 直观对比示例
	// ============================================
	fmt.Println("=== 直观对比示例 ===")

	// 定义一个简单的加法函数
	add := func(a, b int) int {
		return a + b
	}

	// 对比1：赋值函数本身
	fmt.Println("对比1：赋值函数本身")
	f := add
	fmt.Printf("  f 的类型: %T\n", f)
	fmt.Printf("  f 的值: %v (函数地址)\n", f)
	result1 := f(2, 3)
	fmt.Printf("  调用 f(2, 3): %d\n", result1)
	fmt.Println("  可以多次调用: f(5, 6) =", f(5, 6))
	fmt.Println()

	// 对比2：赋值函数返回值
	fmt.Println("对比2：赋值函数返回值")
	v := add(2, 3)
	fmt.Printf("  v 的类型: %T\n", v)
	fmt.Printf("  v 的值: %d\n", v)
	fmt.Println("  注意: v 是一个整数，不能调用")
	// v(5, 6)  // ❌ 编译错误：cannot call non-function v
	fmt.Println()

	// ============================================
	// 3. 传感器切换示例
	// ============================================
	fmt.Println("=== 传感器切换示例 ===")

	// 将函数赋值给变量，实现可互换的传感器逻辑
	fmt.Println("1. 初始使用模拟传感器")
	sensor2 := fakeSensor
	fmt.Printf("   传感器读数: %.2f K\n", sensor2())

	fmt.Println("2. 切换到真实传感器")
	sensor2 = realSensor // 重新赋值，指向不同的函数
	fmt.Printf("   传感器读数: %.2f K\n", sensor2())

	fmt.Println("3. 再次切换回模拟传感器")
	sensor2 = fakeSensor
	fmt.Printf("   传感器读数: %.2f K\n", sensor2())
	fmt.Println()

	// ============================================
	// 4. 为什么这样设计？
	// ============================================
	fmt.Println("=== 为什么这样设计？ ===")
	fmt.Println("✅ 可互换性: 可以在运行时轻松切换传感器实现")
	fmt.Println("✅ 解耦逻辑: 调用方只关心统一的接口 sensor()")
	fmt.Println("✅ 测试友好: 开发时用 fakeSensor，生产环境用 realSensor")
	fmt.Println()

	// ============================================
	// 5. 类型检查
	// ============================================
	fmt.Println("=== 类型检查 ===")

	// 检查函数变量的类型
	var sensorFunc func() kelvin
	sensorFunc = fakeSensor
	fmt.Printf("sensorFunc 的类型: %T\n", sensorFunc)
	fmt.Printf("sensorFunc 可以调用: %.2f K\n", sensorFunc())

	// 检查返回值的类型
	var temp kelvin
	temp = fakeSensor()
	fmt.Printf("temp 的类型: %T\n", temp)
	fmt.Printf("temp 的值: %.2f K\n", temp)
	fmt.Println("temp 不能调用，因为它是 kelvin 类型，不是函数")
	fmt.Println()

	// ============================================
	// 6. 总结
	// ============================================
	fmt.Println("=== 总结 ===")
	fmt.Println("写法对比:")
	fmt.Println("  sensor := fakeSensor      // 赋值函数本身，sensor 是函数变量")
	fmt.Println("  temp := fakeSensor()      // 赋值函数返回值，temp 是 kelvin 值")
	fmt.Println()
	fmt.Println("一句话总结:")
	fmt.Println("  不带括号：传递的是\"函数本身\"，相当于把一个\"工具\"给变量")
	fmt.Println("  带括号：传递的是\"函数运行后的结果\"，相当于把工具生产出来的\"产品\"给变量")
}

