package models

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Code        string
	Description string
	Quantity    int
	OrderID     uint
}
