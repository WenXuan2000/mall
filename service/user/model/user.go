package model

import (
	"gorm.io/gorm"
	"math/big"
)

type UserModer struct {
	gorm.Model
	Id big.Int `gorm:""`
}
