package logic

import (
	"ang-miracle.com/go-zero-mall-demo/service/order/model"
	"context"
	"google.golang.org/grpc/status"

	"ang-miracle.com/go-zero-mall-demo/service/order/rpc/internal/svc"
	"ang-miracle.com/go-zero-mall-demo/service/order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type PaidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPaidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaidLogic {
	return &PaidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PaidLogic) Paid(in *order.PaidRequest) (*order.PaidResponse, error) {
	// 查询订单是否存在
	res, err := l.svcCtx.OrderModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "订单不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	res.Status = 1

	err = l.svcCtx.OrderModel.Update(l.ctx, res)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &order.PaidResponse{}, nil
}
