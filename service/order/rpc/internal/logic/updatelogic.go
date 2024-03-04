package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"mall/service/order/model"

	"mall/service/order/rpc/internal/svc"
	"mall/service/order/rpc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *order.UpdateRequest) (*order.UpdateResponse, error) {
	// todo: add your logic here and delete this line
	var res = &model.Order{}
	if ok, err := model.HaveOderByid(l.svcCtx.OrderModel, in.Id, res); ok {
		return nil, status.Error(100, "订单不存在")
	} else if err != nil {
		return nil, err
	}

	if in.Uid != 0 {
		res.Uid = uint64(in.Uid)
	}
	if in.Pid != 0 {
		res.Pid = uint64(in.Pid)
	}
	if in.Amount != 0 {
		res.Amount = uint(in.Amount)
	}
	if in.Status != 0 {
		res.Status = uint(in.Status)
	}

	if err := l.svcCtx.OrderModel.Save(res).Error; err != nil {
		return nil, status.Error(500, err.Error())
	}
	return &order.UpdateResponse{}, nil
}
