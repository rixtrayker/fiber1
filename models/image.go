package models

type Image struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	ImagePath string `json:"image_path"`
}
