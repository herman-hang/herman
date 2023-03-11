package file

import (
	"bytes"
	"fmt"
	"github.com/herman-hang/herman/app/common"
	FileConstant "github.com/herman-hang/herman/app/constants/file"
	"github.com/tencentyun/cos-go-sdk-v5"
	"golang.org/x/net/context"
	"io/ioutil"
	"net/http"
	"net/url"
)

// TencentCOS 腾讯云COS结构体
type TencentCOS struct {
	client *cos.Client
	bucket string
}

// NewTencentCOS 实例化一个腾讯云COS对象
// @param string path 文件存储目录
// @return *TencentCOS error 返回腾讯云COS对象和一个错误信息
func NewTencentCOS(region string, appId string, secretId string, secretKey string, bucket string) (*TencentCOS, error) {
	u, _ := url.Parse(fmt.Sprintf("https://%s-%s.cos.%s.myqcloud.com", bucket, appId, region))
	c := cos.NewClient(&cos.BaseURL{BucketURL: u}, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  secretId,
			SecretKey: secretKey,
		},
	})
	return &TencentCOS{client: c, bucket: bucket}, nil
}

// Upload 文件上传
// @param string key 文件key
// @param content 文件流
// @return void
func (t *TencentCOS) Upload(key string, content []byte) error {
	_, err := t.client.Object.Put(context.Background(), key, bytes.NewReader(content), nil)
	if err != nil {
		common.Log.Error(FileConstant.UploadFail)
	}
	return nil
}

// Download 文件下载
// @param string key 文件key
// @return []byte, error 文件流和错误信息
func (t *TencentCOS) Download(key string) ([]byte, error) {
	res, err := t.client.Object.Get(context.Background(), key, nil)
	if err != nil {
		panic(FileConstant.Download)
	}
	return ioutil.ReadAll(res.Body)
}
