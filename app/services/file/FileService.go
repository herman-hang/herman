package file

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	FileConstant "github.com/herman-hang/herman/app/constants/file"
	"github.com/herman-hang/herman/app/models"
	"github.com/herman-hang/herman/app/repositories"
	"github.com/herman-hang/herman/servers/settings"
	"io/ioutil"
	"mime/multipart"
	"path/filepath"
	"strings"
)

// Storage 文件存储接口实现
type Storage interface {
	Upload(key string, content []byte) error
	Download(key string) ([]byte, error)
	Delete(key string) error
}

// Upload 文件上传
// @param ctx *gin.Context 上下文
// @param files []*multipart.FileHeader 文件对象切片
// @return void
func Upload(ctx *gin.Context, files []*multipart.FileHeader) {
	// 存储驱动
	drive := settings.Config.FileStorage.Drive
	admin, _ := ctx.Get("admin")
	info := admin.(*models.Admin)
	filePath := ""
	for _, fileItem := range files {
		// hash值
		fileHash, content := calculateHash(fileItem)
		// 文件扩展名
		fileExt := strings.ToLower(filepath.Ext(fileItem.Filename))
		// 文件类型
		fileType := fileItem.Header.Get("Content-Type")
		switch drive {
		case "local":
			filePath = settings.Config.FileStorage.Local.Path + "/" + fileHash
			fileDrive := NewLocalOSS(settings.Config.FileStorage.Local.Path)
			if err := fileDrive.Upload(fileHash, content); err != nil {
				panic(FileConstant.UploadFail)
			}
		case "oss":
			oss := settings.Config.FileStorage.Oss
			fileDrive, err := NewAliOSS(oss.Endpoint, oss.AccessKeyId, oss.AccessKeySecret, oss.Bucket)
			if err != nil {
				panic(FileConstant.NewObjectFail)
			}
			if err := fileDrive.Upload(fileHash, content); err != nil {
				panic(FileConstant.UploadFail)
			}
		case "cos":
			cos := settings.Config.FileStorage.Cos
			fileDrive, err := NewTencentCOS(cos.Region, cos.AppId, cos.SecretId, cos.SecretKey, cos.Bucket)
			if err != nil {
				panic(FileConstant.NewObjectFail)
			}
			if err := fileDrive.Upload(fileHash, content); err != nil {
				panic(FileConstant.UploadFail)
			}
		case "qiniu":
			qiniu := settings.Config.FileStorage.Qiniu
			fileDrive := NewQiniu(qiniu.SecretKey, qiniu.SecretKey, qiniu.Bucket, qiniu.Domain)
			if err := fileDrive.Upload(fileHash, content); err != nil {
				panic(FileConstant.UploadFail)
			}
		default:
			panic(FileConstant.ConfigFileDriveError)
		}
		// 记录文件上传
		record(drive, info.Id, fileItem.Filename, fileExt, fileType, filePath, fileHash, fileItem.Size)
	}
}

// record 记录文件上传
// @param string drive 存储驱动
// @param uint creatorId 上传者ID
// @param string fileName 文件名
// @param string fileExt 文件扩展名
// @param string fileType 文件类型
// @param string filePath 文件路径
// @param string hash 文件hash值
// @param int64 fileSize 文件大小
// @return map[string]interface{} 返回一个map
func record(drive string,
	creatorId uint,
	fileName string,
	fileExt string,
	fileType string,
	filePath string,
	hash string,
	fileSize int64) map[string]interface{} {
	info, err := repositories.File.Insert(map[string]interface{}{
		"drive":      drive,
		"creator_id": creatorId,
		"file_name":  fileName,
		"file_ext":   fileExt,
		"file_type":  fileType,
		"file_path":  filePath,
		"hash":       hash,
		"file_size":  fileSize,
	})
	if err != nil {
		panic(FileConstant.RecordFileFail)
	}
	return info
}

// calculateHash 计算文件Hash值
// @param *multipart.FileHeader file 文件对象
// @return hash content 返回一个hash值和一个文件流
func calculateHash(file *multipart.FileHeader) (hash string, content []byte) {
	fp, err := file.Open()
	if err != nil {
		panic(FileConstant.OpenFileFail)
	}
	defer func(fp multipart.File) {
		if err := fp.Close(); err != nil {
			panic(FileConstant.CloseFileFail)
		}
	}(fp)
	content, err = ioutil.ReadAll(fp)
	if err != nil {
		panic(FileConstant.ReadFileFail)
	}
	hashed := sha256.New()
	hashed.Write(content)
	return hex.EncodeToString(hashed.Sum(nil)), content
}
