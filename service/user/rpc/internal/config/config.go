package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Salt string

	Mysql struct {
		Datasource string
	}

	CacheRedis cache.CacheConf
}
