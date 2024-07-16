package data

import (
	"leegoo/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/syndtr/goleveldb/leveldb"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewOssClient, NewOssBuckets, NewData, NewGreeterRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	LevelDb   *leveldb.DB
	OssPocket *OssBuckets
}

// NewData .
func NewData(c *conf.Data, ossPocket *OssBuckets, logger log.Logger) (*Data, func(), error) {
	//leveldb
	db, err := leveldb.OpenFile(c.Leveldb.GetPath(), nil)
	if err != nil {
		panic(err)
	}
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		db.Close()
	}
	return &Data{LevelDb: db, OssPocket: ossPocket}, cleanup, nil
}
