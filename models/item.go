package models

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	BarCode       string    `form:"barcode" json:"barcode" gorm:"type:varchar(50)"`
	ExpireDate    time.Time `form:"expire date" json:"expire_date"`
	SubCategoryID uint      `json:"sub_category_id"`
	Amount        uint      `json:"amount"`
	Price         float32   `json:"price"`
}
