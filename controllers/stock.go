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
		http.Error(w, "Both 'from_date' and 'to_date' are required", http.StatusBadRequest)
		return
	}

	startDate, err := time.Parse("02-01-2006", fromDate)
	if err != nil {
		http.Error(w, "Invalid 'from_date' format. Use DD-MM-YYYY", http.StatusBadRequest)
		return
	}

	endDate, err := time.Parse("02-01-2006", toDate)
	if err != nil {
		http.Error(w, "Invalid 'to_date' format. Use DD-MM-YYYY", http.StatusBadRequest)
		return
	}

	var registers []models.Register
	if err := database.DB.Where("date BETWEEN ? AND ?", startDate, endDate).Find(&registers).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(registers); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GET TOTAL QUANTITY OF PRODUCTS ON THE STOCK
func GetTotalQuantity(w http.ResponseWriter, r *http.Request) {
	var total int64

	if err := database.DB.Table("registers").Select("SUM(quantity)").Scan(&total).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int64{"total_quantity": total})
}

// GET THE SUM OF THE TOTAL SPENT ON PRODUCTS IN THE CURRENT MONTH
func GetTotalSpent(w http.ResponseWriter, r *http.Request) {
	var total float64

	now := time.Now()

	firstDayMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	lastDayMonth := firstDayMonth.AddDate(0, 1, -1).Add(time.Hour*23 + time.Minute*59 + time.Second*59)

	if err := database.DB.Table("registers").
		Where("created_at BETWEEN ? AND ?", firstDayMonth.Format("2006-01-02"), lastDayMonth.Format("2006-01-02")).
		Select("SUM(total_spent)").Scan(&total).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]float64{"total_spent": total})
}
