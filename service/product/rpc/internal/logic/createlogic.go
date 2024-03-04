package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"mall/service/product/model"

	"mall/service/product/rpc/internal/svc"
	"mall/service/product/rpc/types/product"

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

func (l *CreateLogic) Create(in *product.CreateRequest) (*product.CreateResponse, error) {
	// todo: add your logic here and delete this line
	if model.HavaProductByName(l.svcCtx.ProductModel, in.Name) {
		return nil, status.Error(100, "重复商品加入")
	}
	newprodcut := model.Product{
		Name:   in.Name,
		Desc:   in.Desc,
		Stock:  in.Stock,
		Amount: in.Amount,
		Status: in.Status,
	}
	res := l.svcCtx.ProductModel.Create(&newprodcut)
	if res.Error != nil {
		return nil, status.Error(500, res.Error.Error())
	}
	return &product.CreateResponse{Id: int64(newprodcut.ID)}, nil
}
