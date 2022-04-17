package models

import "time"

type Price struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	Images    []Image
}
