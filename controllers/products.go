package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"stock_api/database"
	"stock_api/models"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// PRODUCTS FUNCTIONS

func GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	if err := database.DB.Find(&products).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}


func GetProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	var product models.Product

	// GET PRODUCT BY ID
	if err := database.DB.First(&product, id).Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			http.Error(w, "Product not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// RETURNS THE PRODUCT AS JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}


func CreateProduct(w http.ResponseWriter, r *http.Request) {
	bodyBytes, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewReader(bodyBytes)) // Reseta o Body para permitir nova leitura
	fmt.Println("Requisição recebida:", string(bodyBytes))

	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := database.DB.Create(&product).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}


func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// GET ID BY URL
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	// GET PRODUCT ON DB
	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// UPDATE PRODUCT
	if err := database.DB.Save(&product).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// RETURNS PRODUCT UPDATED
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)

}


func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// GET ID BY URL
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	// GET PRODUCT ON DB
	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err := database.DB.Delete(&product).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Product deleted"})
}