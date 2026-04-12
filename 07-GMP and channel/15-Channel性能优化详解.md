# Channel 性能优化详解（接收侧为主）

> **07-GMP and channel** · 课表 **`15-Channel性能优化详解.md`**  
> 前置：[14-学GMP必学Channel总览](./14-学GMP必学Channel总览.md)、[08-sudog核心数据结构与作用详解](./08-sudog核心数据结构与作用详解.md)。  
> 不同 Go 版本 **runtime** 细节会变，以下记 **原理 + 可验证实践**；结论以 **`go test -bench` / 读 `$GOROOT/src/runtime/chan.go`** 为准。

---

## 一、编译器 + runtime（接收侧）

### 1. 简单接收走专用入口，绕过胖 `selectgo`

编译器对 **单路 `select`、简单 `<-ch`、`v, ok := <-ch`** 常直接生成对 **`runtime.chanrecv1` / `runtime.chanrecv2`** 的调用，而不是走通用 **`selectgo`** 多路复用逻辑，从而 **少分支、少一层调度包装**（仍以本机反汇编/`go tool compile -S` 验证习惯为准）。

```go
// 源码
v := <-ch

// 编译后（概念伪代码）
runtime.chanrecv1(ch, &v)
```

`chanrecv1` / `chanrecv2` 内部再进入 **`chanrecv`** 主逻辑。

### 2. 逃逸分析与 GC（别把「channel 在栈上」当默认结论）

- **`make(chan T)`** 创建的 **`hchan`** 在多数实现里 **长期落在堆上**（需稳定地址、与 runtime 队列指针互指）。  
- **`-gcflags=-m`** 更有价值的是：观察 **送进 channel 的值、闭包、与 ch 一起逃逸的结构体**，而不是死记「ch 一定栈分配」。

```bash
go build -gcflags=-m ./...
# 或测试包
go test -gcflags=-m -run=^$ ./...
```

### 3. 无缓冲：发送方已在等时的「直传」路径（接收侧最大甜头之一）

当 **无缓冲 channel** 且 **发送方已阻塞在 `sendq`** 时，`chanrecv` 可走 **直接配对**：把元素从 **发送方 sudog 记录的栈槽** 拷到 **接收方栈**，**不必经过环形缓冲区**（与「先写入环、再读出环」相比少一次中间驻留）。

概念步骤（仍受竞态与队列状态影响，这里是教学抽象）：

1. 对 **`hchan` 加锁**（自旋 + mutex 等实现细节见版本）。  
2. **`sendq` 非空** → 取队头 **sudog（发送 G）**。  
3. **elem 拷贝**：发送侧 → 接收侧（常见为 **栈到栈** 类路径）。  
4. **唤醒发送 G**，解锁，接收返回。

### 4. 唤醒与公平性

- **`recvq` / `sendq`** 多为 **FIFO** 公平队列。  
- 是否「合并多次 **`goready`**」属于实现细节；工程上记：**少阻塞、少无谓唤醒** 总是方向。

---

## 二、接收侧性能直觉（原理 → 选型）

### 1. 无缓冲：低并发、强同步时往往最省

- **条件**：收发 **容易同时就绪**，配对紧密。  
- **收益**：可走 **直传**，少缓冲环的读写。  
- **代价**：高并发下 **阻塞/唤醒频繁**，尾延迟抖动可能大。

### 2. 小缓冲：高吞吐读的常见甜点区

- **作用**：收发 **节奏解耦**；缓冲非空时接收可 **连续读** 而少等发送。  
- **经验起点**（非教条）：**64 / 128 / 256** 与 **消息体大小、生产者 burst** 强相关，**必须 bench**。

### 3. 大缓冲对接收侧的常见伤害

| 问题 | 直觉 |
|------|------|
| 内存 | **buf × 元素大小** 线性涨 |
| 延迟 | 数据在环里 **排队更久** |
| 观测 | **掩盖消费者跟不上**（队列一直很长） |
| 缓存 | 大环 **跨 cache line** 扫描更贵 |

---

## 三、使用层：接收侧实战（模板）

### 1. 分片 + 每片单 worker（降 `hchan` 锁竞争）

**问题**：大量 goroutine **抢读同一个 `ch`** → **`hchan.lock` 热点**。  
**思路**：**N 个 channel + 哈希/轮询分流**，每片 **单 reader**。

```go
const shards = 8

var chs [shards]chan int

func init() {
	for i := range chs {
		chs[i] = make(chan int, 128)
	}
}

func send(v int) {
	i := v % shards
	if i < 0 {
		i += shards // Go 负数取模可能为负，归一化到 [0,shards)
	}
	chs[i] <- v
}

func startWorkers(process func(int)) {
	for i := range chs {
		go func(ch <-chan int) {
			for v := range ch {
				process(v)
			}
		}(chs[i])
	}
}
```

### 2. 批量接收（摊销「一次唤醒处理的消息数」）

适用于 **高频小消息**、下游 **批量刷盘/刷 DB**。

```go
func batchReader(ch <-chan int, batchSize int, flushEvery time.Duration, processBatch func([]int)) {
	batch := make([]int, 0, batchSize)
	timer := time.NewTimer(flushEvery)
	defer timer.Stop()

	for {
		select {
		case v, ok := <-ch:
			if !ok {
				if len(batch) > 0 {
					processBatch(batch)
				}
				return
			}
			batch = append(batch, v)
			if len(batch) >= batchSize {
				processBatch(batch)
				batch = batch[:0]
				if !timer.Stop() {
					<-timer.C
				}
				timer.Reset(flushEvery)
			}
		case <-timer.C:
			if len(batch) > 0 {
				processBatch(batch)
				batch = batch[:0]
			}
			timer.Reset(flushEvery)
		}
	}
}
```

> **注意**：`timer.Reset` 与 `select` 组合要防 **旧 tick 误触发**；生产可换 **`time.AfterFunc`** 或 **`context`** 驱动 flush。

### 3. 非阻塞读：`select + default`

```go
func tryRecv(ch <-chan int) (int, bool) {
	select {
	case v := <-ch:
		return v, true
	default:
		return 0, false
	}
}
```

适用于 **主循环不能卡死**、**背压/降级**；注意 **丢消息** 语义是否符合业务。

### 4. 大消息：传指针，减少元素拷贝

```go
// 大结构体：值会复制整个 T
// ch <- BigStruct{...}

// 常见优化：传指针（注意并发下 **只读或所有权转移**）
ch <- &BigStruct{...}
```

---

## 四、Benchmark 对照（接收吞吐）

```go
func BenchmarkUnbufferedRecv(b *testing.B) {
	ch := make(chan int)
	go func() {
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		close(ch)
	}()
	b.ResetTimer()
	for range ch {
	}
}

func BenchmarkBuffered128Recv(b *testing.B) {
	ch := make(chan int, 128)
	go func() {
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		close(ch)
	}()
	b.ResetTimer()
	for range ch {
	}
}
```

常见形态（**非保证**）：**小缓冲** 在「**一端狂发、一端狂收**」时优于纯无缓冲；**超大缓冲** 往往 **不占优**（内存与缓存行为变差）。以你机器 **`go test -benchmem -count=5`** 为准。

逃逸辅助：

```bash
go test -bench . -gcflags=-m
```

---

## 五、自检答案（课表 15）

### 1. 「优化读」和「优化发」对称吗？

- **底层**：**`chansend` / `chanrecv`** 镜像逻辑（`sendq` ↔ `recvq`、满/空对称），锁与拷贝结构相似。  
- **不对称点（接收侧）**：**`for range ch`**、**`v, ok := <-ch` 的关闭语义**、以及业务上 **多读单写/多写单读** 模式不同，热点常不一样。

### 2. 何时加缓冲反而更慢（接收视角）？

1. **收发天然强同步、低并行**：无缓冲 **直传** 更短路径；小缓冲可能多 **环读写** 一次。  
2. **缓冲过大**：内存、缓存、延迟、观测性全面变差。  
3. **消费者极快、生产者很慢**：缓冲长期空，**纯占内存与分支**。  
4. **单 goroutine 自发自收**：缓冲收益常为 **零**。

---

## 六、接收侧口诀

- **低并发、强同步**：无缓冲往往更干净。  
- **高并发、高吞吐**：**小缓冲 + 分片 + 单 reader**。  
- **高频小消息**：**批量 flush**。  
- **大对象**：**指针 + 明确所有权**。  
- **大缓冲**：多半在 **掩盖设计问题**，用 **指标与 bench** 说话。

---

## 延伸阅读

- [11-Channel核心数据结构hchan详解](./11-Channel核心数据结构hchan详解.md)  
- [12-Channel环形队列实现原理](./12-Channel环形队列实现原理.md)  
- [10-Channel阻塞协程的原理与现象](./10-Channel阻塞协程的原理与现象.md)  
- [09-GMP与sudog四者联动关系](./09-GMP与sudog四者联动关系.md)
