package models

import (
	"gorm.io/gorm"
)

//the point of this model is for storing items sold and saving which cashier is responsible
//for the sale(reduction in inventory) of an item and which customer took it
//these include cashier(employee), buyer(customer), item
type ItemSale struct {
	gorm.Model
	ItemID     int
	Item       Item
	EmployeeID int
	Employee   Employee
	//null if its walk in customer
	CustomerID int `gorm:"default:null"`
	Customer   Customer
	Amount     uint `form:"amount" json:"amount"`
	//Price is used because price might be updated later for each item but we want the price at
	//the time of sale
	Price float32 `json:"price"`
}
