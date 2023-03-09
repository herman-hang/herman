package file

import (
	FileConstant "github.com/herman-hang/herman/app/constants/file"
	"os"
	"path/filepath"
	"time"
)

// LocalOSS 本地存储对象结构体
type LocalOSS struct {
	path string
}

// NewLocalOSS 实例化一个本地存储对象
// @param string path 文件存储目录
// @return *LocalOSS 返回本地存储对象
func NewLocalOSS(path string) *LocalOSS {
	return &LocalOSS{
		path: path,
	}
}

// Upload 文件上传
// @param string key 文件key
// @param content 文件流
// @return error 返回一个错误信息
func (l *LocalOSS) Upload(key string, content []byte) error {
	create, err := os.Create(l.path + "/" + key)
	if err != nil {
		return err
	}
	defer create.Close()

	if _, err := create.Write(content); err != nil {
		return err
	}
	return nil
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
