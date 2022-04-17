package models

import "time"

type Order struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	Name      string    `json:"name"`
	Email     string    `json:"serial" gorm:"uniqueIndex"`
	Products  []Product `gorm:"many2many:order_products;"`
	Coupons  []Coupon
}
