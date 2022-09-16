package controllers

import (
	"fp-back-user/app/common"
	"github.com/gin-gonic/gin"
)

type AppClientUserController struct {
}

func NewAppClientUserController() common.BaseController {
	return &AppClientUserController{}
}

func (a AppClientUserController) Name() string {
	return "user-app-client"
}

func (a AppClientUserController) RegisterRoute(group *gin.RouterGroup) {
	group.GET("/user/client", a.UserLoginController)
	group.GET("/user/reset/password", a.UserResetPasswordController)
	group.GET("/user/personal/info", a.UserInfo)
}
