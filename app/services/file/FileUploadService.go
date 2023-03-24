package file

import (
	"crypto/sha256"
	"encoding/hex"
	FileConstant "github.com/herman-hang/herman/app/constants/file"
	"github.com/herman-hang/herman/app/repositories"
	"github.com/herman-hang/herman/kernel/core"
	"github.com/herman-hang/herman/servers/settings"
	"io/ioutil"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"
)

// Exec 执行文件上传
// @param files []*multipart.FileHeader 文件对象切片
// @param creatorId uint 创建者id
// @return fileInfos existFileInfos 待记录的文件信息,已存在的文件信息
func Exec(files []*multipart.FileHeader, creatorId uint) (fileInfos []map[string]interface{}, existFileInfos []map[string]interface{}) {
	for _, fileItem := range files {
		// hash值
		fileHash, content := calculateHash(fileItem)
		// 文件扩展名
		fileExt := strings.ToLower(filepath.Ext(fileItem.Filename))
		// 文件类型
		fileType := fileItem.Header.Get("Content-Type")
		// 判断当前文件是否存在,不存在则上传
		find, _ := repositories.File().Find(map[string]interface{}{"hash": fileHash}, []string{"id", "file_name", "file_type", "file_ext", "file_size"})
		if len(find) > 0 {
			existFileInfos = append(existFileInfos, find)
		} else {
			// 适配驱动上传
			filePath := adaptiveUpload(fileHash, content)
			// 待记录的文件信息
			fileInfos = append(fileInfos, map[string]interface{}{
				"drive":     settings.Config.FileStorage.Drive,
				"creatorId": creatorId,
				"fileName":  fileItem.Filename,
				"fileExt":   fileExt,
				"fileType":  fileType,
				"filePath":  filePath,
				"hash":      fileHash,
				"fileSize":  fileItem.Size,
			})
		}

	}
	return fileInfos, existFileInfos
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

// adaptiveUpload 适配驱动上传文件
// @param string fileHash 文件hash值
// @param []byte content 文件流
// @return filePath 返回一个文件路径
func adaptiveUpload(fileHash string, content []byte) (filePath string) {
	folderName := time.Now().Format("2006/01/02")
	switch settings.Config.FileStorage.Drive {
	case "local":
		filePath = mkdir(settings.Config.FileStorage.Local.Path) + "/" + fileHash
		fileDrive := NewLocalOSS()
		go func() {
			if err := fileDrive.Upload(filePath, content); err != nil {
				core.Log.Error(FileConstant.UploadFail)
			}
		}()
	case "oss":
		aliOss := settings.Config.FileStorage.Oss
		filePath = filepath.Join(aliOss.Path, folderName) + "/" + fileHash
		fileDrive, err := NewAliOSS(aliOss.Endpoint, aliOss.AccessKeyId, aliOss.AccessKeySecret, aliOss.Bucket)
		if err != nil {
			panic(FileConstant.NewObjectFail)
		}
		go func() {
			if err := fileDrive.Upload(filePath, content); err != nil {
				core.Log.Error(FileConstant.UploadFail)
			}
		}()
	case "cos":
		cos := settings.Config.FileStorage.Cos
		filePath = filepath.Join(cos.Path, folderName) + "/" + fileHash
		fileDrive, err := NewTencentCOS(cos.Region, cos.AppId, cos.SecretId, cos.SecretKey, cos.Bucket)
		if err != nil {
			panic(FileConstant.UploadFail)
		}
		go func() {
			if err := fileDrive.Upload(filePath, content); err != nil {
				core.Log.Error(FileConstant.UploadFail)
			}
		}()
	case "qiniu":
		qiniu := settings.Config.FileStorage.Qiniu
		filePath = filepath.Join(qiniu.Path, folderName) + "/" + fileHash
		fileDrive := NewQiniu(qiniu.AccessKey, qiniu.SecretKey, qiniu.Bucket, qiniu.Domain)
		go func() {
			if err := fileDrive.Upload(filePath, content); err != nil {
				core.Log.Error(FileConstant.UploadFail)
			}
		}()
	default:
		panic(FileConstant.ConfigFileDriveError)
	}

	return strings.ReplaceAll(filePath, "\\", "/")
}
