package models

import (
	"gorm.io/gorm"
)

//the point of this model is for storing items added and removed and saving which cashier is responsible
//for the sale(reduction in inventory) of an item and which purchaser is responsible for adding to stock
type ItemUpdate struct {
	gorm.Model
	ItemID     int
	Item       Item
	EmployeeID int
	Employee   Employee
	Amount     int `form:"amount Increase" json:"amount"`
}
