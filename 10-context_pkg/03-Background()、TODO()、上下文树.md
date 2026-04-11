# 03 - Background()、TODO()、上下文树（精简背诵版）

## 1. Background 与 TODO 的区别

- **`context.Background()`**  
  真正的**根 Context**，永不取消、永不超时。  
  用在：main、init、顶层逻辑、服务启动入口。

- **`context.TODO()`**  
  只是**占位 Context**，底层与 `Background` 同属空实现，但**语义不同**：  
  「这里本该传 ctx，但暂时没传 / 还没设计好」。  
  生产代码**不应该长期保留** `TODO`。

**面试一句话：**

> Background 是正式根节点；TODO 是临时占位，提醒后续要补 ctx。

---

## 2. 上下文树（核心心智模型）

Context 是**树形结构**：

- 子 Context 从父 Context 派生（`WithCancel` / `WithTimeout` / `WithDeadline` / `WithValue`）
- **父取消 → 所有子自动级联取消**
- **子取消 → 不影响父**

结构示意：

```text
Background
  └─ HTTP Request ctx
       ├─ WithTimeout(500ms)
       │    └─ WithValue(traceID)
       └─ WithValue(userID)
```

一句话总结：

> 上下文树 = 生命周期继承 + 取消向下传播。

---

## 3. 中间件在做什么

中间件本质就三步：

1. 从**父 ctx** 派生出**子 ctx**
2. 给子 ctx 附加：超时、取消、traceID、userID 等
3. 把新 ctx 传给下一层 handler / 业务函数

也就是：

**包装 ctx → 传递 ctx → 控制整个请求生命周期**

---

## 4. 一句话串起来

- Background 是根，TODO 是占位；
- Context 是树，父管子、子不管父；
- 中间件就是不断派生子 ctx 往下传。

---

## 延伸阅读（同目录）

| 主题 | 文档 |
|------|------|
| 总览与接口 | [02-context是什么、作用、核心接口.md](./02-context是什么、作用、核心接口.md) |
| WithCancel | [04-WithCancel 手动取消.md](./04-WithCancel%20手动取消.md) |
| WithTimeout / Deadline | [05-WithTimeout 超时控制.md](./05-WithTimeout%20超时控制.md) |
| WithValue | [06-WithValue 数据透传.md](./06-WithValue%20数据透传.md) |
| 中间件实战 | [07-context在中间件中的实战.md](./07-context在中间件中的实战.md) |

背诵版一页汇总四个创建函数时，可按：**Cancel / Timeout / Deadline / Value** 对照上面四篇展开。
