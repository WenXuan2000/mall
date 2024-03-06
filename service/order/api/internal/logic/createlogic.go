package logic

import (
	"context"
	"github.com/dtm-labs/dtmgrpc"
	"google.golang.org/grpc/status"
	"mall/service/order/rpc/types/order"
	"mall/service/product/rpc/types/product"

	"mall/service/order/api/internal/svc"
	"mall/service/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateRequest) (resp *types.CreateResponse, err error) {
	orderRpcBusiServer, err := l.svcCtx.Config.OrderRpc.BuildTarget()
	if err != nil {
		return nil, status.Error(100, "订单创建异常")
	}
	productRpcBusiServer, err := l.svcCtx.Config.ProductRpc.BuildTarget()
	if err != nil {
		return nil, status.Error(100, "订单创建异常")
	}
	var dtmServer = "etcd://etcd:2379/dtmservice"
	gid := dtmgrpc.MustGenGid(dtmServer)
	saga := dtmgrpc.NewSagaGrpc(dtmServer, gid).Add(orderRpcBusiServer+"/order.Order/Create", orderRpcBusiServer+"/order.Order/CreateRevert", &order.CreateRequest{
		Uid:    req.Uid,
		Pid:    req.Pid,
		Amount: req.Amount,
		Status: 0,
	}).Add(productRpcBusiServer+"/product.Product/DecrStock", productRpcBusiServer+"/product.Product/DecrStockRevert", &product.DecrStockRequest{
		Id:  req.Pid,
		Num: 1,
	})
	err = saga.Submit()
	if err != nil {
		return nil, err
	}

	return &types.CreateResponse{}, nil
}
