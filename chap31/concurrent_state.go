// 示例：并发状态管理
// 演示 Go 语言中如何安全地管理并发状态
// 包括互斥锁、应答通道和服务循环的使用

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// ============================================
	// 1. 竞态条件示例（危险！）
	// ============================================
	fmt.Println("=== 1. 竞态条件示例（危险！）===")
	
	fmt.Println("演示：多个 goroutine 同时修改共享变量，没有同步机制")
	
	var unsafeCounter int
	var wg sync.WaitGroup
	
	// 启动 10 个 goroutine，每个都增加计数器
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				unsafeCounter++  // 竞态条件！
			}
		}()
	}
	
	wg.Wait()
	fmt.Printf("期望值: 10000, 实际值: %d (可能不正确！)\n", unsafeCounter)
	fmt.Println("说明：由于竞态条件，结果可能每次运行都不同")
	fmt.Println()
	
	// ============================================
	// 2. 使用互斥锁保护共享状态
	// ============================================
	fmt.Println("=== 2. 使用互斥锁保护共享状态 ===")
	
	// 2.1 基本互斥锁使用
	fmt.Println("\n2.1 基本互斥锁使用:")
	
	var mu sync.Mutex
	var safeCounter int
	
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mu.Lock()
				safeCounter++
				mu.Unlock()
			}
		}()
	}
	
	wg.Wait()
	fmt.Printf("期望值: 10000, 实际值: %d (正确！)\n", safeCounter)
	fmt.Println()
	
	// 2.2 使用 defer 确保解锁（正确方式）
	fmt.Println("2.2 使用 defer 确保解锁（正确方式）:")
	
	safeCounter = 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 将加锁解锁提取到函数中，在函数级别使用 defer
			incrementWithDefer := func() {
				mu.Lock()
				defer mu.Unlock()  // 正确：在函数级别使用 defer
				safeCounter++
			}
			for j := 0; j < 1000; j++ {
				incrementWithDefer()
			}
		}()
	}
	
	wg.Wait()
	fmt.Printf("计数器值: %d\n", safeCounter)
	fmt.Println("说明：defer 应该在函数级别使用，而不是在循环内")
	fmt.Println()
	
	// 2.3 封装在结构体中的互斥锁
	fmt.Println("2.3 封装在结构体中的互斥锁:")
	
	counter := NewSafeCounter()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}(i)
	}
	
	wg.Wait()
	fmt.Printf("计数器值: %d\n", counter.Value())
	fmt.Println()
	
	// ============================================
	// 3. 使用通道保护共享状态
	// ============================================
	fmt.Println("=== 3. 使用通道保护共享状态 ===")
	
	// 3.1 基本通道模式
	fmt.Println("\n3.1 基本通道模式（应答通道）:")
	
	channelCounter := NewChannelCounter()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				channelCounter.Increment()
			}
		}(i)
	}
	
	wg.Wait()
	fmt.Printf("计数器值: %d\n", channelCounter.Get())
	fmt.Println()
	
	// 3.2 带请求类型的通道服务
	fmt.Println("3.2 带请求类型的通道服务:")
	
	service := NewCounterService()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				service.Add(1)
			}
		}(i)
	}
	
	wg.Wait()
	fmt.Printf("计数器值: %d\n", service.Get())
	service.Stop()
	fmt.Println()
	
	// ============================================
	// 4. 服务循环示例
	// ============================================
	fmt.Println("=== 4. 服务循环示例 ===")
	
	// 4.1 简单的服务循环
	fmt.Println("\n4.1 简单的服务循环:")
	
	requestChan := make(chan string, 10)
	responseChan := make(chan string, 10)
	
	// 启动服务循环
	go serviceLoop(requestChan, responseChan)
	
	// 发送请求
	requestChan <- "请求1"
	requestChan <- "请求2"
	requestChan <- "请求3"
	
	// 接收响应
	for i := 0; i < 3; i++ {
		response := <-responseChan
		fmt.Printf("收到响应: %s\n", response)
	}
	
	close(requestChan)
	time.Sleep(100 * time.Millisecond)  // 等待服务循环结束
	fmt.Println()
	
	// 4.2 带状态的服务循环
	fmt.Println("4.2 带状态的服务循环:")
	
	stateService := NewStateService()
	
	// 启动多个客户端
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			stateService.SetState(id, id*10)
			state := stateService.GetState(id)
			fmt.Printf("客户端 %d: 设置状态为 %d, 获取状态为 %d\n", id, id*10, state)
		}(i)
	}
	
	wg.Wait()
	stateService.Stop()
	fmt.Println()
	
	// ============================================
	// 5. 互斥锁 vs 通道性能对比
	// ============================================
	fmt.Println("=== 5. 互斥锁 vs 通道性能对比 ===")
	
	iterations := 100000
	
	// 互斥锁方式
	start := time.Now()
	mutexCounter := NewSafeCounter()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				mutexCounter.Increment()
			}
		}()
	}
	wg.Wait()
	mutexTime := time.Since(start)
	fmt.Printf("互斥锁方式: %d 次操作，耗时 %v\n", iterations*10, mutexTime)
	
	// 通道方式
	start = time.Now()
	channelCounter2 := NewChannelCounter()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				channelCounter2.Increment()
			}
		}()
	}
	wg.Wait()
	channelTime := time.Since(start)
	fmt.Printf("通道方式: %d 次操作，耗时 %v\n", iterations*10, channelTime)
	fmt.Println("说明：互斥锁通常性能更好，但通道更符合 Go 语言哲学")
	fmt.Println()
	
	// ============================================
	// 6. 常见错误示例（已注释）
	// ============================================
	fmt.Println("=== 6. 常见错误示例 ===")
	fmt.Println("以下代码会导致问题，已注释：")
	
	fmt.Println("\n❌ 错误1：忘记解锁（导致死锁）")
	fmt.Println("   mu.Lock()")
	fmt.Println("   counter++")
	fmt.Println("   // 忘记 mu.Unlock()")
	
	fmt.Println("\n❌ 错误2：重复加锁（导致死锁）")
	fmt.Println("   mu.Lock()")
	fmt.Println("   mu.Lock()  // 死锁！")
	
	fmt.Println("\n❌ 错误3：在持有锁时调用可能阻塞的函数")
	fmt.Println("   mu.Lock()")
	fmt.Println("   result := <-someChannel  // 可能永远阻塞")
	fmt.Println("   mu.Unlock()")
	
	fmt.Println("\n✅ 正确做法：")
	fmt.Println("   - 总是使用 defer 解锁")
	fmt.Println("   - 不要在持有锁时调用未知函数")
	fmt.Println("   - 保持锁的粒度小")
}

// ============================================
// 互斥锁方式：线程安全的计数器
// ============================================

// SafeCounter 使用互斥锁保护的计数器
type SafeCounter struct {
	mu      sync.Mutex
	counter int
}

// NewSafeCounter 创建新的安全计数器
func NewSafeCounter() *SafeCounter {
	return &SafeCounter{}
}

// Increment 增加计数器
func (sc *SafeCounter) Increment() {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.counter++
}

// Add 增加指定值
func (sc *SafeCounter) Add(n int) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.counter += n
}

// Value 获取计数器值
func (sc *SafeCounter) Value() int {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.counter
}

// ============================================
// 通道方式：使用应答通道的计数器
// ============================================

// ChannelCounter 使用通道保护的计数器
type ChannelCounter struct {
	add     chan int
	get     chan chan int
	counter int
}

// NewChannelCounter 创建新的通道计数器
func NewChannelCounter() *ChannelCounter {
	cc := &ChannelCounter{
		add: make(chan int),
		get: make(chan chan int),
	}
	go cc.run()
	return cc
}

// run 运行计数器服务循环
func (cc *ChannelCounter) run() {
	for {
		select {
		case n := <-cc.add:
			cc.counter += n
		case resp := <-cc.get:
			resp <- cc.counter
		}
	}
}

// Increment 增加计数器
func (cc *ChannelCounter) Increment() {
	cc.add <- 1
}

// Add 增加指定值
func (cc *ChannelCounter) Add(n int) {
	cc.add <- n
}

// Get 获取计数器值
func (cc *ChannelCounter) Get() int {
	resp := make(chan int)
	cc.get <- resp
	return <-resp
}

// ============================================
// 服务循环示例
// ============================================

// serviceLoop 简单的服务循环
func serviceLoop(requests <-chan string, responses chan<- string) {
	for req := range requests {
		// 模拟处理请求
		response := fmt.Sprintf("处理完成: %s", req)
		responses <- response
	}
	close(responses)
}

// ============================================
// 带请求类型的通道服务
// ============================================

// CounterRequest 计数器请求类型
type CounterRequest struct {
	operation string
	value     int
	response  chan int
}

// CounterService 使用请求-响应模式的计数器服务
type CounterService struct {
	requests chan CounterRequest
	done     chan struct{}
	counter  int
}

// NewCounterService 创建新的计数器服务
func NewCounterService() *CounterService {
	cs := &CounterService{
		requests: make(chan CounterRequest),
		done:     make(chan struct{}),
	}
	go cs.run()
	return cs
}

// run 运行服务循环
func (cs *CounterService) run() {
	for req := range cs.requests {
		switch req.operation {
		case "add":
			cs.counter += req.value
			req.response <- cs.counter
		case "get":
			req.response <- cs.counter
		}
		close(req.response)
	}
	close(cs.done)
}

// Add 增加计数器
func (cs *CounterService) Add(n int) {
	resp := make(chan int)
	cs.requests <- CounterRequest{"add", n, resp}
	<-resp
}

// Get 获取计数器值
func (cs *CounterService) Get() int {
	resp := make(chan int)
	cs.requests <- CounterRequest{"get", 0, resp}
	return <-resp
}

// Stop 停止服务
func (cs *CounterService) Stop() {
	close(cs.requests)
	<-cs.done
}

// ============================================
// 带状态的服务循环
// ============================================

// StateRequest 状态请求
type StateRequest struct {
	operation string
	key       int
	value     int
	response  chan int
}

// StateService 状态服务
type StateService struct {
	requests chan StateRequest
	done     chan struct{}
	state    map[int]int
}

// NewStateService 创建新的状态服务
func NewStateService() *StateService {
	ss := &StateService{
		requests: make(chan StateRequest),
		done:     make(chan struct{}),
		state:    make(map[int]int),
	}
	go ss.run()
	return ss
}

// run 运行状态服务循环
func (ss *StateService) run() {
	for req := range ss.requests {
		switch req.operation {
		case "set":
			ss.state[req.key] = req.value
			req.response <- req.value
		case "get":
			value, ok := ss.state[req.key]
			if !ok {
				value = 0
			}
			req.response <- value
		}
		close(req.response)
	}
	close(ss.done)
}

// SetState 设置状态
func (ss *StateService) SetState(key, value int) {
	resp := make(chan int)
	ss.requests <- StateRequest{"set", key, value, resp}
	<-resp
}

// GetState 获取状态
func (ss *StateService) GetState(key int) int {
	resp := make(chan int)
	ss.requests <- StateRequest{"get", key, 0, resp}
	return <-resp
}

// Stop 停止服务
func (ss *StateService) Stop() {
	close(ss.requests)
	<-ss.done
}
