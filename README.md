# 基于Go-Zero框架的商城系统后端
该系统基于：https://github.com/nivin-studio/go-zero-mall

系统运行环境：https://github.com/nivin-studio/gonivinck

请使用make命令运行各服务

下面是一个启动用户rpc服务的例子：
```bash
make userrpc
```
## 与原项目不同之处：
1. 利用Redis完成Token 自动续期方案。
2. 修复原始项目中的一些代码错误符合最新的版本(2024.3)。
3. 在main分支有gorm的应用，但是由于需要用redis，这里将gorm替换成go-zero自带的数据库操作方法。

## 项目技术栈

1. **Golang:** 主要编程语言，用于开发应用程序和服务.

2. **DTM (Distributed Transaction Manager):** 用于管理分布式事务的组件，确保在分布式系统中的事务一致性.

3. **Etcd:** 用作分布式键值存储系统，用于配置管理和服务发现.

4. **Redis:** 用作内存数据库，用于缓存和提高数据访问速度.

5. **MySQL:** 关系型数据库，用于持久化存储应用程序的数据.

6. **Prometheus:** 用于监控和警报的开源系统，收集和存储应用程序和系统的指标.

7. **Grafana:** 可视化工具，与Prometheus集成，用于创建和共享动态监控仪表板.

8. **Jaeger:** 分布式追踪系统，用于监控和分析应用程序性能和请求流程.

9. **docker-compose:** 用于编排和管理容器化应用程序的工具，通过定义和运行多个Docker容器来简化部署过程.
