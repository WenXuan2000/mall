package logic

import (
	"context"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/status"
	"mall/service/product/model"

	"mall/service/product/rpc/internal/svc"
	"mall/service/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DetailLogic) Detail(in *product.DetailRequest) (*product.DetailResponse, error) {
	// todo: add your logic here and delete this line
	newproduct := &model.Product{}
	// Go 中的 errors.Is 函数会检查传给它的错误链中的任何错误是否与目标错误相匹配。然而，在 GORM 中，gorm.ErrRecordNotFound 错误经常被其他错误包裹。
	//如果 Take 不直接返回 gorm.ErrRecordNotFound，而是将其封装在另一个错误中，errors.Is 就可能无法正常工作。在这种情况下，您可以使用类型断言来检查特定的错误类型。下面是一个示例
	if err := l.svcCtx.ProductModel.Where("id=?", in.Id).Take(newproduct).Error; gorm.IsRecordNotFoundError(err) {
		return nil, status.Error(100, err.Error())
	} else if err != nil {
		return nil, status.Error(500, err.Error())
	}
	return &product.DetailResponse{
		Id:     int64(newproduct.ID),
		Name:   newproduct.Name,
		Desc:   newproduct.Name,
		Stock:  newproduct.Stock,
		Amount: newproduct.Amount,
		Status: newproduct.Status,
	}, nil
}
