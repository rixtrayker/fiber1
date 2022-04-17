package models

type Product struct {
	ID           uint    `json:"id" gorm:"primaryKey"`
	Name         string  `json:"name"`
	SerialNumber string  `json:"serial_number"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	Images       []Image `json:"images" gorm:"many2many:product_images"`
}
