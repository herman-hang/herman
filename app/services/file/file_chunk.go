package file

import (
	FileConstant "github.com/herman-hang/herman/app/constants/file"
	"github.com/herman-hang/herman/app/repositories"
)

// IsExist 判断切片是否存在
// @param map[string]interface{} data 文件信息
// @return []map[string]interface{} 返回该文件的切片
func IsExist(data map[string]interface{}) ([]map[string]interface{}, error) {
	var existInfo []map[string]interface{}
	// 查询文件是否存在
	info, err := repositories.File().Find(map[string]interface{}{
		"hash": data["hash"],
	}, []string{"id", "file_name", "file_ext", "file_size"})
	if err != nil {
		return nil, err
	}
	if len(info) > 0 {
		// 查询分片数据
		chunkInfo, err := repositories.FileChunk().FindChunk(info["id"].(uint))
		if err == nil {
			existInfo = chunkInfo
		}
	}
	return existInfo, nil
}

// Chunk 文件分片
// @param map[string]interface{} data 文件信息
// @param uint fileId 文件ID
// @return []map[string]interface{} 返回该文件的切片
func Chunk(data map[string]interface{}, fileId uint) []map[string]interface{} {
	var (
		i         uint64
		fileSize  uint64
		fileChunk uint64
		chunkSize uint64
		info      []map[string]interface{}
	)
	// 文件大小
	fileSize = data["fileSize"].(uint64)
	// 计算每片大小10MB
	fileChunk = 10 * 1024 * 1024
	// 分片方案实现
	if fileSize%fileChunk == 0 {
		chunkSize = fileSize / fileChunk
	} else {
		chunkSize = fileSize/fileChunk + 1
	}
	for i = 1; i <= chunkSize; i++ {
		if i == chunkSize && fileChunk*chunkSize > fileSize {
			fileChunk = fileSize - (chunkSize-1)*fileChunk
		}
		info = append(info, map[string]interface{}{
			"file_id":      fileId,
			"chunk_number": i,
			"chunk_size":   fileChunk,
			"state":        FileConstant.NotUploadState,
			"progress":     int((float64(i) / float64(chunkSize)) * 100),
		})
	}
	return info
}

// DataFactory 数据工厂
// @param []map[string]interface{} data 文件的切片信息
// @return []map[string]interface{} 返回该文件加工后的切片
func DataFactory(data []map[string]interface{}) []map[string]interface{} {
	var info []map[string]interface{}
	for _, v := range data {
		info = append(info, map[string]interface{}{
			"file_id":      v["file_id"],
			"chunk_number": v["chunk_number"],
			"chunk_size":   v["chunk_size"],
			"progress":     v["progress"],
		})
	}
	return info
}
