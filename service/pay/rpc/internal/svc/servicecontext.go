package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"mall/service/order/rpc/orderclient"
	"mall/service/order/rpc/types/order"
	"mall/service/pay/model"
	"mall/service/pay/rpc/internal/config"
	"mall/service/user/rpc/types/user"
	"mall/service/user/rpc/userclient"
)

type ServiceContext struct {
	Config   config.Config
	PayModel model.PayModel
	UserRpc  user.UserClient
	OrderRpc order.OrderClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:   c,
		PayModel: model.NewPayModel(conn, c.CacheRedis),
		UserRpc:  userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		OrderRpc: orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
	}
}
