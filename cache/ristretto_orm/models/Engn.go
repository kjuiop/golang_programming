package models

import (
	"github.com/dgraph-io/ristretto"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type Engn struct {
	db    *xorm.Engine
	cache *ristretto.Cache
}

func NewEngn() (*Engn, error) {
	engn, err := xorm.NewEngine("mysql", "root:1234@tcp(127.0.0.1:3306)/tortee")
	if err != nil {
		return nil, err
	}
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,
		MaxCost:     1 << 30,
		BufferItems: 64,
	})
	if err != nil {
		return nil, err
	}

	e := Engn{
		db:    engn,
		cache: cache,
	}

	return &e, nil
}

func (engn *Engn) CloseEngn() error {
	engn.cache.Close()
	return engn.db.Close()
}
