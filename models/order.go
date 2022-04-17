package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	Name      string    `json:"name"`
	Email     string    `json:"serial" gorm:"uniqueIndex"`
	Products  []Product `gorm:"many2many:order_products"`
	PaymentId int       `json:"payment_id"`
	Payment   Payment   `gorm:"foreignKey:PaymentId"`
	UserId    int       `json:"user_id"`
	User      User      `gorm:"foreignKey:UserId"`
	Coupons   []Coupon  `gorm:"many2many:order_coupons"`
}
