package svc

import (
	"gorm.io/gorm"
	"mall/common/modeInit"
	"mall/service/product/model"
	"mall/service/product/rpc/internal/config"
)

type ServiceContext struct {
	Config       config.Config
	ProductModel *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := modeInit.InitGorm(c.Mysql.DataSource)
	db.AutoMigrate(&model.Product{})
	return &ServiceContext{
		Config:       c,
		ProductModel: db,
	}
}
