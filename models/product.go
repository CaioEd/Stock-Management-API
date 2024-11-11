package models


type Product struct {
	ID				uint			`json:"id"`
	Name			string 			`json:"name"`
	Description		string 			`json:"description"`
	Category		string 			`json:"category"`
	Quantity		int 			`json:"quantity"`
	Price			float64 		`json:"price"`
}
