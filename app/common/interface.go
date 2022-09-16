package common

import "github.com/gin-gonic/gin"

type BaseController interface {
	Name() string
	RegisterRoute(*gin.RouterGroup)
}
