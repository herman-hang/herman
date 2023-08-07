package file

import (
	"errors"
	"github.com/gin-gonic/gin"
	FileConstant "github.com/herman-hang/herman/application/constants/admin/file"
	"github.com/herman-hang/herman/application/models"
	"github.com/herman-hang/herman/application/repositories"
	"github.com/herman-hang/herman/kernel/app"
	"github.com/herman-hang/herman/kernel/core"
	"gorm.io/gorm"
	"io/ioutil"
	"mime/multipart"
	"os"
)

// Upload 文件上传
// @param ctx *gin.Context 上下文
// @param files []*multipart.FileHeader 文件对象切片
// @return existFileInfos []map[string]interface{} 文件切片信息
func Upload(ctx *gin.Context, files []*multipart.FileHeader) (existFileInfos []map[string]interface{}) {
	// 获取登录信息
	info, _ := ctx.Get("admin")
	admin := info.(*models.Admin)
	// 执行文件上传
	fileInfos, existFileInfos := Exec(files, admin.Id)
	for _, info := range fileInfos {
		// 保存文件信息
		fileInfo, err := repositories.File().Insert(info)
		if err != nil {
			panic(FileConstant.RecordFileFail)
		}
		existFileInfos = append(existFileInfos, map[string]interface{}{
			"id":       fileInfo["id"],
			"fileName": fileInfo["fileName"],
			"fileType": fileInfo["fileType"],
			"fileExt":  fileInfo["fileExt"],
			"fileSize": fileInfo["fileSize"],
		})
	}
	return existFileInfos
}

// Download 文件下载
// @param ctx *gin.Context 上下文
// @param data map[string]interface{} 请求参数
// @return void
func Download(ctx *gin.Context, data map[string]interface{}) {
	info, err := repositories.File().Find(map[string]interface{}{
		"id": data["id"],
	}, []string{"id", "drive", "file_name", "file_path", "hash"})
	if len(info) == 0 {
		panic(FileConstant.NotExist)
	}
	if err != nil {
		panic(FileConstant.DownloadFail)
	}
	// 返回文件流
	stream := adaptiveDownload(info)
	// 响应文件流
	response(ctx, stream, info)
}

// Preview 图片预览
// @param ctx *gin.Context 上下文
// @param data map[string]interface{} 请求参数
// @return void
func Preview(ctx *gin.Context, data map[string]interface{}) {
	info, err := repositories.File().Find(map[string]interface{}{
		"id": data["id"],
	}, []string{"id", "drive", "file_name", "file_type", "file_ext", "file_path", "file_size", "hash"})
	if len(info) == 0 {
		panic(FileConstant.NotExist)
	}
	if err != nil {
		panic(FileConstant.PreviewFail)
	}
	// 判断是否为图片
	if info["fileType"].(string) != "image/jpeg" &&
		info["fileType"].(string) != "image/png" &&
		info["fileType"].(string) != "image/gif" {
		panic(FileConstant.NotImage)
	}
	// 返回文件流
	stream := adaptiveDownload(info)
	// 响应文件流
	ctx.Header("Content-Type", info["fileType"].(string))
	ctx.Header("Connection", "keep-alive")
	response(ctx, stream, info)
}

// Prepare 生成分片上传方案
// @param ctx *gin.Context 上下文
// @param data map[string]interface{} 请求参数
// @return info []map[string]interface{} 文件切片信息
func Prepare(ctx *gin.Context, data map[string]interface{}) []map[string]interface{} {
	var info []map[string]interface{}
	if app.Config.FileStorage.Drive != "local" {
		panic(FileConstant.NotSupport)
	}
	// 判断是否已经存在
	info, err := IsExist(data)
	if err == nil && len(info) > 0 {
		return info
	}
	// 不存在，则进行方案制作
	err = core.Db().Transaction(func(tx *gorm.DB) error {
		// 获取登录信息
		origin, _ := ctx.Get("admin")
		admin := origin.(*models.Admin)
		data["drive"] = app.Config.FileStorage.Drive
		data["creatorId"] = admin.Id
		file, err := repositories.File(tx).Insert(data)
		if err != nil {
			return errors.New(FileConstant.PrepareFail)
		}
		// 进行分片
		mapSlice := Chunk(data, file["id"].(uint))
		if err := repositories.FileChunk(tx).Create(mapSlice); err != nil {
			return errors.New(FileConstant.CreateFail)
		}
		// 数据加工
		info = DataFactory(mapSlice)
		return nil
	})
	if err != nil {
		panic(err.Error())
	}

	return info
}

// ChunkUpload 分片上传
// @param data map[string]interface{} 请求参数
// @param file *multipart.FileHeader 文件对象
// @return void
func ChunkUpload(data map[string]interface{}, file *multipart.FileHeader) {
	data["state"] = FileConstant.UploadState
	// 判断分片是否已经上传
	info, _ := repositories.FileChunk().Find(data, []string{"id", "hash"})
	if len(info) > 0 {
		panic(FileConstant.ChunkExist)
	}
	fp, err := file.Open()
	if err != nil {
		panic(FileConstant.OpenFileFail)
	}
	content, err := ioutil.ReadAll(fp)
	if err != nil {
		panic(FileConstant.ReadFileFail)
	}
	// 上传分片
	filePath := adaptiveUpload(info["hash"].(string), content)
	err = repositories.FileChunk().Update([]uint{info["id"].(uint)}, map[string]interface{}{
		"filePath": filePath,
		"state":    FileConstant.UploadState,
	})
	if err != nil {
		panic(FileConstant.UploadFail)
	}
}

// Merge 分片合并
// @param data map[string]interface{} 请求参数
// @return void
func Merge(data map[string]interface{}) {
	fileInfo, _ := repositories.File().Find(map[string]interface{}{
		"id": data["id"],
	}, []string{"id", "hash"})
	if len(fileInfo) == 0 {
		panic(FileConstant.NotExist)
	}
	chunkFile, _ := repositories.FileChunk().FindChunk(data["id"].(uint))
	if len(chunkFile) == 0 {
		panic(FileConstant.ChunkNotExist)
	}
	filePath := createFile(fileInfo["hash"].(string))
	// 创建目标文件
	file, err := os.Create(filePath)
	if err != nil {
		panic(FileConstant.CreateFileFail)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(FileConstant.CloseFileFail)
		}
	}(file)
	// 合并文件
	go execMerge(chunkFile, file)
	// 删除分片文件
	go removeChunkFile(chunkFile)
}
