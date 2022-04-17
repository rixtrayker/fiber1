package models

import "time"

type Coupon struct {
	ID             uint `json:"id" gorm:"primaryKey"`
	CreatedAt      time.Time
	ExpirationDate time.Time `json:"expiration_date"`
	Price          float64   `json:"price"`
	Text           string    `json:"text"`
	Description    string    `json:"description"`
}
