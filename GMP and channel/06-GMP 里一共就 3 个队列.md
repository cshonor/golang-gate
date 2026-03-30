我给你一次性把 GMP 里所有队列、全局队列、放哪、长啥样全部讲清楚，非常直白，不绕。
 
一、GMP 里一共就 3 个队列
 
全是放 Goroutine（G） 的：
 
1. P 本地队列 runq（环形队列）
2. 全局 G 队列 sched.runq（链表队列）
3. 等待队列（各种阻塞队列）
 
 
 
二、你问的重点：
 
全局 Goroutine 队列在哪里？
 
它在：
 
runtime 包的全局变量 sched 里面
 
go  
var sched struct {
    // 全局可运行 G 队列
    runq     gList
    runqsize int32

    // 一堆调度相关的锁、状态
}
 
 
- sched.runq：就是全局 G 队列
- 它是一个链表，不是环形队列
- 所有 P 共享
- 访问必须加锁
 
 
 
三、GMP 里哪些地方用到了 队列 / 环形队列
 
1）P 本地 runq —— 环形队列
 
go  
runq [256]guintptr
runqhead uint32
runqtail uint32
 
 
- 数组实现
- 环形
- 无锁
- 最快
 
2）全局 G 队列 sched.runq —— 链表队列
 
- 链表
- 有锁
- 所有 P 共用
 
3）channel 缓冲 buf —— 环形队列
 
go  
buf      unsafe.Pointer
sendx    uint
recvx    uint
 
 
- 放数据，不放 G
- 也是环形队列
 
 
 
四、什么时候用哪个队列？（超级关键）
 
1. 新建 G
→ 优先放 P 本地 runq
→ 满了 → 把本地一半 G 挪到 全局队列
2. M 取 G 执行
1. 先从 本地 P runq 拿
2. 没有 → 去偷别的 P 的 runq
3. 还没有 → 去全局队列拿一批
 
 
 
五、你现在最该记住的三句话
 
1. P 本地 runq：环形队列，无锁，最快
2. 全局 G 队列：在 sched 全局变量里，链表，有锁
3. M 必须绑 P 才能执行 G，同一时刻 1P ↔ 1M
 
 
 
你现在是不是全局队列、本地队列、环形队列全都通了？
我可以再给你画一张极简 GMP 队列流程图，面试直接背。