// 示例：温度类型和方法
// 演示如何声明新类型并为类型绑定方法，展示函数与方法的区别

package main

import "fmt"

// --------------------------
// 1. 声明新类型
// --------------------------
type Kelvin float64
type Celsius float64
type Fahrenheit float64

// --------------------------
// 2. 为类型绑定方法
// --------------------------

// 开尔文转摄氏度
func (k Kelvin) ToCelsius() Celsius {
	return Celsius(k - 273.15)
}

// 摄氏度转开尔文
func (c Celsius) ToKelvin() Kelvin {
	return Kelvin(c + 273.15)
}

// 摄氏度转华氏度
func (c Celsius) ToFahrenheit() Fahrenheit {
	return Fahrenheit(c*1.8 + 32)
}

// 华氏度转摄氏度
func (f Fahrenheit) ToCelsius() Celsius {
	return Celsius((f - 32) / 1.8)
}

func main() {
	// 打印函数与方法的对比
	fmt.Println("==== 函数与方法对比 ====")
	fmt.Println("  定义方式\tfunc 名(参数) 返回值\tfunc (接收者 类型) 名(参数) 返回值")
	fmt.Println("  绑定关系\t独立，不绑定类型\t必须绑定到接收者类型")
	fmt.Println("  调用方式\t函数名(参数)\t\t接收者.方法名(参数)")
	fmt.Println("  核心作用\t通用代码复用\t\t为特定类型增加行为")
	fmt.Println()

	// 声明新类型示例
	fmt.Println("==== 声明新类型 ====")
	var k Kelvin = 300.0
	var c Celsius = 26.85
	var f Fahrenheit = 80.0

	fmt.Printf("Kelvin: %.2f K\n", k)
	fmt.Printf("Celsius: %.2f °C\n", c)
	fmt.Printf("Fahrenheit: %.2f °F\n", f)
	fmt.Println()

	// 方法调用示例
	fmt.Println("==== 方法调用示例 ====")
	fmt.Printf("%.2f K = %.2f °C\n", k, k.ToCelsius())
	fmt.Printf("%.2f °C = %.2f K\n", c, c.ToKelvin())
	fmt.Printf("%.2f °C = %.2f °F\n", c, c.ToFahrenheit())
	fmt.Printf("%.2f °F = %.2f °C\n", f, f.ToCelsius())
}

