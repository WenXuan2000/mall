package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"mall/common/modeInit"
	"mall/service/order/model"
	"mall/service/order/rpc/internal/config"
	"mall/service/product/rpc/productclient"
	"mall/service/product/rpc/types/product"
	"mall/service/user/rpc/types/user"
	"mall/service/user/rpc/userclient"
)

type ServiceContext struct {
	Config     config.Config
	OrderModel *gorm.DB
	UserRpc    user.UserClient
	ProductRpc product.ProductClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := modeInit.InitGorm(c.Mysql.DataSource)
	db.AutoMigrate(&model.Order{})
	return &ServiceContext{
		Config:     c,
		OrderModel: db,
		UserRpc:    userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		ProductRpc: productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}
