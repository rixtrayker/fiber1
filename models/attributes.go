package models

import "time"

type Price struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	Price     float64 `json:"price"`
}
