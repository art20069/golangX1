package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Name  string `gorm:"type:varchar(200)"`
	Email string `gorm:"type:varchar(200)"`
	Tel   string `gorm:"type:varchar(200)"`
	// somchaiOrder.OrderItems => dataItems
	// OrderItems []OrderItem
	//---------------------------------------//
	// somchaiOrder.Products => dataItems
	Products []OrderItem `gorm:"foreignKey:OrderID"`
}
