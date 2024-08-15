package logic

import (
	"context"
	"fmt"
	"mall/service/product/rpc/internal/svc"
	"mall/service/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type DecrStockWODTMLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDecrStockWODTMLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DecrStockWODTMLogic {
	return &DecrStockWODTMLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DecrStockWODTMLogic) DecrStockWODTM(in *product.DecrStockRequest) (*product.DecrStockResponse, error) {
	// todo: add your logic here and delete this line
	// 更新产品库存
	p, _ := l.svcCtx.ProductModel.FindOne(l.ctx, in.Id)
	if p.Stock-1 < 0 {
		return nil, fmt.Errorf("库存不足")
	}
	p.Stock -= 1
	err := l.svcCtx.ProductModel.Update(l.ctx, p)
	if err != nil {
		return nil, err
	}

	// 这种情况是库存不足，不再重试，走回滚
	//if err == dtmcli.ErrFailure {
	//	return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	//}
	if err != nil {
		return nil, err
	}

	return &product.DecrStockResponse{}, nil
}
