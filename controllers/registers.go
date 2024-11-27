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

// REGISTERS FUNCTIONS

func GetRegisters(w http.ResponseWriter, r *http.Request) {
	var registers []models.Register
	if err := database.DB.Find(&registers).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(registers)
}


func GetRegisterByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	var register models.Register

	// GET REGISTER BY ID
	if err := database.DB.First(&register, id).Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			http.Error(w, "Register not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// RETURNS THE REGISTER AS JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(register)
}


func CreateRegister(w http.ResponseWriter, r *http.Request) {
	bodyBytes, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewReader(bodyBytes)) // Reset Body
	fmt.Println("Request received:", string(bodyBytes))

	var register models.Register
	if err := json.NewDecoder(r.Body).Decode(&register); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := database.DB.Create(&register).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(register)
}


func UpdateRegister(w http.ResponseWriter, r *http.Request) {
	// GET ID BY URL
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	// GET REGISTER ON DB
	var register models.Register
	if err := database.DB.First(&register, id).Error; err != nil {
		http.Error(w, "Register not found", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&register); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// UPDATE REGISTER
	if err := database.DB.Save(&register).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// RETURNS REGISTER UPDATED
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(register)

}
