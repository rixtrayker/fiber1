package models

type Image struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Path string `json:"path"`
}
