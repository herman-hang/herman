package admin

import (
	AdminController "github.com/fp/fp-gin-framework/app/controllers/admin"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.RouterGroup) {
	router.POST("/login", AdminController.Login)
}
