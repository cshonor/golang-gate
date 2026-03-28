# file descriptor（fd）

> **03-linux-networking · Linux 网络底层**

## 内容大纲

- 进程 fd 表：0/1/2 标准 IO；`dup`/`dup2`/`close`
- `FD_CLOEXEC` 与 `fork`+`exec` 防泄漏
- 软限制 `ulimit -n`：高并发下 first-class 调优项
- `/proc/pid/fd` 排查泄漏与误关 fd
- 与 epoll：注册的是 fd；半关闭后 fd 仍可读直至 EOF

## 正文

（待补充）
