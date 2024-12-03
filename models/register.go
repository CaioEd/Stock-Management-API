package models


import "gorm.io/gorm"

type Register struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Quantity    float64 `json:"quantity"`
	Price       float64 `json:"price"`
	Total_Spent float64 `json:"total_spent"`
	Date        string `json:"register_date"`
}

func (r *Register) BeforeSave(tx *gorm.DB) (err error) {
	r.Total_Spent = r.Quantity * r.Price
	return
}