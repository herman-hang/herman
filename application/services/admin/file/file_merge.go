package file

import (
	"fmt"
	FileConstant "github.com/herman-hang/herman/application/constants/admin/file"
	"github.com/herman-hang/herman/kernel/app"
	"io"
	"os"
)

// createFile 创建文件
func createFile(hash string) string {
	filePath := mkdir(app.Config.FileStorage.Local.Path) + "/" + hash
	// 删除已经存在的文件
	if _, err := os.Stat(filePath); err == nil {
		err := os.Remove(filePath)
		if err != nil {
			panic(FileConstant.RemoveFileFail)
		}
	}
	return filePath
}

// execMerge 执行文件合并
// @param chunkFile []map[string]interface{} 文件切片信息
// @param file *os.File 文件对象
// @return void
func execMerge(chunkFile []map[string]interface{}, file *os.File) {
	for _, item := range chunkFile {
		if item["state"] != FileConstant.UploadState {
			panic(fmt.Sprintf(FileConstant.ChunkNotUpload, item["chunk_number"]))
		}
		// 打开源文件
		sf, err := os.Open(item["chunk_path"].(string))
		if err != nil {
			panic(FileConstant.OpenFileFail)
		}
		defer sf.Close()
		_, err = io.Copy(file, sf)
		if err != nil {
			panic(FileConstant.MergeFail)
		}
	}
}

// removeChunkFile 删除分片文件
// @param chunkFile []map[string]interface{} 文件切片信息
// @return void
func removeChunkFile(chunkFile []map[string]interface{}) {
	for _, item := range chunkFile {
		path := item["chunk_path"].(string)
		// 删除已经存在的文件
		if _, err := os.Stat(path); err == nil {
			err := os.Remove(path)
			if err != nil {
				panic(FileConstant.RemoveFileFail)
			}
		}
	}
}
