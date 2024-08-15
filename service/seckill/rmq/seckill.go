package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"

	"mall/service/seckill/rmq/internal/service"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/conf"
	"mall/service/seckill/rmq/internal/config"
)

var configFile = flag.String("f", "etc/seckill.yaml", "the etc file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	logx.SetLevel(logx.ErrorLevel)
	srv := service.NewService(c)
	//fmt.Println(c.Kafka)
	queue := kq.MustNewQueue(c.Kafka, kq.WithHandle(srv.Consume))
	defer queue.Stop()

	fmt.Println("seckill started!!!")
	queue.Start()
}
