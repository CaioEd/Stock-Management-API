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

// CATEGORIES FUNCTIONS

func GetCategories(w http.ResponseWriter, r *http.Request) {
	var categories []models.Category
	if err := database.DB.Find(&categories).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}


func GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	var categories models.Category

	// GET USER BY ID
	if err := database.DB.First(&categories, id).Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			http.Error(w, "Category not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// RETURNS THE USER AS JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}


func CreateCategory(w http.ResponseWriter, r *http.Request) {
	bodyBytes, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewReader(bodyBytes)) // Reset body
	fmt.Println("Request received:", string(bodyBytes))

	var categories models.Category
	if err := json.NewDecoder(r.Body).Decode(&categories); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := database.DB.Create(&categories).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}


func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	// GET ID BY URL
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	// GET CATEGORY ON DB
	var categories models.Category
	if err := database.DB.First(&categories, id).Error; err != nil {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&categories); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// UPDATE CATEGORY
	if err := database.DB.Save(&categories).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// RETURNS CATEGORY UPDATED
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)

}