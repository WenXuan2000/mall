package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"mall/service/order/model"

	"mall/service/order/rpc/internal/svc"
	"mall/service/order/rpc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveLogic {
	return &RemoveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveLogic) Remove(in *order.RemoveRequest) (*order.RemoveResponse, error) {
	// todo: add your logic here and delete this line
	var res = &model.Order{}
	if ok, err := model.HaveOderByid(l.svcCtx.OrderModel, in.Id, res); ok {
		return nil, status.Error(100, "订单不存在")
	} else if err != nil {
		return nil, err
	}
	if err := l.svcCtx.OrderModel.Delete(res).Error; err != nil {
		return nil, status.Error(500, err.Error())
	}
	return &order.RemoveResponse{}, nil
}
