package svc

import (
	"ang-miracle.com/go-zero-mall-demo/service/pay/api/internal/config"
	"ang-miracle.com/go-zero-mall-demo/service/pay/rpc/pay"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	PayRpc pay.Pay
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		PayRpc: pay.NewPay(zrpc.MustNewClient(c.PayRpc)),
	}
}
