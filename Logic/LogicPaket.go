package logic

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// Struktur Paket sesuai dengan tabel database
type Paket struct {
	IDPaket      string  `json:"id_paket"`
	Penerima     string  `json:"penerima"`
	Pengirim     string  `json:"pengirim"`
	TelpPenerima string  `json:"telp_penerima"`
	TelpPengirim string  `json:"telp_pengirim"`
	NamaBarang   string  `json:"nama_barang"`
	Harga        float64 `json:"harga"`
	Berat        float64 `json:"berat"`
	Status       int     `json:"status"`
	IDTujuan     string  `json:"id_tujuan"`
}

// Fungsi untuk generate ID unik dengan format "DMR-XXXXXXXXXX" (10 digit angka acak)
func generateID() string {
	const prefix = "DMR-"
	return fmt.Sprintf("%s%010d", prefix, rand.Int63()%10000000000)
}

func init() {
	// Seed untuk memastikan angka acak berbeda setiap eksekusi
	rand.Seed(time.Now().UnixNano())
}

// Create Paket
func CreatePaket(w http.ResponseWriter, r *http.Request) {
	var paket Paket

	if err := json.NewDecoder(r.Body).Decode(&paket); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate IDPaket dan IDTujuan jika belum ada
	if paket.IDPaket == "" {
		paket.IDPaket = generateID()
	}
	if paket.IDTujuan == "" {
		paket.IDTujuan = generateID()
	}

	// Query INSERT untuk tabel paket
	query := "INSERT INTO paket (id_paket, penerima, pengirim, telp_penerima, telp_pengirim, nama_barang, harga, berat, status, id_tujuan) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := db.Exec(query, paket.IDPaket, paket.Penerima, paket.Pengirim, paket.TelpPenerima, paket.TelpPengirim, paket.NamaBarang, paket.Harga, paket.Berat, paket.Status, paket.IDTujuan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Paket created successfully"})
}

// Get Paket
func GetPaket(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id_paket, penerima, pengirim, telp_penerima, telp_pengirim, nama_barang, harga, berat, status, id_tujuan FROM paket")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var pakets []Paket
	for rows.Next() {
		var paket Paket
		if err := rows.Scan(&paket.IDPaket, &paket.Penerima, &paket.Pengirim, &paket.TelpPenerima, &paket.TelpPengirim, &paket.NamaBarang, &paket.Harga, &paket.Berat, &paket.Status, &paket.IDTujuan); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		pakets = append(pakets, paket)
	}

	json.NewEncoder(w).Encode(pakets)
}

// Update Paket
func UpdatePaket(w http.ResponseWriter, r *http.Request) {
	var paket Paket

	if err := json.NewDecoder(r.Body).Decode(&paket); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Pastikan semua kolom dalam query UPDATE
	query := "UPDATE paket SET penerima = ?, pengirim = ?, telp_penerima = ?, telp_pengirim = ?, nama_barang = ?, harga = ?, berat = ?, status = ?, id_tujuan = ? WHERE id_paket = ?"
	_, err := db.Exec(query, paket.Penerima, paket.Pengirim, paket.TelpPenerima, paket.TelpPengirim, paket.NamaBarang, paket.Harga, paket.Berat, paket.Status, paket.IDTujuan, paket.IDPaket)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Paket updated successfully"})
}

// Delete Paket
func DeletePaket(w http.ResponseWriter, r *http.Request) {
	var paket Paket

	if err := json.NewDecoder(r.Body).Decode(&paket); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "DELETE FROM paket WHERE id_paket = ?"
	_, err := db.Exec(query, paket.IDPaket)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Paket deleted successfully"})
}
