# gotolist

基于 `DDD(Domain Driven Design)` 领域驱动设计 架构实现的 todolist。

## 功能

接口文档位于 `docs/`，可直接将 `docs` 下的 JSON 文件导入 Postman 进行联调。

- 用户模块
  - 注册
  - 登陆
- 备忘录模块
  - 创建
  - 更新
  - 列表
  - 删除
  - 详情
- 公共能力
  - 结构化日志
  - JWT 鉴权
  - CORS 跨域
  - Docker 一键启动环境
  - 领域事件（EventBus）
  - 事件发件箱（Outbox）
  - 依赖注入容器（Container）

## 架构

![DDD架构图](./docs/structure.png)

```
./gotolist
├── application         // 应用层: 编排领域层完成用例
│   ├── task            // task 应用层
│   └── user            // user 应用层
├── cmd                 // 程序启动入口
├── conf                // 配置文件
├── consts              // 常量与枚举
├── docs                // 接口文档（Postman JSON 等）
├── domain              // 领域层: 纯业务逻辑
│   ├── task
│   │   ├── entity      // task 实体 + 充血模型
│   │   ├── event       // task 领域事件定义
│   │   ├── repository  // task 持久化接口
│   │   └── service     // task 业务逻辑
│   └── user
│       ├── entity      // user 实体 + 充血模型
│       ├── event       // user 领域事件定义
│       ├── repository  // user 持久化接口
│       └── service     // user 业务逻辑
├── infrastructure      // 基础架构层: 数据来源与基础服务
│   ├── auth            // JWT 鉴权服务
│   ├── common
│   │   ├── context     // 上下文（用户信息等）传递
│   │   └── log         // 日志服务
│   ├── container       // 依赖注入容器与生命周期监听
│   ├── encrypt         // 加密服务（密码等）
│   ├── eventbus        // 领域事件总线
│   └── persistence     // 持久化
│       ├── dbs         // 数据库连接与迁移
│       ├── outbox      // 事件发件箱
│       ├── task        // task DAO
│       └── user        // user DAO
├── interfaces          // 接口层: 对接不同端
│   ├── adapter
│   │   └── initialize  // 路由初始化
│   ├── controller      // HTTP Controller
│   ├── midddleware     // 中间件（CORS / JWT）
│   └── types           // 请求/响应 DTO 与转换器
└── logs                // 日志文件
```

整体遵循 DDD 的四层职责：

- 接口层（Interfaces）: 接收不同端的请求，转成领域输入；返回结果时再转回 DTO。
- 应用层（Application）: 编排领域对象完成用例，不写业务规则。
- 领域层（Domain）: 核心业务逻辑。`entity` 是充血模型，承载行为；`repository` 仅定义接口；`service` 聚合实体完成业务。
- 基础架构层（Infrastructure）: 实现领域层定义的接口，提供数据库、加密、事件总线、容器等基础能力。

## 启动

1. 启动基础环境（MySQL 等）

```shell
docker-compose -f docker-compose.yml up
```

2. 启动服务

```shell
go run cmd/main.go
```

3. 调试接口

将 `docs/` 下的 Postman JSON 导入 Postman，按模块调试用户与备忘录相关接口。