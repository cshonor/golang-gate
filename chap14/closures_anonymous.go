// 示例：闭包和匿名函数
// 演示匿名函数（函数字面量）的定义、使用和闭包的工作原理

package main

import (
	"fmt"
	"time"
)

func main() {
	// ============================================
	// 1. 匿名函数（函数字面量）基础
	// ============================================
	fmt.Println("=== 匿名函数（函数字面量）基础 ===")

	// 方式1：直接调用匿名函数（立即执行）
	func() {
		fmt.Println("这是一个匿名函数，立即执行")
	}()

	// 方式2：赋值给变量
	greet := func(name string) {
		fmt.Printf("Hello, %s!\n", name)
	}
	greet("Alice")
	greet("Bob")

	// 方式3：带返回值的匿名函数
	double := func(x int) int {
		return x * 2
	}
	fmt.Printf("double(5) = %d\n", double(5))

	// 方式4：多参数匿名函数
	add := func(a, b int) int {
		return a + b
	}
	fmt.Printf("add(3, 4) = %d\n", add(3, 4))

	fmt.Println()

	// ============================================
	// 2. 闭包基础：捕获外部变量
	// ============================================
	fmt.Println("=== 闭包基础：捕获外部变量 ===")

	// 闭包：匿名函数会保留对外部作用域变量的引用
	x := 10
	addX := func(y int) int {
		return x + y // 捕获外部变量 x
	}
	fmt.Printf("x = %d, addX(5) = %d\n", x, addX(5))

	// 修改外部变量，闭包会看到变化
	x = 20
	fmt.Printf("修改 x = %d 后, addX(5) = %d\n", x, addX(5))

	fmt.Println()

	// ============================================
	// 3. 闭包示例：计数器
	// ============================================
	fmt.Println("=== 闭包示例：计数器 ===")

	// 创建计数器闭包
	counter := func() func() int {
		count := 0
		// 匿名函数捕获了外部变量 count
		return func() int {
			count++
			return count
		}
	}()

	fmt.Printf("counter() = %d\n", counter())
	fmt.Printf("counter() = %d\n", counter())
	fmt.Printf("counter() = %d\n", counter())

	// 创建多个独立的计数器
	counter1 := func() func() int {
		count := 0
		return func() int {
			count++
			return count
		}
	}()

	counter2 := func() func() int {
		count := 0
		return func() int {
			count++
			return count
		}
	}()

	fmt.Println("\n多个独立的计数器:")
	fmt.Printf("counter1() = %d\n", counter1())
	fmt.Printf("counter2() = %d\n", counter2())
	fmt.Printf("counter1() = %d\n", counter1())
	fmt.Printf("counter2() = %d\n", counter2())

	fmt.Println()

	// ============================================
	// 4. 闭包示例：累加器
	// ============================================
	fmt.Println("=== 闭包示例：累加器 ===")

	// 创建累加器闭包
	accumulator := func(initial int) func(int) int {
		sum := initial
		return func(x int) int {
			sum += x
			return sum
		}
	}

	acc1 := accumulator(10)
	fmt.Printf("acc1(5) = %d\n", acc1(5))
	fmt.Printf("acc1(3) = %d\n", acc1(3))
	fmt.Printf("acc1(7) = %d\n", acc1(7))

	acc2 := accumulator(0)
	fmt.Printf("acc2(10) = %d\n", acc2(10))
	fmt.Printf("acc2(20) = %d\n", acc2(20))

	fmt.Println()

	// ============================================
	// 5. 闭包示例：延迟执行
	// ============================================
	fmt.Println("=== 闭包示例：延迟执行 ===")

	// 闭包可以捕获函数参数
	delayedPrint := func(message string) func() {
		return func() {
			fmt.Println("延迟执行:", message)
		}
	}

	printer1 := delayedPrint("第一条消息")
	printer2 := delayedPrint("第二条消息")

	// 稍后执行
	printer1()
	printer2()

	fmt.Println()

	// ============================================
	// 6. 闭包示例：函数工厂
	// ============================================
	fmt.Println("=== 闭包示例：函数工厂 ===")

	// 创建乘法器工厂
	createMultiplier := func(factor int) func(int) int {
		return func(x int) int {
			return x * factor
		}
	}

	doubleFunc := createMultiplier(2)
	tripleFunc := createMultiplier(3)
	quadrupleFunc := createMultiplier(4)

	fmt.Printf("doubleFunc(5) = %d\n", doubleFunc(5))
	fmt.Printf("tripleFunc(5) = %d\n", tripleFunc(5))
	fmt.Printf("quadrupleFunc(5) = %d\n", quadrupleFunc(5))

	// 创建前缀生成器
	createPrefixer := func(prefix string) func(string) string {
		return func(s string) string {
			return prefix + s
		}
	}

	helloPrefixer := createPrefixer("Hello, ")
	goodbyePrefixer := createPrefixer("Goodbye, ")

	fmt.Printf("%s\n", helloPrefixer("Alice"))
	fmt.Printf("%s\n", goodbyePrefixer("Bob"))

	fmt.Println()

	// ============================================
	// 7. 闭包示例：回调函数
	// ============================================
	fmt.Println("=== 闭包示例：回调函数 ===")

	// 模拟异步操作
	processAsync := func(data string, callback func(string)) {
		// 模拟处理时间
		time.Sleep(100 * time.Millisecond)
		result := fmt.Sprintf("处理完成: %s", data)
		callback(result)
	}

	// 使用闭包作为回调
	processAsync("数据1", func(result string) {
		fmt.Printf("回调1收到: %s\n", result)
	})

	// 闭包可以捕获外部变量
	processedCount := 0
	processAsync("数据2", func(result string) {
		processedCount++
		fmt.Printf("回调2收到: %s (已处理 %d 次)\n", result, processedCount)
	})

	fmt.Println()

	// ============================================
	// 8. 闭包的注意事项：循环变量捕获
	// ============================================
	fmt.Println("=== 闭包的注意事项：循环变量捕获 ===")

	fmt.Println("问题示例：")
	// 错误：所有闭包都捕获了同一个 i
	var funcs1 []func() int
	for i := 0; i < 3; i++ {
		funcs1 = append(funcs1, func() int {
			return i // 所有闭包都引用同一个 i
		})
	}
	fmt.Println("错误方式（所有闭包返回相同的值）:")
	for _, f := range funcs1 {
		fmt.Printf("  %d\n", f()) // 都输出 3
	}

	fmt.Println("\n正确方式（每个闭包捕获不同的值）:")
	// 正确：通过参数传递，每个闭包捕获不同的值
	var funcs2 []func() int
	for i := 0; i < 3; i++ {
		j := i // 创建局部变量副本
		funcs2 = append(funcs2, func() int {
			return j // 每个闭包捕获自己的 j 副本
		})
	}
	for _, f := range funcs2 {
		fmt.Printf("  %d\n", f()) // 输出 0, 1, 2
	}

	// 或者通过参数传递
	var funcs3 []func() int
	for i := 0; i < 3; i++ {
		func(i int) {
			funcs3 = append(funcs3, func() int {
				return i
			})
		}(i)
	}
	fmt.Println("通过参数传递:")
	for _, f := range funcs3 {
		fmt.Printf("  %d\n", f())
	}

	fmt.Println()

	// ============================================
	// 9. 闭包的实际应用：配置函数
	// ============================================
	fmt.Println("=== 闭包的实际应用：配置函数 ===")

	// 创建配置器闭包
	createConfig := func(defaultValue int) func(...int) int {
		value := defaultValue
		return func(opts ...int) int {
			if len(opts) > 0 {
				value = opts[0]
			}
			return value
		}
	}

	config := createConfig(100)
	fmt.Printf("默认值: %d\n", config())
	fmt.Printf("设置新值: %d\n", config(200))
	fmt.Printf("当前值: %d\n", config())

	fmt.Println()

	// ============================================
	// 10. 闭包的实际应用：中间件模式
	// ============================================
	fmt.Println("=== 闭包的实际应用：中间件模式 ===")

	// 创建中间件链
	createMiddleware := func(name string) func(func()) func() {
		return func(next func()) func() {
			return func() {
				fmt.Printf("中间件 %s: 开始\n", name)
				next()
				fmt.Printf("中间件 %s: 结束\n", name)
			}
		}
	}

	handler := func() {
		fmt.Println("  处理请求")
	}

	middleware1 := createMiddleware("日志")
	middleware2 := createMiddleware("认证")

	// 组合中间件
	wrappedHandler := middleware1(middleware2(handler))
	wrappedHandler()

	fmt.Println()

	// ============================================
	// 11. 总结
	// ============================================
	fmt.Println("=== 总结 ===")
	fmt.Println("1. 匿名函数：没有名字的函数，可以直接调用或赋值给变量")
	fmt.Println("2. 闭包：匿名函数捕获外部作用域的变量，形成'函数+引用环境'的组合")
	fmt.Println("3. 闭包的用途：")
	fmt.Println("   - 创建有状态的函数（计数器、累加器）")
	fmt.Println("   - 实现回调函数")
	fmt.Println("   - 函数工厂模式")
	fmt.Println("   - 中间件模式")
	fmt.Println("4. 注意事项：")
	fmt.Println("   - 循环变量捕获问题：需要创建局部变量副本")
	fmt.Println("   - 闭包会持有外部变量的引用，可能导致内存泄漏")
	fmt.Println("   - 在并发场景下需要注意竞态条件")
}

