// 示例：同时声明一组变量和一组常量
// 演示使用 var (...) 和 const (...) 的语法

package main

import "fmt"

func main() {
	// ============================================
	// 1. 使用 var (...) 同时声明一组变量
	// ============================================
	fmt.Println("=== 使用 var (...) 声明一组变量 ===")
	
	// 方式1：使用 var (...) 同时声明多个变量
	var (
		earthWeight float64 = 164.0  // 地球体重
		earthAge    int     = 41     // 地球年龄
		name        string  = "万正鹏" // 姓名
		isStudent   bool    = true   // 是否是学生
	)
	
	fmt.Println("地球体重:", earthWeight)
	fmt.Println("地球年龄:", earthAge)
	fmt.Println("姓名:", name)
	fmt.Println("是学生:", isStudent)
	
	fmt.Println()
	
	// ============================================
	// 2. 使用 var (...) 时可以不指定类型（自动推断）
	// ============================================
	fmt.Println("=== var (...) 自动推断类型 ===")
	
	var (
		weight = 70.0    // 自动推断为 float64
		height = 175.0   // 自动推断为 float64
		age    = 25      // 自动推断为 int
		city   = "北京"  // 自动推断为 string
	)
	
	fmt.Println("体重:", weight, "kg")
	fmt.Println("身高:", height, "cm")
	fmt.Println("年龄:", age, "岁")
	fmt.Println("城市:", city)
	
	fmt.Println()
	
	// ============================================
	// 3. 使用 const (...) 同时声明一组常量
	// ============================================
	fmt.Println("=== 使用 const (...) 声明一组常量 ===")
	
	// 使用 const (...) 同时声明多个常量
	const (
		lightSpeed      = 3e8        // 光速：3×10^8 米/秒
		marsGravity     = 0.3783     // 火星重力系数
		earthGravity    = 9.8        // 地球重力
		pi              = 3.14159    // 圆周率
		earthDaysPerYear = 365       // 地球一年天数
		marsDaysPerYear  = 687       // 火星一年天数
	)
	
	fmt.Println("光速:", lightSpeed, "m/s")
	fmt.Println("火星重力系数:", marsGravity)
	fmt.Println("地球重力:", earthGravity, "m/s²")
	fmt.Println("圆周率:", pi)
	fmt.Println("地球一年:", earthDaysPerYear, "天")
	fmt.Println("火星一年:", marsDaysPerYear, "天")
	
	fmt.Println()
	
	// ============================================
	// 4. 对比：分开声明 vs 括号声明
	// ============================================
	fmt.Println("=== 对比：分开声明 vs 括号声明 ===")
	
	// 方式1：分开声明（比较繁琐）
	var a int = 10
	var b string = "hello"
	var c float64 = 3.14
	
	fmt.Println("分开声明:", a, b, c)
	
	// 方式2：使用括号同时声明（更简洁）
	var (
		d int     = 20
		e string  = "world"
		f float64 = 6.28
	)
	
	fmt.Println("括号声明:", d, e, f)
	
	fmt.Println()
	
	// ============================================
	// 5. 实际应用：火星计算（使用括号声明）
	// ============================================
	fmt.Println("=== 实际应用：火星计算 ===")
	
	// 使用 const (...) 声明所有常量
	const (
		marsGravity2     = 0.3783
		earthDaysPerYear2 = 365
		marsDaysPerYear2  = 687
	)
	
	// 使用 var (...) 声明所有变量
	var (
		earthWeight2 = 164.0
		earthAge2    = 41
	)
	
	// 计算
	marsWeight2 := earthWeight2 * marsGravity2
	marsAge2 := earthAge2 * earthDaysPerYear2 / marsDaysPerYear2
	
	fmt.Println("地球体重:", earthWeight2, "磅")
	fmt.Println("火星体重:", marsWeight2, "磅")
	fmt.Println("地球年龄:", earthAge2, "年")
	fmt.Println("火星年龄:", marsAge2, "年")
	
	fmt.Println()
	
	// ============================================
	// 6. 混合使用：部分指定类型，部分自动推断
	// ============================================
	fmt.Println("=== 混合使用 ===")
	
	var (
		score    float64 = 95.5  // 指定类型
		grade    = "A"           // 自动推断为 string
		passed   = true          // 自动推断为 bool
		attempts int    = 3      // 指定类型
	)
	
	fmt.Println("分数:", score)
	fmt.Println("等级:", grade)
	fmt.Println("通过:", passed)
	fmt.Println("尝试次数:", attempts)
	
	fmt.Println()
	
	// ============================================
	// 7. 常量组中的特殊用法：iota
	// ============================================
	fmt.Println("=== 常量组中的 iota（进阶） ===")
	
	// iota 是 Go 语言中的常量计数器
	// 在 const (...) 中，iota 从 0 开始，每行自动递增
	const (
		Sunday    = iota  // 0
		Monday            // 1
		Tuesday           // 2
		Wednesday         // 3
		Thursday          // 4
		Friday            // 5
		Saturday          // 6
	)
	
	fmt.Println("Sunday:", Sunday)
	fmt.Println("Monday:", Monday)
	fmt.Println("Tuesday:", Tuesday)
	
	// 注意：iota 是进阶内容，这里只是简单演示
}

