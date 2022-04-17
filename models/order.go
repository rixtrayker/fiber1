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
	Payment   Payment
	UserId    int  `json:"user_id"`
	User      User `gorm:"foreignKey:UserID"`
	Coupons   []Coupon
}
