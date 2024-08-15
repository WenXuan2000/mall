package logic

import (
	"context"
	"encoding/json"
	"github.com/eapache/go-resiliency/batcher"
	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/limit"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mall/service/product/rpc/types/product"
	"strconv"
	"time"

	"mall/service/seckill/rpc/internal/svc"
	"mall/service/seckill/rpc/types/seckill"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	limitPeriod       = 10
	limitQuota        = 1
	seckillUserPrefix = "seckill#u#"
	localCacheExpire  = time.Second * 60

	batcherSize     = 100
	batcherBuffer   = 100
	batcherWorker   = 10
	batcherInterval = time.Second
)

type SeckillOrderLogic struct {
	ctx        context.Context
	svcCtx     *svc.ServiceContext
	limiter    *limit.PeriodLimit
	localCache *collection.Cache
	batcher    *batcher.Batcher
	logx.Logger
}
type KafkaData struct {
	Uid int64 `json:"uid"`
	Pid int64 `json:"pid"`
}

func NewSeckillOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SeckillOrderLogic {
	localCache, err := collection.NewCache(localCacheExpire)
	if err != nil {
		panic(err)
	}
	return &SeckillOrderLogic{
		ctx:        ctx,
		svcCtx:     svcCtx,
		Logger:     logx.WithContext(ctx),
		localCache: localCache,
		limiter:    limit.NewPeriodLimit(limitPeriod, limitQuota, svcCtx.BizRedis, seckillUserPrefix),
	}
}

func (l *SeckillOrderLogic) SeckillOrder(in *seckill.SeckillOrderRequest) (*seckill.SeckillOrderResponse, error) {
	// todo: add your logic here and delete this line
	code, _ := l.limiter.Take(strconv.FormatInt(in.UserId, 10))
	if code == limit.OverQuota {
		return nil, status.Errorf(codes.OutOfRange, "Number of requests exceeded the limit")
	}
	p, err := l.svcCtx.ProductRPC.Detail(l.ctx, &product.DetailRequest{Id: in.ProductId})
	if err != nil {
		return nil, err
	}
	if p.Stock <= 0 {
		return nil, status.Errorf(codes.OutOfRange, "Insufficient stock")
	}
	kd, err := json.Marshal(&KafkaData{Uid: in.UserId, Pid: in.ProductId})
	if err != nil {
		return nil, err
	}
	if err := l.svcCtx.KafkaPusher.Push(string(kd)); err != nil {
		return nil, err
	}
	return &seckill.SeckillOrderResponse{}, nil
}
