package models

type Register struct {
	ID				uint			`json:"id"`
	Name			string 			`json:"name"`
	Category		string 			`json:"category"`
	Quantity		string			`json:"quantity"`
	Price			string 			`json:"price"`
	Date			string			`json:"register_date"`
}