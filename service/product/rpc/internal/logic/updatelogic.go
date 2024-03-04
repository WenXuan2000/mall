package logic

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"mall/service/product/model"
	"mall/service/product/rpc/internal/svc"
	"mall/service/product/rpc/types/product"

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

func (l *UpdateLogic) Update(in *product.UpdateRequest) (*product.UpdateResponse, error) {
	// todo: add your logic here and delete this line
	var newproduct = &model.Product{}
	if err := l.svcCtx.ProductModel.Where("id=?", in.Id).First(&newproduct).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, status.Error(100, "没有该物品")
	} else if err != nil {
		return nil, status.Error(500, err.Error())
	}
	if in.Name != "" {
		newproduct.Name = in.Name
	}
	if in.Desc != "" {
		newproduct.Desc = in.Desc
	}
	if in.Stock != 0 {
		newproduct.Stock = in.Stock
	}
	if in.Amount != 0 {
		newproduct.Amount = in.Amount
	}
	if in.Status != 0 {
		newproduct.Status = in.Status
	}
	if err := l.svcCtx.ProductModel.Save(&newproduct).Error; err != nil {
		return nil, status.Error(500, "无法保存")
	}
	return &product.UpdateResponse{}, nil
}
