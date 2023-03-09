package file

import (
	"github.com/gin-gonic/gin"
	"github.com/herman-hang/herman/app"
	FileService "github.com/herman-hang/herman/app/services/file"
	FileValidate "github.com/herman-hang/herman/app/validates/file"
)

func UploadFile(ctx *gin.Context) {
	context := app.Request{Context: ctx}
	FileService.Upload(ctx, FileValidate.Check(ctx))
	context.Json(nil)
}
