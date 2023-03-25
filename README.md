# Herman框架

## 1. 序言
### 介绍 

Herman基于Gin，Casbin，Kafka，Mysql，Redis，Zap，Cobra，Grom开发，专注于后端快速上手的一款开源，简洁，轻量框架。

### 项目结构

```
├─app --------------------------------------------------------- 应用程序目录
│  ├─constants ------------------------------------------------ 常量存放目录
│  ├─controllers ---------------------------------------------- 控制器目录
│  ├─models --------------------------------------------------- 数据模型目录
│  ├─repositories --------------------------------------------- 仓储层目录
│  ├─services ------------------------------------------------- 服务处理目录
│  ├─utils ---------------------------------------------------- 工具类目录
│  ├─validates ------------------------------------------------ 验证器目录
│  ├─request.go ----------------------------------------------- 请求对象库
│  └─response.go ---------------------------------------------- 响应对象库
├─cmd --------------------------------------------------------- 命令管理目录
├─jobs -------------------------------------------------------- 队列作业目录
├─kernel ------------------------------------------------------ 框架核心目录
├─middlewares ------------------------------------------------- 中间件目录
├─config ------------------------------------------------------ 配置文件目录
├─database ---------------------------------------------------- 数据库相关目录
│  ├─migrations ----------------------------------------------- 数据迁移目录
│  └─seeders -------------------------------------------------- 数据填充目录
├─runtime ----------------------------------------------------- 运行目录
│  └─logs ----------------------------------------------------- 日志记录目录
├─resources --------------------------------------------------- 资源目录
│  ├─css ------------------------------------------------------ CSS文件目录
│  ├─defaultImages -------------------------------------------- 验证码资源目录
│  ├─fonts ---------------------------------------------------- 验证码字体目录
│  ├─images --------------------------------------------------- 图片文件目录
│  ├─js ------------------------------------------------------- JS文件目录
│  └─views ---------------------------------------------------- 视图文件目录
├─routers ----------------------------------------------------- 路由文件目录
├─server ------------------------------------------------------ GO服务目录
│  └─settings ------------------------------------------------- 核心配置目录
├─storages ---------------------------------------------------- 文件存储目录
├─tests ------------------------------------------------------- 测试目录
├─.air.toml --------------------------------------------------- Air热重载配置文件
├─.gitignore -------------------------------------------------- gitignore文件
├─go.mod ------------------------------------------------------ go.mod文件
├─go.sum ------------------------------------------------------ go.sum文件
├─config.yaml.debug ------------------------------------------- 开发环境配置文件
├─config.yaml.test -------------------------------------------- 测试环境配置文件
├─config.yaml.release ----------------------------------------- 正式环境配置文件
├─Dockerfile -------------------------------------------------- Dodcker镜像配置
├─docker-compose.yaml ----------------------------------------- Dodcker容器编排文件
├─LICENSE ----------------------------------------------------- 许可证文件
├─Makefile ---------------------------------------------------- Makefile文件
├─main.go ----------------------------------------------------- 入口文件
└─README.md --------------------------------------------------- Readme文件
```

### 开发规范

#### （1）目录与文件命名

- 目录名称采用小驼峰命名（首字母小写）
- .go文件采用下划线命名，例如：`user`，`user_login`
- 数据库迁移文件采用下划线命名，例如：`1_init.down.sql`，`1_init.up.sql`，1为版本号，init为自定义名称，down代表回滚，up代表迁移。
- 资源文件(图片，CSS文件，JS文件等)均采用蛇形命名，例如CSS文件：`test.css`，`test_user.css`，以此类推。

#### （2）函数、方法、结构体

- 函数和方法命名可以大驼峰（首字母大写）和小驼峰（首字母小写）命名，具体看业务需求，如果只需在本包调用则小驼峰即可，否则需要大驼峰

- 结构体名称、字段名、json标签一律使用大驼峰命名，示例：

  ```
  type Users struct {
     Id           uint       `json:"id" gorm:"primary_key"`
     User         string     `json:"user"`
     Password     string     `json:"password"`
     Nickname     string     `json:"nickname"`
     Sex          string     `json:"sex"`
     Age          int        `json:"age"`
     Region       string     `json:"region"`
     Phone        string     `json:"phone"`
     Email        string     `json:"email"`
     Introduction string     `json:"introduction"`
     Status       string     `json:"status"`
     CreatedAt    time.Time  `json:"createdAt"`
     UpdatedAt    time.Time  `json:"updatedAt"`
     DeletedAt    *time.Time `json:"deletedAt" sql:"index"`
  }
  ```

#### （3）变量与常量

- 全局变量和函数方法规范相似，如果需要跨包调用，则采用大驼峰（首字母大写），否则采用小驼峰（首字母小写）
- 常量使用大驼峰命名（首字母大写），例如：`Success`，`TokenNotExit`

#### （4）数据表与字段

- 数据表名没有前缀，表名不能出现大写字母，建议以蛇形定义，例如：`user`，`user_role`
- 字段名称采用蛇形命名，不能出现大写字母，例如：`user_id`，`user_name`

### 安装

#### （1）修改环境文件

为了项目开发管理灵活性，根目录分别有3个环境文件，分别为`config.yaml.debug`开发环境文件，`config.yaml.test`测试环境文件，`config.yaml.release`正式环境文件，如果当前使用的环境为开发环境，则修改`config.yaml.debug`为`config.yaml`，例如：

```shell
cp config.yaml.debug config.yaml
```

以此类推。

#### （2）配置MySQL和Redis

项目启动依赖于Mysql和Redis，所以在启动之前，必须配置好MySQL和Redis的服务连接参数，否则程序无法启动。

```yaml
# 数据库配置
mysql:
  # 连接IP地址
  host: 127.0.0.1
  # 连接端口号
  port: 3306
  # 连接用户名
  user: root
  # 连接密码
  password: root
  # 连接数据库名称
  dbname: herman
  # 最大连接数
  max_open_conn: 100
  # 最大连接空闲数，建议和max_open_conn一致
  max_idle_conn: 10

# Redis配置
redis:
  # 连接IP地址
  host: 127.0.0.1
  # 连接端口号
  port: 6380
  # 连接用户名
  username:
  # 连接密码
  password:
  # 默认数据库，默认是0
  db: 0
  # 最大连接数
  pool_size: 100
```

#### （3）安装依赖

执行该命令之前，Go环境必须已经安装。

```go
go mod download
```

#### （4）程序启动

（1）编译后启动（推荐正式环境使用）

```shell
go build -o herman . # 项目编译成二进制文件herman
herman server --host=0.0.0.0 --port=8000 --migrate=true # host和port是可选的，但是migtate首次运行程序是必须的，会自动帮你迁移数据表到数据库
```

（2）非编译启动

```shell
go run main.go server --host=0.0.0.0 --port=8000 --migrate=true # 首次非编译启动程序，host和port也是可选的
```

（3）热重载启动（推荐开发环境使用）

该启动方式运行之前，**必须要完成数据库迁移工作**，否则会发生致命错误。使用该功能要求Go 的版本为 1.16及以上。

```bash
go install github.com/cosmtrek/air@latest # 如果已经安装则无需操作此步
```
您可以执行以下命令初始化，把默认配置添加到当前路径下的`.air.toml` 文件。

```bash
air init
```

配置示例：

```toml
# [Air](https://github.com/cosmtrek/air) TOML 格式的配置文件

# 工作目录
# 使用 . 或绝对路径，请注意 `tmp_dir` 目录必须在 `root` 目录下
root = "."
tmp_dir = "tmp"

[build]
# 只需要写你平常编译使用的shell命令。你也可以使用 `make`
cmd = "go build -o ./tmp/herman.exe ."
# 由`cmd`命令得到的二进制文件名
bin = "tmp\\herman.exe server"
# 监听以下文件扩展名的文件.
include_ext = ["go", "tpl", "tmpl", "html"]
# 忽略这些文件扩展名或目录
exclude_dir = ["assets", "tmp", "vendor", "frontend/node_modules"]
# 监听以下指定目录的文件
include_dir = []
# 排除以下文件
exclude_file = []
# 如果文件更改过于频繁，则没有必要在每次更改时都触发构建。可以设置触发构建的延迟时间
delay = 1000 # ms
# 发生构建错误时，停止运行旧的二进制文件。
stop_on_error = true
# air的日志文件名，该日志文件放置在你的`tmp_dir`中
log = "air_errors.log"

[log]
# 显示日志时间
time = true

[color]
# 自定义每个部分显示的颜色。如果找不到颜色，使用原始的应用程序日志。
main = "magenta"
watcher = "cyan"
build = "yellow"
runner = "green"

[misc]
# 退出时删除tmp目录
clean_on_exit = true
```

热重载启动：

```bash
air
```

## 2. 架构

### 生命周期

使用一门技术，了解它的生命周期是必不可少，只有你去了解它，使用起来才会更加自信。以下是Herman处理一个HTTP请求的流程：

- 应用入口：Golang大部分项目都是`main.go`文件入口，Herman也不例外，在没有编译成二进制的前提下，入口从`main.go`文件开始。
- 服务注册：进入cobra的init函数进行配置，日志初始化；框架版本，服务，数据库迁移，JWT令牌注册。
- 设置运行模式：采用Gin框架设置当前的运行模式。
- Gin框架启动：注册核心中间件，初始化路由，监听HTTP请求。
- 控制器：接收请求上下文，处理请求参数，验证器和服务调用，以及响应返回。
- 验证器：接收控制器处理好的参数进行验证，验证通过的参数返回给验证器。
- 服务层：接收验证器验证通过的参数，调用仓储层获取数据库数据，然后把数据返回给控制器。
- 仓储层：这一层又被称为数据库模型与数据交互的桥梁，主要操作数据库模型，再次封装增删改查，实现代码高度解耦。
- 模型层：与数据库数据表一对一绑定，数据表字段与模型结构体绑定。
- 请求响应：所有逻辑处理完成，数据由控制器响应返回。

### 容器

Golang虽然是一门面向过程的语言，但是Herman也引入了容器的概念，对项目核心的对象，比如Redis，MySQL，Casbin等都存放在`/kernel/core/Container.go`文件中。

```go
package core

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Engine *gin.Engine
	Log    *zap.SugaredLogger
	Db     *gorm.DB
	Redis  *redis.Client
	Casbin *casbin.CachedEnforcer
)
```

### 中间件

中间件分为**前置中间件**和**后置中间件**的，主要存放在`/middlewares`，比如以下定义的中间件：

```go
// ServerHandler 服务管理中间件
// @return gin.HandlerFunc
func ServerHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		Reload() // 请求前执行
		ctx.Next()
		Close() // 请求后执行
	}
}
```

中间件定义调用均采用Gin框架提供的API，所以调用中间可以在路由，也可以在别处，具体看业务要求，以下在程序启动后调用：

```go
func NewServer(host string, port uint) {
	// 设置gin框架运行模式
	gin.SetMode(settings.Config.Mode)
	// 启动gin框架
	engine := gin.New()
	// 注册中间件
	engine.Use(log.GinLogger()).Use(middlewares.CatchError()).Use(middlewares.ServerHandler())
	// 初始化路由
	core.Engine = routers.InitRouter(engine)
	// 启动服务
	Run(host, port)
}
```

你也可以在路由中调用，比如鉴权中间件：

```go
// Jwt 鉴权
// @return gin.HandlerFunc 返回一个中间件上下文
func Jwt(guard string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if VerifyRoute(ctx.Request.URL.Path, ctx.Request.Method, MiddlewareConstant.ExcludeRoute) {
			return
		}
		claims := utils.JwtVerify(ctx, guard)
		switch guard {
		case "user", "mobile": // 前台和移动端（用户）
			// 用户信息存储在请求中
			ctx.Set("user", repositories.User().GetUserInfo(claims.Uid))
		case "admin": // 管理员后台
			ctx.Set("admin", repositories.Admin().GetAdminInfo(claims.Uid))
		case "merchant": // 商家后台

		default:
			panic(MiddlewareConstant.GuardError)
		}
		ctx.Next()
	}
}
```

路由定义中调用：

```go
// 后台模块
adminRouter := api.Group("/admin", middlewares.Jwt("admin"), middlewares.CheckPermission())
{
    admin.Router(adminRouter)
}
```

### 命令行

命令行核心采用cobra实现，主要存放在`cmd`目录，命令注册在`/kernel/cobra/cobra.go`文件，比如以下例子：

```go
// HermanVersionCmd 获取herman版本号
var (
	HermanVersionCmd = &cobra.Command{
		Use:          "version",
		Short:        "Get herman version",
		Example:      "herman version",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf(`Herman version: %v`, color.GreenString(settings.Version))
			return nil
		},
	}
)
```

编写好之后，进行注册：

```go
// rootCmd 定义命令行
var rootCmd = &cobra.Command{Use: "herman"}

// 注册命令行
func init() {
	// 执行命令前初始化操作
	cobra.OnInitialize(settings.InitConfig, servers.ZapLogs, func() {
		if command.IsMigrate {
			// 数据库迁移
			_ = command.Migrate("up")
		}
		// 如果执行的是数据库迁移命令，则不需要加载初始化操作
		if !command.MigrationStatus {
			middlewares.Reload()
		}
	})

	// 注册框架版本命令
	rootCmd.AddCommand(command.HermanVersionCmd)
}
```

官方已经内置了几个命令：

（1）查看框架版本号

```shell
herman version # Herman version: 1.0.0
```

（2）数据库迁移

```shell
herman migrate --direction=up --number=1 # 表示迁移1个版本
```

这里每个参数需要绑定

```go
// init 命令参数绑定
// @return void
func init() {
	// 迁移状态
	MigrationCmd.Flags().BoolVarP(&MigrationStatus, "status", "s", false, "Database migration status")
	// 迁移方式，up和down
	MigrationCmd.Flags().StringVarP(&direction, "direction", "d", "up", "Database migration")
	// 执行指定数据库版本，主要在出现Error: Dirty database version XX.使用
	MigrationCmd.Flags().UintVarP(&version, "version", "v", 0, "Database version")
	// 执行迁移的版本次数，比如回滚1个版本，可以执行herman -d down -n 1，不指定则全部迁移
	MigrationCmd.Flags().UintVarP(&number, "number", "n", 0, "Database migration steps")
}
```

命令绑定之后，可以随意结合，都是可选的，根据业务需求执行。

（3）随机生成JWT令牌

```shell
herman jwt:secret
```

（4）框架服务启动

```shell
herman server --host=0.0.0.0 --port=8000 --migrate=true # 启动服务并做数据库迁移
```

如果框架已经迁移过数据库，也可以这样启动服务：

```shell
herman server # 默认端口为8000
```

cobra扩展文档：https://cobra.dev/

### 队列

队列采用kafka，主要存放在`jobs`目录，比如以下短信发送例子：

```go
// SendSms 发送短信队列
// @param string topic 消息主题
// @return void
func SendSms(topic string) {
	var data map[string]interface{}
	// 调用消费者对数据进行消费，并返回结构体
	kafkaConsumer := ExecConsumer(topic)
	for {
		// 从通道取出消费的数据
		message := <-kafkaConsumer.MessageQueue
		// 将取出的JSON数据转为map
		if err := json.Unmarshal(message, &data); err != nil {
			core.Log.Errorf("Consumer sms json data failed, err:%v", err)
		}
		execSend(data)
	}
}
```

服务层调用：

```go
jobs.Dispatch(data,jobs.SendSms)
```

其中，以下代码为固定写法，目的是取出消费的数据：

```go
	for {
		// 从通道取出消费的数据
		message := <-kafkaConsumer.MessageQueue
		// 将取出的JSON数据转为map
		if err := json.Unmarshal(message, &data); err != nil {
			core.Log.Errorf("Consumer sms json data failed, err:%v", err)
		}
		execSend(data)
	}
```

### 缓存

目前框架只支持Redis缓存，对象挂载在`/kernel/core/container.go`中，使用前要先设置上下文：

```go
// 设置上下文
ctx := context.Background()
```

设置一个key值

```go
val, err := core.Redis.Get(ctx, "key").Result()
fmt.Println(val)
```

取出一个key值

```go
get := core.Redis.Get(ctx, "key")
fmt.Println(get.Val(), get.Err())
```

设置一个key值并设置过期时间

```go
core.Redis.Set(ctx, "key", 1, time.Minute*30)
```

更多Redis扩展：https://redis.uptrace.dev/zh/guide/

### 日志

日志集成Zap，这是一个强大的日志库，它在Herman中起到很关键的作用。

```go
// 记录一个日志
core.Log.info(data)
// 记录一个日志并换行
core.Log.infoln(data)
// 调式
core.Log.Debug(data)
// 记录一个错误
core.Log.Error(data)
// 记录一个错误并换行
core.Log.Errorln(data)
// 记录一个错误并终止进程
core.Log.Fatal(data)
```

更多API文档：https://pkg.go.dev/go.uber.org/zap

### 辅助函数

辅助函数又称工具类，主要存放在`/app/utils`中，如果项目中存在一些碎片化的代码，想把它做进一步封装，就可以在该目录下创建一个工具文件，在里面完成相应的封装。比如我下方做了一个验证码工厂：

```go
// Factory 初始化滑块验证码
// @return factory 返回一个验证码工厂
func Factory() (factory *CaptchaService.CaptchaServiceFactory) { // 行为校验配置模块（具体参数可从业务系统配置文件自定义）
	// 行为校验初始化
	factory = CaptchaService.NewCaptchaServiceFactory(
		CaptchaConfig.BuildConfig(settings.Config.Captcha.CacheType,
			settings.Config.Captcha.ResourcePath,
			&CaptchaConfig.WatermarkConfig{
				Text: settings.Config.Captcha.Text,
			},
			nil, nil, settings.Config.Captcha.CacheExpireSec))
	// 注册内存缓存
	factory.RegisterCache(Constant.MemCacheKey, CaptchaService.NewMemCacheService(CaptchaConstant.CacheMaxNumber))
	// 注册自定义配置redis数据库
	factory.RegisterCache(Constant.RedisCacheKey, CaptchaService.NewConfigRedisCacheService([]string{fmt.Sprintf("%s:%d",
		settings.Config.Redis.Host,
		settings.Config.Redis.Port,
	)},
		settings.Config.Redis.UserName,
		settings.Config.Redis.Password,
		false,
		settings.Config.Redis.Db,
	))
	// 注册文字点选验证码服务
	factory.RegisterService(Constant.ClickWordCaptcha, CaptchaService.NewClickWordCaptchaService(factory))
	// 注册滑动拼图验证码服务
	factory.RegisterService(Constant.BlockPuzzleCaptcha, CaptchaService.NewBlockPuzzleCaptchaService(factory))

	return factory
}

```

封装好之后，在框架那个地方都可以调用，非常方便。

### 权限模型

Casbin是一种轻量级的开源访问控制框架，支持多种访问控制模型，如RBAC, ABAC和ACL。框架中已经采用了RBAC，适配GORM来做角色资源管理，可以灵活管理角色的权限。核心封装代码在`/kernel/casbin/casbin.go`。框架Casbin的对象挂载在容器`/kernel/core/container.go`，调用：

```go
success, _ := core.Casbin.Enforce(info.User, ctx.Request.URL.Path, ctx.Request.Method)
```

更多学习：https://casbin.org/zh/docs/category/the-basics

### 配置

框架的所有配置都是通过读取根目录下的`config.yaml`文件所得，并且存放在`config`目录中，调用方式：

```go
settings.Config
```

比如获取MySQL的配置

```go
settings.Config.Mysql
```

当然，如果你不想创建配置文件作映射，也可以直接获取环境文件`config.yaml`的配置，但是不建议这么操作。

```go
viper.Get("app")
```

## 3. 路由

路由沿用了Gin集成的功能，所有路由定义都在`/routers/router.go`，例子：

```go
func InitRouter(rootEngine *gin.Engine) *gin.Engine {
	// 测试路由
	rootEngine.GET("/", func(context *gin.Context) {
		response := app.Request{Context: context}
		response.Success(app.D(map[string]interface{}{
			"message": "Welcome to Herman!",
		}))
	})
	// 设置路由前缀
	api := rootEngine.Group(settings.Config.AppPrefix)
	// 获取验证码
	api.GET("/captcha", CaptchaController.GetCaptcha)
	// 检查验证码正确性
	api.POST("/captcha/check", CaptchaController.CheckCaptcha)

	// 用户模块
	userRouter := api.Group("/user", middlewares.Jwt("user"))
	{
		mobile.Router(userRouter)
	}

	// 后台模块
	adminRouter := api.Group("/admin", middlewares.Jwt("admin"), middlewares.CheckPermission())
	{
		admin.Router(adminRouter)
	}

	return rootEngine
}
```

## 4. 控制器

控制器层面的责任非常明确，只负责**接收上下文**，**获取参数**，**调用**以及**响应**，不做其他任何操作。调用这里包括调用验证器验证参数，调用服务层处理逻辑，然后响应返回。

```go
// AddAdmin 管理员添加
// @param *gin.Context ctx 上下文
// @return void
func AddAdmin(ctx *gin.Context) {
	context := app.Request{Context: ctx} // 上下文二次封装
	data := context.Params() // 获取参数
	AdminService.Add(AdminValidate.Add.Check(data)) // 调用验证器验证参数，然后调用服务层处理逻辑
	context.Json(nil) // 响应返回
}
```

## 5. 验证器

验证器定义：

```go
// Add 重写验证器结构体，切记不使用引用，而是拷贝
var Add = validates.Validates{Validate: AddValidate{}}

// AddValidate 管理员添加验证规则
type AddValidate struct {
	User         string       `json:"user" validate:"required,min=5,max=15" label:"用户名"`
	Password     string       `json:"password" validate:"required,min=6,max=15" label:"密码"`
	Roles        []role.Roles `json:"roles" validate:"required" label:"选择角色"`
	Photo        string       `json:"photo" validate:"omitempty,url,max=255" label:"头像"`
	Name         string       `json:"name" validate:"omitempty,max=20" label:"真实姓名"`
	Card         string       `json:"card" validate:"omitempty,max=20" label:"身份证号码"`
	Sex          uint8        `json:"sex" validate:"required,oneof=1 2 3" label:"性别"`
	Age          uint8        `json:"age" validate:"required,min=0,max=120" label:"年龄"`
	Region       string       `json:"region" validate:"omitempty,max=255" label:"住址"`
	Phone        string       `json:"phone" validate:"omitempty,len=11" label:"手机号码"`
	Email        string       `json:"email" validate:"omitempty,email" label:"邮箱"`
	Introduction string       `json:"introduction" validate:"omitempty" label:"简介"`
	State        uint8        `json:"state" validate:"required,oneof=1 2" label:"状态"`
	Sort         uint         `json:"sort" validate:"omitempty" label:"排序"`
}
```

如果有对验证器公共结构体进行重写，那么就可以使用结构体的公共方法check，上面控制器的例子就是使用了验证器的公共方法。

```go
// Check 验证方法
// @param map[string]interface{} data 待验证数据
// @return void
func (base Validates) Check(data map[string]interface{}) (toMap map[string]interface{}) {
	// map赋值给结构体
	if err := mapstructure.WeakDecode(data, &base.Validate); err != nil {
		panic(constants.MapToStruct)
	}
	if err := Validate(base.Validate); err != nil {
		panic(err.Error())
	}

	toMap, err := utils.ToMap(base.Validate, "json")

	if err != nil {
		panic(constants.StructToMap)
	}
	return toMap
}
```

如果没有额外的业务扩展，这样是非常便利了，只关注验证规则如何去定义就可以。当然，如果你需要做一些验证扩展也是可以的，比如管理员登录：

- 控制器

```go
// Login 管理员登录
// @param *gin.Context ctx 上下文
// @return void
func Login(ctx *gin.Context) {
	context := app.Request{Context: ctx}
	data := context.Params()
	context.Json(AdminService.Login(AdminValidate.Login(data)), AdminConstant.LoginSuccess)
}
```

- 验证器

```go

// CaptchaLoginValidate 管理员登录验证结构体
type CaptchaLoginValidate struct {
	User        string `json:"user" validate:"required,min=5,max=15" label:"用户名"`
	Password    string `json:"password" validate:"required,min=6,max=15" label:"密码"`
	CaptchaType int    `json:"captchaType" validate:"required,numeric,oneof=1 2" label:"验证码类型"`
	Token       string `json:"token" validate:"required" label:"验证码Token"`
	PointJson   string `json:"pointJson" validate:"required" label:"验证码PointJson"`
}

// ExcludeCaptchaLoginValidate 管理员登录排除验证码相关验证结构体
type ExcludeCaptchaLoginValidate struct {
	User     string `json:"user" validate:"required,min=5,max=15" label:"用户名"`
	Password string `json:"password" validate:"required,min=6,max=15" label:"密码"`
}

// Login 登录验证器
// @param map[string]interface{} data 待验证数据
// @return toMap 返回验证通过的数据
func Login(data map[string]interface{}) (toMap map[string]interface{}) {
	// 判断是否需要验证码
	if !settings.Config.Captcha.Switch {
		return excludeCaptchaLogin(data)
	}
	return captchaLogin(data)
}

// captchaLogin 验证码登录验证器
// @param map[string]interface{} data 待验证数据
// @return toMap 返回验证通过的数据
func captchaLogin(data map[string]interface{}) (toMap map[string]interface{}) {
	var login CaptchaLoginValidate
	// map赋值给结构体
	if err := mapstructure.WeakDecode(data, &login); err != nil {
		panic(constants.MapToStruct)
	}

	if err := validates.Validate(login); err != nil {
		panic(err.Error())
	}

	// 验证码二次验证
	err := utils.Factory().GetService(fmt.Sprintf("%s", data["captchaType"])).Verification(fmt.Sprintf("%s", data["token"]),
		fmt.Sprintf("%s", data["PointJson"]))
	if err != nil {
		panic(CaptchaConstant.CheckCaptchaError)
	}

	toMap, err = utils.ToMap(&login, "json")
	if err != nil {
		panic(constants.StructToMap)
	}

	return toMap
}

// excludeCaptchaLogin 排除验证码登录验证器
// @param map[string]interface{} data 待验证数据
// @return toMap 返回验证通过的数据
func excludeCaptchaLogin(data map[string]interface{}) (toMap map[string]interface{}) {
	var login ExcludeCaptchaLoginValidate
	// map赋值给结构体
	if err := mapstructure.WeakDecode(data, &login); err != nil {
		panic(constants.MapToStruct)
	}

	if err := validates.Validate(login); err != nil {
		panic(err.Error())
	}

	toMap, err := utils.ToMap(&login, "json")
	if err != nil {
		panic(constants.StructToMap)
	}

	return toMap
}

```

业务需要扩展验证器，可以直接在验证器文件中自定义规则即可，比如上面的例子就是把管理员登录是否需要验证码做了2种场景验证。

## 6. 服务

服务层主要责任是逻辑处理，服务层没有什么约束，可以调用仓储层，工具类等等，但是这里值得注意的是，如果需要开启数据库事务的，必须要在这一层开启，然后在事务中进行多维度调用。例子如下：

```go
err := core.Db.Transaction(func(tx *gorm.DB) error {
   // casbin重新初始化
   _, _ = casbin.InitEnforcer(casbin.GetAdminPolicy(), tx)
   // 判断角色Key是否存在
   if isExist, _ := repositories.Role(tx).KeyIsExist(data["role"].(string)); isExist {
      return errors.New(RoleConstant.KeyExist)
   }
   roles := data["roles"]
   rules := data["rules"]
   delete(data, "roles")
   delete(data, "rules")
   // 添加角色信息
   roleInfo, err := repositories.Role(tx).Insert(data)
   if err != nil {
      return errors.New(RoleConstant.AddFail)
   }
   // 添加策略
   if err := AddPolicies(roles.([]role.Roles), rules.([]role.Rules), roleInfo); err != nil {
      return err
   }
   return nil
})
```

如果有使用到Casbin，那么Casbin的Db也需要更新，比如上面代码的：

```go
_, _ = casbin.InitEnforcer(casbin.GetAdminPolicy(), tx)
```

如果遇到错误，必须要把错误返回，否则事务不会进行回滚，比如上面代码：

```go
roleInfo, err := repositories.Role(tx).Insert(data)
if err != nil {
    return errors.New(RoleConstant.AddFail)
}
```

有错误返回，数据库会进行回滚。

## 7. 仓储

仓储层是位于Service层和Model层之间，是对Model层进一步封装。仓储层公共方法已有**新增**，**更新**，**删除**，**根据查询条件获取详情**，**查询数据是否存在**，**获取列表数据**，**获取全部数据**。代码如下：

```go
package repositories

import (
	"github.com/herman-hang/herman/app/constants"
	"github.com/herman-hang/herman/app/utils"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
)

// BaseRepository 公共仓储层
type BaseRepository struct {
	Model interface{}
	Db    *gorm.DB
}

// PageInfo 分页结构体
type PageInfo struct {
	Page     int64  `json:"page"`     // 页码
	PageSize int64  `json:"pageSize"` // 每页大小
	Keywords string `json:"keywords"` // 关键字
}

// Insert 新增
// @param map[string]interface{} data 待添加数据
// @return toMap err 查询数据，错误信息
func (base *BaseRepository) Insert(data map[string]interface{}) (toMap map[string]interface{}, err error) {
	// 初始化ID，让ID持续自增
	data["id"] = constants.InitId
	if err := mapstructure.WeakDecode(data, base.Model); err != nil {
		return nil, err
	}
	if err := base.Db.Create(base.Model).Error; err != nil {
		return nil, err
	}
	// 模型拷贝
	tempStruct := base.Model
	toMap, err = utils.ToMap(tempStruct, "json")
	if err != nil {
		return nil, err
	}
	return toMap, nil
}

// Find 根据查询条件获取详情
// @param map[string]interface{} condition 查询条件
// @param []string fields 查询指定字段
// @return data err 详情数据，错误信息
func (base *BaseRepository) Find(condition map[string]interface{}, fields ...[]string) (info map[string]interface{}, err error) {
	data := make(map[string]interface{})
	info = make(map[string]interface{})
	if len(fields) > 0 {
		if err := base.Db.Model(&base.Model).Where(condition).Select(fields[0]).Find(&data).Error; err != nil {
			return nil, err
		}
	} else {
		if err := base.Db.Model(&base.Model).Where(condition).Find(&data).Error; err != nil {
			return nil, err
		}
	}
	if len(data) > 0 {
		for k, v := range data {
			// 下划线转为小驼峰
			info[utils.UnderscoreToLowerCamelCase(k)] = v
		}
	}
	return info, nil
}

// Update 更新
// @param []uint ids 查询条件
// @param map[string]interface{} attributes 待更新数据
// @return error 错误信息
func (base *BaseRepository) Update(ids []uint, data map[string]interface{}) error {
	var attributes = make(map[string]interface{})
	// 驼峰转下划线
	for k, v := range data {
		k := utils.ToSnakeCase(k)
		attributes[k] = v
	}
	if err := base.Db.Model(&base.Model).Where("id IN (?)", ids).Updates(attributes).Error; err != nil {
		return err
	}
	return nil
}

// Delete 删除
// @param []uint ids 主键ID
// @return error 错误信息
func (base *BaseRepository) Delete(ids []uint) error {
	if err := base.Db.Delete(&base.Model, ids).Error; err != nil {
		return err
	}
	return nil
}

// IsExist 查询数据是否存在
// @param map[string]interface{} condition 查询条件
// @return bool 返回一个bool值
func (base *BaseRepository) IsExist(condition map[string]interface{}) bool {
	data := make(map[string]interface{})
	err := base.Db.Model(&base.Model).Where(condition).Find(&data).Error
	if err != nil && len(data) > constants.LengthByZero {
		return true
	}
	return false
}

// GetList 获取列表数据
// @param string query 查询条件
// @param []string fields 查询指定字段
// @param string order 排序条件
// @param map[string]interface{} pageInfo 列表分页和关键词数据
// @return list total pageNum err 返回列表，总条数，总页码数，错误信息
func (base *BaseRepository) GetList(query string, fields []string, order string, pageInfo ...map[string]interface{}) (data map[string]interface{}, err error) {
	var (
		page    PageInfo
		total   int64
		pageNum int64
		list    []map[string]interface{}
	)
	if len(pageInfo) > 0 {
		if err := mapstructure.WeakDecode(pageInfo[0], &page); err != nil {
			panic(constants.MapToStruct)
		}
	}
	// 总条数
	base.Db.Model(&base.Model).Count(&total)
	// 计算总页数
	if page.PageSize != 0 && total%page.PageSize != 0 {
		pageNum = total / page.PageSize
		pageNum++
	}
	// 示例 query = fmt.Sprintf(" dns like '%%%s' ", createDbnameInfo.DNS)
	err = base.Db.Model(&base.Model).
		Select(fields).
		Where(query).
		Order(order).
		Limit(int(page.PageSize)).
		Offset(int((page.Page - 1) * page.PageSize)).
		Find(&list).Error
	if err != nil {
		return nil, err
	}
	data = map[string]interface{}{
		"list":     list,          // 数据
		"total":    total,         // 总条数
		"pageNum":  pageNum,       // 总页数
		"pageSize": page.PageSize, // 每页大小
		"page":     page.Page,     // 当前页码
	}
	return data, nil
}

// GetAllData 获取全部数据
// @param []string fields 查询指定字段
// @return list err 返回列表，错误信息
func (base *BaseRepository) GetAllData(fields []string) (data []map[string]interface{}, err error) {
	if len(fields) > 0 {
		if err := base.Db.Model(&base.Model).Select(fields).Find(&data).Error; err != nil {
			return nil, err
		}
	} else {
		if err := base.Db.Model(&base.Model).Find(&data).Error; err != nil {
			return nil, err
		}
	}
	return data, nil
}

```

在仓储层中要使用以上方法，你需要根据Model创建对应的子仓储，然后继承公共的`BaseRepository`结构体才能使用，比如一下是管理员的仓储层。

```go
package repositories

import (
	AdminConstant "github.com/herman-hang/herman/app/constants/admin"
	"github.com/herman-hang/herman/app/models"
	"github.com/herman-hang/herman/kernel/core"
	"gorm.io/gorm"
)

// AdminRepository 管理员表仓储层
type AdminRepository struct {
	BaseRepository
}

// Admin 实例化管理员表仓储层
// @param *gorm.DB tx 事务
// @return AdminRepository 返回管理员表仓储层
func Admin(tx ...*gorm.DB) *AdminRepository {
	if len(tx) > 0 && tx[0] != nil {
		return &AdminRepository{BaseRepository{Model: new(models.Admin), Db: tx[0]}}
	}

	return &AdminRepository{BaseRepository{Model: new(models.Admin), Db: core.Db}}
}

// GetAdminInfo 获取管理员信息
// @param interface{} attributes 管理员id或者管理员user
// @return admin 返回当前管理员的信息
func (u AdminRepository) GetAdminInfo(attributes interface{}) (admin *models.Admin) {
	var err error
	switch attributes.(type) {
	case uint:
		err = core.Db.Where("id = ?", attributes).Find(&admin).Error
	case string:
		err = core.Db.Where("user = ?", attributes).Find(&admin).Error

	}
	if err != nil {
		panic(AdminConstant.GetAdminInfoFail)
	}

	return admin
}
```

重点在这里：

```go
// AdminRepository 管理员表仓储层
type AdminRepository struct {
	BaseRepository
}
```

另外，每个子仓储都必须写上实例化方法：

```go
// Admin 实例化管理员表仓储层
// @param *gorm.DB tx 事务
// @return AdminRepository 返回管理员表仓储层
func Admin(tx ...*gorm.DB) *AdminRepository {
	if len(tx) > 0 && tx[0] != nil {
		return &AdminRepository{BaseRepository{Model: new(models.Admin), Db: tx[0]}}
	}

	return &AdminRepository{BaseRepository{Model: new(models.Admin), Db: core.Db}}
}
```

根据不同的仓储层，实例化结构体要随之改变。比如上面的是管理员仓储层，那么实例化的就是管理员仓储层的结构体，完成了上面的操作之后，这些方法就能在Service层调用了，比如：

```go
// 获取管理员信息
admin := repositories.Admin().GetAdminInfo(fmt.Sprintf("%s", data["user"]))
```

我这里举例只调用了子仓储的一个方法，你还可以通过`repositories.Admin()`去调用公共方法。当然，如果上面已经封装好的方法仍然无法满足你的需求，你可以在子仓储中使用GORM模型进行扩展，GORM更多方法：https://gorm.io/zh_CN/docs/

## 8. 数据库模型

每个模型对应一张数据表，结构体成员采用大驼峰命名，json标签的反射字段采用小驼峰命名，gorm标签的column属性命名与数据库字段对应。例子如下：

```go
package models

import (
	"gorm.io/gorm"
	"time"
)

// Admin 管理员结构体
type Admin struct {
	Id           uint           `json:"id" gorm:"column:id;primary_key;comment:管理员ID"`
	User         string         `json:"user" gorm:"column:user;comment:管理员用户名"`
	Password     string         `json:"password" gorm:"column:password;comment:管理员密码"`
	Photo        string         `json:"photo" gorm:"column:photo;comment:管理员头像"`
	Name         string         `json:"name" gorm:"column:name;comment:真实姓名"`
	Card         string         `json:"card" gorm:"column:card;comment:身份证号码"`
	Sex          uint8          `json:"sex" gorm:"column:sex;default:3;comment:性别(1为女,2为男，3为保密)"`
	Age          uint8          `json:"age" gorm:"column:age;default:0;comment:年龄"`
	Region       string         `json:"region" gorm:"column:region;comment:地区"`
	Phone        string         `json:"phone" gorm:"column:phone;comment:手机号码"`
	Email        string         `json:"email" gorm:"column:email;comment:邮箱"`
	Introduction string         `json:"introduction" gorm:"column:introduction;comment:简介"`
	State        uint8          `json:"state" gorm:"column:state;default:2;comment:状态(1已停用,2已启用)"`
	Sort         uint           `json:"sort" gorm:"column:sort;default:0;comment:排序"`
	LoginOutIp   string         `json:"loginOutIp" gorm:"column:login_out_ip;comment:上一次登录IP地址"`
	LoginTotal   uint           `json:"loginTotal" gorm:"column:login_total;default:0;comment:登录总数"`
	LoginOutAt   time.Time      `json:"loginOutAt" gorm:"column:login_out_at;default:1970-01-01 00:00:00;comment:上一次登录时间"`
	CreatedAt    time.Time      `json:"createdAt" gorm:"column:created_at;comment:创建时间"`
	UpdatedAt    time.Time      `json:"updatedAt" gorm:"column:updated_at;comment:更新时间"`
	DeletedAt    gorm.DeletedAt `json:"deletedAt" gorm:"column:deleted_at;index;comment:删除时间"`
}

// TableName 设置用户表名
func (Admin) TableName() string {
	return "admin"
}

```

其中TableName()是必须的，需要返回一个`string`类型为数据表名称。

## 9. 响应

统一响应方法在`/app/response.go`，代码如下：

```go
package app

import (
	"fmt"
	"github.com/herman-hang/herman/app/constants"
	"github.com/herman-hang/herman/app/utils"
	"net/http"
)

// Response 响应信息结构体
type Response struct {
	HttpCode int         `json:"-"`
	Code     int         `json:"code"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data"`
}

// Option 定义配置选项函数（关键）
type Option func(*Response)

// C 设置JSON结构状态码
// @param int code 状态码
// @return Option 返回配置选项函数
func C(code int) Option {
	return func(this *Response) {
		this.Code = code
	}
}

// M 设置响应信息
// @param string message 自定义响应信息
// @return Option 返回配置选项函数
func M(message string) Option {
	return func(this *Response) {
		this.Message = message
	}
}

// D 设置响应参数
// @param interface{} data 响应数据
// @return Option 返回配置选项函数
func D(data interface{}) Option {
	return func(this *Response) {
		this.Data = data
	}
}

// H 设置HTTP响应状态码
// @param int HttpCode HTTP状态码，比如：200，500等
// @return Option 返回配置选项函数
func H(HttpCode int) Option {
	return func(this *Response) {
		this.HttpCode = HttpCode
	}
}

// Success 方法一：响应函数
// @param *Gin g 上下文结构体
// @param Option opts 接收多个配置选项函数参数，可以是C，M，D，H
func (r *Request) Success(opts ...Option) {
	defaultResponse := &Response{
		HttpCode: http.StatusOK,
		Code:     http.StatusOK,
		Message:  constants.Success,
		Data:     nil,
	}

	// 依次调用opts函数列表中的函数，为结构体成员赋值
	for _, o := range opts {
		o(defaultResponse)
	}
	// 响应http请求
	r.Context.JSON(defaultResponse.HttpCode, defaultResponse)
	return
}

// Json 方法二：响应函数（所有字段转小驼峰写法）
// @param interface{} data 接收响应参数
// @param args 第一个参数为message，第二个参数为code
func (r *Request) Json(data interface{}, args ...interface{}) {
	var jsonString []byte
	// 将数据转为json格式返回
	camelJson, _ := utils.CamelJSON(data)
	switch len(args) {
	case 0:
		jsonString = []byte(fmt.Sprintf(`{"code":%d,"message":"%s","data":%s}`, http.StatusOK, constants.Success, camelJson))
	case 1:
		jsonString = []byte(fmt.Sprintf(`{"code":%d,"message":"%s","data":%s}`, http.StatusOK, args[0], camelJson))
	case 2:
		jsonString = []byte(fmt.Sprintf(`{"code":%d,"message":"%s","data":%s}`, args[1], args[0], camelJson))
	}
	// 响应http请求
	r.Context.Data(http.StatusOK, "application/json", jsonString)
}

```

目前响应json有2种方法，下面是其中一种：

```go
	// 测试路由
	rootEngine.GET("/", func(context *gin.Context) {
		response := app.Request{Context: context}
		response.Success(app.D(map[string]interface{}{
			"message": "Welcome to Herman!",
		}))
	})
```

其中`response.Success()`参数中可以接收4个参数，每一个参数都是响应方法中的一个函数。比如上面就只调用了一个函数`app.D()`，根据业务需求，你还可以在`response.Success()`追加其他函数进去，比如`app.H()`，`app.C()`，`app.M()`。

另一种方法就是直接Json，比如：

```go
func Login(ctx *gin.Context) {
   context := app.Request{Context: ctx}
   data := context.Params()
   context.Json(AdminService.Login(AdminValidate.Login(data)), AdminConstant.LoginSuccess)
}
```

`context.Json()`方法第一个参数是data，第二个参数是message，第三个参数是code，必须严格根据这个顺序来传入，否则出错。

## 10. 测试

单元测试核心代码位于`/kernel/core/test/test_suite.go`，单元测试比较推荐使用套件测试，每个模块需要在`/tests`目录下进行创建，这个模块建议和控制器一一对应。值得注意的是，单元测试支持多应用测试，在做HTTP测试的时候，登录方法都需要封装在`/kernel/core/test/test_suite.go`中，比如框架中的管理员登录：

```go
// AdminLogin 管理员登录
// @return void
func (s *SuiteCase) AdminLogin() {
	var (
		response app.Response
		loginUri = s.AppPrefix + "/admin/login"
	)
	// map转json
	_, _, w := s.Request("POST", loginUri, map[string]interface{}{
		"user":     "admin",
		"password": "123456",
	})
	// json转struct
	_ = json.Unmarshal(w.Body.Bytes(), &response)
	s.Authorization = response.Data.(string)
}
```

封装后登录之后，需要在`SetupSuite()`方法中作逻辑处理，如下：

```go
// SetupSuite 测试套件前置函数
// @return void
func (s *SuiteCase) SetupSuite() {
	settings.InitConfig()
	servers.ZapLogs()
	middlewares.Reload()
	gin.SetMode(settings.Config.Mode)
	e := gin.Default()
	e.Use(middlewares.CatchError())
	core.Engine = routers.InitRouter(e)
	s.AppPrefix = settings.Config.AppPrefix
	switch s.Guard {
	case "admin":
		s.AdminLogin()
	default:
		panic(MiddlewareConstant.GuardError)
	}
}
```

这样就可以在单测里面调用来的，每个单元测试都有一个测试套件方法，如下：

```go
// TestAdminTestSuite 管理员测试套件
// @return void
func TestAdminTestSuite(t *testing.T) {
   suite.Run(t, &AdminTestSuite{SuiteCase: test.SuiteCase{Guard: "admin"}})
}
```

这里根据业务是必须定义的，实例化结构体要根据业务需求随之应变。`Guard: "admin"`表示使用管理员登录方法，整个套件测试例子如下：

```go
package admin

import (
   "fmt"
   "github.com/brianvoe/gofakeit/v6"
   "github.com/herman-hang/herman/app/repositories"
   "github.com/herman-hang/herman/kernel/core/test"
   "github.com/herman-hang/herman/database/seeders/admin"
   "github.com/herman-hang/herman/database/seeders/role"
   "github.com/stretchr/testify/suite"
   "testing"
)

// 管理员测试套件结构体
type AdminTestSuite struct {
   test.SuiteCase
}

var (
   AdminLoginUri = "/admin/login"  // 管理员登录URI
   AdminUri      = "/admin/admins" // 管理员URI
)

// TestLogin 测试管理员登录
// @return void
func (base *AdminTestSuite) TestLogin() {
   base.Assert([]test.Case{
      {
         Method:  "POST",
         Uri:     base.AppPrefix + AdminLoginUri,
         Params:  map[string]interface{}{"user": "admin", "password": "123456"},
         Code:    200,
         Message: "登录成功",
      },
   })
}

// TestAddAdmin 测试添加管理员
// @return void
func (base *AdminTestSuite) TestAddAdmin() {
   roleInfo, _ := repositories.Role().Insert(role.Role())
   adminInfo := admin.Admin()
   adminInfo["roles"] = []map[string]interface{}{
      {
         "name": roleInfo["name"].(string),
         "role": roleInfo["role"].(string),
      },
   }
   base.Assert([]test.Case{
      {
         Method:  "POST",
         Uri:     base.AppPrefix + AdminUri,
         Params:  adminInfo,
         Code:    200,
         Message: "操作成功",
      },
   })
}

// TestModifyAdmin 测试修改管理员
// @return void
func (base *AdminTestSuite) TestModifyAdmin() {
   roleInfo, _ := repositories.Role().Insert(role.Role())
   adminInfo := admin.Admin()
   adminInfo["roles"] = []map[string]interface{}{
      {
         "name": roleInfo["name"].(string),
         "role": roleInfo["role"].(string),
      },
   }
   info, _ := repositories.Admin().Insert(adminInfo)
   base.Assert([]test.Case{
      {
         Method: "PUT",
         Uri:    base.AppPrefix + AdminUri,
         Params: map[string]interface{}{
            "id":           info["id"],
            "user":         gofakeit.Username(),
            "password":     gofakeit.Password(false, false, true, false, false, 10),
            "photo":        gofakeit.ImageURL(100, 100),
            "roles":        adminInfo["roles"],
            "name":         gofakeit.Name(),
            "card":         "450981200008272525",
            "sex":          gofakeit.RandomInt([]int{1, 2, 3}),
            "age":          gofakeit.Number(18, 60),
            "region":       gofakeit.Country(),
            "phone":        "18888888888",
            "email":        gofakeit.Email(),
            "introduction": gofakeit.Sentence(10),
            "state":        gofakeit.RandomInt([]int{1, 2}),
            "sort":         gofakeit.Number(1, 100),
         },
         Code:    200,
         Message: "操作成功",
      },
   })
}

// TestDeleteAdmin 测试根据ID获取管理员详情
// @return void
func (base *AdminTestSuite) TestFindAdmin() {
   roleInfo, _ := repositories.Role().Insert(role.Role())
   adminInfo := admin.Admin()
   adminInfo["roles"] = []map[string]interface{}{
      {
         "name": roleInfo["name"].(string),
         "role": roleInfo["role"].(string),
      },
   }
   info, _ := repositories.Admin().Insert(adminInfo)
   base.Assert([]test.Case{
      {
         Method:  "GET",
         Uri:     base.AppPrefix + AdminUri + "/" + fmt.Sprintf("%d", info["id"]),
         Params:  nil,
         Code:    200,
         Message: "操作成功",
      },
   })
}

// TestGetAdminList 测试删除管理员
// @return void
func (base *AdminTestSuite) TestRemoveAdmin() {
   roleInfo, _ := repositories.Role().Insert(role.Role())
   adminInfo := admin.Admin()
   adminInfo["roles"] = []map[string]interface{}{
      {
         "name": roleInfo["name"].(string),
         "role": roleInfo["role"].(string),
      },
   }
   info, _ := repositories.Admin().Insert(adminInfo)
   base.Assert([]test.Case{
      {
         Method: "DELETE",
         Uri:    base.AppPrefix + AdminUri,
         Params: map[string]interface{}{
            "id": []uint{info["id"].(uint)},
         },
         Code:    200,
         Message: "操作成功",
      },
   })
}

// TestGetAdminList 测试获取管理员列表
// @return void
func (base *AdminTestSuite) TestListAdmin() {
   roleInfo, _ := repositories.Role().Insert(role.Role())
   adminInfo := admin.Admin()
   adminInfo["roles"] = []map[string]interface{}{
      {
         "name": roleInfo["name"].(string),
         "role": roleInfo["role"].(string),
      },
   }
   _, _ = repositories.Admin().Insert(adminInfo)
   base.Assert([]test.Case{
      {
         Method:  "GET",
         Uri:     base.AppPrefix + AdminUri,
         Params:  map[string]interface{}{"page": 1, "pageSize": 2, "keywords": ""},
         Code:    200,
         Message: "操作成功",
         IsList:  true,
         Fields: []string{
            "id",
            "user",
            "photo",
            "sort",
            "state",
            "phone",
            "email",
            "name",
            "card",
            "introduction",
            "sex",
            "age",
            "region",
            "createdAt",
         },
      },
   })
}

// TestAdminTestSuite 管理员测试套件
// @return void
func TestAdminTestSuite(t *testing.T) {
   suite.Run(t, &AdminTestSuite{SuiteCase: test.SuiteCase{Guard: "admin"}})
}
```

## 11. 数据库迁移

目前数据库迁移功能已经非常强大，支持数据库版本更新，版本回滚，强制删除，更新回滚指定版本等等，数据库迁移文件位于`/database/migrations`目录下，命名格式`版本号_文件描述_迁移属性.sql`，其中版本号可以自定义，推荐用自然数自增，避免出现不可预估的问题，迁移文件必须是成对存在的，有迁移文件就必须有回滚文件，比如创建了一个迁移文件`1_init.up.sql`，那么回滚文件`1_init.down.sql`就必须存在，并且文件内容(DDL)不能为空，否则在操作数据库迁移时会出错。根据需求可以执行以下命令：

- 数据库迁移

```shell
herman migrate --status=true --direction=up
```

简写：

```shell
herman migrate -s true
```

或：

```shell
herman migrate -s true -d up
```

- 数据库回滚

```shell
herman migrate --status=true --direction=down
```

简写：

```shell
herman migrate -s true -d down
```

- 强制执行指定版本的文件

```shell
herman migrate --status=true --direction=force --version=1 # 强制执行版本号为1的迁移文件
```

简写：

```shell
herman migrate -s true -d force -v 1 # 强制执行版本号为1的迁移文件
```

- 迁移1个版本

```shell
herman migrate --status=true --direction=up --number=1
```

简写：

```shell
herman migrate -s true -d up -n 1
```

- 回滚1个版本

```shell
herman migrate --status=true --direction=down --number=1
```

简写：

```shell
herman migrate -s true -d down -n 1
```

- 强制删除数据库

```shell
herman migrate --status=true --direction=drop
```

简写：

```shell
herman migrate -s true -d drop
```

## 12. 数据填充

数据填充位于`/database/seeders`，该目录下的每个模块都是对应于控制器。数据填充采用的是`brianvoe/gofakeit`，填充例子如下：

```go
package admin

import (
   "github.com/brianvoe/gofakeit/v6"
)

// Admin 管理员填充器
func Admin() map[string]interface{} {
   return map[string]interface{}{
      "user":         gofakeit.Username(),
      "password":     gofakeit.Password(false, false, true, false, false, 10),
      "photo":        gofakeit.ImageURL(100, 100),
      "name":         gofakeit.Name(),
      "card":         "450981200008272525",
      "sex":          gofakeit.RandomInt([]int{1, 2, 3}),
      "age":          gofakeit.Number(18, 60),
      "region":       gofakeit.Country(),
      "phone":        "18888888888",
      "email":        gofakeit.Email(),
      "introduction": gofakeit.Sentence(10),
      "state":        gofakeit.RandomInt([]int{1, 2}),
      "sort":         gofakeit.Number(1, 100),
   }
}
```

更多规则：https://github.com/brianvoe/gofakeit

## 13. License

Apache License Version 2.0 see http://www.apache.org/licenses/LICENSE-2.0.html
