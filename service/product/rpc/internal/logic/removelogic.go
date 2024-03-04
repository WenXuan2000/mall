package logic

import (
	"context"
	"errors"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"mall/service/product/model"

	"mall/service/product/rpc/internal/svc"
	"mall/service/product/rpc/types/product"

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

func (l *RemoveLogic) Remove(in *product.RemoveRequest) (*product.RemoveResponse, error) {
	// todo: add your logic here and delete this line
	var newproduct = &model.Product{}
	if err := l.svcCtx.ProductModel.Where("id=?", in.Id).First(&newproduct).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, status.Error(100, "没有该物品")
	} else if err != nil {
		return nil, status.Error(500, err.Error())
	}
	if err := l.svcCtx.ProductModel.Delete(&newproduct).Error; err != nil {
		return nil, status.Error(500, err.Error())
	}
	return &product.RemoveResponse{}, nil
}
