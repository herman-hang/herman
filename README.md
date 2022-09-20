# fp-gin-framework(搭建中)

基于Gin框架搭建的Web项目基础框架。

## 1. 项目结构

```
├─app --------------------------------------------------------- 应用程序目录
│  ├─command -------------------------------------------------- 命令管理目录
│  ├─common --------------------------------------------------- 公共模块目录
│  ├─constants ------------------------------------------------ 常量存放目录
│  ├─controllers ---------------------------------------------- 控制器目录
│  ├─jobs ----------------------------------------------------- 队列作业目录
│  ├─middlewares ---------------------------------------------- 中间件目录
│  ├─models --------------------------------------------------- 数据模型目录
│  ├─services ------------------------------------------------- 服务处理目录
│  ├─utils ---------------------------------------------------- 工具类目录
│  └─validates ------------------------------------------------ 验证器目录
├─config ------------------------------------------------------ 配置文件目录
│  └─config.yaml ---------------------------------------------- 配置文件
├─database ---------------------------------------------------- 数据库相关目录
│  ├─migrations ----------------------------------------------- 数据迁移目录
│  ├─seeders -------------------------------------------------- 数据填充目录
│  └─sqls ----------------------------------------------------- 数据库更新SQL文件目录
├─logs -------------------------------------------------------- 日志处理目录
├─resources --------------------------------------------------- 资源目录
│  ├─css ------------------------------------------------------ CSS文件目录
│  ├─images --------------------------------------------------- 图片文件目录
│  ├─js ------------------------------------------------------- JS文件目录
│  └─views ---------------------------------------------------- 视图文件目录
├─routers ----------------------------------------------------- 路由文件目录
├─server ------------------------------------------------------ GO服务目录
├─settings ---------------------------------------------------- 全局设置配置目录
├─storage ----------------------------------------------------- 驱动存储目录
├─tests ------------------------------------------------------- 测试目录
├─.air.toml --------------------------------------------------- Air热重载配置文件
├─.gitignore -------------------------------------------------- gitignore文件
├─go.mod ------------------------------------------------------ go.mod文件
├─go.sum ------------------------------------------------------ go.sum文件
├─main.go ----------------------------------------------------- 程序入口文件
└─README.md --------------------------------------------------- Readme文件
```

## 2. 项目热重载启动

```bash
# 优先在当前路径查找 `.air.toml` 后缀的文件，如果没有找到，则使用默认的
air -c .air.toml
```
您可以运行以下命令初始化，把默认配置添加到当前路径下的`.air.toml` 文件。

```bash
air init
```

在这之后，你只需执行 `air` 命令，无需添加额外的变量，它就能使用 `.air.toml` 文件中的配置了。

```bash
air
```
