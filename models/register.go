package models

<<<<<<< HEAD
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
=======
type Register struct {
	ID				uint			`json:"id"`
	Name			string 			`json:"name"`
	Category		string 			`json:"category"`
	Quantity		string			`json:"quantity"`
	Price			string 			`json:"price"`
	Date			string			`json:"register_date"`
>>>>>>> 257318786fdcaf0255bea40fb0a0f81e4b3f94ce
}