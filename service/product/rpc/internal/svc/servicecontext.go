package svc

import (
	"ang-miracle.com/go-zero-mall-demo/service/product/model"
	"ang-miracle.com/go-zero-mall-demo/service/product/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	ProductModel model.ProductModel // 依赖
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		ProductModel: model.NewProductModel(conn, c.CacheRedis), // 依赖注入
	}
}
