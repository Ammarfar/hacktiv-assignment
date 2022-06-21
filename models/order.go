package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerName string
	Items        []Item
	OrderedAt    time.Time
}
