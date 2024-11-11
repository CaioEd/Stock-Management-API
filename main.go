package main

import (
	"log"
	"net/http"
	"stock_api/database"
	"stock_api/routes"
)


func main() {
	// INICIALIZA O BANCO DE DADOS
	database.Init()
	defer database.Close()

	// ROUTES CONFIGURATION
	r := routes.NewRouter()

	log.Println("API rodando na porta 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Error initializing server: ", err)
	}
}
