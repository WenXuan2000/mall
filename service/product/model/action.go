package model

import (
	"errors"
	"gorm.io/gorm"
)

func HavaProductByName(db *gorm.DB, name string) bool {
	var productexit = &Product{}
	if err := db.Where("name=?", name).First(&productexit).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
func HavaProduct(db *gorm.DB, id int64) bool {
	var productexit = Product{}
	if err := db.Where("id=?", id).First(&productexit).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
