// 示例：Go 语言的组合与转发
// 演示如何用组合替代继承，通过结构体嵌套和方法转发实现代码复用

package main

import "fmt"

func main() {
	// ============================================
	// 1. 核心思想：组合优于继承
	// ============================================
	fmt.Println("=== 1. 核心思想：组合优于继承 ===")

	fmt.Println("Go语言的设计哲学：用组合替代继承")
	fmt.Println("通过结构体嵌套和方法转发，实现代码复用和行为扩展")
	fmt.Println("完全不需要传统的类继承体系")
	fmt.Println()

	fmt.Println("组合（Composition）：")
	fmt.Println("  一个结构体可以嵌入其他结构体")
	fmt.Println("  直接复用其字段和方法")
	fmt.Println("  就像把多个功能模块'拼'在一起")
	fmt.Println()

	fmt.Println("转发（Forwarding）：")
	fmt.Println("  当外部调用当前结构体的方法时")
	fmt.Println("  内部把请求转发给嵌入的结构体的对应方法")
	fmt.Println("  实现行为的复用和代理")
	fmt.Println()

	// ============================================
	// 2. 组合合并多个结构
	// ============================================
	fmt.Println("=== 2. 组合合并多个结构 ===")

	fmt.Println("示例：汽车 = 引擎 + 车轮")
	fmt.Println()

	// 基础结构：引擎
	fmt.Println("基础结构：引擎")
	engine := Engine{}
	fmt.Printf("  引擎类型: %T\n", engine)
	engine.Start()
	fmt.Println()

	// 基础结构：车轮
	fmt.Println("基础结构：车轮")
	wheels := Wheels{}
	fmt.Printf("  车轮类型: %T\n", wheels)
	wheels.Rotate()
	fmt.Println()

	// 组合多个结构：汽车
	fmt.Println("组合多个结构：汽车")
	car := Car{}
	fmt.Printf("  汽车类型: %T\n", car)
	fmt.Println("  汽车可以直接调用嵌入结构体的方法:")
	car.Start()   // 直接调用嵌入的Engine的Start方法
	car.Rotate()  // 直接调用嵌入的Wheels的Rotate方法
	fmt.Println()

	// 说明组合的优势
	fmt.Println("组合的优势：")
	fmt.Println("  ✅ Car 天然拥有了 Engine 和 Wheels 的能力")
	fmt.Println("  ✅ 比继承更灵活，可以自由选择要组合哪些模块")
	fmt.Println("  ✅ 不需要定义继承关系")
	fmt.Println()

	// ============================================
	// 3. 方法转发（自动转发 vs 手动转发）
	// ============================================
	fmt.Println("=== 3. 方法转发（自动转发 vs 手动转发）===")

	// 自动转发（嵌入结构体）
	fmt.Println("自动转发（嵌入结构体）:")
	autoCar := Car{}
	fmt.Println("  直接调用嵌入结构体的方法:")
	autoCar.Start()
	autoCar.Rotate()
	fmt.Println()

	// 手动转发（添加额外逻辑）
	fmt.Println("手动转发（添加额外逻辑）:")
	manualCar := CarWithForwarding{}
	fmt.Println("  调用手动转发的方法:")
	manualCar.Start()   // 会先打印准备信息，再转发
	manualCar.Rotate()  // 直接转发
	fmt.Println()

	// 对比说明
	fmt.Println("对比说明：")
	fmt.Println("  自动转发：直接嵌入，方法自动可用")
	fmt.Println("  手动转发：可以添加额外逻辑，完全控制调用流程")
	fmt.Println()

	// ============================================
	// 4. 更复杂的组合示例
	// ============================================
	fmt.Println("=== 4. 更复杂的组合示例 ===")

	// 示例1：电脑系统
	fmt.Println("示例1：电脑系统")
	computer := Computer{}
	fmt.Println("  电脑组件:")
	computer.CPU.Process()
	computer.Memory.Store()
	computer.Storage.Save()
	fmt.Println("  电脑可以直接使用组件的方法:")
	computer.Process()  // 转发给CPU
	computer.Store()    // 转发给Memory
	computer.Save()     // 转发给Storage
	fmt.Println()

	// 示例2：机器人
	fmt.Println("示例2：机器人")
	robot := Robot{
		Body:   Body{},
		Brain:  Brain{},
		Arms:   Arms{},
		Legs:   Legs{},
	}
	fmt.Println("  机器人组件:")
	robot.Body.Move()
	robot.Brain.Think()
	robot.Arms.Grab()
	robot.Legs.Walk()
	fmt.Println("  机器人可以直接使用组件的方法:")
	robot.Move()  // 转发给Body
	robot.Think() // 转发给Brain
	robot.Grab()  // 转发给Arms
	robot.Walk()  // 转发给Legs
	fmt.Println()

	// ============================================
	// 5. 方法转发的高级用法
	// ============================================
	fmt.Println("=== 5. 方法转发的高级用法 ===")

	// 示例1：添加日志
	fmt.Println("示例1：添加日志")
	loggedCar := CarWithLogging{}
	loggedCar.Start()
	loggedCar.Rotate()
	fmt.Println()

	// 示例2：条件转发
	fmt.Println("示例2：条件转发")
	smartCar := SmartCar{}
	smartCar.Start()  // 会检查状态
	smartCar.Start()  // 第二次调用会提示已启动
	fmt.Println()

	// 示例3：多级转发
	fmt.Println("示例3：多级转发")
	hybridCar := HybridCar{}
	hybridCar.Start()  // 会先启动引擎，再启动电机
	fmt.Println()

	// ============================================
	// 6. 组合 vs 继承对比
	// ============================================
	fmt.Println("=== 6. 组合 vs 继承对比 ===")

	fmt.Println("传统继承方式（Go不支持）:")
	fmt.Println("  class Car extends Engine, Wheels { ... }")
	fmt.Println("  问题：")
	fmt.Println("    - 只能继承一个父类（单继承限制）")
	fmt.Println("    - 继承关系固定，难以修改")
	fmt.Println("    - 容易产生'脆弱基类'问题")
	fmt.Println()

	fmt.Println("Go的组合方式:")
	fmt.Println("  type Car struct { Engine; Wheels }")
	fmt.Println("  优势：")
	fmt.Println("    - 可以组合多个结构体（多组合）")
	fmt.Println("    - 组合关系灵活，可以动态调整")
	fmt.Println("    - 每个模块独立，避免脆弱基类问题")
	fmt.Println()

	// ============================================
	// 7. 实际应用场景
	// ============================================
	fmt.Println("=== 7. 实际应用场景 ===")

	// 场景1：HTTP服务器（模拟标准库）
	fmt.Println("场景1：HTTP服务器（模拟标准库）")
	server := HTTPServer{
		Router: Router{},
		Logger: Logger{},
	}
	server.HandleRequest("/api/users")
	server.Log("Request processed")
	fmt.Println()

	// 场景2：数据库连接池
	fmt.Println("场景2：数据库连接池")
	pool := ConnectionPool{
		Pool:   Pool{},
		Config: Config{},
	}
	pool.Initialize()
	pool.GetConnection()
	fmt.Println()

	// 场景3：中间件链
	fmt.Println("场景3：中间件链")
	middleware := MiddlewareChain{
		Auth:    AuthMiddleware{},
		Logging: LoggingMiddleware{},
		Cache:   CacheMiddleware{},
	}
	middleware.Process()
	fmt.Println()

	// ============================================
	// 8. 彻底抛弃类继承
	// ============================================
	fmt.Println("=== 8. 彻底抛弃类继承 ===")

	fmt.Println("在Go里，你不需要再纠结'父类、子类、继承链'这些概念")
	fmt.Println()

	fmt.Println("复用代码：")
	fmt.Println("  ✅ 用结构体嵌套（组合）替代继承")
	fmt.Println()

	fmt.Println("扩展行为：")
	fmt.Println("  ✅ 用方法转发和接口实现替代重写")
	fmt.Println()

	fmt.Println("解耦设计：")
	fmt.Println("  ✅ 组合让每个模块更独立")
	fmt.Println("  ✅ 依赖更清晰")
	fmt.Println("  ✅ 避免了继承带来的'脆弱基类'问题")
	fmt.Println()

	// ============================================
	// 9. 组合的最佳实践
	// ============================================
	fmt.Println("=== 9. 组合的最佳实践 ===")

	fmt.Println("1. 优先使用组合:")
	fmt.Println("   ✅ 需要复用代码时，优先考虑组合")
	fmt.Println("   ✅ 通过嵌入结构体实现代码复用")
	fmt.Println()

	fmt.Println("2. 合理使用转发:")
	fmt.Println("   ✅ 需要添加额外逻辑时，使用手动转发")
	fmt.Println("   ✅ 需要完全控制时，使用手动转发")
	fmt.Println("   ✅ 简单场景可以使用自动转发")
	fmt.Println()

	fmt.Println("3. 保持模块独立:")
	fmt.Println("   ✅ 每个结构体应该职责单一")
	fmt.Println("   ✅ 组合的结构体之间应该低耦合")
	fmt.Println("   ✅ 避免循环依赖")
	fmt.Println()

	fmt.Println("4. 接口配合使用:")
	fmt.Println("   ✅ 组合 + 接口 = 强大的设计")
	fmt.Println("   ✅ 接口定义行为契约")
	fmt.Println("   ✅ 组合实现具体功能")
	fmt.Println()

	// ============================================
	// 10. 总结
	// ============================================
	fmt.Println("=== 10. 总结 ===")
	fmt.Println()
	fmt.Println("1. 核心思想:")
	fmt.Println("   ✅ 组合优于继承")
	fmt.Println("   ✅ 用组合替代继承")
	fmt.Println("   ✅ 通过结构体嵌套和方法转发实现代码复用")
	fmt.Println()
	fmt.Println("2. 组合合并多个结构:")
	fmt.Println("   ✅ 嵌入结构体，直接复用其字段和方法")
	fmt.Println("   ✅ 比继承更灵活，可以自由选择组合模块")
	fmt.Println()
	fmt.Println("3. 方法转发:")
	fmt.Println("   ✅ 自动转发：直接嵌入，方法自动可用")
	fmt.Println("   ✅ 手动转发：可以添加额外逻辑，完全控制")
	fmt.Println()
	fmt.Println("4. 彻底抛弃类继承:")
	fmt.Println("   ✅ 复用代码：用组合替代继承")
	fmt.Println("   ✅ 扩展行为：用方法转发和接口实现")
	fmt.Println("   ✅ 解耦设计：组合让模块更独立")
	fmt.Println()
	fmt.Println("5. 最佳实践:")
	fmt.Println("   ✅ 优先使用组合")
	fmt.Println("   ✅ 合理使用转发")
	fmt.Println("   ✅ 保持模块独立")
	fmt.Println("   ✅ 接口配合使用")
	fmt.Println()
}

// ============================================
// 基础结构体定义
// ============================================

// Engine 引擎结构体
type Engine struct{}

func (e *Engine) Start() {
	fmt.Println("    Engine started")
}

// Wheels 车轮结构体
type Wheels struct{}

func (w *Wheels) Rotate() {
	fmt.Println("    Wheels rotating")
}

// Car 汽车结构体（组合Engine和Wheels）
type Car struct {
	Engine // 嵌入Engine，直接复用其方法
	Wheels // 嵌入Wheels，直接复用其方法
}

// CarWithForwarding 带手动转发的汽车
type CarWithForwarding struct {
	engine Engine // 注意：小写，不嵌入
	wheels Wheels // 注意：小写，不嵌入
}

// Start 手动转发Start方法，并添加额外逻辑
func (c *CarWithForwarding) Start() {
	fmt.Println("    Preparing to start car...")
	c.engine.Start() // 转发给engine的Start方法
	fmt.Println("    Car started successfully")
}

// Rotate 手动转发Rotate方法
func (c *CarWithForwarding) Rotate() {
	c.wheels.Rotate()
}

// CPU CPU结构体
type CPU struct{}

func (c *CPU) Process() {
	fmt.Println("    CPU processing")
}

// Memory 内存结构体
type Memory struct{}

func (m *Memory) Store() {
	fmt.Println("    Memory storing")
}

// Storage 存储结构体
type Storage struct{}

func (s *Storage) Save() {
	fmt.Println("    Storage saving")
}

// Computer 电脑结构体（组合多个组件）
type Computer struct {
	CPU
	Memory
	Storage
}

// Process 转发给CPU
func (c *Computer) Process() {
	c.CPU.Process()
}

// Store 转发给Memory
func (c *Computer) Store() {
	c.Memory.Store()
}

// Save 转发给Storage
func (c *Computer) Save() {
	c.Storage.Save()
}

// Body 身体结构体
type Body struct{}

func (b *Body) Move() {
	fmt.Println("    Body moving")
}

// Brain 大脑结构体
type Brain struct{}

func (b *Brain) Think() {
	fmt.Println("    Brain thinking")
}

// Arms 手臂结构体
type Arms struct{}

func (a *Arms) Grab() {
	fmt.Println("    Arms grabbing")
}

// Legs 腿部结构体
type Legs struct{}

func (l *Legs) Walk() {
	fmt.Println("    Legs walking")
}

// Robot 机器人结构体（组合多个组件）
type Robot struct {
	Body
	Brain
	Arms
	Legs
}

// Move 转发给Body
func (r *Robot) Move() {
	r.Body.Move()
}

// Think 转发给Brain
func (r *Robot) Think() {
	r.Brain.Think()
}

// Grab 转发给Arms
func (r *Robot) Grab() {
	r.Arms.Grab()
}

// Walk 转发给Legs
func (r *Robot) Walk() {
	r.Legs.Walk()
}

// CarWithLogging 带日志的汽车
type CarWithLogging struct {
	Engine
	Wheels
}

// Start 添加日志的转发
func (c *CarWithLogging) Start() {
	fmt.Println("    [LOG] Starting car...")
	c.Engine.Start()
	fmt.Println("    [LOG] Car started")
}

// Rotate 添加日志的转发
func (c *CarWithLogging) Rotate() {
	fmt.Println("    [LOG] Rotating wheels...")
	c.Wheels.Rotate()
	fmt.Println("    [LOG] Wheels rotated")
}

// SmartCar 智能汽车（条件转发）
type SmartCar struct {
	Engine
	started bool
}

// Start 条件转发
func (c *SmartCar) Start() {
	if c.started {
		fmt.Println("    Car is already started")
		return
	}
	fmt.Println("    Starting smart car...")
	c.Engine.Start()
	c.started = true
	fmt.Println("    Smart car started")
}

// HybridCar 混合动力汽车（多级转发）
type HybridCar struct {
	Engine
	ElectricMotor
}

// Start 多级转发
func (c *HybridCar) Start() {
	fmt.Println("    Starting hybrid car...")
	c.Engine.Start()
	c.ElectricMotor.Start()
	fmt.Println("    Hybrid car started")
}

// ElectricMotor 电机结构体
type ElectricMotor struct{}

func (e *ElectricMotor) Start() {
	fmt.Println("    Electric motor started")
}

// Router 路由器结构体
type Router struct{}

func (r *Router) HandleRequest(path string) {
	fmt.Printf("    Routing request to: %s\n", path)
}

// Logger 日志结构体
type Logger struct{}

func (l *Logger) Log(message string) {
	fmt.Printf("    [LOG] %s\n", message)
}

// HTTPServer HTTP服务器（组合Router和Logger）
type HTTPServer struct {
	Router
	Logger
}

// Pool 连接池结构体
type Pool struct{}

func (p *Pool) Initialize() {
	fmt.Println("    Pool initialized")
}

func (p *Pool) GetConnection() {
	fmt.Println("    Connection retrieved from pool")
}

// Config 配置结构体
type Config struct{}

func (c *Config) Load() {
	fmt.Println("    Config loaded")
}

// ConnectionPool 连接池（组合Pool和Config）
type ConnectionPool struct {
	Pool
	Config
}

// AuthMiddleware 认证中间件
type AuthMiddleware struct{}

func (a *AuthMiddleware) Authenticate() {
	fmt.Println("    Authenticating...")
}

// LoggingMiddleware 日志中间件
type LoggingMiddleware struct{}

func (l *LoggingMiddleware) Log() {
	fmt.Println("    Logging...")
}

// CacheMiddleware 缓存中间件
type CacheMiddleware struct{}

func (c *CacheMiddleware) Cache() {
	fmt.Println("    Caching...")
}

// MiddlewareChain 中间件链（组合多个中间件）
type MiddlewareChain struct {
	Auth    AuthMiddleware
	Logging LoggingMiddleware
	Cache   CacheMiddleware
}

// Process 处理请求
func (m *MiddlewareChain) Process() {
	m.Auth.Authenticate()
	m.Logging.Log()
	m.Cache.Cache()
	fmt.Println("    Request processed through middleware chain")
}

