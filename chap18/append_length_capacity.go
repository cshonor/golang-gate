// 示例：Go 语言切片的 append 函数与长度/容量
// 演示 append 函数的使用、切片的长度和容量、扩容机制

package main

import "fmt"

func main() {
	// ============================================
	// 1. append 函数基础
	// ============================================
	fmt.Println("=== 1. append 函数基础 ===")

	fmt.Println("append 函数的作用:")
	fmt.Println("  - 向切片中添加元素")
	fmt.Println("  - 是Go语言实现动态扩容的核心函数")
	fmt.Println("  - 它是一个可变参数函数，可以一次性添加一个或多个元素")
	fmt.Println()

	// 基础示例
	fmt.Println("基础示例:")
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	fmt.Printf("  原始切片: %v\n", dwarfs)
	fmt.Printf("  长度: %d, 容量: %d\n", len(dwarfs), cap(dwarfs))

	// 添加一个元素
	dwarfs = append(dwarfs, "Orcus")
	fmt.Printf("  添加一个元素后: %v\n", dwarfs)
	fmt.Printf("  长度: %d, 容量: %d\n", len(dwarfs), cap(dwarfs))

	// 添加多个元素
	dwarfs = append(dwarfs, "Salacia", "Quaoar", "Sedna")
	fmt.Printf("  添加多个元素后: %v\n", dwarfs)
	fmt.Printf("  长度: %d, 容量: %d\n", len(dwarfs), cap(dwarfs))
	fmt.Println()

	// ============================================
	// 2. append 函数的特点
	// ============================================
	fmt.Println("=== 2. append 函数的特点 ===")

	fmt.Println("特点1：容量足够时，直接在底层数组末尾添加")
	slice1 := make([]int, 3, 5) // 长度3，容量5
	slice1[0] = 1
	slice1[1] = 2
	slice1[2] = 3
	fmt.Printf("  原始: %v, len=%d, cap=%d\n", slice1, len(slice1), cap(slice1))
	slice1 = append(slice1, 4)
	fmt.Printf("  添加后: %v, len=%d, cap=%d (容量足够，未扩容)\n", slice1, len(slice1), cap(slice1))
	fmt.Println()

	fmt.Println("特点2：容量不足时，自动创建容量更大的新底层数组")
	slice2 := []int{1, 2, 3}
	fmt.Printf("  原始: %v, len=%d, cap=%d\n", slice2, len(slice2), cap(slice2))
	slice2 = append(slice2, 4, 5, 6, 7)
	fmt.Printf("  添加后: %v, len=%d, cap=%d (容量不足，已扩容)\n", slice2, len(slice2), cap(slice2))
	fmt.Println()

	// ============================================
	// 3. 切片的长度（Length）与容量（Capacity）
	// ============================================
	fmt.Println("=== 3. 切片的长度（Length）与容量（Capacity）===")

	fmt.Println("长度（len）：切片中当前包含的元素数量")
	fmt.Println("容量（cap）：切片底层数组从切片起始索引开始，到数组末尾的元素总数")
	fmt.Println("           代表了切片最多能容纳多少元素（不扩容的情况下）")
	fmt.Println()

	// 示例1：从数组创建切片
	fmt.Println("示例1：从数组创建切片")
	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice3 := arr[2:7] // 从索引2到7（不包含7）
	fmt.Printf("  数组: %v\n", arr)
	fmt.Printf("  切片 arr[2:7]: %v\n", slice3)
	fmt.Printf("  长度: %d (切片中的元素数量)\n", len(slice3))
	fmt.Printf("  容量: %d (从索引2到数组末尾的元素数量)\n", cap(slice3))
	fmt.Println()

	// 示例2：字面量创建的切片
	fmt.Println("示例2：字面量创建的切片")
	slice4 := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	fmt.Printf("  切片: %v\n", slice4)
	fmt.Printf("  长度: %d\n", len(slice4))
	fmt.Printf("  容量: %d (初始长度等于容量)\n", cap(slice4))
	fmt.Println()

	// 示例3：使用 make 创建切片
	fmt.Println("示例3：使用 make 创建切片")
	slice5 := make([]int, 3, 10) // 长度3，容量10
	fmt.Printf("  make([]int, 3, 10): %v\n", slice5)
	fmt.Printf("  长度: %d\n", len(slice5))
	fmt.Printf("  容量: %d\n", cap(slice5))
	fmt.Println()

	// ============================================
	// 4. 扩容机制
	// ============================================
	fmt.Println("=== 4. 扩容机制 ===")

	fmt.Println("当调用 append 时，如果新元素数量超出了切片的容量，Go会自动扩容")
	fmt.Println("扩容策略通常是将容量翻倍（对于小切片），以减少频繁扩容带来的性能开销")
	fmt.Println("扩容后，新切片会指向一个全新的底层数组，原数组会被垃圾回收")
	fmt.Println()

	// 演示扩容过程
	fmt.Println("演示扩容过程:")
	slice6 := []int{1, 2, 3}
	fmt.Printf("  初始: %v, len=%d, cap=%d\n", slice6, len(slice6), cap(slice6))

	for i := 4; i <= 10; i++ {
		oldCap := cap(slice6)
		slice6 = append(slice6, i)
		newCap := cap(slice6)
		if newCap != oldCap {
			fmt.Printf("  添加 %d: %v, len=%d, cap=%d (扩容: %d -> %d)\n",
				i, slice6, len(slice6), newCap, oldCap, newCap)
		} else {
			fmt.Printf("  添加 %d: %v, len=%d, cap=%d (未扩容)\n",
				i, slice6, len(slice6), newCap)
		}
	}
	fmt.Println()

	// ============================================
	// 5. 不同初始容量下的扩容表现
	// ============================================
	fmt.Println("=== 5. 不同初始容量下的扩容表现 ===")

	// 场景1：小切片（容量 < 1024）
	fmt.Println("场景1：小切片（容量 < 1024）")
	smallSlice := []int{1}
	fmt.Printf("  初始: len=%d, cap=%d\n", len(smallSlice), cap(smallSlice))
	for i := 0; i < 10; i++ {
		oldCap := cap(smallSlice)
		smallSlice = append(smallSlice, i+2)
		newCap := cap(smallSlice)
		if newCap != oldCap {
			fmt.Printf("  添加元素后: len=%d, cap=%d (扩容: %d -> %d, 倍数: %.2f)\n",
				len(smallSlice), newCap, oldCap, newCap, float64(newCap)/float64(oldCap))
		}
	}
	fmt.Println()

	// 场景2：使用 make 指定初始容量
	fmt.Println("场景2：使用 make 指定初始容量")
	madeSlice := make([]int, 0, 5) // 长度0，容量5
	fmt.Printf("  初始: len=%d, cap=%d\n", len(madeSlice), cap(madeSlice))
	for i := 1; i <= 10; i++ {
		oldCap := cap(madeSlice)
		madeSlice = append(madeSlice, i)
		newCap := cap(madeSlice)
		if newCap != oldCap {
			fmt.Printf("  添加 %d: len=%d, cap=%d (扩容: %d -> %d)\n",
				i, len(madeSlice), newCap, oldCap, newCap)
		}
	}
	fmt.Println()

	// 场景3：一次性添加多个元素
	fmt.Println("场景3：一次性添加多个元素")
	multiSlice := []int{1, 2, 3}
	fmt.Printf("  初始: len=%d, cap=%d\n", len(multiSlice), cap(multiSlice))
	oldCap := cap(multiSlice)
	multiSlice = append(multiSlice, 4, 5, 6, 7, 8, 9, 10)
	newCap := cap(multiSlice)
	fmt.Printf("  一次性添加7个元素后: len=%d, cap=%d (扩容: %d -> %d)\n",
		len(multiSlice), newCap, oldCap, newCap)
	fmt.Println()

	// ============================================
	// 6. append 的返回值
	// ============================================
	fmt.Println("=== 6. append 的返回值 ===")

	fmt.Println("重要：append 返回新的切片，必须接收返回值")
	fmt.Println()

	// 错误示例（已注释）
	fmt.Println("错误示例（已注释）:")
	fmt.Println("  slice := []int{1, 2, 3}")
	fmt.Println("  append(slice, 4)  // ❌ 错误：没有接收返回值")
	fmt.Println("  fmt.Println(slice) // 输出：[1 2 3]，未改变")
	fmt.Println()

	// 正确示例
	fmt.Println("正确示例:")
	correctSlice := []int{1, 2, 3}
	fmt.Printf("  原始: %v\n", correctSlice)
	correctSlice = append(correctSlice, 4) // ✅ 正确：接收返回值
	fmt.Printf("  添加后: %v\n", correctSlice)
	fmt.Println()

	// ============================================
	// 7. append 与底层数组的关系
	// ============================================
	fmt.Println("=== 7. append 与底层数组的关系 ===")

	fmt.Println("扩容前：多个切片可能共享同一个底层数组")
	baseArray := [5]int{1, 2, 3, 4, 5}
	sliceA := baseArray[1:4]
	sliceB := baseArray[2:5]
	fmt.Printf("  底层数组: %v\n", baseArray)
	fmt.Printf("  sliceA[1:4]: %v, len=%d, cap=%d\n", sliceA, len(sliceA), cap(sliceA))
	fmt.Printf("  sliceB[2:5]: %v, len=%d, cap=%d\n", sliceB, len(sliceB), cap(sliceB))
	fmt.Println()

	fmt.Println("扩容后：新切片指向新的底层数组")
	sliceA = append(sliceA, 99) // 容量足够，不扩容
	fmt.Printf("  sliceA append(99) 后（容量足够）:\n")
	fmt.Printf("    sliceA: %v, len=%d, cap=%d\n", sliceA, len(sliceA), cap(sliceA))
	fmt.Printf("    sliceB: %v (未受影响)\n", sliceB)
	fmt.Printf("    底层数组: %v (被修改了)\n", baseArray)
	fmt.Println()

	// 扩容示例
	fmt.Println("扩容示例:")
	sliceC := []int{1, 2, 3}
	sliceD := sliceC // 共享底层数组
	fmt.Printf("  初始: sliceC=%v, sliceD=%v\n", sliceC, sliceD)
	fmt.Printf("  sliceC: len=%d, cap=%d\n", len(sliceC), cap(sliceC))
	sliceC = append(sliceC, 4, 5, 6) // 容量不足，扩容
	fmt.Printf("  sliceC append(4,5,6) 后（容量不足，扩容）:\n")
	fmt.Printf("    sliceC: %v, len=%d, cap=%d (指向新数组)\n", sliceC, len(sliceC), cap(sliceC))
	fmt.Printf("    sliceD: %v (仍指向原数组)\n", sliceD)
	fmt.Println("  说明：扩容后，sliceC 指向新数组，sliceD 仍指向原数组")
	fmt.Println()

	// ============================================
	// 8. append 的常见用法
	// ============================================
	fmt.Println("=== 8. append 的常见用法 ===")

	// 用法1：添加单个元素
	fmt.Println("用法1：添加单个元素")
	slice7 := []int{1, 2, 3}
	slice7 = append(slice7, 4)
	fmt.Printf("  append(slice, 4): %v\n", slice7)
	fmt.Println()

	// 用法2：添加多个元素
	fmt.Println("用法2：添加多个元素")
	slice8 := []int{1, 2, 3}
	slice8 = append(slice8, 4, 5, 6)
	fmt.Printf("  append(slice, 4, 5, 6): %v\n", slice8)
	fmt.Println()

	// 用法3：添加另一个切片（使用展开运算符）
	fmt.Println("用法3：添加另一个切片（使用展开运算符 ...）")
	slice9 := []int{1, 2, 3}
	slice10 := []int{4, 5, 6}
	slice9 = append(slice9, slice10...) // 注意：... 是展开运算符
	fmt.Printf("  append(slice9, slice10...): %v\n", slice9)
	fmt.Println()

	// 用法4：在开头添加元素（需要创建新切片）
	fmt.Println("用法4：在开头添加元素")
	slice11 := []int{2, 3, 4}
	slice11 = append([]int{1}, slice11...) // 在开头添加1
	fmt.Printf("  append([]int{1}, slice...): %v\n", slice11)
	fmt.Println()

	// 用法5：插入元素到指定位置
	fmt.Println("用法5：插入元素到指定位置")
	slice12 := []int{1, 2, 4, 5}
	index := 2
	value := 3
	// 在索引2处插入3
	slice12 = append(slice12[:index], append([]int{value}, slice12[index:]...)...)
	fmt.Printf("  在索引 %d 处插入 %d: %v\n", index, value, slice12)
	fmt.Println()

	// ============================================
	// 9. append 与性能考虑
	// ============================================
	fmt.Println("=== 9. append 与性能考虑 ===")

	fmt.Println("性能优化建议:")
	fmt.Println("  1. 如果知道大概的元素数量，使用 make 预分配容量")
	fmt.Println("  2. 避免频繁的小容量扩容")
	fmt.Println("  3. 批量添加时，一次性添加多个元素")
	fmt.Println()

	// 对比：预分配 vs 不预分配
	fmt.Println("对比：预分配 vs 不预分配")
	// 不预分配
	noPrealloc := []int{}
	for i := 0; i < 1000; i++ {
		noPrealloc = append(noPrealloc, i)
	}
	fmt.Printf("  不预分配: 最终 len=%d, cap=%d (可能多次扩容)\n", len(noPrealloc), cap(noPrealloc))

	// 预分配
	prealloc := make([]int, 0, 1000)
	for i := 0; i < 1000; i++ {
		prealloc = append(prealloc, i)
	}
	fmt.Printf("  预分配: 最终 len=%d, cap=%d (无需扩容)\n", len(prealloc), cap(prealloc))
	fmt.Println()

	// ============================================
	// 10. 总结
	// ============================================
	fmt.Println("=== 10. 总结 ===")
	fmt.Println()
	fmt.Println("1. append 函数:")
	fmt.Println("   ✅ 向切片中添加元素")
	fmt.Println("   ✅ 可变参数函数，可添加一个或多个元素")
	fmt.Println("   ✅ 容量足够时直接添加，容量不足时自动扩容")
	fmt.Println("   ✅ 必须接收返回值")
	fmt.Println()
	fmt.Println("2. 长度与容量:")
	fmt.Println("   ✅ 长度：当前元素数量")
	fmt.Println("   ✅ 容量：底层数组从切片起始到末尾的元素数量")
	fmt.Println("   ✅ 使用 len() 和 cap() 获取")
	fmt.Println()
	fmt.Println("3. 扩容机制:")
	fmt.Println("   ✅ 容量不足时自动扩容")
	fmt.Println("   ✅ 小切片通常容量翻倍")
	fmt.Println("   ✅ 扩容后指向新底层数组")
	fmt.Println()
	fmt.Println("4. 最佳实践:")
	fmt.Println("   ✅ 预知容量时使用 make 预分配")
	fmt.Println("   ✅ 批量添加时一次性添加多个元素")
	fmt.Println("   ✅ 注意 append 的返回值")
	fmt.Println("   ✅ 注意扩容后底层数组的变化")
	fmt.Println()
}

