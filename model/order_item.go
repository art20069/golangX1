package model

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	SKU      string  `gorm:"type:varchar(200); not null"`
	Name     string  `gorm:"type:varchar(200); not null"`
	Image    string  `gorm:"type:varchar(200); not null"`
	Price    float64 `gorm:"type:varchar(200); not null"`
	Quantity uint    `gorm:"type:varchar(200); not null"`
	OrderID  uint
}
