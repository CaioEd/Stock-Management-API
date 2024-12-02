package models

type Product struct {
	ID				uint			`json:"id"`
	Name			string 			`json:"name"`
	Description		string 			`json:"description"`
	Category		string 			`json:"category"`
}
