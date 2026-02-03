// 示例：Go 语言切片的进阶知识点
// 切片的复合字面量、数组 vs 切片的类型区别、函数参数通用性、切片副本与底层数组

package main

import (
	"fmt"
	"strings"
)

func main() {
	// ============================================
	// 1. 切片的复合字面量
	// ============================================
	fmt.Println("=== 1. 切片的复合字面量 ===")

	fmt.Println("可以直接用字面量创建切片，而无需先定义数组")
	fmt.Println("语法: []Type{value1, value2, ...}")
	fmt.Println()

	// 使用复合字面量创建切片
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	fmt.Printf("切片字面量: %v\n", dwarfs)
	fmt.Printf("类型: %T\n", dwarfs)
	fmt.Printf("长度: %d, 容量: %d\n", len(dwarfs), cap(dwarfs))
	fmt.Println()

	fmt.Println("说明:")
	fmt.Println("  这会在底层创建一个包含5个元素的数组")
	fmt.Println("  再生成一个指向该数组的切片")
	fmt.Println()

	// 对比：数组字面量 vs 切片字面量
	fmt.Println("对比：数组字面量 vs 切片字面量")
	arrayLiteral := [...]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	sliceLiteral := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	fmt.Printf("  数组字面量: %T - %v\n", arrayLiteral, arrayLiteral)
	fmt.Printf("  切片字面量: %T - %v\n", sliceLiteral, sliceLiteral)
	fmt.Println()

	// 其他类型的切片字面量
	fmt.Println("其他类型的切片字面量:")
	intSlice := []int{1, 2, 3, 4, 5}
	floatSlice := []float64{3.14, 2.71, 1.41}
	boolSlice := []bool{true, false, true}
	fmt.Printf("  整数切片: %v (类型: %T)\n", intSlice, intSlice)
	fmt.Printf("  浮点数切片: %v (类型: %T)\n", floatSlice, floatSlice)
	fmt.Printf("  布尔切片: %v (类型: %T)\n", boolSlice, boolSlice)
	fmt.Println()

	// ============================================
	// 2. 数组 vs 切片的类型区别
	// ============================================
	fmt.Println("=== 2. 数组 vs 切片的类型区别 ===")

	fmt.Println("核心区别:")
	fmt.Println("  - 数组的类型包含长度信息，例如 [5]string")
	fmt.Println("  - 切片的类型不包含长度信息，仅为 []string")
	fmt.Println()

	// 演示类型区别
	dwarfArray := [...]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dwarfsSlice := dwarfArray[:]
	fmt.Printf("数组类型: %T\n", dwarfArray)  // 输出：array [5]string
	fmt.Printf("切片类型: %T\n", dwarfsSlice) // 输出：slice []string
	fmt.Println()

	// 类型不匹配示例
	fmt.Println("类型不匹配示例:")
	array5 := [5]string{"a", "b", "c", "d", "e"}
	array3 := [3]string{"x", "y", "z"}
	fmt.Printf("  [5]string: %T\n", array5)
	fmt.Printf("  [3]string: %T\n", array3)
	fmt.Println("  说明: [5]string 和 [3]string 是不同的类型，不能互相赋值")
	fmt.Println()

	// 切片类型统一
	fmt.Println("切片类型统一:")
	slice1 := []string{"a", "b", "c"}
	slice2 := []string{"x", "y"}
	slice3 := []string{"only"}
	fmt.Printf("  切片1: %T (长度: %d)\n", slice1, len(slice1))
	fmt.Printf("  切片2: %T (长度: %d)\n", slice2, len(slice2))
	fmt.Printf("  切片3: %T (长度: %d)\n", slice3, len(slice3))
	fmt.Println("  说明: 所有 []string 切片都是同一类型，无论长度如何")
	fmt.Println()

	// ============================================
	// 3. 切片的威力：函数参数与通用性
	// ============================================
	fmt.Println("=== 3. 切片的威力：函数参数与通用性 ===")

	fmt.Println("切片比数组更通用，因为它的类型不绑定长度")
	fmt.Println("可以将任意长度的同类型切片传给同一个函数")
	fmt.Println("而数组必须严格匹配长度")
	fmt.Println()

	// 示例函数：处理任意长度的字符串切片
	fmt.Println("示例函数：处理任意长度的字符串切片")
	worlds1 := []string{"  Mercury  ", "  Venus  ", "  Earth  "}
	worlds2 := []string{"  Mars  ", "  Jupiter  "}
	worlds3 := []string{"  Saturn  "}

	fmt.Printf("处理前 worlds1: %v\n", worlds1)
	hyperspace(worlds1)
	fmt.Printf("处理后 worlds1: %v\n", worlds1)
	fmt.Println()

	fmt.Printf("处理前 worlds2: %v\n", worlds2)
	hyperspace(worlds2)
	fmt.Printf("处理后 worlds2: %v\n", worlds2)
	fmt.Println()

	fmt.Printf("处理前 worlds3: %v\n", worlds3)
	hyperspace(worlds3)
	fmt.Printf("处理后 worlds3: %v\n", worlds3)
	fmt.Println()

	// 对比：数组函数 vs 切片函数
	fmt.Println("对比：数组函数 vs 切片函数")
	fmt.Println("  数组函数:")
	fmt.Println("    func processArray(arr [5]string) { ... }")
	fmt.Println("    只能处理长度为5的数组")
	fmt.Println()
	fmt.Println("  切片函数:")
	fmt.Println("    func processSlice(slice []string) { ... }")
	fmt.Println("    可以处理任意长度的切片")
	fmt.Println()

	// 演示通用性
	fmt.Println("演示通用性:")
	differentLengths := [][]string{
		{"a"},
		{"a", "b"},
		{"a", "b", "c"},
		{"a", "b", "c", "d", "e"},
	}
	for i, slice := range differentLengths {
		fmt.Printf("  切片 %d (长度 %d): ", i+1, len(slice))
		hyperspace(slice)
		fmt.Printf("%v\n", slice)
	}
	fmt.Println()

	// ============================================
	// 4. 切片副本与底层数组
	// ============================================
	fmt.Println("=== 4. 切片副本与底层数组 ===")

	fmt.Println("当你把一个切片赋值给另一个变量（或作为参数传递）")
	fmt.Println("它们会共享同一个底层数组")
	fmt.Println("修改其中一个切片的元素，会影响所有共享该数组的切片")
	fmt.Println()

	// 演示切片赋值共享底层数组
	fmt.Println("演示切片赋值共享底层数组:")
	original := []string{"Mercury", "Venus", "Earth", "Mars"}
	copy := original
	fmt.Printf("  原始切片: %v\n", original)
	fmt.Printf("  副本切片: %v\n", copy)
	fmt.Println()

	// 修改副本的元素
	fmt.Println("修改副本的元素:")
	copy[0] = "修改的元素"
	fmt.Printf("  原始切片: %v (也被修改了)\n", original)
	fmt.Printf("  副本切片: %v\n", copy)
	fmt.Println("  说明: 两个切片共享同一个底层数组")
	fmt.Println()

	// 演示函数参数传递
	fmt.Println("演示函数参数传递:")
	testSlice := []string{"  A  ", "  B  ", "  C  "}
	fmt.Printf("  调用函数前: %v\n", testSlice)
	hyperspace(testSlice)
	fmt.Printf("  调用函数后: %v (被修改了)\n", testSlice)
	fmt.Println("  说明: 函数参数传递切片时，也共享底层数组")
	fmt.Println()

	// 修改切片的指针（不共享）
	fmt.Println("修改切片的指针（不共享）:")
	worlds := []string{"Mercury", "Venus", "Earth", "Mars"}
	worlds2_copy := worlds
	fmt.Printf("  原始: %v\n", worlds)
	fmt.Printf("  副本: %v\n", worlds2_copy)
	fmt.Println()

	// 修改指针（重新切片）
	worlds = worlds[1:] // 修改指针，不共享
	fmt.Printf("  修改原始切片的指针后:\n")
	fmt.Printf("    原始: %v (指向新位置)\n", worlds)
	fmt.Printf("    副本: %v (未受影响)\n", worlds2_copy)
	fmt.Println("  说明: 只修改切片的指针（如 worlds = worlds[1:]），不会影响其他切片的指针")
	fmt.Println()

	// 详细说明
	fmt.Println("详细说明:")
	fmt.Println("  1. 切片赋值: 复制切片头（指针、长度、容量），共享底层数组")
	fmt.Println("  2. 修改元素: 会影响所有共享底层数组的切片")
	fmt.Println("  3. 修改指针: 只影响当前切片，不影响其他切片")
	fmt.Println()

	// ============================================
	// 5. 数组 vs 切片的完整对比
	// ============================================
	fmt.Println("=== 5. 数组 vs 切片的完整对比 ===")

	fmt.Println("特性对比表:")
	fmt.Println("  特性              数组              切片")
	fmt.Println("  ──────────────────────────────────────────────")
	fmt.Println("  类型包含长度      ✅ [5]string      ❌ []string")
	fmt.Println("  长度固定          ✅ 是              ❌ 否（动态）")
	fmt.Println("  值类型            ✅ 是              ❌ 否（引用类型）")
	fmt.Println("  赋值/传递         ✅ 复制整个数组    ❌ 复制切片头")
	fmt.Println("  共享底层数组      ❌ 否              ✅ 是")
	fmt.Println("  函数参数通用性    ❌ 需匹配长度      ✅ 任意长度")
	fmt.Println("  内存占用          ⚠️  固定           ⚠️  动态")
	fmt.Println()

	// 类型检查示例
	fmt.Println("类型检查示例:")
	var arr5 [5]string
	var arr3 [3]string
	var slice []string

	fmt.Printf("  arr5 类型: %T\n", arr5)
	fmt.Printf("  arr3 类型: %T\n", arr3)
	fmt.Printf("  slice 类型: %T\n", slice)
	fmt.Println("  说明:")
	fmt.Println("    - [5]string 和 [3]string 是不同的类型（编译时就会报错）")
	fmt.Println("    - [5]string 和 []string 是不同的类型")
	fmt.Println("    - []string 和 []string 是相同的类型（无论长度如何）")
	fmt.Println("    - 数组类型包含长度，切片类型不包含长度")
	fmt.Println()

	// 函数参数示例
	fmt.Println("函数参数示例:")
	fmt.Println("  数组函数:")
	fmt.Println("    func process(arr [5]string) { ... }")
	fmt.Println("    只能接受长度为5的数组")
	fmt.Println()
	fmt.Println("  切片函数:")
	fmt.Println("    func process(slice []string) { ... }")
	fmt.Println("    可以接受任意长度的切片")
	fmt.Println()

	// ============================================
	// 6. 实际应用示例
	// ============================================
	fmt.Println("=== 6. 实际应用示例 ===")

	// 示例1：处理不同长度的数据
	fmt.Println("示例1：处理不同长度的数据")
	processData([]int{1, 2, 3})
	processData([]int{10, 20, 30, 40, 50})
	processData([]int{100})
	fmt.Println()

	// 示例2：字符串处理
	fmt.Println("示例2：字符串处理")
	names := []string{"  Alice  ", "  Bob  ", "  Charlie  "}
	fmt.Printf("处理前: %v\n", names)
	trimNames(names)
	fmt.Printf("处理后: %v\n", names)
	fmt.Println()

	// 示例3：切片共享的注意事项
	fmt.Println("示例3：切片共享的注意事项")
	data := []int{1, 2, 3, 4, 5}
	sub1 := data[1:4]
	sub2 := data[2:5]
	fmt.Printf("  原始数据: %v\n", data)
	fmt.Printf("  子切片1: %v\n", sub1)
	fmt.Printf("  子切片2: %v\n", sub2)
	sub1[0] = 999
	fmt.Printf("  修改 sub1[0] 后:\n")
	fmt.Printf("    原始数据: %v (也被修改了)\n", data)
	fmt.Printf("    子切片1: %v\n", sub1)
	fmt.Printf("    子切片2: %v (也看到了变化)\n", sub2)
	fmt.Println()

	// ============================================
	// 7. 总结
	// ============================================
	fmt.Println("=== 总结 ===")
	fmt.Println()
	fmt.Println("1. 切片的复合字面量:")
	fmt.Println("   ✅ 可以直接用 []Type{values} 创建切片")
	fmt.Println("   ✅ 无需先定义数组")
	fmt.Println("   ✅ 底层会自动创建数组")
	fmt.Println()
	fmt.Println("2. 数组 vs 切片的类型区别:")
	fmt.Println("   ✅ 数组类型包含长度: [5]string")
	fmt.Println("   ✅ 切片类型不包含长度: []string")
	fmt.Println("   ✅ 数组长度不同是不同类型")
	fmt.Println("   ✅ 切片长度不同是同一类型")
	fmt.Println()
	fmt.Println("3. 切片的威力（函数参数通用性）:")
	fmt.Println("   ✅ 可以处理任意长度的同类型切片")
	fmt.Println("   ✅ 数组必须严格匹配长度")
	fmt.Println("   ✅ 切片更适合作为函数参数")
	fmt.Println()
	fmt.Println("4. 切片副本与底层数组:")
	fmt.Println("   ✅ 切片赋值时共享底层数组")
	fmt.Println("   ✅ 修改元素会影响所有共享的切片")
	fmt.Println("   ✅ 修改指针只影响当前切片")
	fmt.Println()
	fmt.Println("5. 最佳实践:")
	fmt.Println("   ✅ 优先使用切片而不是数组")
	fmt.Println("   ✅ 函数参数使用切片类型")
	fmt.Println("   ✅ 注意切片共享底层数组的特性")
	fmt.Println("   ✅ 需要独立副本时使用 copy() 函数")
	fmt.Println()
}

// ============================================
// 辅助函数
// ============================================

// hyperspace 处理字符串切片，去除空格
// 这个函数可以处理任何长度的 []string 切片
func hyperspace(worlds []string) {
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}

// processData 处理整数切片
// 可以处理任意长度的 []int 切片
func processData(data []int) {
	fmt.Printf("  处理数据 (长度: %d): %v\n", len(data), data)
}

// trimNames 去除字符串切片中每个元素的前后空格
func trimNames(names []string) {
	for i := range names {
		names[i] = strings.TrimSpace(names[i])
	}
}

