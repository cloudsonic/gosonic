package ledis

import (
	"sync"

	"github.com/cloudsonic/sonic-server/conf"
	"github.com/siddontang/ledisdb/config"
	"github.com/siddontang/ledisdb/ledis"
)

var (
	_ledisInstance *ledis.Ledis
	_dbInstance    *ledis.DB
	once           sync.Once
)

func Db() *ledis.DB {
	once.Do(func() {
		config := config.NewConfigDefault()
		config.DataDir = conf.Sonic.DbPath
		l, _ := ledis.Open(config)
		instance, err := l.Select(0)
		if err != nil {
			panic(err)
		}
		_ledisInstance = l
		_dbInstance = instance
	})
	return _dbInstance
}

func dropDb() {
	Db()
	_ledisInstance.FlushAll()
}
