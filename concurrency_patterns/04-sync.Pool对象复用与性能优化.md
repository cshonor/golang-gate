# 04 - sync.Pool：对象复用与性能优化

## 1. sync.Pool 解决什么

减少“频繁创建/销毁临时对象”的堆分配，降低 GC 压力。

常见对象：

- `bytes.Buffer`
- `[]byte` 临时缓冲
- 编解码/序列化中间对象

## 2. 最小用法

```go
var bufPool = sync.Pool{
    New: func() any { return new(bytes.Buffer) },
}

func Use() {
    b := bufPool.Get().(*bytes.Buffer)
    b.Reset()
    defer bufPool.Put(b)
    // 使用 b
}
```

## 3. 面试必背：Pool 的语义

- Pool 里的对象 **可能会被 GC 清掉**
- Pool 不是“强缓存”，更像“尽力复用”

所以 Pool 适合：

- 临时对象
- 可丢弃、可重建

不适合：

- 必须长期存在的资源（连接、文件句柄）

## 4. 最佳实践

- `Get` 后立刻 `Reset/清理`，防止脏数据泄漏
- `Put` 前把大字段释放/截断，避免把超大 buffer 长期留在池里

