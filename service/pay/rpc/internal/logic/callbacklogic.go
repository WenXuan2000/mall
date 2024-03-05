package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"mall/service/order/rpc/types/order"
	"mall/service/pay/model"
	"mall/service/user/rpc/types/user"

	"mall/service/pay/rpc/internal/svc"
	"mall/service/pay/rpc/types/pay"

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
	// 查询用户是否存在
	_, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		Id: in.Uid,
	})
	if err != nil {
		return nil, err
	}

	// 查询订单是否存在
	_, err = l.svcCtx.OrderRpc.Detail(l.ctx, &order.DetailRequest{
		Id: in.Oid,
	})
	if err != nil {
		return nil, err
	}

	// 查询支付是否存在
	res, err := l.svcCtx.PayModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "支付不存在")
		}
		return nil, status.Error(500, err.Error())
	}
	// 支付金额与订单金额不符
	if in.Amount != res.Amount {
		return nil, status.Error(100, "支付金额与订单金额不符")
	}

	res.Source = in.Source
	res.Status = in.Status

	err = l.svcCtx.PayModel.Update(l.ctx, res)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// 更新订单支付状态
	_, err = l.svcCtx.OrderRpc.Paid(l.ctx, &order.PaidRequest{
		Id: in.Oid,
	})
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &pay.CallbackResponse{}, nil
}

/*
查询用户信息： 通过调用 l.svcCtx.UserRpc.UserInfo 查询指定用户（通过 in.Uid 指定）的信息。如果查询出错，直接返回错误。

查询订单信息： 通过调用 l.svcCtx.OrderRpc.Detail 查询指定订单（通过 in.Oid 指定）的详细信息。如果查询出错，直接返回错误。

查询支付信息： 通过调用 l.svcCtx.PayModel.FindOne 查询指定支付（通过 in.Id 指定）的详细信息。如果支付信息不存在，返回支付不存在的错误。如果其他错误发生，返回对应的错误。

检查支付金额： 检查支付金额是否与订单金额一致。如果不一致，返回支付金额与订单金额不符的错误。

更新支付信息： 更新支付信息的来源和状态，然后通过 l.svcCtx.PayModel.Update 方法将更新后的支付信息写回数据库。

更新订单支付状态： 通过调用 l.svcCtx.OrderRpc.Paid 更新订单的支付状态。

返回响应： 如果以上步骤都成功，返回一个空的支付回调响应，表示支付回调处理成功。

这段代码的目的是确保支付回调请求的有效性，包括检查用户、订单和支付信息的存在性，以及支付金额是否与订单金额一致。如果所有检查都通过，将更新支付和订单的状态，并返回成功的响应。
*/
