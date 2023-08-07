package file

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	FileConstant "github.com/herman-hang/herman/application/constants/admin/file"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
)

// AliOSS 阿里云OSS结构体
type AliOSS struct {
	client *oss.Client
	bucket *oss.Bucket
}

// NewAliOSS 实例化一个阿里云OSS对象
// @param string path 文件存储目录
// @return *TencentCOS error 返回阿里云OSS和一个错误信息
func NewAliOSS(endpoint string, accessKeyId string, accessKeySecret string, bucketName string) (*AliOSS, error) {
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		return nil, errors.Wrap(err, FileConstant.CreateOSSClientFail)
	}
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return nil, errors.Wrap(err, FileConstant.GetOSSBucketFail)
	}
	return &AliOSS{client: client, bucket: bucket}, nil
}

// Upload 文件上传
// @param string key 文件key
// @param content 文件流
// @return void
func (a *AliOSS) Upload(key string, content []byte) error {
	return a.bucket.PutObject(key, bytes.NewReader(content))
}

// Download 文件下载
// @param string key 文件key
// @return []byte, error 文件流和错误信息
func (a *AliOSS) Download(key string) ([]byte, error) {
	body, err := a.bucket.GetObject(key)
	if err != nil {
		return nil, err
	}
	// 数据读取完成后，获取的流必须关闭，否则会造成连接泄漏，导致请求无连接可用，程序无法正常工作。
	defer func(body io.ReadCloser) {
		if err := body.Close(); err != nil {
			panic(FileConstant.CloseFileFail)
		}
	}(body)
	data, err := ioutil.ReadAll(body)
	return data, nil
}
