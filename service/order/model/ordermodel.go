package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OrderModel = (*customOrderModel)(nil)

type (
	// OrderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrderModel.
	OrderModel interface {
		orderModel
		FindAllByUid(ctx context.Context, uid int64) ([]*Order, error)
	}

	customOrderModel struct {
		*defaultOrderModel
	}
)

// NewOrderModel returns a model for the database table.
func NewOrderModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) OrderModel {
	return &customOrderModel{
		defaultOrderModel: newOrderModel(conn, c, opts...),
	}
}
func (m *customOrderModel) FindAllByUid(ctx context.Context, uid int64) ([]*Order, error) {
	var resp []*Order

	query := fmt.Sprintf("select %s from %s where `uid` = ?", orderRows, m.table)
	// QueryRowsNoCacheCtx 的命名中的 "NoCache" 表示该函数不使用缓存机制，而是直接向数据库发出查询请求，并返回结果行数。这样的函数可能在某些场景下更适合，例如对实时性要求较高、数据频繁变动的情况，或者希望避免缓存带来的一致性问题等情况。
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, uid)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
