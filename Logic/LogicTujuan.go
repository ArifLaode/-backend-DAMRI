package logic

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type Tujuan struct {
	IDTujuan      string  `json:"id_tujuan"`
	Nama          string  `json:"nama"`
	SekitarLokasi string  `json:"sekitar"`
	Jarak         float64 `json:"jarak"`
	Harga         float64 `json:"harga"`
	Koordinat     string  `json:"koordinat"`
}

func CreateTujuan(w http.ResponseWriter, r *http.Request) {
	var tujuan Tujuan

	if err := json.NewDecoder(r.Body).Decode(&tujuan); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tujuan.IDTujuan = uuid.New().String()

	query := "INSERT INTO tujuan (id_tujuan, nama, sekitar, jarak, harga, koordinat) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := db.Exec(query, tujuan.IDTujuan, tujuan.Nama, tujuan.SekitarLokasi, tujuan.Jarak, tujuan.Harga, tujuan.Koordinat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Tujuan created successfully"})
}

func GetTujuan(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id_tujuan, nama, sekitar, jarak, harga, koordinat FROM tujuan")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tujuans []Tujuan
	for rows.Next() {
		var tujuan Tujuan
		if err := rows.Scan(&tujuan.IDTujuan, &tujuan.Nama, &tujuan.SekitarLokasi, &tujuan.Jarak, &tujuan.Harga, &tujuan.Koordinat); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tujuans = append(tujuans, tujuan)
	}

	json.NewEncoder(w).Encode(tujuans)
}

func UpdateTujuan(w http.ResponseWriter, r *http.Request) {
	var tujuan Tujuan

	if err := json.NewDecoder(r.Body).Decode(&tujuan); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "UPDATE tujuan SET nama = ?, sekitar = ?, harga = ?, jarak = ?, koordinat = ? WHERE id_tujuan = ?"
	_, err := db.Exec(query, tujuan.Nama, tujuan.SekitarLokasi, tujuan.Harga, tujuan.Jarak, tujuan.Koordinat, tujuan.IDTujuan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Tujuan updated successfully"})
}

func DeleteTujuan(w http.ResponseWriter, r *http.Request) {
	var tujuan Tujuan

	if err := json.NewDecoder(r.Body).Decode(&tujuan); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "DELETE FROM tujuan WHERE id_tujuan = ?"
	_, err := db.Exec(query, tujuan.IDTujuan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Tujuan deleted successfully"})
}
