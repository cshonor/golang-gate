### 01 - defer 基本用法
#### 1. defer 做什么
`defer` 是 Go 语言关键字，用于**注册延迟执行的函数调用**，保证该调用**一定会在当前函数退出前执行**（无论函数是正常 `return`、发生 `panic` 还是提前返回）。

核心作用：**资源安全管理 + 收尾逻辑兜底**。
常见用途：
- 文件/流关闭：`file.Close()`
- 互斥锁释放：`mu.Unlock()`
- 数据库事务回滚：`tx.Rollback()`
- 连接释放、日志收尾、异常捕获

#### 2. 核心特性（必记）
- **执行顺序**：多个 `defer` 按**后进先出（LIFO）** 执行（栈结构）。
- **参数求值时机**：`defer` 语句执行时，**参数立即计算并固定**，函数体延迟执行。
- **返回值修改**：仅能修改**具名返回值**（匿名返回值不可修改）。
- **异常安全**：函数 `panic` 时，`defer` 仍会执行，可配合 `recover` 捕获异常。

#### 3. 标准模板（资源管理最佳实践）
**拿到资源后，第一时间 defer 释放**，是最稳的资源管理习惯。

```go
import "os"

func readFile(path string) ([]byte, error) {
    // 1. 打开资源
    f, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    // 2. 立即 defer 关闭（关键！）
    defer f.Close()

    // 3. 业务逻辑
    data := make([]byte, 1024)
    n, err := f.Read(data)
    if err != nil {
        return nil, err
    }
    return data[:n], nil
}
```

#### 4. 常见场景示例
##### 场景1：互斥锁安全释放
```go
import "sync"

var mu sync.Mutex
var count int

func increment() {
    mu.Lock()         // 加锁
    defer mu.Unlock() // 立即 defer 解锁，避免死锁
    count++
}
```

##### 场景2：数据库事务安全回滚
```go
import "database/sql"

func transfer(tx *sql.Tx, from, to string, amount float64) error {
    // 开启事务后立即 defer 回滚（异常时兜底）
    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
        }
    }()

    // 执行转账逻辑
    _, err := tx.Exec("UPDATE accounts SET balance = balance - ? WHERE id = ?", amount, from)
    if err != nil {
        return err
    }
    _, err = tx.Exec("UPDATE accounts SET balance = balance + ? WHERE id = ?", amount, to)
    if err != nil {
        return err
    }

    // 无异常则提交事务
    return tx.Commit()
}
```

##### 场景3：参数求值 vs 闭包延迟引用
```go
func demo() {
    i := 0
    defer fmt.Println(i) // 参数立即求值，打印 0
    i++

    defer func() {
        fmt.Println(i) // 闭包延迟引用，打印 1
    }()
}
// 输出：1 0（后进先出）
```

#### 5. 避坑要点
- **循环中慎用 defer**：循环内 `defer` 会累积到函数结束才执行，可能导致资源泄漏/性能问题。
- **仅修改具名返回值**：匿名返回值无法被 `defer` 闭包修改。
- **defer 后必须是函数调用**：不能是函数定义（如 `defer func(){} ` 错误，需 `defer func(){}()`）。

