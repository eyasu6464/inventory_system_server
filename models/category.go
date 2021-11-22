package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name          string        `form:"name" json:"name" gorm:"type:varchar(500)"`
	SubCategories []SubCategory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
