package file

import (
	"github.com/gin-gonic/gin"
	"github.com/herman-hang/herman/app"
	FileService "github.com/herman-hang/herman/app/services/file"
	FileValidate "github.com/herman-hang/herman/app/validates/file"
)

// UploadFile 上传文件
// @param ctx *gin.Context 上下文
// @return void
func UploadFile(ctx *gin.Context) {
	context := app.Request{Context: ctx}
	files := FileValidate.Check(ctx)
	FileService.Upload(context.Context, files)
	context.Json(nil)
}
