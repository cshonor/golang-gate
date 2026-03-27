# 02 - Background()、TODO()、上下文树

## 1. Background 与 TODO 的区别

- `context.Background()`：根 context，通常用于 main/初始化/顶层
- `context.TODO()`：临时占位，表示“这里未来要传入正确的 ctx”，生产代码别长期留

面试一句话：

> Background 是真正的根；TODO 是“还没想好/还没接入”的占位符。

## 2. 上下文树（最关键的心智模型）

Context 是一棵树：

- 子 ctx 由父 ctx 派生（WithCancel/WithTimeout/WithValue）
- **父取消会传递到子**（一层层往下）
- 子取消 **不影响父**（一般如此）

```text
request ctx
  ├─ withTimeout ctx
  │    └─ withValue(traceID) ctx
  └─ withValue(userID) ctx
```

## 3. 你在中间件里做了什么

中间件本质就是：

- 从父 ctx 派生一个子 ctx（加 timeout / 加 value）
- 再把这个子 ctx 往下传

