package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"mall/service/order/api/internal/config"
	"mall/service/order/rpc/orderclient"
	"mall/service/order/rpc/types/order"
	"mall/service/product/rpc/productclient"
	"mall/service/product/rpc/types/product"
)

type ServiceContext struct {
	Config config.Config

	OrderRpc   order.OrderClient
	ProductRpc product.ProductClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		OrderRpc:   orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
		ProductRpc: productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}
