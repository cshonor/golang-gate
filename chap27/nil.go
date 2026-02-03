// 示例：Go 语言的 nil
// 演示 nil 的本质、问题、Go的改进和常见陷阱

package main

import "fmt"

func main() {
	// ============================================
	// 1. nil 的本质
	// ============================================
	fmt.Println("=== 1. nil 的本质 ===")

	fmt.Println("在Go里，nil是多个类型的'零值'")
	fmt.Println("它不是一个单一的类型")
	fmt.Println("而是指针、切片、映射、通道、函数和接口的默认零值")
	fmt.Println("这和很多语言里单一的NULL不同")
	fmt.Println("也是Go的nil容易让人困惑的原因之一")
	fmt.Println()

	// 演示不同类型的nil
	fmt.Println("演示不同类型的nil:")
	var ptr *int
	var slice []int
	var m map[string]int
	var ch chan int
	var fn func()
	var iface interface{}

	fmt.Printf("  指针 *int: %v, nil=%t\n", ptr, ptr == nil)
	fmt.Printf("  切片 []int: %v, nil=%t\n", slice, slice == nil)
	fmt.Printf("  映射 map[string]int: %v, nil=%t\n", m, m == nil)
	fmt.Printf("  通道 chan int: %v, nil=%t\n", ch, ch == nil)
	fmt.Printf("  函数 func(): %v, nil=%t\n", fn, fn == nil)
	fmt.Printf("  接口 interface{}: %v, nil=%t\n", iface, iface == nil)
	fmt.Println()

	// ============================================
	// 2. "十亿美元错误"的背景
	// ============================================
	fmt.Println("=== 2. '十亿美元错误'的背景 ===")

	fmt.Println("Tony Hoare在2009年的演讲中提到")
	fmt.Println("他在1965年发明的空引用（NULL）")
	fmt.Println("这个设计导致了无数空指针异常")
	fmt.Println("他本人称其为'十亿美元的错误'")
	fmt.Println()

	fmt.Println("Go语言正是为了避免这个问题")
	fmt.Println("对nil做了更安全的设计")
	fmt.Println()

	// ============================================
	// 3. Go 对 nil 的改进
	// ============================================
	fmt.Println("=== 3. Go 对 nil 的改进 ===")

	fmt.Println("改进1：类型安全")
	fmt.Println("  nil必须和特定类型绑定")
	fmt.Println("  不能像某些语言那样随意赋值给不同类型的变量")
	fmt.Println()

	// 演示类型安全
	fmt.Println("演示类型安全:")
	var intPtr *int
	var stringPtr *string
	fmt.Printf("  *int 指针: %v\n", intPtr)
	fmt.Printf("  *string 指针: %v\n", stringPtr)
	// intPtr = stringPtr // ❌ 错误：类型不匹配
	fmt.Println("  // intPtr = stringPtr  // ❌ 错误：类型不匹配")
	fmt.Println()

	fmt.Println("改进2：默认零值")
	fmt.Println("  声明变量时如果不初始化，会自动赋予对应类型的零值")
	fmt.Println("  减少了nil出现的场景")
	fmt.Println()

	// 演示默认零值
	fmt.Println("演示默认零值:")
	var num int
	var str string
	var b bool
	fmt.Printf("  int 零值: %d\n", num)
	fmt.Printf("  string 零值: %q\n", str)
	fmt.Printf("  bool 零值: %t\n", b)
	fmt.Println("  说明：基本类型有零值，不需要nil")
	fmt.Println()

	fmt.Println("改进3：接口的nil特性")
	fmt.Println("  接口的nil比较特殊")
	fmt.Println("  只有当接口的类型和值都为nil时，整个接口才是nil")
	fmt.Println("  这是常见的坑点")
	fmt.Println()

	// ============================================
	// 4. 处理没有值的情况
	// ============================================
	fmt.Println("=== 4. 处理没有值的情况 ===")

	fmt.Println("学会在代码中检查nil，避免空指针panic")
	fmt.Println()

	// 示例1：安全使用指针
	fmt.Println("示例1：安全使用指针")
	var safePtr *int
	if safePtr == nil {
		fmt.Println("  指针是nil，需要初始化")
		safePtr = new(int)
		*safePtr = 42
	}
	fmt.Printf("  安全使用: %d\n", *safePtr)
	fmt.Println()

	// 示例2：安全使用切片
	fmt.Println("示例2：安全使用切片")
	var safeSlice []int
	if safeSlice == nil {
		fmt.Println("  切片是nil，可以安全使用")
		safeSlice = make([]int, 0)
	}
	fmt.Printf("  切片长度: %d\n", len(safeSlice))
	fmt.Println()

	// 示例3：安全使用映射
	fmt.Println("示例3：安全使用映射")
	var safeMap map[string]int
	if safeMap == nil {
		fmt.Println("  映射是nil，需要初始化")
		safeMap = make(map[string]int)
	}
	safeMap["key"] = 100
	fmt.Printf("  映射: %v\n", safeMap)
	fmt.Println()

	// ============================================
	// 5. 理解nil引发的问题
	// ============================================
	fmt.Println("=== 5. 理解nil引发的问题 ===")

	// 问题1：接口的nil陷阱
	fmt.Println("问题1：接口的nil陷阱")
	var nilPtr *int
	var nilInterface interface{} = nilPtr
	fmt.Printf("  nilPtr == nil: %t\n", nilPtr == nil)
	fmt.Printf("  nilInterface == nil: %t (注意：false!)\n", nilInterface == nil)
	fmt.Printf("  nilInterface的类型: %T\n", nilInterface)
	fmt.Println("  说明：接口包含类型信息，即使值是nil，接口也不是nil")
	fmt.Println()

	// 问题2：nil切片 vs 空切片
	fmt.Println("问题2：nil切片 vs 空切片")
	var nilSlice []int
	emptySlice := []int{}
	madeSlice := make([]int, 0)

	fmt.Printf("  nil切片: %v, len=%d, cap=%d, nil=%t\n", nilSlice, len(nilSlice), cap(nilSlice), nilSlice == nil)
	fmt.Printf("  空切片: %v, len=%d, cap=%d, nil=%t\n", emptySlice, len(emptySlice), cap(emptySlice), emptySlice == nil)
	fmt.Printf("  make切片: %v, len=%d, cap=%d, nil=%t\n", madeSlice, len(madeSlice), cap(madeSlice), madeSlice == nil)
	fmt.Println()

	fmt.Println("  说明：")
	fmt.Println("    - nil切片：未初始化，可以安全使用（append等）")
	fmt.Println("    - 空切片：已初始化但为空，不是nil")
	fmt.Println("    - 两者在大多数操作中行为相同")
	fmt.Println()

	// 问题3：nil映射的操作
	fmt.Println("问题3：nil映射的操作")
	var nilMap map[string]int
	fmt.Printf("  nil映射: %v, nil=%t\n", nilMap, nilMap == nil)
	// nilMap["key"] = 1 // ❌ 错误：会panic
	fmt.Println("  // nilMap[\"key\"] = 1  // ❌ 错误：会panic")
	value, exists := nilMap["key"]
	fmt.Printf("  读取nil映射: value=%d, exists=%t (可以读取，返回零值)\n", value, exists)
	fmt.Println()

	// 问题4：nil通道
	fmt.Println("问题4：nil通道")
	var nilChan chan int
	fmt.Printf("  nil通道: %v, nil=%t\n", nilChan, nilChan == nil)
	fmt.Println("  // <-nilChan  // ❌ 错误：会阻塞")
	fmt.Println("  // nilChan <- 1  // ❌ 错误：会阻塞")
	fmt.Println("  说明：nil通道会永远阻塞")
	fmt.Println()

	// ============================================
	// 6. 接口的nil陷阱详解
	// ============================================
	fmt.Println("=== 6. 接口的nil陷阱详解 ===")

	fmt.Println("这是Go中最常见的nil陷阱")
	fmt.Println()

	// 示例1：接口nil的判断
	fmt.Println("示例1：接口nil的判断")
	var nilInt *int
	var intInterface interface{} = nilInt

	fmt.Printf("  nilInt == nil: %t\n", nilInt == nil)
	fmt.Printf("  intInterface == nil: %t (false!)\n", intInterface == nil)
	fmt.Printf("  intInterface的类型: %T\n", intInterface)
	fmt.Println()

	// 正确的nil检查
	fmt.Println("正确的nil检查:")
	if intInterface == nil {
		fmt.Println("  接口是nil")
	} else {
		fmt.Println("  接口不是nil（即使值是nil）")
		// 需要检查值是否为nil
		if intInterface.(*int) == nil {
			fmt.Println("    但接口的值是nil")
		}
	}
	fmt.Println()

	// 示例2：函数返回接口
	fmt.Println("示例2：函数返回接口")
	result := getNilInterface()
	if result == nil {
		fmt.Println("  返回的接口是nil")
	} else {
		fmt.Printf("  返回的接口不是nil: %T, %v\n", result, result)
	}
	fmt.Println()

	// ============================================
	// 7. nil切片与空切片的区别
	// ============================================
	fmt.Println("=== 7. nil切片与空切片的区别 ===")

	fmt.Println("虽然行为相似，但有一些细微差别")
	fmt.Println()

	// 对比
	fmt.Println("对比:")
	nilSlice2 := []int(nil)
	emptySlice2 := []int{}
	madeSlice2 := make([]int, 0)

	fmt.Printf("  nil切片: %v, len=%d, nil=%t\n", nilSlice2, len(nilSlice2), nilSlice2 == nil)
	fmt.Printf("  空切片: %v, len=%d, nil=%t\n", emptySlice2, len(emptySlice2), emptySlice2 == nil)
	fmt.Printf("  make切片: %v, len=%d, nil=%t\n", madeSlice2, len(madeSlice2), madeSlice2 == nil)
	fmt.Println()

	// 都可以使用append
	fmt.Println("都可以使用append:")
	nilSlice2 = append(nilSlice2, 1, 2, 3)
	emptySlice2 = append(emptySlice2, 1, 2, 3)
	madeSlice2 = append(madeSlice2, 1, 2, 3)
	fmt.Printf("  append后nil切片: %v\n", nilSlice2)
	fmt.Printf("  append后空切片: %v\n", emptySlice2)
	fmt.Printf("  append后make切片: %v\n", madeSlice2)
	fmt.Println()

	// JSON序列化的区别
	fmt.Println("JSON序列化的区别:")
	fmt.Println("  nil切片: null")
	fmt.Println("  空切片: []")
	fmt.Println("  说明：在JSON序列化时有区别")
	fmt.Println()

	// ============================================
	// 8. nil的常见陷阱和避坑方法
	// ============================================
	fmt.Println("=== 8. nil的常见陷阱和避坑方法 ===")

	fmt.Println("陷阱1：接口nil判断")
	fmt.Println("  问题：接口包含类型信息，值nil不等于接口nil")
	fmt.Println("  避坑：检查接口的值是否为nil")
	fmt.Println()

	fmt.Println("陷阱2：nil映射写入")
	fmt.Println("  问题：向nil映射写入会panic")
	fmt.Println("  避坑：使用前检查nil或使用make初始化")
	fmt.Println()

	fmt.Println("陷阱3：nil通道操作")
	fmt.Println("  问题：nil通道会永远阻塞")
	fmt.Println("  避坑：使用前检查nil或使用make初始化")
	fmt.Println()

	fmt.Println("陷阱4：nil指针解引用")
	fmt.Println("  问题：解引用nil指针会panic")
	fmt.Println("  避坑：使用前检查nil")
	fmt.Println()

	// ============================================
	// 9. 实际应用示例
	// ============================================
	fmt.Println("=== 9. 实际应用示例 ===")

	// 示例1：安全的函数返回
	fmt.Println("示例1：安全的函数返回")
	value1, err := safeGetValue("key1")
	if err != nil {
		fmt.Printf("  错误: %v\n", err)
	} else {
		fmt.Printf("  值: %d\n", value1)
	}
	fmt.Println()

	// 示例2：nil检查工具函数
	fmt.Println("示例2：nil检查工具函数")
	var testPtr *int
	if isNil(testPtr) {
		fmt.Println("  指针是nil")
	}
	testPtr = new(int)
	if isNil(testPtr) {
		fmt.Println("  指针是nil")
	} else {
		fmt.Println("  指针不是nil")
	}
	fmt.Println()

	// 示例3：处理可能为nil的接口
	fmt.Println("示例3：处理可能为nil的接口")
	var nilValue *string
	processInterface(nilValue)
	value2 := "hello"
	processInterface(&value2)
	fmt.Println()

	// ============================================
	// 10. 总结
	// ============================================
	fmt.Println("=== 10. 总结 ===")
	fmt.Println()
	fmt.Println("1. nil的本质:")
	fmt.Println("   ✅ nil是多个类型的零值")
	fmt.Println("   ✅ 指针、切片、映射、通道、函数、接口的零值")
	fmt.Println()
	fmt.Println("2. Go对nil的改进:")
	fmt.Println("   ✅ 类型安全：nil必须和特定类型绑定")
	fmt.Println("   ✅ 默认零值：减少nil出现的场景")
	fmt.Println("   ✅ 接口nil特性：需要特别注意")
	fmt.Println()
	fmt.Println("3. 处理没有值的情况:")
	fmt.Println("   ✅ 总是检查nil")
	fmt.Println("   ✅ 使用前初始化")
	fmt.Println("   ✅ 避免空指针panic")
	fmt.Println()
	fmt.Println("4. nil引发的问题:")
	fmt.Println("   ✅ 接口nil陷阱：值nil不等于接口nil")
	fmt.Println("   ✅ nil切片vs空切片：行为相似但有区别")
	fmt.Println("   ✅ nil映射写入：会panic")
	fmt.Println("   ✅ nil通道操作：会阻塞")
	fmt.Println()
	fmt.Println("5. 避坑方法:")
	fmt.Println("   ✅ 使用前检查nil")
	fmt.Println("   ✅ 使用make初始化")
	fmt.Println("   ✅ 理解接口nil的特殊性")
	fmt.Println("   ✅ 区分nil切片和空切片")
	fmt.Println()
}

// ============================================
// 辅助函数
// ============================================

// getNilInterface 返回一个包含nil值的接口
func getNilInterface() interface{} {
	var nilPtr *int
	return nilPtr // 返回的不是nil接口，而是包含nil值的接口
}

// safeGetValue 安全获取值
func safeGetValue(key string) (int, error) {
	data := make(map[string]int)
	data["key1"] = 100
	value, exists := data[key]
	if !exists {
		return 0, fmt.Errorf("键 %s 不存在", key)
	}
	return value, nil
}

// isNil 检查值是否为nil（通用方法）
func isNil(v interface{}) bool {
	if v == nil {
		return true
	}
	switch val := v.(type) {
	case *int:
		return val == nil
	case []int:
		return val == nil
	case map[string]int:
		return val == nil
	case chan int:
		return val == nil
	case func():
		return val == nil
	default:
		return false
	}
}

// processInterface 处理可能为nil的接口
func processInterface(v interface{}) {
	if v == nil {
		fmt.Println("  接口是nil")
		return
	}
	// 类型断言检查
	if strPtr, ok := v.(*string); ok {
		if strPtr == nil {
			fmt.Println("  接口的值是nil")
		} else {
			fmt.Printf("  接口的值: %s\n", *strPtr)
		}
	}
}

