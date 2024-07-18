package data

import (
	"leegoo/internal/conf"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/syndtr/goleveldb/leveldb"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	LevelDb   *leveldb.DB
	OssPocket *OssPocket
}

type OssPocket struct {
	client  *oss.Client
	buckets map[string]*oss.Bucket
}

func newOssPocket(ossConfig *conf.Data_Oss) (*OssPocket, error) {
	client, err := oss.New(ossConfig.Endpoint, ossConfig.AccessKeyId, ossConfig.AccessKeySecret)
	if err != nil {
		return nil, err
	}

	return &OssPocket{
		client:  client,
		buckets: make(map[string]*oss.Bucket),
	}, nil
}

func (b *OssPocket) Bucket(name string) (*oss.Bucket, error) {
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

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	//leveldb
	db, err := leveldb.OpenFile(c.Leveldb.GetPath(), nil)
	if err != nil {
		panic(err)
	}

	// oss
	ossPocket, err := newOssPocket(c.Oss)
	if err != nil {
		panic(err)
	}

	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		db.Close()
	}
	return &Data{LevelDb: db, OssPocket: ossPocket}, cleanup, nil
}
