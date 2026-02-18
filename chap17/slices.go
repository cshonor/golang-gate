// 示例：Go 语言的切片（Slice）
// 切片是指向底层数组的"窗口"，是引用类型，可以动态操作数组的一部分
string 切片是按「字节」切，不是按「字符」切。
英文没事，中文、 emoji 一切就乱码。
[数字]类型  → 数组（固定大小）
​
-  []类型    → 切片（可变大小）
package main

import "fmt"

func main() {
	// ============================================
	// 1. 核心概念：切片 vs 数组
	// ============================================
	fmt.Println("=== 核心概念：切片 vs 数组 ===")

	fmt.Println("特性对比:")
	fmt.Println("  数组:")
	fmt.Println("    - 类型：值类型")
	fmt.Println("    - 长度：固定，声明时确定")
	fmt.Println("    - 赋值/传递：复制整个数组")
	fmt.Println("    - 用途：存储固定长度的同类型元素")
	fmt.Println()
	fmt.Println("  切片:")
	fmt.Println("    - 类型：引用类型")
	fmt.Println("    - 长度：动态，可随时改变")
	fmt.Println("    - 赋值/传递：复制切片头（指针、长度、容量），底层数组共享")
	fmt.Println("    - 用途：动态操作数组的一部分，是Go中最常用的集合类型")
	fmt.Println()

	// ============================================
	// 2. 从数组切割创建切片（窗口）
	// ============================================
	fmt.Println("=== 从数组切割创建切片（窗口）===")

	// 定义底层数组
	planets := [...]string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}//不定长要加，
	fmt.Printf("底层数组: %v\n", planets)
	fmt.Printf("数组长度: %d\n", len(planets))
	fmt.Println()

	// 切割数组创建切片（左闭右开区间）
	terrestrial := planets[0:4]   // 类地行星：索引0-3（不包含4）
	gasGiants := planets[4:6]     // 气态巨行星：索引4-5（不包含6）
	iceGiants := planets[6:8]      // 冰巨行星：索引6-7（不包含8）

	fmt.Println("切割数组创建切片:")
	fmt.Printf("  类地行星 terrestrial[0:4]: %v\n", terrestrial)
	fmt.Printf("  气态巨行星 gasGiants[4:6]: %v\n", gasGiants)
	fmt.Printf("  冰巨行星 iceGiants[6:8]: %v\n", iceGiants)
	fmt.Println()

	// 切片的长度和容量
	fmt.Println("切片的长度和容量:")
	fmt.Printf("  terrestrial: len=%d, cap=%d\n", len(terrestrial), cap(terrestrial))
	fmt.Printf("  gasGiants: len=%d, cap=%d\n", len(gasGiants), cap(gasGiants))
	fmt.Printf("  iceGiants: len=%d, cap=%d\n", len(iceGiants), cap(iceGiants))
	fmt.Println()

	// ============================================
	// 3. 切片的本质：共享底层数组
	// ============================================
	fmt.Println("=== 切片的本质：共享底层数组 ===")

	fmt.Println("修改切片元素会影响底层数组:")
	terrestrial[2] = "蓝色星球"
	fmt.Printf("  修改后类地行星: %v\n", terrestrial)
	fmt.Printf("  修改后底层数组: %v\n", planets)
	fmt.Println("  说明：切片和底层数组共享同一块内存")
	fmt.Println()

	// 多个切片共享底层数组
	fmt.Println("多个切片共享底层数组:")
	slice1 := planets[1:5]
	slice2 := planets[2:6]
	fmt.Printf("  slice1[1:5]: %v\n", slice1)
	fmt.Printf("  slice2[2:6]: %v\n", slice2)
	slice1[0] = "修改的元素"
	fmt.Printf("  修改 slice1[0] 后:")
	fmt.Printf("    slice1: %v\n", slice1)
	fmt.Printf("    slice2: %v\n", slice2)
	fmt.Printf("    底层数组: %v\n", planets)
	fmt.Println("  说明：修改一个切片会影响共享同一底层数组的其他切片")
	fmt.Println()

	// ============================================
	// 4. 切片的创建方式
	// ============================================
	fmt.Println("=== 切片的创建方式 ===")

	// 方式1：从数组切割
	fmt.Println("方式1：从数组切割")
	arr := [5]int{1, 2, 3, 4, 5}
	sliceFromArray := arr[1:4]
	fmt.Printf("  数组: %v\n", arr)
	fmt.Printf("  切片 arr[1:4]: %v\n", sliceFromArray)
	fmt.Println()

	// 方式2：使用 make 函数
	fmt.Println("方式2：使用 make 函数")
	// make([]类型, 长度, 容量)
	sliceMake := make([]string, 4, 8)
	fmt.Printf("  make([]string, 4, 8): len=%d, cap=%d\n", len(sliceMake), cap(sliceMake))
	fmt.Printf("  切片内容: %v\n", sliceMake)
	fmt.Println("  说明：创建长度为4，容量为8的字符串切片")
	fmt.Println()

	// 方式3：字面量创建
	fmt.Println("方式3：字面量创建")
	sliceLiteral := []string{"a", "b", "c"}
	fmt.Printf("  []string{\"a\", \"b\", \"c\"}: %v\n", sliceLiteral)
	fmt.Printf("  len=%d, cap=%d\n", len(sliceLiteral), cap(sliceLiteral))
	fmt.Println()

	// 方式4：空切片
	fmt.Println("方式4：空切片")
	var emptySlice []int
	emptySlice2 := []int{}
	emptySlice3 := make([]int, 0)
	fmt.Printf("  var emptySlice []int: %v, len=%d, cap=%d, nil=%t\n",
		emptySlice, len(emptySlice), cap(emptySlice), emptySlice == nil)
	fmt.Printf("  []int{}: %v, len=%d, cap=%d, nil=%t\n",
		emptySlice2, len(emptySlice2), cap(emptySlice2), emptySlice2 == nil)
	fmt.Printf("  make([]int, 0): %v, len=%d, cap=%d, nil=%t\n",
		emptySlice3, len(emptySlice3), cap(emptySlice3), emptySlice3 == nil)
	fmt.Println()

	// ============================================
	// 5. 切片的长度与容量
	// ============================================
	fmt.Println("=== 切片的长度与容量 ===")

	fmt.Println("长度（Length）：切片中当前元素的数量")
	fmt.Println("容量（Capacity）：从切片的第一个元素到底层数组末尾的元素数量")
	fmt.Println()

	// 演示长度和容量的区别
	demoArray := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	demoSlice := demoArray[2:7]
	fmt.Printf("底层数组: %v\n", demoArray)
	fmt.Printf("切片 demoArray[2:7]: %v\n", demoSlice)
	fmt.Printf("  长度 len: %d (切片中的元素数量)\n", len(demoSlice))
	fmt.Printf("  容量 cap: %d (从索引2到数组末尾的元素数量)\n", cap(demoSlice))
	fmt.Println()

	// 不同切割位置的容量
	fmt.Println("不同切割位置的容量:")
	slice1_cap := demoArray[0:5]
	slice2_cap := demoArray[5:10]
	slice3_cap := demoArray[2:8]
	fmt.Printf("  [0:5]: len=%d, cap=%d\n", len(slice1_cap), cap(slice1_cap))
	fmt.Printf("  [5:10]: len=%d, cap=%d\n", len(slice2_cap), cap(slice2_cap))
	fmt.Printf("  [2:8]: len=%d, cap=%d\n", len(slice3_cap), cap(slice3_cap))
	fmt.Println()

	// ============================================
	// 6. 切片的简写形式
	// ============================================
	fmt.Println("=== 切片的简写形式 ===")

	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("数组: %v\n", numbers)

	// 省略起始索引（从0开始）
	sliceStart := numbers[:3]
	fmt.Printf("  numbers[:3]: %v (等价于 numbers[0:3])\n", sliceStart)

	// 省略结束索引（到数组末尾）
	sliceEnd := numbers[2:]
	fmt.Printf("  numbers[2:]: %v (等价于 numbers[2:5])\n", sliceEnd)

	// 省略两个索引（整个数组）
	sliceAll := numbers[:]
	fmt.Printf("  numbers[:]: %v (等价于 numbers[0:5])\n", sliceAll)
	fmt.Println()

	// ============================================
	// 7. 切片的动态扩容
	// ============================================
	fmt.Println("=== 切片的动态扩容 ===")

	fmt.Println("当切片长度超过容量时，Go会自动扩容底层数组")
	fmt.Println("扩容策略：通常是原来的2倍")
	fmt.Println()

	// 演示扩容
	growthSlice := make([]int, 0, 2)
	fmt.Printf("初始: len=%d, cap=%d, %v\n", len(growthSlice), cap(growthSlice), growthSlice)

	growthSlice = append(growthSlice, 1)
	fmt.Printf("添加1个元素: len=%d, cap=%d, %v\n", len(growthSlice), cap(growthSlice), growthSlice)

	growthSlice = append(growthSlice, 2)
	fmt.Printf("添加2个元素: len=%d, cap=%d, %v\n", len(growthSlice), cap(growthSlice), growthSlice)

	growthSlice = append(growthSlice, 3)
	fmt.Printf("添加3个元素: len=%d, cap=%d, %v (容量自动扩容)\n", len(growthSlice), cap(growthSlice), growthSlice)

	growthSlice = append(growthSlice, 4, 5)
	fmt.Printf("添加4-5个元素: len=%d, cap=%d, %v\n", len(growthSlice), cap(growthSlice), growthSlice)
	fmt.Println()

	// ============================================
	// 8. 切片的引用类型特性
	// ============================================
	fmt.Println("=== 切片的引用类型特性 ===")

	fmt.Println("切片赋值时，只复制切片头（指针、长度、容量），不复制底层数组")
	original := []int{1, 2, 3, 4, 5}
	copySlice := original
	fmt.Printf("  原始切片: %v\n", original)
	fmt.Printf("  复制切片: %v\n", copySlice)
	copySlice[0] = 999
	fmt.Printf("  修改复制切片后:\n")
	fmt.Printf("    原始切片: %v (也被修改了)\n", original)
	fmt.Printf("    复制切片: %v\n", copySlice)
	fmt.Println("  说明：两个切片共享同一个底层数组")
	fmt.Println()

	// ============================================
	// 9. 切片的实际应用示例
	// ============================================
	fmt.Println("=== 切片的实际应用示例 ===")

	// 示例1：行星分类
	fmt.Println("示例1：行星分类")
	allPlanets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}
	terrestrialPlanets := allPlanets[0:4]
	gasGiantPlanets := allPlanets[4:6]
	iceGiantPlanets := allPlanets[6:8]
	fmt.Printf("  类地行星: %v\n", terrestrialPlanets)
	fmt.Printf("  气态巨行星: %v\n", gasGiantPlanets)
	fmt.Printf("  冰巨行星: %v\n", iceGiantPlanets)
	fmt.Println()

	// 示例2：字符串切片
	fmt.Println("示例2：字符串切片")
	text := "Hello, World!"
	textSlice := []byte(text)
	fmt.Printf("  字符串: %q\n", text)
	fmt.Printf("  字节切片: %v\n", textSlice)
	fmt.Printf("  切片长度: %d\n", len(textSlice))
	fmt.Println()

	// ============================================
	// 10. 切片常用操作清单
	// ============================================
	fmt.Println("=== 切片常用操作清单 ===")
	fmt.Println()
	fmt.Println("1. 创建切片:")
	fmt.Println("   ✅ 从数组切割: arr[1:4]")
	fmt.Println("   ✅ 使用 make: make([]int, 5, 10)")
	fmt.Println("   ✅ 字面量: []int{1, 2, 3}")
	fmt.Println()
	fmt.Println("2. 获取信息:")
	fmt.Println("   ✅ 长度: len(slice)")
	fmt.Println("   ✅ 容量: cap(slice)")
	fmt.Println()
	fmt.Println("3. 访问元素:")
	fmt.Println("   ✅ 索引访问: slice[0]")
	fmt.Println("   ✅ 范围访问: slice[1:3]")
	fmt.Println()
	fmt.Println("4. 修改元素:")
	fmt.Println("   ✅ 直接赋值: slice[0] = value")
	fmt.Println("   ⚠️  注意：会影响共享底层数组的其他切片")
	fmt.Println()
	fmt.Println("5. 追加元素:")
	fmt.Println("   ✅ append(slice, element)")
	fmt.Println("   ✅ append(slice, elem1, elem2, ...)")
	fmt.Println("   ✅ 自动扩容：容量不足时自动扩容")
	fmt.Println()
	fmt.Println("6. 遍历切片:")
	fmt.Println("   ✅ for i, v := range slice")
	fmt.Println("   ✅ for i := 0; i < len(slice); i++")
	fmt.Println()
	fmt.Println("7. 注意事项:")
	fmt.Println("   ⚠️  切片是引用类型，赋值时共享底层数组")
	fmt.Println("   ⚠️  修改切片会影响底层数组和其他共享的切片")
	fmt.Println("   ⚠️  需要独立副本时，使用 copy() 函数")
	fmt.Println("   ⚠️  空切片和 nil 切片的区别")
	fmt.Println()

	// ============================================
	// 11. 总结
	// ============================================
	fmt.Println("=== 总结 ===")
	fmt.Println("1. 切片的本质:")
	fmt.Println("   - 切片是指向底层数组的'窗口'")
	fmt.Println("   - 包含：指向底层数组的指针、长度、容量")
	fmt.Println("   - 是引用类型，赋值时共享底层数组")
	fmt.Println()
	fmt.Println("2. 切片的创建:")
	fmt.Println("   - 从数组切割: arr[start:end]")
	fmt.Println("   - 使用 make: make([]type, len, cap)")
	fmt.Println("   - 字面量: []type{values}")
	fmt.Println()
	fmt.Println("3. 长度与容量:")
	fmt.Println("   - 长度：当前元素数量")
	fmt.Println("   - 容量：从第一个元素到底层数组末尾的元素数量")
	fmt.Println()
	fmt.Println("4. 动态扩容:")
	fmt.Println("   - 当长度超过容量时自动扩容")
	fmt.Println("   - 扩容策略：通常是原来的2倍")
	fmt.Println()
	fmt.Println("5. 共享底层数组:")
	fmt.Println("   - 多个切片可以共享同一个底层数组")
	fmt.Println("   - 修改一个切片会影响其他共享的切片")
	fmt.Println("   - 需要注意避免意外修改")
}

