package svc

import (
	"gorm.io/gorm"
	"mall/common/modeInit"
	"mall/service/user/model"
	"mall/service/user/rpc/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	UserModel *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := modeInit.InitGorm(c.Mysql.DataSource)
	db.AutoMigrate(&model.User{})
	return &ServiceContext{
		Config:    c,
		UserModel: db,
	}
}
