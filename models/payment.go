package models

import "time"

type Payment struct {
	ID           uint `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time
	SerialNumber string  `json:"serial_number"`
	Cost         float64 `json:"cost"`
}
