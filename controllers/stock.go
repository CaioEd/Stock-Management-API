package controllers

import (
	"encoding/json"
	"net/http"
	"stock_api/database"
	"stock_api/models"
	"time"
)

// FILTER BY DATE FUNCTION
func GetRegistersByDate(w http.ResponseWriter, r *http.Request) {
	fromDate := r.URL.Query().Get("from_date")
	toDate := r.URL.Query().Get("to_date")

	if fromDate == "" || toDate == "" {
		http.Error(w, "Dates are required", http.StatusBadRequest)
		return
	}

	startDate, err := time.Parse("02-01-2006", fromDate)
	if err != nil {
		http.Error(w, "Error", http.StatusBadRequest)
		return
	}

	endDate, err := time.Parse("02-01-2006", toDate)
	if err != nil {
		http.Error(w, "Error", http.StatusBadRequest)
		return
	}

	var registers []models.Register
	if err := database.DB.Where("date BETWEEN ? AND ?", startDate, endDate).Find(&registers).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(registers); err != nil {
		http.Error(w, err.Error(), http	.StatusInternalServerError)
	}
}
