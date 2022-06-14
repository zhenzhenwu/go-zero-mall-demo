package svc

import (
	"ang-miracle.com/go-zero-mall-demo/service/order/model"
	"ang-miracle.com/go-zero-mall-demo/service/order/rpc/internal/config"
	"ang-miracle.com/go-zero-mall-demo/service/product/rpc/product"
	"ang-miracle.com/go-zero-mall-demo/service/user/rpc/user"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	OrderModel model.OrderModel

	UserRpc user.User
	ProductRpc product.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		OrderModel: model.NewOrderModel(conn, c.CacheRedis),
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)),
		ProductRpc: product.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}
