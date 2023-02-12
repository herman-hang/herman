package utils

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/herman/app/common"
	UserConstant "github.com/herman/app/constants/user"
	"github.com/herman/servers/settings"
	"strings"
	"time"
)

// Claims 用户信息类，作为生成token的参数
type Claims struct {
	Uid   uint   `json:"uid"`
	Guard string `json:"guard"`
	// jwt-go提供的标准claim
	jwt.StandardClaims
}

// GenerateToken 生成token
// @param UserClaims claims jwt信息结构体
// @return string 返回token
func GenerateToken(claims *Claims) string {
	// token有效时间（纳秒）
	var effectTime = settings.Config.EffectTime * time.Hour
	//设置token有效期
	claims.ExpiresAt = time.Now().Add(effectTime).Unix()
	//生成token
	sign, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(settings.Config.JwtConfig.JwtSecret))
	if err != nil {
		panic(UserConstant.CreateTokenFail)
	}

	return "Bearer " + sign
}

// JwtVerify 验证token
// @param *gin.Context ctx 上下文
// @param string guard 看守器
// @return map[string]interface{} 返回解析token的用户信息
func JwtVerify(ctx *gin.Context, guard string) *Claims {
	// 过滤是否验证token
	token := ctx.GetHeader("Authorization")

	if len(token) == UserConstant.LengthByZero {
		panic(UserConstant.TokenNotExit)
	}

	parts := strings.SplitN(token, " ", UserConstant.SplitByTwo)

	if !(len(parts) == UserConstant.SplitByTwo && parts[0] == "Bearer") {
		panic(UserConstant.TokenError)
	}

	return ParseToken(parts[1], ctx, guard)
}

// ParseToken 解析Token
// @param string tokenString 旧token
// @param *gin.Context ctx 上下文
// @param string guard 看守器
// @return Claims 返回配置好的jwt结构体信息
func ParseToken(tokenString string, ctx *gin.Context, guard string) *Claims {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(settings.Config.JwtConfig.JwtSecret), nil
	})
	if !token.Valid || err != nil {
		panic(UserConstant.TokenExpires)
	}

	claims, ok := token.Claims.(*Claims)

	if !ok || claims.Guard != guard {
		panic(UserConstant.TokenNotValid)
	}

	timeRecord := claims.ExpiresAt - time.Now().Unix()
	// token小于10分钟则刷新token
	if (timeRecord / 60) < UserConstant.Minute {
		ttl, err := common.Redis.TTL(context.Background(), fmt.Sprintf("%s%d", "user_token:", claims.Uid)).Result()
		if err != nil {
			panic(UserConstant.TokenRefreshFail)
		}

		// 如果redis的有效期大于10分钟，说明已经刷新token，直接返回即可，避免再一次签发token
		if ttl.Minutes() > UserConstant.Minute {
			return claims
		}

		newToken := Refresh(token)
		ctx.Header("x-new-token", newToken)

		err = common.Redis.Set(context.Background(),
			fmt.Sprintf("%s%d", "user_token:", claims.Uid),
			newToken,
			time.Duration(timeRecord)*time.Second).Err()
		if err != nil {
			panic(UserConstant.TokenSaveFail)
		}
	}

	return claims
}

// Refresh 更新token
// @param *jwt.Token token Token实例
// @return newToken 返回新token
func Refresh(token *jwt.Token) (newToken string) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		panic(UserConstant.TokenNotValid)
	}

	jwt.TimeFunc = time.Now
	claims.StandardClaims.ExpiresAt = time.Now().Add(2 * time.Hour).Unix()

	newToken = GenerateToken(claims)

	if len(newToken) == UserConstant.LengthByZero {
		panic(UserConstant.TokenRefreshFail)
	}

	return newToken
}
