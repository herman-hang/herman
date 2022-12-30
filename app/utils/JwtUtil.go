package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/fp/fp-gin-framework/app/common"
	UserConstant "github.com/fp/fp-gin-framework/app/constants/user"
	"github.com/fp/fp-gin-framework/servers/settings"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

// UserClaims 用户信息类，作为生成token的参数
type UserClaims struct {
	Uid    uint   `json:"user_id"`
	Issuer string `json:"user"`
	// jwt-go提供的标准claim
	jwt.StandardClaims
}

// GenerateToken 生成token
// @param UserClaims claims jwt信息结构体
// @return string 返回token
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
// @param *gin.Context ctx 上下文
// @return map[string]interface{} 返回解析token的用户信息
func JwtVerify(ctx *gin.Context) *UserClaims {
	//过滤是否验证token
	token := ctx.GetHeader("Authorization")

	if token == "" {
		panic(UserConstant.TokenNotExit)
	}

	parts := strings.SplitN(token, " ", 2)

	if !(len(parts) == 2 && parts[0] == "Bearer") {
		panic(UserConstant.TokenError)
	}

	return ParseToken(parts[1], ctx)
}

// ParseToken 解析Token
// @param string tokenString 旧token
// @param *gin.Context ctx 上下文
// @return UserClaims 返回配置好的jwt结构体信息
func ParseToken(tokenString string, ctx *gin.Context) (claims *UserClaims) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(settings.Config.JwtSecret), nil
	})

	if err != nil {
		panic(UserConstant.TokenParseError)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok || !token.Valid {
		panic(UserConstant.TokenNotValid)
	}

	if !claims.VerifyExpiresAt(time.Now().Unix(), false) {
		panic(UserConstant.TokenExpires)
	}

	timeRecord := claims.ExpiresAt - time.Now().Unix()
	// token小于10分钟则刷新token
	if (timeRecord / 60) < 10 {
		_, err := common.Redis.Get(fmt.Sprintf("%v%v", "user-token:", claims.Uid)).Result()
		if err != nil {
			newToken := Refresh(tokenString)
			ctx.Header("x-new-token", newToken)

			err = common.Redis.Set(fmt.Sprintf("%v%v", "user-token:", claims.Uid), newToken, time.Duration(timeRecord)*time.Second).Err()
			if err != nil {
				panic(UserConstant.TokenSaveFail)
			}
		}
	}

	return claims
}

// Refresh 更新token
// @param string tokenString 旧token
// @return newToken 返回新token
func Refresh(tokenString string) (newToken string) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}

	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(settings.Config.JwtSecret), nil
	})

	if err != nil {
		panic(UserConstant.TokenNotValid)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok || !token.Valid {
		panic(UserConstant.TokenNotValid)
	}

	jwt.TimeFunc = time.Now
	claims.StandardClaims.ExpiresAt = time.Now().Add(2 * time.Hour).Unix()

	newToken = GenerateToken(claims)

	if newToken == "" {
		panic(UserConstant.TokenRefreshFail)
	}

	return newToken
}
