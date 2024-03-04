package model

import (
	"fmt"
	jgorm "github.com/jinzhu/gorm"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func HaveOderByid(OrderModel *gorm.DB, Id int64, order *Order) (ok bool, err error) {
	if err := OrderModel.Where("id=?", Id).Take(&order).Error; jgorm.IsRecordNotFoundError(err) {
		return true, status.Error(100, "订单不存在")
	} else if err != nil {
		return true, status.Error(500, err.Error())
	}
	return false, nil
}

func (m *customOrderModel) FindAllByUid(ctx context.Context, uid int64) ([]*Order, error) {
	var resp []*Order

	query := fmt.Sprintf("select %s from %s where `uid` = ?", orderRows, m.table)
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

func FindAllByUid(OrderModel *gorm.DB, Id int64) ([]*Order, error) {
	var resp []*Order

}
