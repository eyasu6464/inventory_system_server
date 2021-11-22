package models

import (
	"gorm.io/gorm"
)

//the point of this model is for storing items sold
//these include cashier(employee), buyer(customer), item
type ItemPurchased struct {
	gorm.Model
	ItemID     int
	Item       Item
	EmployeeID int
	Employee   Employee
	VenderID   int `gorm:"default:null"`
	Vender     Vender
	Amount     uint `form:"amount" json:"amount"`
	//PurchasedPrice is used to store the price at which the item is acquired from vender
	//for calculating profit if we want
	PurchasedPrice float32 `json:"purchased_price"`
}
