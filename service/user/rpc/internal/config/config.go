package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {		// 依赖
		DataSource string
	}
	CacheRedis cache.CacheConf // 依赖
	Salt string
}
