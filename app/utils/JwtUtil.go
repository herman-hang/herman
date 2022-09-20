package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

// UserClaims 用户信息类，作为生成token的参数
type UserClaims struct {
	ID           string `json:"id"`
	User         string `json:"user"`
	Password     string `json:"password"`
	Nickname     string `json:"nickname"`
	Sex          string `json:"sex"`
	Age          int    `json:"age"`
	Region       string `json:"region"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	Introduction string `json:"introduction"`
	Status       string `json:"status"`
	// jwt-go提供的标准claim
	jwt.StandardClaims
}

var (
	// 自定义的token秘钥
	secret = []byte("16849841325189456f487")
	// 该路由下不校验token
	noVerify = []interface{}{
		"/api/v1/user/login",
		"/test",
	}
	// token有效时间（纳秒）
	effectTime = 2 * time.Hour
)

// GenerateToken 生成token
func GenerateToken(claims *UserClaims) string {
	//设置token有效期，也可不设置有效期，采用redis的方式
	//   1)将token存储在redis中，设置过期时间，token如没过期，则自动刷新redis过期时间，
	//   2)通过这种方式，可以很方便的为token续期，而且也可以实现长时间不登录的话，强制登录
	//本例只是简单采用 设置token有效期的方式，只是提供了刷新token的方法，并没有做续期处理的逻辑
	claims.ExpiresAt = time.Now().Add(effectTime).Unix()
	//生成token
	sign, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
	if err != nil {
		panic(err)
	}

	return sign
}

// JwtVerify 验证token
func JwtVerify(ctx *gin.Context) {
	//过滤是否验证token
	//文档里我没给出utils.IsContainArr的代码，这个东西不重要，你直接删掉这段都行，这只是一个url过滤的逻辑
	/*	if utils.IsContainArr(noVerify, c.Request.RequestURI) {
		return
	}*/
	token := ctx.GetHeader("Authorization")
	if token == "" {
		panic("token not exist !")
	}
	// 验证token，并存储在请求中
	ctx.Set("user", parseToken(token))
}

// 解析Token
func parseToken(tokenString string) *UserClaims {
	//解析token
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		panic(err)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		panic("token is valid")
	}

	return claims
}

// Refresh 更新token
func Refresh(tokenString string) string {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}

	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		panic(err)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		panic("token is valid")
	}

	jwt.TimeFunc = time.Now
	claims.StandardClaims.ExpiresAt = time.Now().Add(2 * time.Hour).Unix()

	return GenerateToken(claims)
}
