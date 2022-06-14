package svc

import (
	"ang-miracle.com/go-zero-mall-demo/service/user/api/internal/config"
	"ang-miracle.com/go-zero-mall-demo/service/user/rpc/user"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct { // logic所依赖的资源池
	Config config.Config

	UserRpc user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
