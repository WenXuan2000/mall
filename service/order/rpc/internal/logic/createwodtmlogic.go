package logic

import (
	"context"
	"fmt"
	"mall/service/order/model"
	"mall/service/user/rpc/types/user"

	"mall/service/order/rpc/internal/svc"
	"mall/service/order/rpc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateWODTMLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateWODTMLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateWODTMLogic {
	return &CreateWODTMLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateWODTMLogic) CreateWODTM(in *order.CreateRequest) (*order.CreateResponse, error) {
	_, err1 := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		Id: in.Uid,
	})
	if err1 != nil {
		return &order.CreateResponse{}, fmt.Errorf("用户不存在")
	}
	newOrder := model.Order{
		Uid:    in.Uid,
		Pid:    in.Pid,
		Amount: in.Amount,
		Status: 0,
	}
	_, err1 = l.svcCtx.OrderModel.Insert(l.ctx, &newOrder)
	if err1 != nil {
		return &order.CreateResponse{}, fmt.Errorf("订单创建失败")
	}
	return &order.CreateResponse{}, nil
}
