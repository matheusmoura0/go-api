package model

type Product struct {
	ID    uint    `json:"id_product" gorm:"primaryKey"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
