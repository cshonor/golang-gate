# IO 多路复用：select

> **02-io-models · IO 模型全解**

## 内容大纲

- API：`fd_set`、FD_SETSIZE 上限、用户态/内核态来回拷贝 bitmap
- 时间复杂度：O(n) 扫描就绪集合；n 大时线性退化
- 精度：`timeval` 与超时行为；被信号中断（EINTR）处理
- 可移植性：几乎所有 POSIX 实现；教学与兼容层仍有价值
- 对比 poll/epoll：为何 select 仍是「理解复用」的第一课

## 正文

（待补充）
