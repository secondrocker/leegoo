package data

import (
	nUrl "net/url"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type OssData struct {
	data *Data
}

func NewOssData(data *Data) *OssData {
	return &OssData{
		data: data,
	}
}

func (d *OssData) GetUrl(bucketName string, path string) (string, error) {
	bucket, err := d.data.OssPocket.Bucket(bucketName)
	if err != nil {
		return "", err
	}
	url, err := bucket.SignURL(path, oss.HTTPGet, 3600)
	if err != nil {
		url, err = nUrl.QueryUnescape(url)
	}
	return url, err
}

func (d *OssData) Restore(bucketName string, path string) (string, error) {
	bucket, err := d.data.OssPocket.Bucket(bucketName)
	if err != nil {
		return "", err
	}
	err = bucket.RestoreObject(path)
	if err != nil {
		return "", err
	}
	url, err := bucket.SignURL(path, oss.HTTPGet, 3600)
	if err != nil {
		url, err = nUrl.QueryUnescape(url)
	}
	return url, err
}

func (d *OssData) Delete(bucketName string, path string) error {
	bucket, err := d.data.OssPocket.Bucket(bucketName)
	if err != nil {
		return err
	}
	return bucket.DeleteObject(path)
}
