package controllers

import (
	"encoding/json"
	"net/http"
	"stock_api/database"
	"stock_api/models"
<<<<<<< HEAD
	"time"
)

=======
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// PRODUCTS FUNCTIONS
>>>>>>> 257318786fdcaf0255bea40fb0a0f81e4b3f94ce

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

	lastDayMonth := time.Date(now.Year(), now.Month(), 1, 23, 59, 59, 999999999, time.UTC).AddDate(0, 1, -1)

	if err := database.DB.Table("registers").
		Where("created_at BETWEEN ? AND ?", firstDayMonth.Format("2006-01-02"), lastDayMonth.Format("2006-01-02")).
		Select("SUM(total_spent)").Scan(&total).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]float64{"total_spent": total})
}
<<<<<<< HEAD
=======


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

// USERS FUNCTIONS

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}


func GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	var user models.User

	// GET USER BY ID
	if err := database.DB.First(&user, id).Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// RETURNS THE USER AS JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}


func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyBytes, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewReader(bodyBytes)) // Reseta o Body para permitir nova leitura
	fmt.Println("Request received:", string(bodyBytes))

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := database.DB.Create(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}


func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// GET ID BY URL
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	// GET USER ON DB
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// UPDATE USER
	if err := database.DB.Save(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// RETURNS USER UPDATED
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}

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

	// GET USER BY ID
	if err := database.DB.First(&register, id).Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// RETURNS THE USER AS JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(register)
}


func CreateRegister(w http.ResponseWriter, r *http.Request) {
	bodyBytes, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewReader(bodyBytes)) // Reseta o Body para permitir nova leitura
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

	// GET USER ON DB
	var register models.Register
	if err := database.DB.First(&register, id).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&register); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// UPDATE USER
	if err := database.DB.Save(&register).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// RETURNS USER UPDATED
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(register)

}
>>>>>>> 257318786fdcaf0255bea40fb0a0f81e4b3f94ce
