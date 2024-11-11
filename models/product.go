package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name			string `json:"name"`
	Description		string `json:"description"`
	Category		string `json:"category"`
	Quantity		int `json:"quantity"`
	Price			float64 `json:"price"`
}