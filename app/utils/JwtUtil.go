package utils

import (
	UserConstant "fp-back-user/app/constants/user"
	"fp-back-user/app/models"
	"fp-back-user/settings"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

// UserClaims 用户信息类，作为生成token的参数
type UserClaims struct {
	ID     uint   `json:"id"`
	Issuer string `json:"user"`
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
	sign, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(settings.Config.JwtSecret))
	if err != nil {
		panic(err.Error())
	}

	return "Bearer " + sign
}

// JwtVerify 验证token
func JwtVerify(ctx *gin.Context) {
	//过滤是否验证token
	token := ctx.GetHeader("Authorization")

	if token == "" {
		panic(UserConstant.TokenNotExit)
	}

	parts := strings.SplitN(token, " ", 2)

	if !(len(parts) == 2 && parts[0] == "Bearer") {
		panic(UserConstant.TokenError)
	}

	// 验证token，并存储在请求中
	ctx.Set("userInfo", models.UserInfo(ParseToken(parts[1], ctx).ID))
}

// ParseToken 解析Token
func ParseToken(tokenString string, ctx *gin.Context) *UserClaims {
	//解析token
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(settings.Config.JwtSecret), nil
	})

	if err != nil {
		panic(UserConstant.TokenParseError)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		panic(UserConstant.TokenNotValid)
	}

	if !claims.VerifyExpiresAt(time.Now().Unix(), false) {
		panic(UserConstant.TokenExpires)
	}

	// token小于10分钟则刷新token
	if ((claims.ExpiresAt - time.Now().Unix()) / 60) < 119 {
		ctx.Header("x-new-token", Refresh(tokenString))
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
		panic(UserConstant.TokenNotValid)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		panic(UserConstant.TokenNotValid)
	}

	jwt.TimeFunc = time.Now
	claims.StandardClaims.ExpiresAt = time.Now().Add(2 * time.Hour).Unix()

	return GenerateToken(claims)
}
