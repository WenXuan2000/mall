package main

import (
	"flag"
	"fmt"
	"mall/service/product/api/internal/config"
	"mall/service/product/api/internal/handler"
	"mall/service/product/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/product.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	//opt := &cos.BucketGetOptions{
	//	Prefix:  "IMG",
	//	MaxKeys: 3,
	//}
	//
	//v, _, err := ctx.Cosclient.Bucket.Get(context.Background(), opt)
	//if err != nil {
	//	panic(err)
	//}
	//
	//for _, cc := range v.Contents {
	//	fmt.Printf("%s, %d\n", cc.Key, cc.Size)
	//}
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
