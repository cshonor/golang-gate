// 示例：Go 语言切片的核心知识点
// 切片的本质、创建语法、默认索引、切片的切片、字符串切片

package main

import "fmt"

func main() {
	// ============================================
	// 1. 切片的本质：数组的视图
	// ============================================
	fmt.Println("=== 1. 切片的本质：数组的视图 ===")

	fmt.Println("核心概念:")
	fmt.Println("  - 切片不是数组，它是数组的一个视图")
	fmt.Println("  - 底层指向原始数组")
	fmt.Println("  - 对切片元素的修改会直接反映到原始数组上")
	fmt.Println("  - 所有基于该数组的切片也都会看到这个变化")
	fmt.Println()

	// 演示切片的本质
	planets := [...]string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	terrestrial := planets[0:4]
	gasGiants := planets[4:6]

	fmt.Printf("原始数组: %v\n", planets)
	fmt.Printf("类地行星切片: %v\n", terrestrial)
	fmt.Printf("气态巨行星切片: %v\n", gasGiants)
	fmt.Println()

	// 修改切片元素，观察对原始数组的影响
	fmt.Println("修改切片 terrestrial[2] = \"蓝色星球\":")
	terrestrial[2] = "蓝色星球"
	fmt.Printf("  类地行星切片: %v\n", terrestrial)
	fmt.Printf("  原始数组: %v (也被修改了)\n", planets)
	fmt.Printf("  气态巨行星切片: %v (未受影响，因为不包含该元素)\n", gasGiants)
	fmt.Println()

	// 多个切片共享同一个底层数组
	fmt.Println("多个切片共享同一个底层数组:")
	slice1 := planets[1:5]
	slice2 := planets[2:6]
	fmt.Printf("  slice1[1:5]: %v\n", slice1)
	fmt.Printf("  slice2[2:6]: %v\n", slice2)
	slice1[0] = "修改的元素"
	fmt.Printf("  修改 slice1[0] 后:\n")
	fmt.Printf("    slice1: %v\n", slice1)
	fmt.Printf("    slice2: %v (也看到了变化)\n", slice2)
	fmt.Printf("    原始数组: %v\n", planets)
	fmt.Println()

	// ============================================
	// 2. 切片的创建语法 [start:end]
	// ============================================
	fmt.Println("=== 2. 切片的创建语法 [start:end] ===")

	fmt.Println("语法: slice := source[start:end]")
	fmt.Println("  - start: 切片的起始索引（包含），默认值为 0")
	fmt.Println("  - end: 切片的结束索引（不包含），默认值为原数组/切片的长度")
	fmt.Println("  - 区间: [start:end) 左闭右开")
	fmt.Println()

	// 示例
	planets2 := [...]string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	terrestrial2 := planets2[0:4]   // 取前4个元素（索引0-3）
	gasGiants2 := planets2[4:6]     // 取第5、6个元素（索引4-5）

	fmt.Printf("数组: %v\n", planets2)
	fmt.Printf("  terrestrial[0:4]: %v (索引0-3，共4个元素)\n", terrestrial2)
	fmt.Printf("  gasGiants[4:6]: %v (索引4-5，共2个元素)\n", gasGiants2)
	fmt.Println()

	// 索引范围说明
	fmt.Println("索引范围说明:")
	fmt.Println("  数组长度: 8 (索引范围: 0-7)")
	fmt.Println("  [0:4]: 包含索引 0, 1, 2, 3 (共4个元素)")
	fmt.Println("  [4:6]: 包含索引 4, 5 (共2个元素)")
	fmt.Println("  [6:8]: 包含索引 6, 7 (共2个元素)")
	fmt.Println()

	// ============================================
	// 3. 切片的默认索引
	// ============================================
	fmt.Println("=== 3. 切片的默认索引 ===")

	planets3 := [...]string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}

	fmt.Println("省略 start: 默认从索引 0 开始")
	sliceStart := planets3[:4]
	fmt.Printf("  planets[:4] 等价于 planets[0:4]: %v\n", sliceStart)
	fmt.Println()

	fmt.Println("省略 end: 默认到数组末尾")
	sliceEnd := planets3[4:]
	fmt.Printf("  planets[4:] 等价于 planets[4:8]: %v\n", sliceEnd)
	fmt.Println()

	fmt.Println("同时省略 start 和 end: 得到包含整个数组的切片")
	sliceAll := planets3[:]
	fmt.Printf("  planets[:] 等价于 planets[0:8]: %v\n", sliceAll)
	fmt.Println()

	// 对比说明
	fmt.Println("默认索引对比:")
	fmt.Println("  完整形式         简写形式")
	fmt.Println("  planets[0:4]  = planets[:4]")
	fmt.Println("  planets[4:8]  = planets[4:]")
	fmt.Println("  planets[0:8]  = planets[:]")
	fmt.Println()

	// ============================================
	// 4. 切片的切片
	// ============================================
	fmt.Println("=== 4. 切片的切片 ===")

	fmt.Println("可以基于一个切片再创建新的切片")
	fmt.Println("新切片依然指向原始数组")
	fmt.Println()

	planets4 := [...]string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}

	// 先创建一个切片
	giants := planets4[4:8]
	fmt.Printf("第一步: giants = planets[4:8]: %v\n", giants)
	fmt.Printf("  giants 指向原始数组的索引 4-7\n")
	fmt.Println()

	// 基于切片再创建切片
	gas := giants[0:2]   // 等价于 planets[4:6]
	ice := giants[2:4]   // 等价于 planets[6:8]

	fmt.Printf("第二步: 基于 giants 创建新切片\n")
	fmt.Printf("  gas = giants[0:2]: %v (等价于 planets[4:6])\n", gas)
	fmt.Printf("  ice = giants[2:4]: %v (等价于 planets[6:8])\n", ice)
	fmt.Println()

	// 验证它们都指向同一个底层数组
	fmt.Println("验证: 所有切片都指向同一个底层数组")
	gas[0] = "修改的Jupiter"
	fmt.Printf("  修改 gas[0] 后:\n")
	fmt.Printf("    gas: %v\n", gas)
	fmt.Printf("    ice: %v (未受影响)\n", ice)
	fmt.Printf("    giants: %v (也看到了变化)\n", giants)
	fmt.Printf("    原始数组: %v\n", planets4)
	fmt.Println()

	// 切片的切片的索引说明
	fmt.Println("切片的切片的索引说明:")
	fmt.Println("  giants = planets[4:8]")
	fmt.Println("    giants[0] 对应 planets[4]")
	fmt.Println("    giants[1] 对应 planets[5]")
	fmt.Println("    giants[2] 对应 planets[6]")
	fmt.Println("    giants[3] 对应 planets[7]")
	fmt.Println()
	fmt.Println("  gas = giants[0:2]")
	fmt.Println("    gas[0] 对应 giants[0] 对应 planets[4]")
	fmt.Println("    gas[1] 对应 giants[1] 对应 planets[5]")
	fmt.Println()

	// ============================================
	// 5. 字符串切片
	// ============================================
	fmt.Println("=== 5. 字符串切片 ===")

	fmt.Println("切分字符串会返回一个新的字符串")
	fmt.Println("底层共享原字符串的字节数组，因此也是高效的")
	fmt.Println()

	// 字符串切片示例
	neptune := "Neptune"
	tune := neptune[3:]
	fmt.Printf("  原字符串: %q\n", neptune)
	fmt.Printf("  切片 neptune[3:]: %q\n", tune)
	fmt.Printf("  说明: 从索引3开始到末尾\n")
	fmt.Println()

	// 更多字符串切片示例
	hello := "Hello, World!"
	fmt.Printf("  原字符串: %q\n", hello)
	fmt.Printf("  hello[0:5]: %q (前5个字符)\n", hello[0:5])
	fmt.Printf("  hello[7:]: %q (从索引7开始)\n", hello[7:])
	fmt.Printf("  hello[:5]: %q (前5个字符)\n", hello[:5])
	fmt.Printf("  hello[7:12]: %q (索引7-11)\n", hello[7:12])
	fmt.Println()

	// 字符串切片的特性
	fmt.Println("字符串切片的特性:")
	fmt.Println("  - 字符串是不可变的，切片操作不会修改原字符串")
	fmt.Println("  - 字符串切片返回新的字符串")
	fmt.Println("  - 底层共享字节数组，内存高效")
	fmt.Println()

	// 演示字符串不可变性
	original := "Hello"
	sliced := original[1:4]
	fmt.Printf("  原字符串: %q\n", original)
	fmt.Printf("  切片: %q\n", sliced)
	fmt.Println("  原字符串未被修改（字符串是不可变的）")
	fmt.Println()

	// ============================================
	// 6. 切片核心知识点总结
	// ============================================
	fmt.Println("=== 切片核心知识点总结 ===")
	fmt.Println()
	fmt.Println("1. 切片的本质:")
	fmt.Println("   ✅ 切片是数组的一个视图，不是数组本身")
	fmt.Println("   ✅ 底层指向原始数组")
	fmt.Println("   ✅ 修改切片会影响原始数组")
	fmt.Println("   ✅ 多个切片可以共享同一个底层数组")
	fmt.Println()
	fmt.Println("2. 切片的创建语法:")
	fmt.Println("   ✅ [start:end] - 左闭右开区间")
	fmt.Println("   ✅ start: 起始索引（包含），默认0")
	fmt.Println("   ✅ end: 结束索引（不包含），默认数组长度")
	fmt.Println()
	fmt.Println("3. 默认索引:")
	fmt.Println("   ✅ [:end] - 从开头到end")
	fmt.Println("   ✅ [start:] - 从start到末尾")
	fmt.Println("   ✅ [:] - 整个数组")
	fmt.Println()
	fmt.Println("4. 切片的切片:")
	fmt.Println("   ✅ 可以基于切片再创建切片")
	fmt.Println("   ✅ 新切片依然指向原始数组")
	fmt.Println("   ✅ 索引是相对于父切片的")
	fmt.Println()
	fmt.Println("5. 字符串切片:")
	fmt.Println("   ✅ 切分字符串返回新字符串")
	fmt.Println("   ✅ 底层共享字节数组，高效")
	fmt.Println("   ✅ 字符串是不可变的")
	fmt.Println()
	fmt.Println("6. 注意事项:")
	fmt.Println("   ⚠️  切片是引用类型，赋值时共享底层数组")
	fmt.Println("   ⚠️  修改切片会影响底层数组和其他共享的切片")
	fmt.Println("   ⚠️  需要独立副本时，使用 copy() 函数")
	fmt.Println("   ⚠️  字符串切片返回新字符串，不会修改原字符串")
	fmt.Println()

	// ============================================
	// 7. 实际应用示例
	// ============================================
	fmt.Println("=== 实际应用示例 ===")

	// 示例1：处理数据子集
	fmt.Println("示例1：处理数据子集")
	data := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	firstHalf := data[:5]
	secondHalf := data[5:]
	fmt.Printf("  数据: %v\n", data)
	fmt.Printf("  前半部分: %v\n", firstHalf)
	fmt.Printf("  后半部分: %v\n", secondHalf)
	fmt.Println()

	// 示例2：字符串处理
	fmt.Println("示例2：字符串处理")
	url := "https://example.com/path/to/resource"
	protocol := url[:5]      // "https"
	domain := url[8:15]     // "example"
	path := url[15:]        // "/path/to/resource"
	fmt.Printf("  URL: %q\n", url)
	fmt.Printf("  协议: %q\n", protocol)
	fmt.Printf("  域名: %q\n", domain)
	fmt.Printf("  路径: %q\n", path)
	fmt.Println()

	// 示例3：多层切片
	fmt.Println("示例3：多层切片")
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	firstRow := matrix[0]
	firstTwoCols := firstRow[:2]
	fmt.Printf("  矩阵第一行: %v\n", firstRow)
	fmt.Printf("  第一行的前两列: %v\n", firstTwoCols)
	fmt.Println()
}

