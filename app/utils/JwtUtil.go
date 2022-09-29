package utils

import (
	UserConstant "fp-back-user/app/constants/user"
	"fp-back-user/settings"
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

// GenerateToken 生成token
func GenerateToken(claims *UserClaims) string {
	// token有效时间（纳秒）
	var effectTime = settings.Config.EffectTime * time.Hour
	//设置token有效期，也可不设置有效期，采用redis的方式
	claims.ExpiresAt = time.Now().Add(effectTime).Unix()
	//生成token
	sign, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(settings.Config.JwtSecret)
	if err != nil {
		panic(err.Error())
	}

	return sign
}

// JwtVerify 验证token
func JwtVerify(ctx *gin.Context) {
	//过滤是否验证token
	token := ctx.GetHeader("Authorization")

	if token == "" {
		panic(UserConstant.TokenNotExit)
	}

	// 验证token，并存储在请求中
	ctx.Set("user", parseToken(token))
}

// 解析Token
func parseToken(tokenString string) *UserClaims {
	//解析token
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return settings.Config.JwtSecret, nil
	})

	if err != nil {
		panic(err.Error())
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		panic(UserConstant.TokenNotValid)
	}

	return claims
}

// Refresh 更新token
func Refresh(tokenString string) string {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}

	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return settings.Config.JwtSecret, nil
	})

	if err != nil {
		panic(err.Error())
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		panic(UserConstant.TokenNotValid)
	}

	jwt.TimeFunc = time.Now
	claims.StandardClaims.ExpiresAt = time.Now().Add(2 * time.Hour).Unix()

	return GenerateToken(claims)
}
