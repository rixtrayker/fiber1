package models

import "time"

type Attribute struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	Images    []Image `gorm:"many2many:attribute_images"`
}
