package app

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"github.com/herman-hang/herman/config"
	"go.uber.org/zap"
)

var (
	Version = "1.6.0"
	Config  = new(config.App)
	Engine  *gin.Engine
	Log     *zap.SugaredLogger
)
