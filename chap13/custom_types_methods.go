// 示例：基于同一基础类型创建多个自定义类型并绑定方法
// 演示方法的本质：方法必须绑定到自定义类型，不能绑定到基础类型

package main

import "fmt"

// ============================================
// 1. 基于同一个基础类型 int 创建多个自定义类型
// ============================================

// Age 年龄类型（基于 int）
type Age int

// Score 分数类型（基于 int）
type Score int

// Distance 距离类型（基于 int，单位：米）
type Distance int

// ============================================
// 2. 为不同的自定义类型绑定不同的方法
// ============================================

// --------------------------
// Age 类型的方法
// --------------------------

// IsAdult 判断是否成年（值接收者）
func (a Age) IsAdult() bool {
	return a >= 18
}

// GetCategory 获取年龄分类（值接收者）
func (a Age) GetCategory() string {
	if a < 13 {
		return "儿童"
	} else if a < 18 {
		return "青少年"
	} else if a < 60 {
		return "成年人"
	} else {
		return "老年人"
	}
}

// --------------------------
// Score 类型的方法
// --------------------------

// GetGrade 获取等级（值接收者）
func (s Score) GetGrade() string {
	if s >= 90 {
		return "优秀"
	} else if s >= 80 {
		return "良好"
	} else if s >= 60 {
		return "及格"
	} else {
		return "不及格"
	}
}

// IsPass 判断是否及格（值接收者）
func (s Score) IsPass() bool {
	return s >= 60
}

// Add 增加分数（指针接收者，修改原始值）
func (s *Score) Add(points int) {
	*s += Score(points)
}

// --------------------------
// Distance 类型的方法
// --------------------------

// ToKilometers 转换为千米（值接收者）
func (d Distance) ToKilometers() float64 {
	return float64(d) / 1000.0
}

// ToMiles 转换为英里（值接收者）
func (d Distance) ToMiles() float64 {
	return float64(d) / 1609.34
}

// Add 增加距离（指针接收者，修改原始值）
func (d *Distance) Add(meters int) {
	*d += Distance(meters)
}

func main() {
	// ============================================
	// 3. 演示：为什么不能直接给基础类型绑定方法
	// ============================================
	fmt.Println("=== 为什么不能直接给基础类型绑定方法 ===")
	fmt.Println("❌ 错误示例：")
	fmt.Println("  func (i int) Double() int { ... }  // 编译错误！")
	fmt.Println()
	fmt.Println("✅ 正确做法：")
	fmt.Println("  1. 先声明自定义类型: type MyInt int")
	fmt.Println("  2. 再为自定义类型绑定方法: func (m MyInt) Double() MyInt { ... }")
	fmt.Println()

	// ============================================
	// 4. 使用 Age 类型的方法
	// ============================================
	fmt.Println("=== Age 类型的方法 ===")
	age1 := Age(25)
	age2 := Age(15)
	age3 := Age(65)

	fmt.Printf("年龄 %d: 是否成年? %t, 分类: %s\n", age1, age1.IsAdult(), age1.GetCategory())
	fmt.Printf("年龄 %d: 是否成年? %t, 分类: %s\n", age2, age2.IsAdult(), age2.GetCategory())
	fmt.Printf("年龄 %d: 是否成年? %t, 分类: %s\n", age3, age3.IsAdult(), age3.GetCategory())
	fmt.Println()

	// ============================================
	// 5. 使用 Score 类型的方法
	// ============================================
	fmt.Println("=== Score 类型的方法 ===")
	score1 := Score(95)
	score2 := Score(75)
	score3 := Score(45)

	fmt.Printf("分数 %d: 等级: %s, 是否及格: %t\n", score1, score1.GetGrade(), score1.IsPass())
	fmt.Printf("分数 %d: 等级: %s, 是否及格: %t\n", score2, score2.GetGrade(), score2.IsPass())
	fmt.Printf("分数 %d: 等级: %s, 是否及格: %t\n", score3, score3.GetGrade(), score3.IsPass())

	// 使用指针接收者方法修改分数
	fmt.Printf("\n修改前: %d\n", score2)
	score2.Add(10) // 增加 10 分
	fmt.Printf("增加 10 分后: %d\n", score2)
	fmt.Println()

	// ============================================
	// 6. 使用 Distance 类型的方法
	// ============================================
	fmt.Println("=== Distance 类型的方法 ===")
	distance1 := Distance(5000)  // 5000 米
	distance2 := Distance(10000) // 10000 米

	fmt.Printf("距离 %d 米 = %.2f 千米 = %.2f 英里\n",
		distance1, distance1.ToKilometers(), distance1.ToMiles())
	fmt.Printf("距离 %d 米 = %.2f 千米 = %.2f 英里\n",
		distance2, distance2.ToKilometers(), distance2.ToMiles())

	// 使用指针接收者方法修改距离
	fmt.Printf("\n修改前: %d 米\n", distance1)
	distance1.Add(2000) // 增加 2000 米
	fmt.Printf("增加 2000 米后: %d 米 = %.2f 千米\n", distance1, distance1.ToKilometers())
	fmt.Println()

	// ============================================
	// 7. 演示：虽然基于同一基础类型，但类型不同
	// ============================================
	fmt.Println("=== 类型安全：不同自定义类型不能直接混用 ===")
	var age Age = 25
	var score Score = 95
	var distance Distance = 5000

	// 错误：不能直接混用
	// var result = age + score  // ❌ 编译错误：mismatched types Age and Score

	// 正确：需要类型转换
	fmt.Printf("Age: %d, Score: %d, Distance: %d\n", age, score, distance)
	fmt.Println("注意：虽然都是基于 int，但它们是不同的类型，不能直接运算")
	fmt.Println("需要类型转换才能运算：int(age) + int(score)")
	fmt.Println()

	// ============================================
	// 8. 值接收者 vs 指针接收者对比
	// ============================================
	fmt.Println("=== 值接收者 vs 指针接收者对比 ===")

	// 值接收者：不修改原始值
	score4 := Score(80)
	fmt.Printf("原始分数: %d\n", score4)
	grade := score4.GetGrade() // 值接收者方法，不修改原始值
	fmt.Printf("调用值接收者方法 GetGrade() 后: %d (未改变), 等级: %s\n", score4, grade)

	// 指针接收者：修改原始值
	score5 := Score(80)
	fmt.Printf("原始分数: %d\n", score5)
	score5.Add(10) // 指针接收者方法，修改原始值
	fmt.Printf("调用指针接收者方法 Add(10) 后: %d (已改变)\n", score5)
	fmt.Println()

	// ============================================
	// 9. 总结
	// ============================================
	fmt.Println("=== 总结 ===")
	fmt.Println("1. 方法必须绑定到自定义类型，不能绑定到基础类型")
	fmt.Println("2. 可以基于同一个基础类型创建多个自定义类型")
	fmt.Println("3. 不同的自定义类型可以绑定不同的方法")
	fmt.Println("4. 值接收者：不修改原始值，适合查询操作")
	fmt.Println("5. 指针接收者：可以修改原始值，适合修改操作")
	fmt.Println("6. 不同的自定义类型不能直接混用，需要类型转换")
	fmt.Println("7. 这种设计提供了类型安全和封装性")
}

