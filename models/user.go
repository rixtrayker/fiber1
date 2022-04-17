package models

import "time"

type User struct {
	ID             uint `json:"id" gorm:"primaryKey"`
	CreatedAt      time.Time
	Name           string `json:"name"`
	Email          string `json:"email" gorm:"uniqueIndex"`
	Password       string `json:"password"`
	ActivationDate time.Time
}
