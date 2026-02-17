// 示例：Go 语言的数组
// 演示数组的声明、初始化、访问、迭代和数组的本质特性

package main

import "fmt"

func main() {
	// ============================================
	// 1. 数组的本质特性
	// ============================================
	fmt.Println("=== 数组的本质特性 ===")

	fmt.Println("1. 定长有序：数组的长度在声明时就已确定，且不可改变")
	fmt.Println("2. 同类型元素：数组中的每个元素都必须是相同的类型")
	fmt.Println("3. 值类型：数组是值类型，赋值或传递给函数时会复制整个数组")
	fmt.Println()

	// ============================================
	// 2. 声明数组
	// ============================================
	fmt.Println("=== 声明数组 ===")

	// 方式1：声明长度为8的字符串数组（零值初始化）

//在 Go 语言里，当你用  var  声明一个数组但不显式初始化时，数组里的每个元素都会被自动设置为对应类型的零值。
- 对于  int  类型，零值是  0 
​
- 对于  float64  类型，零值是  0.0 
​
- 对于  string  类型，零值是空字符串  "" 
​
- 对于布尔类型，零值是  false

	var planets [8]string
	fmt.Printf("planets 类型: %T, 长度: %d\n", planets, len(planets))
	fmt.Printf("planets 零值: %v\n", planets)

	// 方式2：声明并初始化
	var numbers [5]int
	fmt.Printf("numbers 类型: %T, 零值: %v\n", numbers, numbers)

	// 方式3：短变量声明
	var scores [3]float64
	fmt.Printf("scores 类型: %T, 零值: %v\n", scores, scores)

	fmt.Println()

	// ============================================
	// 3. 赋值元素
	// ============================================
	fmt.Println("=== 赋值元素 ===")

	// 为行星数组赋值
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
	// 4. 访问元素
	// ============================================
	fmt.Println("=== 访问元素 ===")

	fmt.Printf("第一颗行星: %s\n", planets[0])
	fmt.Printf("第三颗行星: %s\n", planets[2])
	fmt.Printf("最后一颗行星: %s\n", planets[7])
	fmt.Printf("数组长度: %d\n", len(planets))
	fmt.Println()

	// ============================================
	// 5. 初始化时直接赋值
	// ============================================
	fmt.Println("=== 初始化时直接赋值 ===")

	// 方式1：完整指定长度和值
	dwarfPlanets := [5]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	fmt.Printf("矮行星: %v\n", dwarfPlanets)

	// 方式2：让编译器推断长度
	colors := [...]string{"Red", "Green", "Blue"}
	fmt.Printf("颜色: %v (长度: %d)\n", colors, len(colors))

	// 方式3：指定部分元素，其余为零值
	partialArray := [5]int{1, 2, 3}
	fmt.Printf("部分初始化: %v\n", partialArray)

	// 方式4：指定索引初始化
	indexedArray := [5]string{0: "First", 2: "Third", 4: "Fifth"}
	fmt.Printf("索引初始化: %v\n", indexedArray)

	fmt.Println()

	// ============================================
	// 6. 迭代数组
	// ============================================
	fmt.Println("=== 迭代数组 ===")

	// 方式1：for range 遍历（推荐）
	fmt.Println("太阳系行星列表（for range）:")
	for index, name := range planets {
		fmt.Printf("  %d: %s\n", index+1, name)
	}

	// 方式2：只获取值
	fmt.Println("\n只获取值:")
	for _, name := range planets {
		fmt.Printf("  %s\n", name)
	}

	// 方式3：只获取索引
	fmt.Println("\n只获取索引:")
	for i := range planets {
		fmt.Printf("  索引 %d: %s\n", i, planets[i])
	}

	// 方式4：传统 for 循环
	fmt.Println("\n传统 for 循环:")
	for i := 0; i < len(planets); i++ {
		fmt.Printf("  %d: %s\n", i+1, planets[i])
	}

	fmt.Println()

	// ============================================
	// 7. 数组是值类型
	// ============================================
	fmt.Println("=== 数组是值类型 ===")

	// 数组赋值会复制整个数组
	arr1 := [3]int{1, 2, 3}
	arr2 := arr1 // 复制整个数组
	fmt.Printf("arr1: %v\n", arr1)
	fmt.Printf("arr2: %v\n", arr2)

	// 修改 arr2 不会影响 arr1
	arr2[0] = 100
	fmt.Printf("修改 arr2[0] = 100 后:\n")
	fmt.Printf("arr1: %v (未改变)\n", arr1)
	fmt.Printf("arr2: %v (已改变)\n", arr2)

	// 传递给函数也会复制
	modifyArray(arr1)
	fmt.Printf("传递给函数后 arr1: %v (未改变)\n", arr1)

	fmt.Println()

	// ============================================
	// 8. 数组长度是类型的一部分
	// ============================================
	fmt.Println("=== 数组长度是类型的一部分 ===")

	var arr3 [3]int
	var arr4 [5]int
	fmt.Printf("arr3 类型: %T\n", arr3)
	fmt.Printf("arr4 类型: %T\n", arr4)
	fmt.Println("注意: [3]int 和 [5]int 是不同的类型，不能直接赋值")
	// arr3 = arr4  // ❌ 编译错误：cannot use arr4 (type [5]int) as type [3]int

	fmt.Println()

	// ============================================
	// 9. 多维数组
	// ============================================
	fmt.Println("=== 多维数组 ===")

	// 二维数组
	var matrix [3][3]int
	matrix[0] = [3]int{1, 2, 3}
	matrix[1] = [3]int{4, 5, 6}
	matrix[2] = [3]int{7, 8, 9}

	fmt.Println("二维数组:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}

	// 初始化二维数组
	matrix2 := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("\n初始化二维数组:")
	for _, row := range matrix2 {
		fmt.Printf("  %v\n", row)
	}

	fmt.Println()

	// ============================================
	// 10. 数组与切片的区别
	// ============================================
	fmt.Println("=== 数组与切片的区别 ===")

	fmt.Println("特性对比:")
	fmt.Println("  特性          数组（Array）        切片（Slice）")
	fmt.Println("  长度          固定，声明时确定     动态可变")
	fmt.Println("  类型          长度是类型的一部分   长度不是类型的一部分")
	fmt.Println("  值类型        是，赋值会复制       否，是引用类型")
	fmt.Println("  使用频率      较少                 非常常用")
	fmt.Println("  适用场景      固定大小的数据       动态大小的数据")
	fmt.Println()

	// 数组示例
	arr := [3]int{1, 2, 3}
	fmt.Printf("数组: %v (类型: %T, 长度: %d)\n", arr, arr, len(arr))

	// 切片示例（对比）
	slice := []int{1, 2, 3}
	fmt.Printf("切片: %v (类型: %T, 长度: %d)\n", slice, slice, len(slice))
	fmt.Println("注意: 数组类型包含长度 [3]int，切片类型不包含长度 []int")

	fmt.Println()

	// ============================================
	// 11. 数组的实际应用示例
	// ============================================
	fmt.Println("=== 数组的实际应用示例 ===")

	// 示例1：固定大小的配置
	weekdays := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	fmt.Println("星期:")
	for i, day := range weekdays {
		fmt.Printf("  %d: %s\n", i+1, day)
	}

	// 示例2：固定大小的缓冲区
	var buffer [256]byte
	fmt.Printf("\n缓冲区大小: %d 字节\n", len(buffer))

	// 示例3：固定大小的查找表
	lookupTable := [10]int{0, 1, 4, 9, 16, 25, 36, 49, 64, 81}
	fmt.Println("\n平方数查找表:")
	for i, square := range lookupTable {
		fmt.Printf("  %d^2 = %d\n", i, square)
	}

	fmt.Println()

	// ============================================
	// 12. 数组的局限性
	// ============================================
	fmt.Println("=== 数组的局限性 ===")

	fmt.Println("1. 长度固定：无法动态扩容")
	fmt.Println("2. 值类型：传递大数组会复制整个数组，性能开销大")
	fmt.Println("3. 类型限制：不同长度的数组是不同类型，不能混用")
	fmt.Println("4. 实际开发中，大多数场景使用切片而不是数组")
	fmt.Println()

	// ============================================
	// 13. 总结
	// ============================================
	fmt.Println("=== 总结 ===")
	fmt.Println("1. 数组是定长、有序、同类型元素的集合")
	fmt.Println("2. 数组是值类型，赋值和传递会复制整个数组")
	fmt.Println("3. 数组长度是类型的一部分，[3]int 和 [5]int 是不同的类型")
	fmt.Println("4. 可以用 for range 或传统 for 循环遍历数组")
	fmt.Println("5. 数组适用于固定大小的数据，切片适用于动态大小的数据")
	fmt.Println("6. 实际开发中，切片的使用频率远高于数组")
}

// ============================================
// 辅助函数
// ============================================

// modifyArray 修改数组（演示数组是值类型）
func modifyArray(arr [3]int) {
	arr[0] = 999 // 只修改副本，不影响原数组
	fmt.Printf("函数内部修改后: %v\n", arr)
}

