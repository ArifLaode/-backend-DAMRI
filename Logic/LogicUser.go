package logic

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

type Admin struct {
	IDAdmin  string `json:"id_admin"`
	Nama     string `json:"nama"`
	Password string `json:"password"`
	Role     int    `json:"role"`
}

func SetDB(database *sql.DB) {
	db = database
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	var admin Admin
	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var hashedPassword, err = bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	admin.Password = string(hashedPassword)

	query := ("INSERT INTO admin (id_admin, nama, password, role) VALUE (?,?,?,?)")
	_, err = db.Exec(query, admin.IDAdmin, admin.Nama, admin.Password, admin.Role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Admin telah ditambahkan"})
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id_admin, nama, role FROM admin") // Jangan ambil password
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var admins []Admin
	for rows.Next() {
		var admin Admin
		if err := rows.Scan(&admin.IDAdmin, &admin.Nama, &admin.Role); err != nil { // Scan tanpa password
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		admins = append(admins, admin)
	}

	json.NewEncoder(w).Encode(admins)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var admin Admin // Gunakan struct Admin

	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Hash password sebelum update
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	admin.Password = string(hashedPassword)

	query := "UPDATE admin SET nama =?, password =?, role =? WHERE id_admin =?"
	_, err = db.Exec(query, admin.Nama, admin.Password, admin.Role, admin.IDAdmin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Admin updated successfully"})
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var admin Admin // Gunakan struct Admin

	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "DELETE FROM admin WHERE id_admin =?"
	_, err := db.Exec(query, admin.IDAdmin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Admin deleted successfully"})
}
