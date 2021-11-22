package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `form:"username" json:"username" gorm:"type:varchar(500)"`
	Password string `form:"password" json:"password" gorm:"type:varchar(500)"`
}
