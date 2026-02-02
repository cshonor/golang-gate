// 示例：常量和变量
// 演示 const 和 var 的定义，以及在 fmt.Println 中的使用

package main

import "fmt"

func main() {
	// ============================================
	// 1. 常量的定义：const
	// ============================================
	fmt.Println("=== 常量定义 ===")
	
	// 定义常量：const 关键字
	// 常量一旦定义，值就不能再改变
	const lightSpeed = 3e8        // 光速：3×10^8 米/秒
	const marsGravity = 0.3783   // 火星重力系数
	const pi = 3.14159           // 圆周率
	
	fmt.Println("光速:", lightSpeed)
	fmt.Println("火星重力系数:", marsGravity)
	fmt.Println("圆周率:", pi)
	
	// 注意：常量不能重新赋值，下面的代码会报错
	// lightSpeed = 4e8  // ❌ 错误：不能给常量重新赋值
	
	fmt.Println()
	
	// ============================================
	// 2. 变量的定义：var
	// ============================================
	fmt.Println("=== 变量定义 ===")
	
	// 方式1：使用 var 关键字，指定类型
	var distance float64 = 5.04e10  // 距离：5.04×10^10 米（地球到太阳的距离）
	var earthWeight float64 = 164.0  // 地球体重：164 磅
	
	fmt.Println("距离:", distance)
	fmt.Println("地球体重:", earthWeight)
	
	// 变量可以重新赋值
	earthWeight = 160.0  // ✅ 正确：变量可以重新赋值
	fmt.Println("新的地球体重:", earthWeight)
	
	fmt.Println()
	
	// ============================================
	// 3. 简化的变量定义：:=
	// ============================================
	fmt.Println("=== 简化变量定义 := ===")
	
	// 方式2：使用 := 简化定义（Go 会自动判断类型）
	distance2 := 5.04e10  // 自动判断为 float64
	age := 25             // 自动判断为 int
	name := "万正鹏"       // 自动判断为 string
	
	fmt.Println("距离:", distance2)
	fmt.Println("年龄:", age)
	fmt.Println("名字:", name)
	
	fmt.Println()
	
	// ============================================
	// 4. fmt.Println 中的运算
	// ============================================
	fmt.Println("=== Println 中的运算 ===")
	
	// 在 Println 中可以直接进行运算
	// 计算时间：距离 ÷ 光速
	fmt.Println("光从太阳到地球需要:", distance/lightSpeed, "秒")
	
	// 计算火星体重
	marsWeight := earthWeight * marsGravity
	fmt.Println("火星体重:", marsWeight, "磅")
	
	// 多个运算
	fmt.Println("计算结果:", 10+5, "和", 20-3)
	
	fmt.Println()
	
	// ============================================
	// 5. fmt.Println 自动加空格
	// ============================================
	fmt.Println("=== Println 自动加空格 ===")
	
	// Println 会在每个参数之间自动加一个空格
	fmt.Println(168, "seconds")              // 输出：168 seconds（中间有空格）
	fmt.Println("火星", "体重：", 62.04)      // 输出：火星 体重： 62.04（每个参数之间都有空格）
	
	// 对比：fmt.Print 不会自动加空格
	fmt.Print("使用 Print: ", 168, "seconds\n")  // 输出：使用 Print: 168seconds（没有空格）
	
	fmt.Println()
	
	// ============================================
	// 6. 实际应用：光速计算示例
	// ============================================
	fmt.Println("=== 光速计算示例 ===")
	
	// 定义常量：光速
	const lightSpeed2 = 3e8  // 3×10^8 米/秒
	
	// 定义变量：距离
	var distance3 = 5.04e10  // 地球到太阳的距离：5.04×10^10 米
	
	// 在 Println 中直接计算并打印
	// 注意：distance / lightSpeed 的结果会自动和 "seconds" 拼接，中间有空格
	fmt.Println(distance3/lightSpeed2, "seconds")
	// 输出：168 seconds
	
	// 也可以这样写，更清晰
	time := distance3 / lightSpeed2
	fmt.Println("光从太阳到地球需要", time, "秒")
	
	fmt.Println()
	
	// ============================================
	// 7. 常量和变量的混合使用
	// ============================================
	fmt.Println("=== 常量和变量混合使用 ===")
	
	// 常量
	const earthDaysPerYear = 365
	const marsDaysPerYear = 687
	
	// 变量
	earthAge := 41
	
	// 在 Println 中使用常量和变量进行运算
	marsAge := earthAge * earthDaysPerYear / marsDaysPerYear
	fmt.Println("地球年龄:", earthAge, "年")
	fmt.Println("火星年龄:", marsAge, "年")
	fmt.Println("在火星上，你会年轻", earthAge-marsAge, "年")
	
	fmt.Println()
	
	// ============================================
	// 8. 多个变量和常量的定义
	// ============================================
	fmt.Println("=== 多个变量/常量定义 ===")
	
	// 多个常量
	const (
		gravityEarth = 9.8
		gravityMars  = 3.7
		gravityMoon  = 1.6
	)
	
	// 多个变量
	var (
		weight  = 70.0
		height  = 175.0
		age2    = 25
	)
	
	fmt.Println("地球重力:", gravityEarth, "m/s²")
	fmt.Println("火星重力:", gravityMars, "m/s²")
	fmt.Println("月球重力:", gravityMoon, "m/s²")
	fmt.Println("体重:", weight, "kg, 身高:", height, "cm, 年龄:", age2, "岁")
}


