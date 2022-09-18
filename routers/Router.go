package routers

import (
	"fp-back-user/routers/api/user"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.RouterGroup) {
	userRouter := router.Group("/user")

	user.Router(userRouter)
}
