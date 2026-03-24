我给你整理一套极简、能直接写进 MD、面试必背的：
 
G、M、P、sudog 四大数据结构关系图（文字版）
 
 
 
一、总关系（一句话背会）
 
- G（g）：协程本体，存运行现场、栈、状态
- M（m）：OS 线程，真正执行代码
- P（p）：逻辑处理器，存运行队列，控制并行度
- sudog：G 的包装节点，用于阻塞等待队列
 
G 要跑 → 必须绑 M → M 必须绑 P
G 阻塞 → 变成 sudog → 进等待队列
 
 
 
二、各自数据结构核心（只记关键）
 
1. G（runtime.g）—— 协程本体
 
plaintext  
g {
    stack       // 协程栈
    sched       // 调度现场（pc、sp 等，切换用）
    status      // 状态：running / waiting / syscall etc.
    m           // 当前绑定的 M
    waitreason  // 等待原因（chan、mutex、sleep）
}
 
 
2. M（runtime.m）—— 系统线程
 
plaintext  
m {
    g0          // 调度用的栈 G
    curg        // 当前正在运行的 G
    p           // 当前绑定的 P
    nextp       // 即将绑定的 P
}
 
 
3. P（runtime.p）—— 逻辑处理器
 
plaintext  
p {
    runq[]      // 本地 G 运行队列
    runqhead
    runqtail
    m           // 绑定的 M
    gFree       // G 缓存池
}
 
 
4. sudog —— 阻塞等待队列节点
 
plaintext  
sudog {
    g       // 被阻塞的 G
    c       // 对应的 channel
    elem    // 数据地址（收发用）
    next    // 链表下一个
    prev    // 链表上一个
}
 
 
 
 
三、它们之间的关系（最核心）
 
1. M 想运行 G，必须先绑定 P
2. 同一时间，一个 M 只绑一个 P，一个 P 只绑一个 M
3. 一个 P 管理一堆 G（本地队列）
4. G 不阻塞时：直接在 M 上跑
5. G 阻塞时（chan、mutex 等）：
被包装成 sudog → 进入等待队列 → 离开 M
6. 唤醒时：sudog 出队 → G 变回可运行 → 放回 P 队列
 
 
 
四、最精简总结（可直接复制到 MD）
 
- G：协程实体，保存栈、现场、状态。
- M：内核线程，执行载体。
- P：调度上下文，持有 G 队列，控制并行度。
- sudog：阻塞 G 的包装节点，用于等待队列。
- 调度规则：M 绑定 P 才能执行 G；G 阻塞则包装为 sudog 进入等待队列，释放 M。
 
 
 
你接下来想让我整理：
GMP 完整调度流程 还是 G 状态流转图（面试必考）？