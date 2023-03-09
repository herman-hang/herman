package file

import (
	"github.com/gin-gonic/gin"
	FileConstant "github.com/herman-hang/herman/app/constants/file"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

func Upload(ctx *gin.Context, files []*multipart.FileHeader) {
	path := mkdir("./resources/uploads")
	for _, fileItem := range files {
		err := ctx.SaveUploadedFile(fileItem, path)
		if err != nil {
			panic(FileConstant.SaveFail)
		}
	}
}

// mkdir 创建文件目录
// @param basePath string 文件目录
// @return string 返回文件目录
func mkdir(basePath string) string {
	//	1.获取当前时间,并且格式化时间
	folderName := time.Now().Format("2006/01/02")
	folderPath := filepath.Join(basePath, folderName)
	//使用MkdirAll会创建多层级目录
	if err := os.MkdirAll(folderPath, 0755); err != nil {
		panic(FileConstant.MkdirFail)
	}
	return folderPath
}
