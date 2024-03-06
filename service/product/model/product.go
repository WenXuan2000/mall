package model

//
//import "gorm.io/gorm"
//
//type Product struct {
//	gorm.Model
//	Name   string `gorm:"column:name;type:varchar(255);comment:产品名称;NOT NULL" json:"name"`
//	Desc   string `gorm:"column:desc;type:varchar(255);comment:产品描述;NOT NULL" json:"desc"`
//	Stock  int64  `gorm:"column:stock;type:int(10) unsigned;default:0;comment:产品库存;NOT NULL" json:"stock"`
//	Amount int64  `gorm:"column:amount;type:int(10) unsigned;default:0;comment:产品金额;NOT NULL" json:"amount"`
//	Status int64  `gorm:"column:status;type:tinyint(3) unsigned;default:0;comment:产品状态;NOT NULL" json:"status"`
//}
//
//func (m *Product) TableName() string {
//	return "product"
//}
