# epoll / kqueue / IOCP 支持

> **07-go-netpoll · Go netpoll 高并发核心**

## 内容大纲

- Linux：epoll；BSD 与 macOS：kqueue；Windows：IOCP（概念对齐）
- 构建标签与文件拆分：如何在本机源码树定位
- 差异点：LT 与 ET、Windows AFD 细节可略读
- 为何用户态很少直接调 epoll：Go 已内嵌
- 扩展阅读：runtime 网络轮询与定时器、work stealing 的协作

## 正文

（待补充）
