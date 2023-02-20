# Herman框架

## 1. 序言
基于Gin框架开发，专注于后端快速上手的脚手架。

## 2. 项目结构

```
├─app --------------------------------------------------------- 应用程序目录
│  ├─command -------------------------------------------------- 命令管理目录
│  ├─common --------------------------------------------------- 公共模块目录
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
├─bootstrap --------------------------------------------------- 辅助程序目录
├─config ------------------------------------------------------ 配置文件目录
├─database ---------------------------------------------------- 数据库相关目录
│  ├─migrations ----------------------------------------------- 数据迁移目录
│  ├─seeders -------------------------------------------------- 数据填充目录
│  └─sqls ----------------------------------------------------- 数据库更新SQL文件目录
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
├─storage ----------------------------------------------------- 文件存储目录
├─tests ------------------------------------------------------- 测试目录
├─.air.toml --------------------------------------------------- Air热重载配置文件
├─.gitignore -------------------------------------------------- gitignore文件
├─go.mod ------------------------------------------------------ go.mod文件
├─go.sum ------------------------------------------------------ go.sum文件
├─config.yaml ------------------------------------------------- 环境配置文件
├─main.go ----------------------------------------------------- 程序入口文件
└─README.md --------------------------------------------------- Readme文件
```

## 3. 项目热重载启动
使用 Go 的版本为 1.16 或更高:
```bash
go install github.com/cosmtrek/air@latest
```
最简单的方法是执行
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

## 4. 项目开发规范
#### （1）目录与文件命名
- 目录名称采用小驼峰命名（首字母小写）
- .go文件采用大驼峰命名（首字母大写），例如：`User`，`UserController`
- 配置文件采用大驼峰命名（首字母大写），例如：`SmsConfig.go`
- .sql文件以更新版本号命名，例如：`init.sql`，`1.0.0.sql`，`1.1.0.sql`
- 资源文件(图片，CSS文件，JS文件等)均采用蛇形命名，例如CSS文件：`test.css`，`test_user.css`，以此类推

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
- 全局变量如果类型是一个对象，则采用大驼峰命名（首字母大写），例如：`Db`，`Redis`；类型不是对象的则为小驼峰（首字母小写），例如：`user`，`userName`
- 常量使用大驼峰命名（首字母大写），例如：`Success`，`TokenNotExit`

#### （4）数据表与字段
- 数据表名没有前缀，表名不能出现大写字母，建议以蛇形定义，例如：`user`，`user_role`
- 字段名称采用蛇形命名，不能出现大写字母，例如：`user_id`，`user_name`

## 5. 编码示例
#### （1）路由
第一步：首先到routers文件夹下打开Router.go文件（以用户相关路由为例）

```go
// InitRouter 初始化路由
func InitRouter(rootEngine *gin.Engine) {
	api := rootEngine.Group(settings.ConfigJwtConfig.AppPrefix)

	api.Use(middlewares.Jwt())
	{
        // 用户相关路由(路由从这里编写)
		userRouter := api.Group("/user")
		
        // 这里的user为api文件夹下的user.go文件，Router是user包中的一个函数
		user.Router(userRouter)
	}
}
```

第二步：在`根目录/routers/api`文件夹下创建一个`user.go`文件，编写一下内容

```go
// Router 用户相关路由
func Router(router *gin.RouterGroup) {
    // 第一个参数为路由，第二个参数为指定控制器下的函数
	router.POST("/login", UserController.Login)
}
```

这样就写好一些路由了，例如上面的路由为`/api/v1/user/login`，其中`/api/v1`为路由前缀，可以到配置文件设置

#### （2）控制器

控制器主要负责**数据接收，数据验证，函数调用，响应数据返回**，其他业务逻辑全部在service进行，例如：

```go
// 用户登录
func Login(ctx *gin.Context) {
    // 二次封装上下文
	context := app.Request{Context: ctx}
    // 接收数据
	data := context.Params()
    // 响应数据，并以json格式返回
	context.Json(UserService.Login(UserValidate.Login(data)))
}
```

#### （3）验证器

验证规则文档：https://juejin.cn/post/6847902214279659533

以用户登录验证器为例：

```go
package user

import (
	"fp-back-user/app/utils"
	"fp-back-user/app/validates"
	"github.com/mitchellh/mapstructure"
)

// 这里是需要编写的地方，其中验证规则在结构体标签validate里面
type LoginValidate struct {
	User     string `json:"user" validate:"required,min=5,max=15" label:"用户名"`
	Password string `json:"password" validate:"required,min=6,max=15" label:"密码"`
}

// Login 登录验证器
// @param map 待验证数据
func Login(data map[string]interface{}) map[string]interface{} {
	var login LoginValidate

	// map赋值给结构体
	if err := mapstructure.Decode(data, &login); err != nil {
        // 统一的异常捕捉返回
		panic(err.Error())
	}

	if err := validates.Validate(login); err != nil {
        // 统一的异常捕捉返回
		panic(err.Error())
	}

	toMap, err := utils.ToMap(&login, "json")
	if err != nil {
        // 统一的异常捕捉返回
		panic(err.Error())
	}

	return toMap
}
```

#### （4）逻辑服务

服务主要负责判断，模型调用等等逻辑处理，以用户登录service为例：

```go
package user

import (
	"fmt"
	userConstant "fp-back-user/app/constants/user"
	"fp-back-user/app/models"
	"fp-back-user/app/utils"
)

// Login 用户登录
// @param map data 前端请求数据
func Login(data map[string]interface{}) interface{} {
    // 模型调用
	info, err := models.GetUserInfo(fmt.Sprintf("%v", data["user"]))
	if err != nil {
        // 统一的异常捕捉返回
		panic(err.Error())
	}

	// 密码验证
	if !utils.ComparePasswords(info.Password, fmt.Sprintf("%v", data["password"])) {
        // 统一的异常捕捉返回
		panic(userConstant.PasswordError)
	}

	// 返回数据给控制器
	return utils.GenerateToken(&utils.UserClaims{UserId: info.Id, Issuer: info.User})
}

```

#### （5）仓储层

模型主要查找数据库数据返回给service，以获取用户信息为为例：

```go
// User 实例化结构体并重写BaseRepository
var User = &UserRepository{BaseRepository{Model: new(models.Users)}}

type UserRepository struct {
BaseRepository
}

// GetUserInfo 获取用户信息
// @param string user 用户名
// @return *Users error 返回当前user用户的信息和错误信息
func (u UserRepository) GetUserInfo(user string) (*models.Users, error) {
var users models.Users
err := common.Db.Where("user = ?", user).First(&users).Error

if err != nil && err != gorm.ErrRecordNotFound {
panic(UserConstant.GetUserInfoFail)
}

return &users, nil
}
```
#### （7）数据模型

```go
type Users struct {
	Id           uint       `json:"id" gorm:"primary_key" gorm:"comment:主键ID"`
	User         string     `json:"user" gorm:"comment:用户名"`
	Password     string     `json:"password" gorm:"comment:用户密码"`
	Photo        string     `json:"photo" gorm:"comment:用户头像"`
	Nickname     string     `json:"nickname" gorm:"comment:昵称"`
	Name         string     `json:"name" gorm:"comment:真实姓名"`
	Card         string     `json:"card" gorm:"comment:身份证号码"`
	Sex          string     `json:"sex" gorm:"comment:性别(0为女,1为男，2为保密)"`
	Age          int        `json:"age" gorm:"comment:年龄"`
	Region       string     `json:"region" gorm:"comment:地区"`
	Phone        string     `json:"phone" gorm:"comment:手机号码"`
	Email        string     `json:"email" gorm:"comment:邮箱"`
	Introduction string     `json:"introduction" gorm:"comment:简介"`
	Status       string     `json:"status" gorm:"comment:状态(0已停用,1已启用)"`
	SignOutIp    string     `json:"sign_out_ip" gorm:"comment:最后登录IP地址"`
	SignTotal    string     `json:"sign_total" gorm:"comment:登录总数"`
	SignOutAt    string     `json:"sign_out_at" gorm:"comment:最后登录时间"`
	CreatedAt    time.Time  `json:"created_at" gorm:"comment:创建时间"`
	UpdatedAt    time.Time  `json:"updated_at" gorm:"comment:更新时间"`
	DeletedAt    *time.Time `json:"deleted_at" sql:"index" gorm:"comment:删除时间"`
}
```

#### （6）中间件

中间件编写与其他业务逻辑基本相同，编写完成之后可以到路由文件或者main.go文件进行调用即可，以异常捕捉中间件为例：

第一步：编写中间件

```go
// CatchError 异常捕捉
func CatchError() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		this := app.Gin{C: ctx}
		defer func() {
			if err := recover(); err != nil {
				// 没有定义
				this.Response(app.C(constants.Error), app.M(err.(string)))
				this.C.Abort()
			}
		}()
		this.C.Next()
	}
}
```

第二步：在合适代码中调用

```go
	e := gin.New()
	// 注册中间件
	e.Use(middlewares.CatchError())
```

#### （7）工具类

工具类是一个文件有一个独立的功能，比如结构体转换为Map：

```go
// ToMap 结构体转为Map[string]interface{}
// @param interface in 待转结构体
// @param string tagName 根据指定结构体标签作为key
func ToMap(in interface{}, tagName string) (map[string]interface{}, error) {
	out := make(map[string]interface{})
	v := reflect.ValueOf(in)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// 非结构体返回错误提示
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ToMap only accepts struct or struct pointer; got %T", v)
	}

	t := v.Type()
	// 遍历结构体字段,指定tagName值为map中key;字段值为map中value
	for i := 0; i < v.NumField(); i++ {
		fi := t.Field(i)
		if tagValue := fi.Tag.Get(tagName); tagValue != "" {
			out[tagValue] = v.Field(i).Interface()
		}
	}

	return out, nil
}
```

一般文件里不会出现其他的功能，比如结构体转Map里面不应该出现结构体转切片等逻辑。

