package model

//
//import (
//	"gorm.io/gorm"
//)
//
//type User struct {
//	gorm.Model
//	Name     string `gorm:"column:name;type:varchar(255);comment:用户姓名;NOT NULL" json:"name"`
//	Gender   int64  `gorm:"column:gender;type:tinyint(3) unsigned;default:0;comment:用户性别;NOT NULL" json:"gender"`
//	Mobile   string `gorm:"column:mobile;type:varchar(255);comment:用户电话;NOT NULL" json:"mobile"`
//	Password string `gorm:"column:password;type:varchar(255);comment:用户密码;NOT NULL" json:"password"`
//}
//
//func (m *User) TableName() string {
//	return "user"
//}
