package file

import (
	"fmt"
	"github.com/gin-gonic/gin"
	FileConstant "github.com/herman-hang/herman/app/constants/file"
	"github.com/herman-hang/herman/servers/settings"
	"io/ioutil"
	"net/http"
	"os"
)

// adaptiveDownload 适配驱动下载文件
// @param map[string]interface{} info 文件信息
// @return void
func adaptiveDownload(info map[string]interface{}) (data []byte) {
	switch info["drive"].(string) {
	case "local":
		file, err := os.Open(info["filePath"].(string))
		if err != nil { // 处理异常
			panic(FileConstant.OpenFileFail)
		}
		defer func(file *os.File) {
			if err := file.Close(); err != nil {
				panic(FileConstant.CloseFileFail)
			}
		}(file)
		// 读取文件流
		data, err = ioutil.ReadAll(file)
		if err != nil {
			panic(FileConstant.ReadFileFail)
		}
	case "oss": // 阿里云oss
		aliOss := settings.Config.FileStorage.Oss
		fileDrive, err := NewAliOSS(aliOss.Endpoint, aliOss.AccessKeyId, aliOss.AccessKeySecret, aliOss.Bucket)
		if err != nil {
			panic(FileConstant.NewObjectFail)
		}
		data, err = fileDrive.Download(info["filePath"].(string))
		if err != nil {
			panic(FileConstant.Download)
		}
	case "cos": // 腾讯云cos
		cos := settings.Config.FileStorage.Cos
		fileDrive, err := NewTencentCOS(cos.Region, cos.AppId, cos.SecretId, cos.SecretKey, cos.Bucket)
		if err != nil {
			panic(FileConstant.NewObjectFail)
		}
		data, err = fileDrive.Download(info["filePath"].(string))
		if err != nil {
			panic(FileConstant.Download)
		}
	case "qiniu": // 七牛云
		var err error
		qiniu := settings.Config.FileStorage.Qiniu
		fileDrive := NewQiniu(qiniu.SecretKey, qiniu.SecretKey, qiniu.Bucket, qiniu.Domain)
		data, err = fileDrive.Download(info["filePath"].(string))
		if err != nil {
			panic(FileConstant.Download)
		}
	default:
		panic(FileConstant.Download)
	}
	return data
}

// response 响应文件流
// @param *gin.Context ctx 上下文
// @param []byte data 文件流
// @param string fileName 文件名称
// @return void
func response(ctx *gin.Context, data []byte, fileName string) {
	ctx.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	ctx.Data(http.StatusOK, "application/octet-stream", data)
}
