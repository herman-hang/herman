package file

import (
	"crypto/sha256"
	"encoding/hex"
	FileConstant "github.com/herman-hang/herman/app/constants/file"
	"github.com/herman-hang/herman/app/repositories"
	"github.com/herman-hang/herman/servers/settings"
	"io/ioutil"
	"mime/multipart"
	"path/filepath"
	"strings"
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

// adaptiveUpload 适配驱动文件上传
// @param string fileHash 文件hash值
// @param []byte content 文件流
// @return string 返回一个文件路径
func adaptiveUpload(fileHash string, content []byte) (filePath string) {
	// 当使用oss,cos,qiniu时,文件路径为空
	filePath = FileConstant.EmptyString
	switch settings.Config.FileStorage.Drive {
	case "local":
		fileDrive := NewLocalOSS(settings.Config.FileStorage.Local.Path)
		filePath = fileDrive.path
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
	return filePath
}
