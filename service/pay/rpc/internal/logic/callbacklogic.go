package logic

import (
	"context"

	"ang-miracle.com/go-zero-mall-demo/service/pay/rpc/internal/svc"
	"ang-miracle.com/go-zero-mall-demo/service/pay/rpc/pay"

	"github.com/zeromicro/go-zero/core/logx"
)

type CallbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CallbackLogic {
	return &CallbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CallbackLogic) Callback(in *pay.CallbackRequest) (*pay.CallbackResponse, error) {
	// todo: add your logic here and delete this line

	return &pay.CallbackResponse{}, nil
}
