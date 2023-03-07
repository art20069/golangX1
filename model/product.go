package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	SKU        string  `gorm:"uniqueIndex;type:varchar(100);not null"`
	Name       string  `gorm:"type:varchar(100);not null"`
	Desc       string  `gorm:"type:varchar(255)"`
	Price      float64 `gorm:"not null"`
	Status     uint    `gorm:"not null"` // 1 มีสินค้า  2 ไม่มีสินค้า
	Image      string  `gorm:"not null,type:varchar(255)"`
	CategoryID uint    `gorm:"not null"`
	// Category <---> Product (tulip)
	// ex. --> tulip.Category => Category(Name: Flower) ดังนั้นเรากำหนดผูก Category[product] Category[category] ด้านล่าง
	Category Category
}
