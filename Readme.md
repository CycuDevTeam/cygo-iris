# cygo-iris

> 基于iris-go的restful api的业务框架

## 系统要求

- Golang v1.15
- MongoDB
- Redis

## 目录结构

```
.
├── assets          # 静态文件，存放I18N等文件
├── cmd             # 服务命令，存放各种sh脚本
├── controller      # 服务控制器，业务逻辑入口
├── repository      # 持久层
├── route           # 路由
├── service         # 服务层
├── session         # 缓存
├── .env            # 环境变量，用于存放各类服务配置
└── util            # 业务工具库
```

## 使用指南

1. clone该项目至本地
2. 通过全局搜索, 将`cygo-iris`更改为自己的项目名称.
3. ``go mod download`` 下载依赖包
4. 配置`.env`文件，配置各类服务(`cmd/docker-compose`已经预编写了mongodb以及redis服务，有需求可以直接通过``docker-compose up -d``指令启动)
5. ``go run main.go``, happy hack!

## 项目说明

> 以下按照目录结构顺序进行描述

### assets

一般用于存放I18N文件，有其余配置文件也可放置于此。编译Dockerfile时记得复制此文件夹。

### cmd

一般用于存放各种服务指令

### controller

业务逻辑的入口, 目录结构如下

```
controller
├── v1
│   ├── common              
│   │   ├── error_code.go   # 业务逻辑错误码
│   │   └── main.go         # 封装各类常用的response信息
│   └── user                # 业务模块
│       ├── login.go        
│       ├── logout.go       # 接口服务
│       └── register.go
└── v2
    └── common
        ├── error_code.go
        └── main.go
```

v1,v2表示版本控制（若无版本控制需求可省略）。

每个业务模块单独列为一个文件夹，并且每个服务接口单独建立一个文件.

其中，每个接口服务应当严格按照以下格式进行编写，以`login.go`举例

```go
// 在每个服务的顶端，都应当定义服务接口

// LoginRequest 使用JSON进行传递，并使用validate对表单进行验证 
// {
//     "account": "beiyanpiki",
//     "password": "testtest"
// }
type LoginRequest struct {
	Account  string `json:"account" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// LoginResponseData 使用Json进行传递，声明需要返回的JSON
// {
//     "code": 0,
//     "message": "Success",
//     "data": {
//         "uid": "5f7da04ce9e50826845f7b71",
//         "username": "beiyanpiki",
//         "role": 0
//     }
// }
type LoginResponseData struct {
	Uid      string `json:"uid"`
	Username string `json:"username"`
	Role     int    `json:"role"`
}

// 编写业务逻辑，请注意任何需要进行有关数据的CRUD或事务处理，请到Service层进行编写
func Login(ctx iris.Context) {
    // 读取表单并进行验证，如果未通过则返回错误
	var req LoginRequest
	if err := ctx.ReadJSON(&req); err != nil {
		common.FormErrorResponse(ctx, err)
		return
	}
    
    // 根据用户名或账号查询用户
	user := service.GetUserByAccount(req.Account)
    // 校验密码	
    if !user.CheckPassword(req.Password) {
		common.ParamErrorResponse(ctx, "PASSWORD_ERROR")
		return
	}
    
    // 更新session
	service.SetLoginSession(ctx, user.Uid)
	common.SuccessResponse(ctx, LoginResponseData{
		Uid:      user.Uid.Hex(),
		Username: user.Username,
		Role:     user.Role,
	})
}
```

### repository

数据持久层, 其数据只能被`service`层所调用

目录结构如下

```
repository
├── common.go   # Mgo常用查询的封装文件
├── init.go     # Mgo初始化文件
└── user.go     
```

单个模型应当单独建立一个文件，以`user.go`为例。

```go
// 在每个文件的顶部，都应当声明改模型的数据结构
// 并在接下来编写有关该模型的CRUD查询
const (
	NormalUser = 0
	Admin      = 1
	SuperAdmin = 2
)

type User struct {
	// Uid: Primary key (_id)
	Uid      bson.ObjectId `bson:"_id,omitempty"`
	Email    string        `bson:"email"`
	Verified bool          `bson:"verified"`
	Username string        `bson:"username"`
	Password string        `bson:"password"`
	// CreatedTime and LastLogin use timestamp.
	CreatedTime int64 `bson:"created_time"`
	LastLogin   int64 `bson:"last_login"`
	Role        int   `bson:"role"`
	IsBanned    bool  `bson:"is_banned"`
}

func CheckExistByUsername(username string) (bool, error) {
	return Has(UserCollection, bson.M{"username": username})
}
```

### route

仅包含一个route文件，用于控制项目路由

### service

业务逻辑层，处于 `controller` 层和 `repository` 层之间

`service` 只能通过 `repository` 层获取数据. 一般复杂的数据处理、CRUD、事务处理等操作在此处进行

### session

仅包含缓存初始化文件

### util

业务工具库，目录结构如下

```
util
├── i18n            # i18n初始化文件
│   └── init.go
├── log             # log初始化文件
│   └── logger.go
└── validator       # 表单验证器配置
    ├── init.go
    └── regex.go
```

#### validator表单验证器

##### 注册自定义规则

对于`go-playground/validator` (https://godoc.org/github.com/go-playground/validator) 中未内置的规则，我们通常需要自己添加并定义规则，过程如下:

```go
// regex.go
const (
	usernameString = "^[a-zA-Z0-9_-]+$"
)

var (
	usernameRegex = regexp.MustCompile(usernameString)
)

func isUsername(fl validator.FieldLevel) bool {
	return usernameRegex.MatchString(fl.Field().String())
}
```

```go
// init.go

// Add costum validate rule here.
func InitRegexValidator(v *validator.Validate) {
	_ = v.RegisterValidation("is_username", isUsername)
}
```

1. 在`regex.go`,首先编写正则表达式, 如``usernameString = "^[a-zA-Z0-9_-]+$"``
2. 在`regex.go`,定义正则规则, 如``usernameRegex = regexp.MustCompile(usernameString)``
3. 在`regex.go`,编写判断函数，判断字符串是否合法
4. 在`init.go`中注册规则, 重启服务即可.

##### 使用

在定义接口后添加`validate`标签即可，系统会自动验证是否合法

```go
type RegisterRequest struct {
	Email           string `json:"email" validate:"required,email"`
	Username        string `json:"username" validate:"required,min=3,max=20,is_username"` // is_username即在init.go中自定义的验证规则
	Password        string `json:"password" validate:"required,min=8,max=20"`
	PasswordConfirm string `json:"passwordConfirm" validate:"eqfield=Password"`
}
```

##### i18n

如果有国际化需求，我们通常对`actualTag`,`namespace`两个内容进行翻译。具体翻译配置请参考如下(./assert/locale/zh-CN)

```ini
eqfield = "不匹配"
email = "格式错误"
min = "不是正确的长度"
max = "不是正确的长度"
required = "必须填写"
is_username="不是正确的格式"

Email = "邮箱"
Username = "用户名"
Password = "密码"
PasswordConfirm = "密码确认"
```

### .env

环境配置文件，一般仅用于开发环境，生产环境请使用docker的environment进行配置