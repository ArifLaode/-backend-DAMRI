package logic

import (
	"encoding/json"
	"net/http"
)

type SortPaket struct {
	IDSort   int     `json:"id_sort"`
	IDTujuan int     `json:"id_tujuan"`
	Distance float64 `json:"distance"`
}

func CreateSortPaket(w http.ResponseWriter, r *http.Request) {
	var sortPaket SortPaket // Menggunakan struktur global SortPaket

	if err := json.NewDecoder(r.Body).Decode(&sortPaket); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "INSERT INTO sort_paket (id_tujuan, distance) VALUES (?, ?)"
	_, err := db.Exec(query, sortPaket.IDTujuan, sortPaket.Distance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "SortPaket created successfully"})
}

func GetSortPaket(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id_sort, id_tujuan, distance FROM sort_paket")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var sortPakets []SortPaket // Menggunakan slice dari struktur global SortPaket
	for rows.Next() {
		var sortPaket SortPaket
		if err := rows.Scan(&sortPaket.IDSort, &sortPaket.IDTujuan, &sortPaket.Distance); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sortPakets = append(sortPakets, sortPaket)
	}

	json.NewEncoder(w).Encode(sortPakets)
}

func UpdateSortPaket(w http.ResponseWriter, r *http.Request) {
	var sortPaket SortPaket // Menggunakan struktur global SortPaket

	if err := json.NewDecoder(r.Body).Decode(&sortPaket); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "UPDATE sort_paket SET id_tujuan = ?, distance = ? WHERE id_sort = ?"
	_, err := db.Exec(query, sortPaket.IDTujuan, sortPaket.Distance, sortPaket.IDSort)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "SortPaket updated successfully"})
}

func DeleteSortPaket(w http.ResponseWriter, r *http.Request) {
	var sortPaket SortPaket // Menggunakan struktur global SortPaket

	if err := json.NewDecoder(r.Body).Decode(&sortPaket); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "DELETE FROM sort_paket WHERE id_sort = ?"
	_, err := db.Exec(query, sortPaket.IDSort)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "SortPaket deleted successfully"})
}
