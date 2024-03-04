package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"mall/service/order/model"
	"mall/service/product/rpc/types/product"
	"mall/service/user/rpc/types/user"

	"mall/service/order/rpc/internal/svc"
	"mall/service/order/rpc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *order.CreateRequest) (*order.CreateResponse, error) {
	// todo: add your logic here and delete this line

	if _, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{Id: in.Uid}); err != nil {
		return nil, err
	}
	productRes, err := l.svcCtx.ProductRpc.Detail(l.ctx, &product.DetailRequest{
		Id: in.Pid,
	})
	if err != nil {
		return nil, err
	}
	if productRes.Stock <= 0 {
		return nil, status.Error(500, "产品库存不足")
	}
	newOder := model.Order{
		Uid:    uint64(in.Uid),
		Pid:    uint64(in.Pid),
		Amount: uint(in.Amount),
		Status: 0,
	}
	if err1 := l.svcCtx.OrderModel.Create(&newOder); err1 != nil {
		return nil, status.Error(500, err.Error())
	}
	//这里的产品库存更新存在数据一致性问题，在以往的项目中我们会使用数据库的事务进行这一系列的操作来保证数据的一致性。但是因为我们这边把“订单”和“产品”分成了不同的微服务，在实际的项目中他们可能拥有不同的数据库，所以我们要考虑在跨服务的情况下还能保证数据的一致性，这就涉及到了分布式事务的使用，在后面的章节中我们将介绍使用分布式事务来修改这个下单的逻辑。
	_, err = l.svcCtx.ProductRpc.Update(l.ctx, &product.UpdateRequest{
		Id:     productRes.Id,
		Name:   productRes.Name,
		Desc:   productRes.Desc,
		Stock:  productRes.Stock - 1,
		Amount: productRes.Amount,
		Status: productRes.Status,
	})
	if err != nil {
		return nil, err
	}

	return &order.CreateResponse{Id: int64(newOder.ID)}, nil
}
