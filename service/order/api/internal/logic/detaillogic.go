package logic

import (
	"ang-miracle.com/go-zero-mall-demo/service/order/rpc/order"
	"context"

	"ang-miracle.com/go-zero-mall-demo/service/order/api/internal/svc"
	"ang-miracle.com/go-zero-mall-demo/service/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.DetailRequest) (resp *types.DetailResponse, err error) {
	res, err := l.svcCtx.OrderRpc.Detail(l.ctx, &order.DetailRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.DetailResponse{
		Id:     res.Id,
		Uid:    res.Uid,
		Pid:    res.Pid,
		Amount: res.Amount,
		Status: res.Status,
	}, nil
}
