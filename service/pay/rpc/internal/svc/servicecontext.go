package svc

import (
	"ang-miracle.com/go-zero-mall-demo/service/order/rpc/order"
	"ang-miracle.com/go-zero-mall-demo/service/pay/model"
	"ang-miracle.com/go-zero-mall-demo/service/pay/rpc/internal/config"
	"ang-miracle.com/go-zero-mall-demo/service/user/rpc/user"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	PayModel model.PayModel
	UserRpc user.User
	OrderRpc order.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		PayModel: model.NewPayModel(conn, c.CacheRedis),
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)),
		OrderRpc: order.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
	}
}
