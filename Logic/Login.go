package logic

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

// AuthenticatedUser adalah struktur data untuk menyimpan informasi pengguna yang berhasil login
type AuthenticatedUser struct {
	ID       string `json:"id_admin"`
	Role     int    `json:"role"`
	Username string `json:"username"`
}

// LoginUser menangani proses login dan memeriksa kredensial pengguna
func LoginUser(w http.ResponseWriter, r *http.Request) {
	godotenv.Load()
	var credentials map[string]string
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	username, password := credentials["username"], credentials["password"]

	// Verifikasi kredensial
	var idAdmin string
	var hashedPassword string // Ambil hashed password dari database
	var role int
	err := db.QueryRow("SELECT id_admin, password, role FROM admin WHERE nama =?", username).Scan(&idAdmin, &hashedPassword, &role) // Ambil password (hash)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Bandingkan password dengan hash yang benar
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		http.Error(w, "Invalid credentials when compare", http.StatusUnauthorized)
		return
	}

	claims := jwt.MapClaims{
		"id_admin": idAdmin,
		"role":     role,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("Tidak ada env yang ditemukan")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		http.Error(w, "Failed Generate String", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"token":    tokenString, // Tambahkan token di sini!
		"id_admin": idAdmin,
		"role":     role,
		"username": username,
	}

	// Menyusun response untuk pengguna yang berhasil login
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
