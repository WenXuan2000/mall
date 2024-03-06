package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"mall/service/user/api/internal/config"
	"mall/service/user/rpc/userclient"
)

type ServiceContext struct {
	Config  config.Config
	Rds     *redis.Redis
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		Rds:     redis.MustNewRedis(c.Redis),
	}
}
