// 示例：goroutine 和通道
// 演示 Go 语言中 goroutine 和通道的基本用法
// 包括启动 goroutine、通道通信、流水线模式等

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// ============================================
	// 1. 启动 Goroutine
	// ============================================
	fmt.Println("=== 1. 启动 Goroutine ===")
	
	// 1.1 基本启动
	fmt.Println("\n1.1 基本启动:")
	go sayHello("Alice")
	go sayHello("Bob")
	time.Sleep(100 * time.Millisecond)  // 等待 goroutine 完成
	fmt.Println()
	
	// 1.2 启动多个 Goroutine
	fmt.Println("1.2 启动多个 Goroutine:")
	for i := 0; i < 5; i++ {
		go func(id int) {
			fmt.Printf("Goroutine %d 正在运行\n", id)
		}(i)
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println()
	
	// 1.3 使用 WaitGroup 等待 Goroutine
	fmt.Println("1.3 使用 WaitGroup 等待 Goroutine:")
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("任务 %d 完成\n", id)
		}(i)
	}
	wg.Wait()  // 等待所有 goroutine 完成
	fmt.Println("所有任务完成！")
	fmt.Println()
	
	// ============================================
	// 2. 通道基本操作
	// ============================================
	fmt.Println("=== 2. 通道基本操作 ===")
	
	// 2.1 无缓冲通道
	fmt.Println("\n2.1 无缓冲通道:")
	ch1 := make(chan string)
	go func() {
		ch1 <- "Hello from goroutine!"
	}()
	message := <-ch1
	fmt.Printf("收到消息: %s\n", message)
	fmt.Println()
	
	// 2.2 缓冲通道
	fmt.Println("2.2 缓冲通道:")
	ch2 := make(chan int, 3)  // 缓冲区大小为 3
	ch2 <- 1
	ch2 <- 2
	ch2 <- 3
	fmt.Printf("缓冲区已满，可以发送 3 个值\n")
	fmt.Printf("接收值: %d\n", <-ch2)
	fmt.Printf("接收值: %d\n", <-ch2)
	fmt.Printf("接收值: %d\n", <-ch2)
	fmt.Println()
	
	// 2.3 关闭通道
	fmt.Println("2.3 关闭通道:")
	ch3 := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch3 <- i
		}
		close(ch3)  // 关闭通道
	}()
	
	for value := range ch3 {
		fmt.Printf("接收值: %d\n", value)
	}
	fmt.Println("通道已关闭")
	fmt.Println()
	
	// 2.4 检查通道是否关闭
	fmt.Println("2.4 检查通道是否关闭:")
	ch4 := make(chan int)
	go func() {
		ch4 <- 42
		close(ch4)
	}()
	
	value, ok := <-ch4
	fmt.Printf("值: %d, 通道打开: %t\n", value, ok)
	
	value, ok = <-ch4
	fmt.Printf("值: %d, 通道打开: %t (已关闭)\n", value, ok)
	fmt.Println()
	
	// ============================================
	// 3. 通道方向
	// ============================================
	fmt.Println("=== 3. 通道方向 ===")
	
	// 3.1 只发送通道
	fmt.Println("\n3.1 只发送通道:")
	ch5 := make(chan int)
	go sendOnly(ch5)
	value = <-ch5
	fmt.Printf("收到值: %d\n", value)
	fmt.Println()
	
	// 3.2 只接收通道
	fmt.Println("3.2 只接收通道:")
	ch6 := make(chan int)
	go func() {
		ch6 <- 100
		close(ch6)
	}()
	receiveOnly(ch6)
	fmt.Println()
	
	// ============================================
	// 4. Select 语句
	// ============================================
	fmt.Println("=== 4. Select 语句 ===")
	
	// 4.1 多通道监听
	fmt.Println("\n4.1 多通道监听:")
	ch7 := make(chan string)
	ch8 := make(chan string)
	
	go func() {
		time.Sleep(50 * time.Millisecond)
		ch7 <- "来自 ch7"
	}()
	
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch8 <- "来自 ch8"
	}()
	
	select {
	case msg := <-ch7:
		fmt.Printf("收到: %s\n", msg)
	case msg := <-ch8:
		fmt.Printf("收到: %s\n", msg)
	}
	fmt.Println()
	
	// 4.2 超时处理
	fmt.Println("4.2 超时处理:")
	ch9 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		ch9 <- "结果"
	}()
	
	select {
	case result := <-ch9:
		fmt.Printf("收到结果: %s\n", result)
	case <-time.After(1 * time.Second):
		fmt.Println("超时！")
	}
	fmt.Println()
	
	// 4.3 非阻塞操作
	fmt.Println("4.3 非阻塞操作:")
	ch10 := make(chan int, 2)
	ch10 <- 1
	ch10 <- 2
	
	select {
	case ch10 <- 3:
		fmt.Println("发送成功")
	default:
		fmt.Println("通道已满，无法发送（非阻塞）")
	}
	fmt.Println()
	
	// ============================================
	// 5. 通道流水线
	// ============================================
	fmt.Println("=== 5. 通道流水线 ===")
	
	// 5.1 基本流水线
	fmt.Println("\n5.1 基本流水线（生成 → 平方 → 消费）:")
	numbers := generate(2, 3, 4, 5)
	squares := square(numbers)
	
	for result := range squares {
		fmt.Printf("平方结果: %d\n", result)
	}
	fmt.Println()
	
	// 5.2 多阶段流水线
	fmt.Println("5.2 多阶段流水线（生成 → 平方 → 加倍 → 消费）:")
	numbers2 := generate(1, 2, 3, 4, 5)
	squares2 := square(numbers2)
	doubles := double(squares2)
	
	for result := range doubles {
		fmt.Printf("最终结果: %d\n", result)
	}
	fmt.Println()
	
	// ============================================
	// 6. 扇出和扇入模式
	// ============================================
	fmt.Println("=== 6. 扇出和扇入模式 ===")
	
	// 6.1 扇出：一个通道分发给多个 goroutine
	fmt.Println("\n6.1 扇出模式:")
	in := generate(1, 2, 3, 4, 5)
	out1 := make(chan int)
	out2 := make(chan int)
	
	go fanOut(in, out1, out2)
	
	var wg2 sync.WaitGroup
	wg2.Add(2)
	
	go func() {
		defer wg2.Done()
		for v := range out1 {
			fmt.Printf("输出1: %d\n", v)
		}
	}()
	
	go func() {
		defer wg2.Done()
		for v := range out2 {
			fmt.Printf("输出2: %d\n", v)
		}
	}()
	
	wg2.Wait()
	fmt.Println()
	
	// 6.2 扇入：多个通道合并为一个
	fmt.Println("6.2 扇入模式:")
	input1 := generate(1, 3, 5)
	input2 := generate(2, 4, 6)
	merged := fanIn(input1, input2)
	
	for v := range merged {
		fmt.Printf("合并输出: %d\n", v)
	}
	fmt.Println()
	
	// ============================================
	// 7. 工作池模式
	// ============================================
	fmt.Println("=== 7. 工作池模式 ===")
	
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	
	// 启动 3 个 worker
	numWorkers := 3
	var wg3 sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg3.Add(1)
		go worker(i, jobs, results, &wg3)
	}
	
	// 发送任务
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)
	
	// 等待所有 worker 完成
	go func() {
		wg3.Wait()
		close(results)
	}()
	
	// 收集结果
	for result := range results {
		fmt.Printf("任务结果: %d\n", result)
	}
	fmt.Println()
	
	// ============================================
	// 8. 常见错误示例（已注释）
	// ============================================
	fmt.Println("=== 8. 常见错误示例 ===")
	fmt.Println("以下代码会导致问题，已注释：")
	
	fmt.Println("\n❌ 错误1：忘记等待 goroutine")
	fmt.Println("   go doSomething()")
	fmt.Println("   // 程序可能立即退出，goroutine 来不及执行")
	
	fmt.Println("\n❌ 错误2：通道死锁（无缓冲通道）")
	fmt.Println("   ch := make(chan int)")
	fmt.Println("   ch <- 42      // 阻塞，因为没有接收者")
	fmt.Println("   value := <-ch // 永远执行不到")
	
	fmt.Println("\n❌ 错误3：关闭已关闭的通道")
	fmt.Println("   close(ch)")
	fmt.Println("   close(ch)  // panic!")
	
	fmt.Println("\n✅ 正确做法：")
	fmt.Println("   - 使用 WaitGroup 或通道等待 goroutine")
	fmt.Println("   - 确保发送和接收在不同的 goroutine")
	fmt.Println("   - 只关闭一次通道，或使用 sync.Once")
}

// ============================================
// 辅助函数
// ============================================

// sayHello 简单的问候函数
func sayHello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// sendOnly 只发送通道示例
func sendOnly(ch chan<- int) {
	ch <- 42
	// value := <-ch  // 编译错误：不能接收
}

// receiveOnly 只接收通道示例
func receiveOnly(ch <-chan int) {
	value := <-ch
	fmt.Printf("只接收通道收到值: %d\n", value)
	// ch <- 42  // 编译错误：不能发送
}

// ============================================
// 流水线函数
// ============================================

// generate 生成数字
func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// square 计算平方
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// double 加倍
func double(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * 2
		}
		close(out)
	}()
	return out
}

// ============================================
// 扇出和扇入
// ============================================

// fanOut 扇出：一个通道分发给多个通道
func fanOut(in <-chan int, out1, out2 chan<- int) {
	defer close(out1)
	defer close(out2)
	
	for value := range in {
		select {
		case out1 <- value:
		case out2 <- value:
		}
	}
}

// fanIn 扇入：多个通道合并为一个
func fanIn(input1, input2 <-chan int) <-chan int {
	out := make(chan int)
	
	go func() {
		defer close(out)
		for {
			select {
			case v, ok := <-input1:
				if !ok {
					input1 = nil  // 禁用这个 case
				} else {
					out <- v
				}
			case v, ok := <-input2:
				if !ok {
					input2 = nil  // 禁用这个 case
				} else {
					out <- v
				}
			}
			
			// 如果两个通道都关闭了，退出
			if input1 == nil && input2 == nil {
				return
			}
		}
	}()
	
	return out
}

// ============================================
// 工作池模式
// ============================================

// worker 工作池中的 worker
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d 处理任务 %d\n", id, job)
		time.Sleep(50 * time.Millisecond)  // 模拟处理时间
		results <- job * job  // 返回结果
	}
}

