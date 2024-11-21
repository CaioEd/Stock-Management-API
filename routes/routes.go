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
	r.HandleFunc("/product/{id}", controllers.DeleteProduct).Methods("DELETE")

	r.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	r.HandleFunc("/user/{id}", controllers.GetUserByID).Methods("GET")
	r.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")

	return r
}
