package models

import (
	"gorm.io/gorm"
)

type Vender struct {
	gorm.Model
	Name        string `form:"name" json:"name" gorm:"type:varchar(500)"`
	PhoneNumber string `form:"phone number" json:"phone_number" gorm:"type:varchar(50)"`
}
