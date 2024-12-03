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

	r.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	r.HandleFunc("/user/{id}", controllers.GetUserByID).Methods("GET")
	r.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")

	r.HandleFunc("/registers", controllers.GetRegisters).Methods("GET")
	r.HandleFunc("/register/{id}", controllers.GetRegisterByID).Methods("GET")
	r.HandleFunc("/register", controllers.CreateRegister).Methods("POST")
	r.HandleFunc("/register/{id}", controllers.UpdateRegister).Methods("PUT")	
	r.HandleFunc("/register/{id}", controllers.DeleteRegister).Methods("DELETE")
	r.HandleFunc("/registers/date", controllers.GetRegistersByDate).Methods("GET")

	r.HandleFunc("/categories", controllers.GetCategories).Methods("GET")
	r.HandleFunc("/category/{id}", controllers.GetCategoryByID).Methods("GET")
	r.HandleFunc("/category", controllers.CreateCategory).Methods("POST")
	r.HandleFunc("/category/{id}", controllers.UpdateCategory).Methods("PUT")	

	r.HandleFunc("/total_quantity_products", controllers.GetTotalQuantity).Methods("GET")

	r.HandleFunc("/total_spent_current_month", controllers.GetTotalSpent).Methods("GET")


	return r
}
