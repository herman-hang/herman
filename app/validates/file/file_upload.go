package file

import (
	"github.com/gin-gonic/gin"
	"github.com/herman-hang/herman/app/constants"
	FileConstant "github.com/herman-hang/herman/app/constants/file"
	"github.com/herman-hang/herman/app/utils"
	"github.com/herman-hang/herman/app/validates"
	"github.com/mitchellh/mapstructure"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
)

// Merge 重写验证器结构体，切记不使用引用，而是拷贝
var Merge = validates.Validates{Validate: ChunkMergeValidate{}}

// ChunkUploadValidate 分片上传验证器规则
type ChunkUploadValidate struct {
	FileId      uint  `json:"fileId" validate:"required" label:"文件ID"`
	ChunkNumber uint8 `json:"chunkNumber" validate:"required" label:"分片编号"`
	Hash        uint8 `json:"hash" validate:"required" label:"文件块校验和"`
}

// ChunkMergeValidate 分片合并验证器规则
type ChunkMergeValidate struct {
	FileId uint `json:"fileId" validate:"required" label:"文件ID"`
}

// Check 小文件上传验证方法
// @param *gin.Context ctx 上下文对象
// @return void
func Check(ctx *gin.Context) (files []*multipart.FileHeader) {
	// 请求体最大允许的大小不能超过 100MB
	if err := ctx.Request.ParseMultipartForm(100 << 20); err != nil {
		panic(FileConstant.MaxMemory)
	}

	form := ctx.Request.MultipartForm
	files = form.File["files"]
	// 检查文件数量
	if len(files) == FileConstant.CountZero {
		panic(FileConstant.Empty)
	}
	if len(files) > FileConstant.MaxCount {
		panic(FileConstant.SurpassMaxCount)
	}
	if err := Validate(files); err != nil {
		panic(err.Error())
	}

	return files
}

// ChunkCheck 分片上传验证方法
// @param *gin.Context ctx 上下文对象
// @param map[string]interface{} data 请求参数
// @return map[string]interface{} *multipart.FileHeader 转换后的map和文件对象
func ChunkCheck(ctx *gin.Context, data map[string]interface{}) (map[string]interface{}, *multipart.FileHeader) {
	var chunk ChunkUploadValidate
	// 请求体最大允许的大小不能超过 15MB
	if err := ctx.Request.ParseMultipartForm(15 << 20); err != nil {
		panic(FileConstant.MaxMemory)
	}

	// map赋值给结构体
	if err := mapstructure.WeakDecode(data, &chunk); err != nil {
		panic(constants.MapToStruct)
	}
	if err := validates.Validate(chunk); err != nil {
		panic(err.Error())
	}

	toMap, err := utils.ToMap(chunk, "json")

	if err != nil {
		panic(constants.StructToMap)
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		panic(FileConstant.UploadFail)
	}

	// 检查文件大小
	if file.Size > 10<<20 {
		panic(file.Filename + FileConstant.SizeFail)
	}
	return toMap, file
}

// Validate 文件验证器
// @param []*multipart.FileHeader files 文件列表
// @return err 返回错误信息
func Validate(files []*multipart.FileHeader) error {
	for _, file := range files {
		ext := strings.ToLower(filepath.Ext(file.Filename))
		// 检查文件扩展名
		if isUnsafeExtension(ext) {
			panic(file.Filename + FileConstant.ExtFail + ext)
		}

		// 检查文件类型
		if !isSafeFileType(file) {
			panic(file.Filename + FileConstant.TypeFail + file.Header.Get("Content-Type"))
		}

		// 检查文件名
		safeName := filepath.Base(file.Filename)
		if !isSafeName(safeName) {
			panic(file.Filename + FileConstant.NameFail + safeName)
		}

		// 检查文件大小
		if file.Size > 10<<20 {
			panic(file.Filename + FileConstant.SizeFail)
		}
	}
	return nil
}

// isUnsafeExtension 检查文件扩展名
// @param string ext 文件扩展名
// @return bool true:不安全 false:安全
func isUnsafeExtension(ext string) bool {
	unsafe := []string{".exe", ".dll", ".bat", ".sh", ".jsp", ".php"}
	for _, u := range unsafe {
		if strings.Compare(ext, u) == 0 {
			return true
		}
	}
	return false
}

// isSafeFileType 检查文件类型
// @param *multipart.FileHeader f 文件
// @return bool true:安全 false:不安全
func isSafeFileType(f *multipart.FileHeader) bool {
	buf := make([]byte, 512)
	file, err := f.Open()
	if err != nil {
		return false
	}
	defer func(file multipart.File) {
		if err := file.Close(); err != nil {
			panic(FileConstant.CloseFileFail)
		}
	}(file)
	if _, err = file.Read(buf); err != nil {
		return false
	}
	fileType := http.DetectContentType(buf)

	switch fileType {
	case "image/jpeg",
		"image/png",
		"image/gif",
		"text/plain",
		"application/pdf",
		"application/msword",
		"application/vnd.ms-excel":
		return true
	default:
		return false
	}
}

// isSafeName 检查文件名
// @param string name 文件名
// @return bool true:安全 false:不安全
func isSafeName(name string) bool {
	// 不能以路径分隔符开头，不能包含路径分隔符，不能包含“..”
	if strings.HasPrefix(name, "/") || strings.HasPrefix(name, "\\") {
		return false
	}

	if strings.Contains(name, "/") || strings.Contains(name, "\\") || strings.Contains(name, "..") {
		return false
	}

	return true
}
