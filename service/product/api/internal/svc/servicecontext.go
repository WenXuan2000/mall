package svc

import (
	cos "github.com/tencentyun/cos-go-sdk-v5"
	"github.com/zeromicro/go-zero/zrpc"
	"mall/service/product/api/internal/config"
	"mall/service/product/rpc/productclient"
	"net/http"
	"net/url"
)

type ServiceContext struct {
	Config     config.Config
	ProductRpc productclient.Product
	Cosclient  *cos.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	u, _ := url.Parse("https://cover-1302313797.cos.ap-nanjing.myqcloud.com")
	// 用于 Get Service 查询，默认全地域 service.cos.myqcloud.com
	su, _ := url.Parse("https://cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, ServiceURL: su}
	// 1.永久密钥
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  c.TenCOS.SECRETID,  // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
			SecretKey: c.TenCOS.SECRETKEY, // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
		},
	})

	return &ServiceContext{
		Config:     c,
		ProductRpc: productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
		Cosclient:  client,
	}
}
