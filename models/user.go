package models

type User struct {
	ID				uint			`json:"id"`
	Name			string 			`json:"name"`
	Mobile			string 			`json:"mobile"`
	Password		string 			`json:"password"`
	Role			string			`json:"role"`
}