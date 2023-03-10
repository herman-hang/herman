package file

import (
	"bytes"
	"context"
	FileConstant "github.com/herman-hang/herman/app/constants/file"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"io"
	"io/ioutil"
	"net/http"
)

type Qiniu struct {
	accessKey string
	secretKey string
	bucket    string
	domain    string
}

func NewQiniu(accessKey string, secretKey string, bucket string, domain string) *Qiniu {
	return &Qiniu{
		accessKey: accessKey,
		secretKey: secretKey,
		bucket:    bucket,
		domain:    domain,
	}
}

// Upload 文件上传
// @param string key 文件key
// @param content 文件流
// @return void
func (q *Qiniu) Upload(key string, content []byte) error {
	// 构建表单上传的策略
	policy := &storage.PutPolicy{
		Scope: q.bucket,
	}
	mac := qbox.NewMac(q.accessKey, q.secretKey)
	uploadToken := policy.UploadToken(mac)

	cfg := storage.Config{}
	formUploader := storage.NewFormUploader(&cfg)

	ret := storage.PutRet{}
	err := formUploader.Put(context.Background(), &ret, uploadToken, key, bytes.NewReader(content), int64(len(content)), nil)
	if err != nil {
		panic(FileConstant.UploadFail)
	}
	return nil
}

// Download 文件下载
// @param string key 文件key
// @return []byte, error 文件流和错误信息
func (q *Qiniu) Download(key string) ([]byte, error) {
	url := q.domain + "/" + key
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(FileConstant.Download)
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			panic(FileConstant.CloseFileFail)
		}
	}(res.Body)
	return ioutil.ReadAll(res.Body)
}
