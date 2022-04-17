package models

import "time"

type Payment struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	Serial    string  `json:"serial"`
	Cost      float64 `json:"cost"`
}
