package svc

import (
	"ang-miracle.com/go-zero-mall-demo/service/order/api/internal/config"
	"ang-miracle.com/go-zero-mall-demo/service/order/rpc/order"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	OrderRpc order.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		OrderRpc: order.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
	}
}
