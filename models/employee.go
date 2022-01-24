package models

import (
	"gorm.io/gorm"
)

type GenderTypes string
type EmployeeTypes string
type StatusTypes string

const (
	Male       GenderTypes   = "M"
	Female     GenderTypes   = "F"
	Cashier    EmployeeTypes = "Cashier"
	Purchaser  EmployeeTypes = "Purchaser"
	Manager    EmployeeTypes = "Manager"
	Active     StatusTypes   = "Active"
	Inactive   StatusTypes   = "Inactive"
	Terminated StatusTypes   = "Terminated"
)

type Employee struct {
	gorm.Model
	FirstName    string        `form:"first name" json:"first_name" gorm:"type:varchar(150)"`
	LastName     string        `form:"last name" json:"last_name" gorm:"type:varchar(150)"`
	Gender       GenderTypes   `gorm:"type:Gender" form:"gender" json:"gender"`
	EmployeeType EmployeeTypes `gorm:"type:EmployeeType" form:"employee role" json:"employee_role"`
	StatusType   StatusTypes   `gorm:"type:StatusType" form:"status" json:"status"`
	UserID       int
	User         User
}
