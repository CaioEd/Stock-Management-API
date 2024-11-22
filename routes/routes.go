package routes

import (
	"stock_api/controllers"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/products", controllers.GetProducts).Methods("GET")
	r.HandleFunc("/product/{id}", controllers.GetProductByID).Methods("GET")
	r.HandleFunc("/products", controllers.CreateProduct).Methods("POST")
	r.HandleFunc("/products/{id}", controllers.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", controllers.DeleteProduct).Methods("DELETE")

	r.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	r.HandleFunc("/user/{id}", controllers.GetUserByID).Methods("GET")
	r.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")

	r.HandleFunc("/registers", controllers.GetRegisters).Methods("GET")
	r.HandleFunc("/register/{id}", controllers.GetRegisterByID).Methods("GET")
	r.HandleFunc("/register", controllers.CreateRegister).Methods("POST")
	r.HandleFunc("/register/{id}", controllers.UpdateRegister).Methods("PUT")	

	r.HandleFunc("/registers/date", controllers.GetRegistersByDate).Methods("GET")

	return r
}
