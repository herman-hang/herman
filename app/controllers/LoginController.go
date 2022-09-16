package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserLoginController 用户登录的Controller
func (a *AppClientUserController) UserLoginController(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 20000,
		"msg":  "UserLoginController",
	})
}

// UserLoginController 用户登录的Controller
func (a *AppClientUserController) UserResetPasswordController(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 20000,
		"msg":  "UserResetPasswordController",
	})
}
