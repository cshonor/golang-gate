// 示例：Go 语言映射的进阶特性
// 演示逗号与ok语法、映射引用类型、make预分配、计数和words.go实验题

package main

import (
	"fmt"
	"strings"
)

func main() {
	// ============================================
	// 1. 逗号与ok语法
	// ============================================
	fmt.Println("=== 1. 逗号与ok语法 ===")

	fmt.Println("问题：如何区分'键不存在'和'键存在但值为零值'？")
	fmt.Println("答案：使用逗号与ok语法")
	fmt.Println()

	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}

	// 方式1：直接取值（无法区分键不存在和值为零值）
	fmt.Println("方式1：直接取值（无法区分）")
	moon := temperature["Moon"]
	fmt.Printf("  temperature[\"Moon\"] = %d (无法判断是键不存在还是值为0)\n", moon)
	fmt.Println()

	// 方式2：逗号与ok语法（可以区分）
	fmt.Println("方式2：逗号与ok语法（可以区分）")
	if moon, ok := temperature["Moon"]; ok {
		fmt.Printf("  On average the moon is %v° C.\n", moon)
	} else {
		fmt.Println("  Where is the moon?")
	}
	fmt.Println()

	// 演示：键存在但值为0的情况
	fmt.Println("演示：键存在但值为0的情况")
	temperature["Moon"] = 0 // 设置值为0
	if moon, ok := temperature["Moon"]; ok {
		fmt.Printf("  On average the moon is %v° C. (键存在，值为0)\n", moon)
	} else {
		fmt.Println("  Where is the moon?")
	}
	fmt.Println()

	// 注意：第二个变量可以使用任何名字
	fmt.Println("注意：第二个变量可以使用任何名字")
	temp, found := temperature["Venus"]
	if found {
		fmt.Printf("  Venus 温度: %d° C\n", temp)
	} else {
		fmt.Println("  Venus 不存在")
	}
	fmt.Println()

	// 速查19-1答案
	fmt.Println("速查19-1答案:")
	fmt.Println("  1. 映射的类型应为 map[float64]int")
	fmt.Println("  2. 如果设置 Moon 的值为0，ok变量的值将为true")
	fmt.Println()

	// ============================================
	// 2. 映射不会被复制（引用类型特性）
	// ============================================
	fmt.Println("=== 2. 映射不会被复制（引用类型特性）===")

	fmt.Println("映射是引用类型，赋值时不会复制，而是共享底层数据")
	fmt.Println()

	// 代码清单19-2：指向相同数据的映射
	fmt.Println("代码清单19-2：指向相同数据的映射")
	planets := map[string]string{
		"Earth": "Sector ZZ9",
		"Mars":  "Sector ZZ9",
	}
	fmt.Printf("  原始映射: %v\n", planets)

	planetsMarkII := planets // 赋值，共享底层数据
	fmt.Printf("  赋值后 planetsMarkII: %v\n", planetsMarkII)

	// 修改原映射
	planets["Earth"] = "whoops"
	fmt.Printf("  修改 planets[\"Earth\"] 后:\n")
	fmt.Printf("    planets: %v\n", planets)
	fmt.Printf("    planetsMarkII: %v (也被修改了)\n", planetsMarkII)
	fmt.Println()

	// 删除元素
	fmt.Println("删除元素:")
	delete(planets, "Earth")
	fmt.Printf("  删除 planets[\"Earth\"] 后:\n")
	fmt.Printf("    planets: %v\n", planets)
	fmt.Printf("    planetsMarkII: %v (也被删除了)\n", planetsMarkII)
	fmt.Println()

	// 速查19-2答案
	fmt.Println("速查19-2答案:")
	fmt.Println("  1. 因为 planetsMarkII 变量与 planets 变量指向的是相同的底层数据")
	fmt.Println("  2. delete 函数可以从映射中移除指定的元素")
	fmt.Println()

	// 函数参数传递（也是引用）
	fmt.Println("函数参数传递（也是引用）:")
	testMap := map[string]int{"a": 1, "b": 2}
	fmt.Printf("  调用函数前: %v\n", testMap)
	modifyMap(testMap)
	fmt.Printf("  调用函数后: %v (被修改了)\n", testMap)
	fmt.Println()

	// ============================================
	// 3. 使用 make 函数对映射实行预分配
	// ============================================
	fmt.Println("=== 3. 使用 make 函数对映射实行预分配 ===")

	fmt.Println("使用 make 函数可以为映射预分配空间")
	fmt.Println("好处：减少后续扩容时的内存分配和数据复制")
	fmt.Println()

	// 方式1：不预分配
	fmt.Println("方式1：不预分配")
	map1 := make(map[string]int)
	fmt.Printf("  make(map[string]int): len=%d\n", len(map1))
	fmt.Println()

	// 方式2：预分配容量
	fmt.Println("方式2：预分配容量")
	map2 := make(map[string]int, 10) // 预分配10个键的空间
	fmt.Printf("  make(map[string]int, 10): len=%d (初始长度为0)\n", len(map2))
	fmt.Println("  说明：预分配容量可以减少后续扩容的开销")
	fmt.Println()

	// 对比：预分配 vs 不预分配
	fmt.Println("对比：预分配 vs 不预分配")
	// 不预分配
	noPrealloc := make(map[int]int)
	for i := 0; i < 100; i++ {
		noPrealloc[i] = i * 2
	}
	fmt.Printf("  不预分配: len=%d (可能多次扩容)\n", len(noPrealloc))

	// 预分配
	prealloc := make(map[int]int, 100)
	for i := 0; i < 100; i++ {
		prealloc[i] = i * 2
	}
	fmt.Printf("  预分配: len=%d (减少扩容次数)\n", len(prealloc))
	fmt.Println()

	// 速查19-3答案
	fmt.Println("速查19-3答案:")
	fmt.Println("  跟切片一样，为映射指定初始大小能够在映射变得更大的时候减少一些后续工作")
	fmt.Println()

	// ============================================
	// 4. 使用映射进行计数
	// ============================================
	fmt.Println("=== 4. 使用映射进行计数 ===")

	// 代码清单19-3：统计温度出现的频率
	fmt.Println("代码清单19-3：统计温度出现的频率")
	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}
	fmt.Printf("  温度数据: %v\n", temperatures)

	frequency := make(map[float64]int)
	for _, t := range temperatures {
		frequency[t]++ // 如果键不存在，会先初始化为0，然后++
	}
	fmt.Printf("  温度频率: %v\n", frequency)
	fmt.Println()

	// 说明：为什么使用映射而不是切片
	fmt.Println("说明：为什么使用映射而不是切片")
	fmt.Println("  - 如果使用切片，键必须是整数")
	fmt.Println("  - 需要为所有可能的温度值预留空间")
	fmt.Println("  - 映射可以只存储实际出现的温度值")
	fmt.Println("  - 映射更适合这种非结构化数据")
	fmt.Println()

	// 更多计数示例
	fmt.Println("更多计数示例:")
	// 字符计数
	text := "hello world"
	charCount := make(map[rune]int)
	for _, char := range text {
		if char != ' ' {
			charCount[char]++
		}
	}
	fmt.Printf("  文本: %q\n", text)
	fmt.Printf("  字符计数: %v\n", charCount)
	fmt.Println()

	// ============================================
	// 5. 实验题：words.go
	// ============================================
	fmt.Println("=== 5. 实验题：words.go ===")

	fmt.Println("题目要求：")
	fmt.Println("  编写一个函数，统计文本字符串中不同单词的出现频率")
	fmt.Println("  返回一个词频映射")
	fmt.Println("  需要将文本转换为小写字母并移除标点符号")
	fmt.Println("  使用 strings 包中的 Fields、ToLower 和 Trim 函数")
	fmt.Println()

	// 测试文本
	testText := "As far as eye could reach he saw nothing but the stems of the great plants about him receding in the violet shade, and far overhead the multiple transparency of huge leaves filtering the sunshine to the solemn splendour of twilight in which he walked. Whenever he felt able he ran again; the ground continued soft and springy, covered with the same resilient weed which was the first thing his hands had touched in Malacandra. Once or twice a small red creature scuttled across his path, but otherwise there seemed to be no life stirring in the wood; nothing to fear—except the fact of wandering unprovisioned and alone in a forest of unknown vegetation thousands or millions of miles beyond reach or knowledge of man."

	fmt.Println("测试文本（C. S. Lewis,《沉寂的星球》）:")
	fmt.Printf("  前100个字符: %q...\n", testText[:100])
	fmt.Println()

	// 调用函数统计词频
	wordFreq := countWordFrequency(testText)
	fmt.Printf("  总单词数: %d\n", len(wordFreq))
	fmt.Println()

	// 打印出现次数不止一次的单词及其词频
	fmt.Println("出现次数不止一次的单词及其词频:")
	for word, count := range wordFreq {
		if count > 1 {
			fmt.Printf("  %q: %d\n", word, count)
		}
	}
	fmt.Println()

	// ============================================
	// 6. 映射的其他应用场景
	// ============================================
	fmt.Println("=== 6. 映射的其他应用场景 ===")

	// 场景1：集合（使用映射模拟集合）
	fmt.Println("场景1：集合（使用映射模拟集合）")
	set := make(map[float64]bool)
	uniqueTemps := []float64{-28.0, 32.0, -31.0, -29.0, -28.0, 32.0}
	for _, temp := range uniqueTemps {
		set[temp] = true
	}
	fmt.Printf("  原始数据: %v\n", uniqueTemps)
	fmt.Printf("  集合: %v\n", set)

	// 判断32.0是否是集合成员
	if set[32.0] {
		fmt.Println("  32.0是集合成员")
	} else {
		fmt.Println("  32.0不是集合成员")
	}
	fmt.Println()

	// 场景2：索引映射
	fmt.Println("场景2：索引映射（值到索引的映射）")
	items := []string{"apple", "banana", "cherry", "apple", "banana"}
	indexMap := make(map[string][]int)
	for i, item := range items {
		indexMap[item] = append(indexMap[item], i)
	}
	fmt.Printf("  数组: %v\n", items)
	fmt.Printf("  索引映射: %v\n", indexMap)
	fmt.Println()

	// ============================================
	// 7. 小结
	// ============================================
	fmt.Println("=== 7. 小结 ===")
	fmt.Println()
	fmt.Println("1. 逗号与ok语法:")
	fmt.Println("   ✅ 使用 value, ok := map[key] 检查键是否存在")
	fmt.Println("   ✅ 可以区分'键不存在'和'键存在但值为零值'")
	fmt.Println("   ✅ 第二个变量可以使用任何名字")
	fmt.Println()
	fmt.Println("2. 映射不会被复制:")
	fmt.Println("   ✅ 映射是引用类型，赋值时共享底层数据")
	fmt.Println("   ✅ 修改一个映射会影响所有共享的映射")
	fmt.Println("   ✅ 函数参数传递也是引用")
	fmt.Println()
	fmt.Println("3. 使用 make 函数预分配:")
	fmt.Println("   ✅ make(map[KeyType]ValueType, capacity)")
	fmt.Println("   ✅ 预分配可以减少后续扩容的开销")
	fmt.Println("   ✅ 初始长度为0")
	fmt.Println()
	fmt.Println("4. 使用映射进行计数:")
	fmt.Println("   ✅ 非常适合统计频率")
	fmt.Println("   ✅ 比切片更适合非结构化数据")
	fmt.Println("   ✅ map[key]++ 会自动处理键不存在的情况")
	fmt.Println()
	fmt.Println("5. 映射的核心特点:")
	fmt.Println("   ✅ 映射是非结构化数据的多用途收集器")
	fmt.Println("   ✅ 复合字面量是初始化映射的方便手段")
	fmt.Println("   ✅ 使用 range 可以对映射进行迭代")
	fmt.Println("   ✅ 映射在被赋值或传递时共享相同的底层数据")
	fmt.Println("   ✅ 通过组合方式使用收集器可以进一步提升威力")
	fmt.Println()
}

// ============================================
// 辅助函数
// ============================================

// modifyMap 修改映射（演示引用传递）
func modifyMap(m map[string]int) {
	m["c"] = 3
	fmt.Printf("    函数内部添加元素: %v\n", m)
}

// countWordFrequency 统计文本中单词的频率
// 将文本转换为小写，移除标点符号，返回词频映射
func countWordFrequency(text string) map[string]int {
	// 转换为小写
	text = strings.ToLower(text)

	// 移除标点符号（简单处理：只保留字母和空格）
	var cleaned strings.Builder
	for _, char := range text {
		if (char >= 'a' && char <= 'z') || char == ' ' {
			cleaned.WriteRune(char)
		} else {
			cleaned.WriteRune(' ') // 将标点符号替换为空格
		}
	}

	// 使用 Fields 分割单词（会自动处理多个空格）
	words := strings.Fields(cleaned.String())

	// 统计词频
	frequency := make(map[string]int)
	for _, word := range words {
		// 使用 Trim 移除可能的空格（虽然 Fields 已经处理了）
		word = strings.TrimSpace(word)
		if word != "" {
			frequency[word]++
		}
	}

	return frequency
}

