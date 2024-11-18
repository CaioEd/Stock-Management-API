package models


type Product struct {
	ID				uint			`json:"id"`
	Name			string 			`json:"name"`
	Description		string 			`json:"description"`
	Category		string 			`json:"category"`
	Quantity		string			`json:"quantity"`
	Price			string 			`json:"price"`
}
