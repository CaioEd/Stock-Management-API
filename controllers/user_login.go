package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"stock_api/database"
	"stock_api/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

var JwtSecret []byte

func init() {
	// Load env variables
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, falling back to system environment variables")
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET not defined in environment variables")
	}

	JwtSecret = []byte(secret)
}

// User authentication and returns JWT token
func Login(w http.ResponseWriter, r *http.Request) {
	var userLogin struct {
		Mobile   string `json:"mobile"`
		Password string `json:"password"`
	}

	// Decodifica os dados da requisição
	err := json.NewDecoder(r.Body).Decode(&userLogin)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if userLogin.Mobile == "" || userLogin.Password == "" {
		http.Error(w, "Mobile and password are required", http.StatusBadRequest)
		return
	}

	var user models.User
	// Verify if the users exists on database
	err = database.DB.Where("mobile = ?", userLogin.Mobile).First(&user).Error
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Verify is password is correct
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password))
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Create token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":    user.ID,
		"mobile": user.Mobile,
		"role":   user.Role,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(JwtSecret)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"token":   tokenString,
		"expires": time.Now().Add(time.Hour * 24).Format(time.RFC3339),
		"user": map[string]interface{}{
			"id":     user.ID,
			"name":   user.Name,
			"mobile": user.Mobile,
			"role":   user.Role,
		},
	})
}
