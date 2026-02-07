# 第三十章：goroutine 和并发

## 1. 章节概述

### 1.1 核心主题
这是 Go 语言并发编程的入门章节，也是 Go 语言最具特色的核心内容之一。本章介绍如何使用 goroutine 和通道（channel）实现并发编程。

### 1.2 学习目标
1. **学会启动 goroutine**：掌握如何使用 `go` 关键字创建并启动轻量级的并发执行单元。
2. **学会使用通道进行通信**：理解 Go 语言"不要通过共享内存来通信，而要通过通信来共享内存"的并发哲学，掌握通道（channel）的使用方法。
3. **理解通道流水线**：学习如何将多个通道串联起来，构建高效的数据处理流水线，这是 Go 并发模式的重要实践。

### 1.3 开篇比喻：地鼠工厂
- **工厂里的每一只地鼠**：代表一个独立执行的 goroutine
- **位高权重的地鼠**：负责分派任务，对应主 goroutine 或调度器
- **地鼠们通过传递工作成果相互协作**：这正是通道（channel）通信的形象化表达

### 1.4 与后续章节的联系
- **第30章**：介绍了如何启动 goroutine 和使用通道进行通信，解决了"如何并发执行"的问题。
- **第31章**：深入到并发带来的核心挑战——共享状态管理，解决了"如何安全地并发执行"的问题。

## 2. Goroutine（协程）

### 2.1 什么是 Goroutine？
- **Goroutine**：Go 语言中的轻量级线程，由 Go 运行时管理
- **特点**：
  - 启动成本极低：只需要几 KB 的栈空间
  - 由 Go 运行时调度：不是操作系统线程
  - 可以轻松创建成千上万个 goroutine

### 2.2 启动 Goroutine

#### 2.2.1 基本语法
```go
go 函数调用
```

**示例**：
```go
// 启动一个 goroutine 执行函数
go doSomething()

// 启动一个 goroutine 执行匿名函数
go func() {
    fmt.Println("在 goroutine 中执行")
}()
```

#### 2.2.2 重要特点
- **非阻塞**：`go` 关键字会立即返回，不会等待函数执行完成
- **并发执行**：goroutine 和主程序并发运行
- **主程序退出，所有 goroutine 也会退出**

### 2.3 Goroutine 示例

#### 2.3.1 基本启动
```go
func main() {
    go sayHello("Alice")
    go sayHello("Bob")
    time.Sleep(1 * time.Second)  // 等待 goroutine 完成
}

func sayHello(name string) {
    fmt.Printf("Hello, %s!\n", name)
}
```

#### 2.3.2 多个 Goroutine
```go
func main() {
    for i := 0; i < 5; i++ {
        go func(id int) {
            fmt.Printf("Goroutine %d\n", id)
        }(i)
    }
    time.Sleep(1 * time.Second)
}
```

**注意**：如果不等待，主程序可能在这些 goroutine 完成之前就退出了。

## 3. Channel（通道）

### 3.1 什么是 Channel？
- **Channel**：Go 语言中用于 goroutine 之间通信的管道
- **作用**：实现"通过通信来共享内存"，而不是"通过共享内存来通信"
- **类型**：通道是类型化的，只能传递指定类型的数据

### 3.2 创建通道

#### 3.2.1 基本语法
```go
// 创建通道
ch := make(chan 类型)

// 示例
ch := make(chan int)        // 传递 int 类型
ch := make(chan string)     // 传递 string 类型
```

#### 3.2.2 缓冲通道
```go
// 创建带缓冲的通道
ch := make(chan int, 10)  // 缓冲区大小为 10
```

**区别**：
- **无缓冲通道**：发送和接收必须同时准备好，否则会阻塞
- **缓冲通道**：缓冲区未满时可以发送，缓冲区非空时可以接收

### 3.3 通道操作

#### 3.3.1 发送数据
```go
ch <- 值  // 将值发送到通道
```

#### 3.3.2 接收数据
```go
值 := <-ch      // 从通道接收值
值, ok := <-ch  // 接收值，并检查通道是否关闭
```

#### 3.3.3 关闭通道
```go
close(ch)  // 关闭通道
```

**规则**：
- 只有发送方可以关闭通道
- 关闭后的通道不能再发送数据
- 可以从已关闭的通道接收数据（会收到零值和 `false`）

### 3.4 通道方向

#### 3.4.1 只发送通道
```go
func sendOnly(ch chan<- int) {
    ch <- 42  // 只能发送
    // value := <-ch  // 编译错误！
}
```

#### 3.4.2 只接收通道
```go
func receiveOnly(ch <-chan int) {
    value := <-ch  // 只能接收
    // ch <- 42  // 编译错误！
}
```

**好处**：明确通道的用途，提高代码可读性和安全性。

### 3.5 通道示例

#### 3.5.1 基本通信
```go
func main() {
    ch := make(chan string)
    
    go func() {
        ch <- "Hello from goroutine!"
    }()
    
    message := <-ch
    fmt.Println(message)
}
```

#### 3.5.2 多个值传递
```go
func main() {
    ch := make(chan int)
    
    go func() {
        for i := 0; i < 5; i++ {
            ch <- i
        }
        close(ch)
    }()
    
    for value := range ch {
        fmt.Println(value)
    }
}
```

## 4. 通道流水线（Channel Pipeline）

### 4.1 什么是流水线？
- **流水线**：将多个处理步骤串联起来，每个步骤通过通道连接
- **优势**：可以并行处理数据，提高效率
- **模式**：生产者 → 处理器1 → 处理器2 → ... → 消费者

### 4.2 基本流水线模式

#### 4.2.1 三阶段流水线
```go
// 阶段1：生成数据
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

// 阶段2：处理数据（平方）
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

// 阶段3：消费数据
func main() {
    // 构建流水线
    numbers := generate(2, 3, 4, 5)
    squares := square(numbers)
    
    // 消费结果
    for result := range squares {
        fmt.Println(result)
    }
}
```

#### 4.2.2 多阶段流水线
```go
// 可以串联多个处理阶段
numbers := generate(1, 2, 3, 4, 5)
squares := square(numbers)
doubles := double(squares)  // 另一个处理阶段

for result := range doubles {
    fmt.Println(result)
}
```

### 4.3 流水线的优势
- **并行处理**：每个阶段可以独立运行
- **解耦**：各阶段之间通过通道通信，互不干扰
- **可扩展**：容易添加新的处理阶段
- **资源控制**：通过通道缓冲控制并发度

## 5. Select 语句

### 5.1 什么是 Select？
- **Select**：用于在多个通道操作中选择一个可执行的操作
- **作用**：实现非阻塞的通道操作，或者同时监听多个通道

### 5.2 Select 基本语法
```go
select {
case <-ch1:
    // 处理 ch1 的数据
case ch2 <- value:
    // 向 ch2 发送数据
case <-time.After(1 * time.Second):
    // 超时处理
default:
    // 如果所有 case 都不可执行，执行 default
}
```

### 5.3 Select 示例

#### 5.3.1 多通道监听
```go
select {
case msg1 := <-ch1:
    fmt.Println("收到 ch1:", msg1)
case msg2 := <-ch2:
    fmt.Println("收到 ch2:", msg2)
}
```

#### 5.3.2 超时处理
```go
select {
case result := <-ch:
    fmt.Println("收到结果:", result)
case <-time.After(5 * time.Second):
    fmt.Println("超时！")
}
```

#### 5.3.3 非阻塞操作
```go
select {
case ch <- value:
    fmt.Println("发送成功")
default:
    fmt.Println("通道已满，无法发送")
}
```

## 6. 常见模式和最佳实践

### 6.1 等待 Goroutine 完成

#### 6.1.1 使用通道
```go
done := make(chan bool)
go func() {
    // 执行任务
    done <- true
}()
<-done  // 等待完成
```

#### 6.1.2 使用 sync.WaitGroup（更推荐）
```go
var wg sync.WaitGroup
for i := 0; i < 10; i++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done()
        // 执行任务
    }(i)
}
wg.Wait()  // 等待所有 goroutine 完成
```

### 6.2 扇出（Fan-out）和扇入（Fan-in）

#### 6.2.1 扇出：一个通道分发给多个 goroutine
```go
func fanOut(in <-chan int, out1, out2 chan<- int) {
    for value := range in {
        select {
        case out1 <- value:
        case out2 <- value:
        }
    }
}
```

#### 6.2.2 扇入：多个通道合并为一个
```go
func fanIn(input1, input2 <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for {
            select {
            case v := <-input1:
                out <- v
            case v := <-input2:
                out <- v
            }
        }
    }()
    return out
}
```

### 6.3 工作池模式（Worker Pool）
```go
func workerPool(jobs <-chan int, results chan<- int, numWorkers int) {
    var wg sync.WaitGroup
    
    // 启动多个 worker
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for job := range jobs {
                results <- process(job)
            }
        }()
    }
    
    wg.Wait()
    close(results)
}
```

## 7. 注意事项和常见错误

### 7.1 常见错误

#### 7.1.1 忘记等待 Goroutine
```go
// ❌ 错误：主程序可能在 goroutine 完成前退出
go doSomething()
// 程序立即退出

// ✅ 正确：使用通道或 WaitGroup 等待
done := make(chan bool)
go func() {
    doSomething()
    done <- true
}()
<-done
```

#### 7.1.2 通道死锁
```go
// ❌ 错误：无缓冲通道，发送和接收不在不同的 goroutine
ch := make(chan int)
ch <- 42      // 阻塞，因为没有接收者
value := <-ch // 永远执行不到

// ✅ 正确：在不同的 goroutine 中发送和接收
ch := make(chan int)
go func() {
    ch <- 42
}()
value := <-ch
```

#### 7.1.3 关闭已关闭的通道
```go
// ❌ 错误：重复关闭通道会导致 panic
close(ch)
close(ch)  // panic!

// ✅ 正确：只关闭一次，或使用 sync.Once
var once sync.Once
once.Do(func() {
    close(ch)
})
```

### 7.2 最佳实践

#### 7.2.1 明确通道方向
- 使用 `chan<-` 和 `<-chan` 明确通道是只发送还是只接收
- 提高代码可读性和安全性

#### 7.2.2 使用缓冲通道控制并发
```go
// 限制并发数为 10
semaphore := make(chan struct{}, 10)
for i := 0; i < 100; i++ {
    semaphore <- struct{}{}  // 获取信号量
    go func(id int) {
        defer func() { <-semaphore }()  // 释放信号量
        // 执行任务
    }(i)
}
```

#### 7.2.3 使用 context 控制取消
```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

go func() {
    select {
    case <-ctx.Done():
        return  // 收到取消信号
    case result := <-ch:
        // 处理结果
    }
}()
```

## 8. 总结

### 8.1 核心概念
- **Goroutine**：轻量级并发执行单元
- **Channel**：goroutine 之间通信的管道
- **流水线**：多个处理阶段通过通道串联
- **Select**：多通道操作选择

### 8.2 关键要点
1. **并发 vs 并行**：goroutine 提供并发，Go 运行时决定是否并行
2. **通道是类型化的**：只能传递指定类型的数据
3. **无缓冲通道是同步的**：发送和接收必须同时准备好
4. **缓冲通道是异步的**：可以在缓冲区未满时发送
5. **通过通信共享内存**：这是 Go 语言的并发哲学

### 8.3 选择指南
- **简单并发任务** → 使用 goroutine
- **需要通信** → 使用通道
- **复杂数据处理** → 使用流水线模式
- **需要等待多个任务** → 使用 WaitGroup
- **需要超时或取消** → 使用 context

### 8.4 实践建议
- 总是明确通道方向（`chan<-` 或 `<-chan`）
- 使用 `defer close(ch)` 确保通道被关闭
- 避免在持有锁时进行通道操作
- 使用 `select` 实现非阻塞操作
- 使用工作池模式控制并发数

