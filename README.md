# Herman框架

## 1. 序言
### 介绍 

Herman基于Gin，Casbin，Kafka，Mysql，Redis，Zap，Cobra，Grom开发，专注于后端快速上手的一款开源，简洁，轻量框架。

### 项目结构

```
├─app --------------------------------------------------------- 应用程序目录
│  ├─command -------------------------------------------------- 命令管理目录
│  ├─constants ------------------------------------------------ 常量存放目录
│  ├─controllers ---------------------------------------------- 控制器目录
│  ├─jobs ----------------------------------------------------- 队列作业目录
│  ├─middlewares ---------------------------------------------- 中间件目录
│  ├─models --------------------------------------------------- 数据模型目录
│  ├─repositories --------------------------------------------- 仓储层目录
│  ├─services ------------------------------------------------- 服务处理目录
│  ├─utils ---------------------------------------------------- 工具类目录
│  ├─validates ------------------------------------------------ 验证器目录
│  ├─Request.go ----------------------------------------------- 请求对象库
│  └─Response.go ---------------------------------------------- 响应对象库
├─bootstrap --------------------------------------------------- 程序核心目录
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
│  ├─log ------------------------------------------------------ 日志驱动目录
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
├─docker-compose.yaml ----------------------------------------- Dodcker容器编排配置文件
├─LICENSE ----------------------------------------------------- 程序许可证文件
├─Makefile ---------------------------------------------------- 程序Makefile文件
├─main.go ----------------------------------------------------- 程序入口文件
└─README.md --------------------------------------------------- Readme文件
```

### 开发规范

#### （1）目录与文件命名

- 目录名称采用小驼峰命名（首字母小写）
- .go文件采用大驼峰命名（首字母大写），例如：`User`，`UserController`
- 配置文件采用大驼峰命名（首字母大写），例如：`SmsConfig.go`
- 数据库迁移文件采用下划线命名，例如：`1_init.down.sql`，`1_init.up.sql`，1为版本号，init为自定义名称，down代表回滚，up代码更新。
- 资源文件(图片，CSS文件，JS文件等)均采用蛇形命名，例如CSS文件：`test.css`，`test_user.css`，以此类推
- 测试文件命名根据控制器文件加`_test.go`，例如：`UserController_test.go`，`_test.go`是golang强制遵循的规范

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

Golang虽然是一门面向过程的语言，但是也引入了容器的概念，对项目核心的对象，比如Redis，MySQL，Casbin等都已存放在`/bootstrap/core/Container.go`文件中。

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

中间件分为**前置中间件**和**后置中间件**的，主要存放在`/app/middlewares`，比如以下定义的中间件：

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

命令行核心采用cobra实现，主要存放在`/app/command`，命令注册在`/bootstrap/casbin/Casbin.go`文件，比如以下例子：

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
herman version # Herman version: 1.3.0
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

队列采用kafka，主要存放在`/app/jobs`，比如以下短信发送例子：

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

目前框架只支持Redis缓存，对象挂载在`/bootstrap/core/Container.go`中，使用前要先设置上下文：

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

辅助函数又称工具类，主要存放在`/app/utils`中，如果项目中存在一些碎片化的代码，想把它做进一步封装，就可以在该目录下创建一个工具文件，在里面完成响应的封装。比如我下方做了一个验证码工厂：

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

Casbin是一种轻量级的开源访问控制框架，支持多种访问控制模型，如RBAC, ABAC和ACL。框架中已经采用了RBAC，适配GORM来做角色资源管理，可以灵活管理角色的权限。核心封装代码在`/bootstrap/casbin/Casbin.go`。框架Casbin的对象挂载在容器`/bootstrap/core/Container.go`，调用：

```go
success, _ := core.Casbin.Enforce(info.User, ctx.Request.URL.Path, ctx.Request.Method)
```

更多学习：https://casbin.org/zh/docs/category/the-basics

### 配置

框架的所有配置都是通过读取根目录下的`config.yaml`文件所得，并且存放在`/config`目录中，调用方式：

```go
settings.Config
```

比如获取MySQL的配置

```go
settings.Config.Mysql
```

当前，如果你不想创建配置文件作映射，也可以直接获取环境文件`config.yaml`的配置，但是不建议这么操作。

```go
viper.Get("app")
```

## 3. 路由

路由沿用了Gin集成的功能，所有路由定义在`/routers/Router.go`，以下例子：

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

## 5.验证器

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

如果有对验证器公共结构体进行重写，那么久可以使用结构体的公共方法check，上面控制器的例子就是使用了验证器的公共方法。

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

业务需要扩展验证器，可以直接在定义验证器文件中自定义规则即可，比如上面的例子就是把管理员登录是否需要验证码做了2种场景验证。

## 6. 服务



## 7. 仓储



## 8. 数据库模型



## 9. 响应



## 10. 测试



## 11. 数据库迁移



## 12. 数据填充



## 13. License

