package main

// import (
// 	"time"

// 	"gorm.io/driver/sqlite"
// 	"gorm.io/gorm"
// )

// type Category struct {
// 	gorm.Model
// 	Name          string        `form:"name" json:"name" gorm:"type:varchar(500)"`
// 	SubCategories []SubCategory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
// }
// type Customer struct {
// 	gorm.Model
// 	Name        string `form:"name" json:"name" gorm:"type:varchar(500)"`
// 	PhoneNumber string `form:"phone number" json:"phone_number" gorm:"type:varchar(50)"`
// }
// type GenderTypes string
// type EmployeeTypes string
// type StatusTypes string

// const (
// 	Male       GenderTypes   = "M"
// 	Female     GenderTypes   = "F"
// 	Cashier    EmployeeTypes = "Cashier"
// 	Purchaser  EmployeeTypes = "Purchaser"
// 	Manager    EmployeeTypes = "Manager"
// 	Active     StatusTypes   = "Active"
// 	Inactive   StatusTypes   = "Inactive"
// 	Terminated StatusTypes   = "Terminated"
// )

// type Employee struct {
// 	gorm.Model
// 	FirstName    string        `form:"first name" json:"first_name" gorm:"type:varchar(150)"`
// 	LastName     string        `form:"last name" json:"last_name" gorm:"type:varchar(150)"`
// 	Gender       GenderTypes   `form:"gender" json:"gender"`
// 	EmployeeType EmployeeTypes `form:"employee role" json:"employee_role"`
// 	StatusType   StatusTypes   `form:"status" json:"status"`
// 	UserID       int
// 	User         User
// }
// type Item struct {
// 	gorm.Model
// 	BarCode       string    `form:"barcode" json:"barcode" gorm:"type:varchar(50)"`
// 	ExpireDate    time.Time `form:"expire date" json:"expire_date"`
// 	SubCategoryID uint      `json:"sub_category_id"`
// 	Amount        uint      `json:"amount"`
// 	Price         float32   `json:"price"`
// 	//this if for storing which venders can supply this item
// 	Venders []Vender `gorm:"many2many:item_venders;"`
// }
// type ItemSale struct {
// 	gorm.Model
// 	ItemID     int
// 	Item       Item
// 	EmployeeID int
// 	Employee   Employee
// 	CustomerID int `gorm:"default:null"`
// 	Customer   Customer
// 	Amount     uint `form:"amount" json:"amount"`
// 	//PurchasedPrice is used because price might be updated later for each item but we want the price at
// 	//the time of sale
// 	PurchasedPrice float32 `json:"purchased_price"`
// }

// type SubCategory struct {
// 	gorm.Model
// 	Name       string `form:"name" json:"name" gorm:"type:varchar(500)"`
// 	CategoryID uint   `json:"category_id"`
// 	Items      []Item `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
// }

// type User struct {
// 	gorm.Model
// 	Username string `form:"username" json:"username" gorm:"type:varchar(500)"`
// 	Password string `form:"password" json:"password" gorm:"type:varchar(500)"`
// }
// type Vender struct {
// 	gorm.Model
// 	Name        string `form:"name" json:"name" gorm:"type:varchar(500)"`
// 	PhoneNumber string `form:"phone number" json:"phone_number" gorm:"type:varchar(50)"`
// }

// func main() {
// 	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}

// 	// Migrate the schema
// 	db.AutoMigrate(&Category{}, &Customer{}, &Employee{}, &Item{}, &ItemSale{}, &SubCategory{}, &User{}, &Vender{})

// }
