# IO 多路复用：`poll`

> **02-io-models · IO 模型全解**  
> 上篇：[04-IO多路复用-select](./04-IO多路复用-select.md)

---

## 一、核心定义

`poll` 用 **`struct pollfd` 数组** 描述「fd + 关心的事件」，一次 `poll` 阻塞等待，返回后由用户扫描 **`revents`** 得知谁就绪。语义上与 `select` 同属 **level-triggered** 风格（默认行为与具体事件类型有关，对比 epoll 时再细抠）。

相对 `select`，常见卖点是：**不依赖 `fd_set` 位图与 `FD_SETSIZE` 那种上限**（数组长度由你分配），接口对「稀疏大 fd」更自然一些。

### 典型流程

1. 填充 `pollfd` 数组（`fd`、`events`）。  
2. `poll(fds, nfds, timeout)` 阻塞。  
3. 内核设置各元素的 `revents`。  
4. 用户遍历数组，对 `POLLIN`/`POLLOUT` 等就绪者做 IO。  
5. 循环。

---

## 二、特点

### 优点

- **无 `select` 那种典型的 fd_set 上限痛点**（仍受系统资源、进程 ulimit 等约束）。  
- **每次可只传数组**，不必维护 `maxfd+1` 与位图重置的整套习惯（实现细节因库而异）。  
- 仍属成熟 POSIX 接口。

### 缺点

- **fd 很多时，每次 `poll` 仍可能伴随「大量 pollfd 在内核侧扫描」**（复杂度直觉上常写成 O(n)；与内核实现相关）。  
- **返回后用户侧 O(n) 扫描** 仍在。  
- 超大规模下 Linux 上多被 **`epoll`** 替代。

---

## 三、API 备忘（C）

```c
int poll(struct pollfd *fds, nfds_t nfds, int timeout);

struct pollfd {
    int   fd;
    short events;   /* 请求 */
    short revents;  /* 结果 */
};
```

---

## 四、与 `select` 的对比（简表）

| 项 | `select` | `poll` |
|----|----------|--------|
| 数据结构 | `fd_set` 位图 | `pollfd` 数组 |
| 典型 fd 数上限痛点 | 有（`FD_SETSIZE` 等） | 弱化为数组/资源限制 |
| 返回后 | 遍历集合找就绪 | 遍历数组看 `revents` |
| 跨平台 | 极常见 | 常见 |

---

## 五、典型场景

- 中等规模、希望接口比 `select` 清爽的路径。  
- 理解 **「为何 epoll 在 Linux 上更省」** 的过渡。

---

## 六、极简总结

- `poll`：**数组 + `revents`**，解决 `select` 若干工程痛点，但**大规模扫描成本**仍在。  
- Linux 超高并发：**`epoll`**。

---

## 导航

- 上一篇：[04-IO多路复用-select](./04-IO多路复用-select.md)  
- 下一篇：[06-IO多路复用-epoll原理](./06-IO多路复用-epoll原理.md)
