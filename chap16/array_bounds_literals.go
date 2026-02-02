// 示例：数组越界检查和复合字面量初始化
// 演示数组越界检查（编译时错误 vs 运行时恐慌）和复合字面量初始化

package main

import "fmt"

func main() {
	// ============================================
	// 1. 数组越界检查（速查16-2解答）
	// ============================================
	fmt.Println("=== 数组越界检查（速查16-2解答）===")

	var planets [8]string
	planets[0] = "Mercury"
	planets[1] = "Venus"
	planets[2] = "Earth"
	planets[3] = "Mars"
	planets[4] = "Jupiter"
	planets[5] = "Saturn"
	planets[6] = "Uranus"
	planets[7] = "Neptune"

	fmt.Printf("planets 数组长度: %d\n", len(planets))
	fmt.Printf("有效索引范围: 0 到 %d\n", len(planets)-1)
	fmt.Println()

	// 问题：访问 planets[11] 会导致编译时错误还是运行时恐慌？
	fmt.Println("速查16-2解答:")
	fmt.Println("  访问 planets[11] 会导致编译时错误")
	fmt.Println("  原因: planets 是长度为8的数组，索引11是编译期就能确定的常量")
	fmt.Println("  编译器会直接发现并报错")
	fmt.Println()

	// 编译时错误示例（注释掉，避免编译错误）
	fmt.Println("编译时错误示例（已注释）:")
	fmt.Println("  planets[8] = \"Pluto\"  // ❌ 编译错误: index 8 out of bounds [0:8]")
	fmt.Println("  planets[11] = \"Planet\" // ❌ 编译错误: index 11 out of bounds [0:8]")
	fmt.Println()

	// 运行时恐慌示例
	fmt.Println("运行时恐慌示例:")
	fmt.Println("  如果索引是变量，会在运行时触发 panic")
	i := 8
	fmt.Printf("  i = %d\n", i)
	fmt.Println("  planets[i] 会在运行时 panic（索引越界）")
	// 取消注释下面的代码会触发运行时 panic
	// fmt.Println(planets[i]) // panic: runtime error: index out of range [8] with length 8
	fmt.Println()

	// ============================================
	// 2. 编译期检查 vs 运行时检查
	// ============================================
	fmt.Println("=== 编译期检查 vs 运行时检查 ===")

	fmt.Println("情况1：常量索引（编译期检查）")
	fmt.Println("  planets[8]  // ❌ 编译错误，编译器能确定越界")
	fmt.Println("  planets[11] // ❌ 编译错误，编译器能确定越界")
	fmt.Println()

	fmt.Println("情况2：变量索引（运行时检查）")
	fmt.Println("  i := 8")
	fmt.Println("  planets[i]  // ⚠️ 运行时 panic，编译器无法确定")
	fmt.Println()

	fmt.Println("Go 对数组越界的检查:")
	fmt.Println("  1. 编译期检查：如果索引是常量，编译器会直接报错")
	fmt.Println("  2. 运行时恐慌：如果索引是变量，会在运行时触发 panic")
	fmt.Println("  3. 安全性：比 C 语言中直接修改非法内存要安全得多")
	fmt.Println()

	// ============================================
	// 3. 复合字面量初始化数组
	// ============================================
	fmt.Println("=== 复合字面量初始化数组 ===")

	// 基础用法：单行初始化
	dwarfs := [5]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	fmt.Printf("矮行星（单行）: %v\n", dwarfs)
	fmt.Printf("数组长度: %d\n", len(dwarfs))
	fmt.Println()

	// 多行初始化（可读性更好）
	planets2 := [8]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	fmt.Println("行星（多行初始化）:")
	for i, planet := range planets2 {
		fmt.Printf("  %d: %s\n", i+1, planet)
	}
	fmt.Println()

	// ============================================
	// 4. 自动推导长度
	// ============================================
	fmt.Println("=== 自动推导长度 ===")

	// 使用 [...] 让 Go 自动根据初始化值的数量推导数组长度
	planets3 := [...]string{
		"Mercury",
		"Venus",
		"Earth",
	}
	fmt.Printf("自动推导长度: %v (长度: %d)\n", planets3, len(planets3))

	numbers := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("自动推导长度: %v (长度: %d)\n", numbers, len(numbers))
	fmt.Println()

	// ============================================
	// 5. 部分初始化
	// ============================================
	fmt.Println("=== 部分初始化 ===")

	// 只初始化前几个元素，其余为零值
	partialArray := [5]int{1, 2, 3}
	fmt.Printf("部分初始化: %v\n", partialArray)
	fmt.Println("  说明: 未初始化的元素自动为零值（0）")
	fmt.Println()

	// ============================================
	// 6. 索引初始化
	// ============================================
	fmt.Println("=== 索引初始化 ===")

	// 指定索引初始化，未指定的索引为零值
	indexedArray := [5]string{
		0: "First",
		2: "Third",
		4: "Fifth",
	}
	fmt.Printf("索引初始化: %v\n", indexedArray)
	fmt.Println("  说明: 索引1和3未指定，自动为零值（空字符串）")
	fmt.Println()

	// ============================================
	// 7. 不同类型数组的复合字面量
	// ============================================
	fmt.Println("=== 不同类型数组的复合字面量 ===")

	// 整数数组
	intArray := [5]int{10, 20, 30, 40, 50}
	fmt.Printf("整数数组: %v\n", intArray)

	// 浮点数数组
	floatArray := [3]float64{3.14, 2.71, 1.41}
	fmt.Printf("浮点数数组: %v\n", floatArray)

	// 布尔数组
	boolArray := [4]bool{true, false, true, true}
	fmt.Printf("布尔数组: %v\n", boolArray)

	// 混合类型（结构体数组，后面会学到）
	type Point struct {
		X, Y int
	}
	points := [3]Point{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Printf("结构体数组: %v\n", points)
	fmt.Println()

	// ============================================
	// 8. 复合字面量的优势
	// ============================================
	fmt.Println("=== 复合字面量的优势 ===")

	fmt.Println("1. 紧凑：可以在单个步骤中完成声明和赋值")
	fmt.Println("2. 清晰：代码更简洁易读")
	fmt.Println("3. 灵活：支持自动推导长度、部分初始化、索引初始化")
	fmt.Println("4. 安全：编译器会检查元素数量和类型")
	fmt.Println()

	// 对比：传统方式 vs 复合字面量
	fmt.Println("对比：传统方式 vs 复合字面量")
	fmt.Println("传统方式:")
	fmt.Println("  var arr [3]int")
	fmt.Println("  arr[0] = 1")
	fmt.Println("  arr[1] = 2")
	fmt.Println("  arr[2] = 3")
	fmt.Println()
	fmt.Println("复合字面量:")
	fmt.Println("  arr := [3]int{1, 2, 3}")
	fmt.Println("  更简洁！")
	fmt.Println()

	// ============================================
	// 9. 数组安全操作清单
	// ============================================
	fmt.Println("=== 数组安全操作清单 ===")
	fmt.Println()
	fmt.Println("1. 越界检查:")
	fmt.Println("   ✅ 使用常量索引时，编译器会检查越界")
	fmt.Println("   ⚠️  使用变量索引时，需要在运行时检查")
	fmt.Println("   ✅ 使用 len() 函数获取数组长度")
	fmt.Println("   ✅ 访问前检查: if i >= 0 && i < len(arr)")
	fmt.Println()
	fmt.Println("2. 初始化:")
	fmt.Println("   ✅ 使用复合字面量: arr := [3]int{1, 2, 3}")
	fmt.Println("   ✅ 自动推导长度: arr := [...]int{1, 2, 3}")
	fmt.Println("   ✅ 多行初始化提高可读性")
	fmt.Println()
	fmt.Println("3. 迭代:")
	fmt.Println("   ✅ 使用 for range: for i, v := range arr")
	fmt.Println("   ✅ 使用传统 for: for i := 0; i < len(arr); i++")
	fmt.Println("   ✅ 避免硬编码索引范围")
	fmt.Println()
	fmt.Println("4. 最佳实践:")
	fmt.Println("   ✅ 总是使用 len() 获取长度，不要硬编码")
	fmt.Println("   ✅ 在循环中使用 len(arr) 作为边界")
	fmt.Println("   ✅ 使用复合字面量初始化，避免逐个赋值")
	fmt.Println("   ✅ 对于动态大小数据，考虑使用切片而不是数组")
	fmt.Println()

	// ============================================
	// 10. 实际应用示例
	// ============================================
	fmt.Println("=== 实际应用示例 ===")

	// 示例1：安全访问数组元素
	safeAccess := func(arr [5]int, index int) (int, bool) {
		if index >= 0 && index < len(arr) {
			return arr[index], true
		}
		return 0, false
	}

	testArray := [5]int{10, 20, 30, 40, 50}
	value, ok := safeAccess(testArray, 2)
	if ok {
		fmt.Printf("安全访问 arr[2]: %d\n", value)
	} else {
		fmt.Println("索引越界")
	}

	value2, ok2 := safeAccess(testArray, 10)
	if ok2 {
		fmt.Printf("安全访问 arr[10]: %d\n", value2)
	} else {
		fmt.Println("索引越界: arr[10] 不存在")
	}

	fmt.Println()

	// ============================================
	// 11. 总结
	// ============================================
	fmt.Println("=== 总结 ===")
	fmt.Println("1. 数组越界检查:")
	fmt.Println("   - 常量索引：编译时错误")
	fmt.Println("   - 变量索引：运行时 panic")
	fmt.Println("2. 复合字面量初始化:")
	fmt.Println("   - 基础用法: [长度]类型{值1, 值2, ...}")
	fmt.Println("   - 自动推导: [...]类型{值1, 值2, ...}")
	fmt.Println("   - 索引初始化: [长度]类型{索引1: 值1, 索引2: 值2}")
	fmt.Println("3. 最佳实践:")
	fmt.Println("   - 使用 len() 获取长度")
	fmt.Println("   - 使用复合字面量初始化")
	fmt.Println("   - 在循环中检查索引范围")
}

