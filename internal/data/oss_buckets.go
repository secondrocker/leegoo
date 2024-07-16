package data

import (
	"leegoo/internal/conf"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type OssBuckets struct {
	client  *oss.Client
	buckets map[string]*oss.Bucket
}

func NewOssClient(c *conf.Data) (*oss.Client, error) {
	client, err := oss.New(c.Oss.Endpoint, c.Oss.AccessKeyId, c.Oss.AccessKeySecret)
	return client, err
}

func NewOssBuckets(ossClient *oss.Client) *OssBuckets {
	return &OssBuckets{
		client:  ossClient,
		buckets: make(map[string]*oss.Bucket),
	}
}

func (b *OssBuckets) Bucket(name string) (*oss.Bucket, error) {
	bucket, ok := b.buckets[name]
	if ok {
		return bucket, nil
	}
	bucket, err := b.client.Bucket(name)
	if err != nil {
		return nil, err
	}
	b.buckets[name] = bucket
	return bucket, nil
}
