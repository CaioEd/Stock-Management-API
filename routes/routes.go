package routes

import (
	"stock_api/controllers"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/products", controllers.GetProducts).Methods("GET")
	r.HandleFunc("/products", controllers.CreateProduct).Methods("POST")
	// r.HandleFunc("/products/{id}", controllers.UpdateProduct).Methods("PUT")
	// r.HandleFunc("/products{id}", controllers.DeleteProduct).Methods("DELETE")

	return r
}
