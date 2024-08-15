package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/mr"
	"mall/service/product/model"
	"strconv"
	"strings"

	"mall/service/product/rpc/internal/svc"
	"mall/service/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailListLogic {
	return &DetailListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DetailListLogic) DetailList(in *product.DetailListRequest) (*product.DetailListResponse, error) {
	// todo: add your logic here and delete this line
	reproduct := make(map[int64]*product.ProductItem)
	pdis := strings.Split(in.Id, ",")
	pidsint := []int64{}
	for _, i := range pdis {
		temp, _ := strconv.Atoi(i)
		pidsint = append(pidsint, int64(temp))
	}
	productlist, err := mr.MapReduce[int64, *model.Product, []*model.Product](func(source chan<- int64) {
		for _, pid := range pidsint {
			source <- pid
		}
	}, func(id int64, writer mr.Writer[*model.Product], cancel func(error)) {
		p, err := l.svcCtx.ProductModel.FindOne(l.ctx, id)
		if err != nil {
			cancel(err)
			return
		}
		writer.Write(p)
	}, func(pipe <-chan *model.Product, writer mr.Writer[[]*model.Product], cancel func(error)) {
		var products []*model.Product
		for pro := range pipe {
			products = append(products, pro)
		}
		writer.Write(products)
	})
	if err != nil {
		return nil, err
	}
	for _, i := range productlist {
		temp := &product.ProductItem{
			Id:     i.Id,
			Name:   i.Name,
			Desc:   i.Desc,
			Stock:  i.Stock,
			Amount: i.Amount,
			Status: i.Status,
		}
		reproduct[int64(i.Id)] = temp
	}
	//fmt.Println(reproduct)
	return &product.DetailListResponse{
		Products: reproduct,
	}, nil
}

//func (l *ArticlesLogic) articleByIds(ctx context.Context, articleIds []int64) ([]*model.Article, error) {
//	articles, err := mr.MapReduce[int64, *model.Article, []*model.Article](func(source chan<- int64) {
//		for _, aid := range articleIds {
//			source <- aid
//		}
//	}, func(id int64, writer mr.Writer[*model.Article], cancel func(error)) {
//		p, err := l.svcCtx.ArticleModel.FindOne(ctx, id)
//		if err != nil {
//			cancel(err)
//			return
//		}
//		writer.Write(p)
//	}, func(pipe <-chan *model.Article, writer mr.Writer[[]*model.Article], cancel func(error)) {
//		var articles []*model.Article
//		for article := range pipe {
//			articles = append(articles, article)
//		}
//		writer.Write(articles)
//	})
//	if err != nil {
//		return nil, err
//	}
//
//	return articles, nil
//}
