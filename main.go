package main

import (
	"log"
	"net/http"
	"stock_api/database"
	"stock_api/routes"

	"github.com/rs/cors"
)


func main() {
	// INICIALIZA O BANCO DE DADOS
	database.Init()
	defer database.Close()

	// ROUTES CONFIGURATION
	r := routes.NewRouter()

	// CORS CONFIGURATION
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Permite requisições apenas do seu front-end
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	log.Println("API rodando na porta 8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal("Error initializing server: ", err)
	}
}
