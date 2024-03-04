package model

import (
	"gorm.io/gorm"
)

type GOrder struct {
	gorm.Model
	Uid    uint64 `gorm:"column:uid;type:bigint(20) unsigned;default:0;comment:用户ID;NOT NULL" json:"uid"`
	Pid    uint64 `gorm:"column:pid;type:bigint(20) unsigned;default:0;comment:产品ID;NOT NULL" json:"pid"`
	Amount uint   `gorm:"column:amount;type:int(10) unsigned;default:0;comment:订单金额;NOT NULL" json:"amount"`
	Status uint   `gorm:"column:status;type:tinyint(3) unsigned;default:0;comment:订单状态;NOT NULL" json:"status"`
}

func (m *Order) TableName() string {
	return "order"
}
