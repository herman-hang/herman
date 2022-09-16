package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserInfo 用户信息的Controller
func (a *AppClientUserController) UserInfo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 20000,
		"msg":  "UserInfo PersonalINfo",
	})
}
