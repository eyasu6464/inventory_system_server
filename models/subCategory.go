package models

import (
	"gorm.io/gorm"
)

type SubCategory struct {
	gorm.Model
	Name       string `form:"name" json:"name" gorm:"type:varchar(500)"`
	CategoryID uint   `json:"category_id"`
	Items      []Item `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
