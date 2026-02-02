// 示例：数组的值类型特性和多维数组
// 演示数组的复制行为、函数传递和多维数组的使用

package main

import "fmt"

func main() {
	// ============================================
	// 1. 速查16-4：使用 range 迭代数组的优势
	// ============================================
	fmt.Println("=== 速查16-4：使用 range 迭代数组的优势 ===")

	planets := [8]string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}

	fmt.Println("问题1：使用 range 迭代数组可以避免哪些错误？")
	fmt.Println("答案：")
	fmt.Println("  1. 避免索引越界错误：range 会自动遍历所有元素")
	fmt.Println("  2. 代码更简洁：无需手动维护循环变量")
	fmt.Println()

	// 方式1：使用 range（推荐）
	fmt.Println("方式1：使用 range（推荐）")
	for i, planet := range planets {
		fmt.Printf("  %d: %s\n", i+1, planet)
	}

	// 方式2：传统 for 循环
	fmt.Println("\n方式2：传统 for 循环")
	for i := 0; i < len(planets); i++ {
		fmt.Printf("  %d: %s\n", i+1, planets[i])
	}

	fmt.Println()
	fmt.Println("问题2：在什么情况下，使用传统的 for 循环比使用 range 更适合？")
	fmt.Println("答案：")
	fmt.Println("  1. 需要定制迭代过程：逆序遍历、按固定步长遍历")
	fmt.Println("  2. 需要直接操作索引，或在循环中修改索引值")
	fmt.Println()

	// 逆序遍历
	fmt.Println("逆序遍历（传统 for 循环）:")
	for i := len(planets) - 1; i >= 0; i-- {
		fmt.Printf("  %d: %s\n", i+1, planets[i])
	}

	// 按固定步长遍历
	fmt.Println("\n每隔一个元素遍历（传统 for 循环）:")
	for i := 0; i < len(planets); i += 2 {
		fmt.Printf("  %d: %s\n", i+1, planets[i])
	}

	fmt.Println()

	// ============================================
	// 2. 速查16-5：数组是值类型（复制行为）
	// ============================================
	fmt.Println("=== 速查16-5：数组是值类型（复制行为）===")

	// 问题1：planetsMarkII 数组的 Earth 元素为何没有被修改？
	planets1 := [...]string{"Mercury", "Venus", "Earth"}
	planetsMarkII := planets1 // 数组赋值会创建完整副本
	planetsMarkII[2] = "whoops"

	fmt.Println("问题1：planetsMarkII 数组的 Earth 元素为何没有被修改？")
	fmt.Printf("原始数组 planets1: %v\n", planets1)
	fmt.Printf("副本数组 planetsMarkII: %v\n", planetsMarkII)
	fmt.Println("答案：数组是值类型，赋值会创建完整副本，修改副本不影响原数组")
	fmt.Println()

	// 问题2：如何让 main 函数中的 planets 数组被修改？
	fmt.Println("问题2：如何让 main 函数中的 planets 数组被修改？")
	fmt.Println("方法1：让函数返回修改后的数组")
	planets2 := [8]string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	planets2 = terraformReturn(planets2)
	fmt.Printf("方法1结果: %v\n", planets2)

	fmt.Println("\n方法2：使用指针传递数组")
	planets3 := [8]string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	terraformPointer(&planets3)
	fmt.Printf("方法2结果: %v\n", planets3)

	fmt.Println("\n方法3：使用切片（推荐，第17章内容）")
	planets4 := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	terraformSlice(planets4)
	fmt.Printf("方法3结果: %v\n", planets4)

	fmt.Println()

	// ============================================
	// 3. 16.5 数组被复制（值类型特性）
	// ============================================
	fmt.Println("=== 16.5 数组被复制（值类型特性）===")

	// 示例1：数组赋值会创建副本
	fmt.Println("示例1：数组赋值会创建副本")
	planets5 := [...]string{"Mercury", "Venus", "Earth"}
	planetsCopy := planets5
	planetsCopy[2] = "whoops"
	fmt.Printf("原始数组: %v\n", planets5)        // 未改变
	fmt.Printf("副本数组: %v\n", planetsCopy)      // 已改变
	fmt.Println()

	// 示例2：函数参数传递数组也会创建副本
	fmt.Println("示例2：函数参数传递数组也会创建副本")
	planets6 := [8]string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	fmt.Printf("调用函数前: %v\n", planets6)
	terraform(planets6) // 传递数组，函数内部修改的是副本
	fmt.Printf("调用函数后: %v (未改变，因为传递的是副本)\n", planets6)
	fmt.Println()

	// 示例3：数组比较
	fmt.Println("示例3：数组比较（值类型可以比较）")
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	arr3 := [3]int{1, 2, 4}
	fmt.Printf("arr1 == arr2: %t\n", arr1 == arr2) // true，值相同
	fmt.Printf("arr1 == arr3: %t\n", arr1 == arr3) // false，值不同
	fmt.Println()

	// ============================================
	// 4. 速查16-6：多维数组（数独游戏）
	// ============================================
	fmt.Println("=== 速查16-6：多维数组（数独游戏）===")

	fmt.Println("问题：如何声明 9×9 的整数网格？")
	fmt.Println("答案：var grid [9][9]int")
	fmt.Println()

	// 声明 9×9 的数独网格
	var sudoku [9][9]int
	fmt.Printf("数独网格类型: %T\n", sudoku)
	fmt.Printf("数独网格大小: %d×%d\n", len(sudoku), len(sudoku[0]))
	fmt.Println()

	// 初始化数独网格（示例：第一行）
	for i := 0; i < 9; i++ {
		sudoku[0][i] = i + 1
	}
	fmt.Println("初始化第一行:")
	for i := 0; i < 9; i++ {
		fmt.Printf("  sudoku[0][%d] = %d\n", i, sudoku[0][i])
	}
	fmt.Println()

	// ============================================
	// 5. 16.6 由数组组成的数组（多维数组）
	// ============================================
	fmt.Println("=== 16.6 由数组组成的数组（多维数组）===")

	// 示例1：8×8 国际象棋棋盘
	fmt.Println("示例1：8×8 国际象棋棋盘")
	var board [8][8]string
	board[0][0] = "r" // 将"车"放置在[0][0]位置
	board[1][1] = "p" // 将"兵"放置在[1][1]位置

	// 遍历棋盘第一行
	for column := range board[0] {
		board[0][column] = "p" // 第一行全部放置"兵"
	}

	fmt.Println("棋盘第一行:")
	for i := 0; i < 8; i++ {
		fmt.Printf("  board[0][%d] = %q\n", i, board[0][i])
	}
	fmt.Println()

	// 示例2：3×3 矩阵
	fmt.Println("示例2：3×3 矩阵")
	matrix := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("矩阵:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
	fmt.Println()

	// 示例3：遍历多维数组
	fmt.Println("示例3：遍历多维数组")
	fmt.Println("使用 range 遍历二维数组:")
	for i, row := range matrix {
		fmt.Printf("  第 %d 行: %v\n", i+1, row)
	}
	fmt.Println()

	// 示例4：三维数组
	fmt.Println("示例4：三维数组（3×3×3）")
	var cube [3][3][3]int
	// 初始化
	value := 1
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				cube[i][j][k] = value
				value++
			}
		}
	}
	fmt.Printf("三维数组类型: %T\n", cube)
	fmt.Printf("三维数组大小: %d×%d×%d\n", len(cube), len(cube[0]), len(cube[0][0]))
	fmt.Printf("第一层: %v\n", cube[0])
	fmt.Println()

	// ============================================
	// 6. 数组值类型的性能考虑
	// ============================================
	fmt.Println("=== 数组值类型的性能考虑 ===")

	fmt.Println("数组是值类型的优缺点:")
	fmt.Println("优点:")
	fmt.Println("  ✅ 数据安全：不会意外修改原始数据")
	fmt.Println("  ✅ 可以比较：相同类型的数组可以直接用 == 比较")
	fmt.Println()
	fmt.Println("缺点:")
	fmt.Println("  ⚠️  性能开销：大数组复制会消耗内存和时间")
	fmt.Println("  ⚠️  内存占用：每个副本都占用完整的内存空间")
	fmt.Println()
	fmt.Println("建议:")
	fmt.Println("  - 小数组（<100元素）：可以直接传递")
	fmt.Println("  - 大数组：使用指针或切片（引用类型）")
	fmt.Println()

	// 演示大数组复制的开销
	smallArray := [5]int{1, 2, 3, 4, 5}
	largeArray := [1000]int{}
	for i := range largeArray {
		largeArray[i] = i
	}

	fmt.Println("小数组复制（开销小）:")
	copy1 := smallArray
	fmt.Printf("  原始: %v\n", smallArray)
	fmt.Printf("  副本: %v\n", copy1)

	fmt.Println("\n大数组复制（开销大）:")
	copy2 := largeArray
	fmt.Printf("  原始数组长度: %d\n", len(largeArray))
	fmt.Printf("  副本数组长度: %d\n", len(copy2))
	fmt.Println("  注意：复制了 1000 个整数，占用更多内存")
	fmt.Println()

	// ============================================
	// 7. 多维数组的实际应用
	// ============================================
	fmt.Println("=== 多维数组的实际应用 ===")

	// 应用1：数独游戏
	fmt.Println("应用1：数独游戏（9×9网格）")
	var sudokuGrid [9][9]int
	// 初始化示例（第一行）
	for i := 0; i < 9; i++ {
		sudokuGrid[0][i] = (i + 1) % 9
		if sudokuGrid[0][i] == 0 {
			sudokuGrid[0][i] = 9
		}
	}
	fmt.Println("数独第一行:")
	for i := 0; i < 9; i++ {
		fmt.Printf("%d ", sudokuGrid[0][i])
	}
	fmt.Println()

	// 应用2：图像像素（简化示例）
	fmt.Println("\n应用2：图像像素（RGB，3×3图像）")
	type RGB struct {
		R, G, B uint8
	}
	var image [3][3]RGB
	// 初始化白色图像
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			image[i][j] = RGB{255, 255, 255}
		}
	}
	fmt.Printf("图像大小: %d×%d\n", len(image), len(image[0]))
	fmt.Printf("左上角像素: R=%d, G=%d, B=%d\n", image[0][0].R, image[0][0].G, image[0][0].B)
	fmt.Println()

	// ============================================
	// 8. 总结
	// ============================================
	fmt.Println("=== 总结 ===")
	fmt.Println("1. range 迭代数组:")
	fmt.Println("   - 避免索引越界错误")
	fmt.Println("   - 代码更简洁")
	fmt.Println("   - 需要定制迭代时用传统 for 循环")
	fmt.Println()
	fmt.Println("2. 数组是值类型:")
	fmt.Println("   - 赋值会创建完整副本")
	fmt.Println("   - 传递给函数也会复制")
	fmt.Println("   - 修改副本不影响原数组")
	fmt.Println("   - 大数组建议使用指针或切片")
	fmt.Println()
	fmt.Println("3. 多维数组:")
	fmt.Println("   - 声明: var grid [9][9]int")
	fmt.Println("   - 表示数组的数组")
	fmt.Println("   - 适用于固定大小的二维结构（棋盘、矩阵等）")
	fmt.Println()
	fmt.Println("4. 最佳实践:")
	fmt.Println("   - 小数组可以直接传递")
	fmt.Println("   - 大数组使用指针或切片")
	fmt.Println("   - 动态大小数据使用切片而不是数组")
}

// ============================================
// 辅助函数
// ============================================

// terraform 修改数组（值传递，不会修改原数组）
func terraform(p [8]string) {
	for i := range p {
		p[i] = "New " + p[i]
	}
	fmt.Printf("函数内部修改后: %v\n", p)
}

// terraformReturn 返回修改后的数组
func terraformReturn(p [8]string) [8]string {
	for i := range p {
		p[i] = "New " + p[i]
	}
	return p
}

// terraformPointer 使用指针修改数组
func terraformPointer(p *[8]string) {
	for i := range p {
		p[i] = "New " + p[i]
	}
}

// terraformSlice 修改切片（引用类型，会修改原切片）
func terraformSlice(p []string) {
	for i := range p {
		p[i] = "New " + p[i]
	}
}

