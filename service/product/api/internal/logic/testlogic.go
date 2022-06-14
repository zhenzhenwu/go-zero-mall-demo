package logic

import (
	"context"

	"ang-miracle.com/go-zero-mall-demo/service/product/api/internal/svc"
	"ang-miracle.com/go-zero-mall-demo/service/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestLogic {
	return &TestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestLogic) Test() (resp *types.TestRepsonse, err error) {
	// todo: add your logic here and delete this line

	return
}
