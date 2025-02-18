package router

import (
	logic "damri/Logic"
	"net/http"
)

// EnableCors middleware untuk mengaktifkan CORS
func EnableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// SetupRoutes mendefinisikan semua route yang tersedia
func SetupRoutes() {
	mux := http.NewServeMux()

	mux.HandleFunc("/login", logic.LoginUser)
	mux.HandleFunc("/admin/read", logic.RoleMiddleware(logic.GetUser, 1, 2))
	mux.HandleFunc("/admin/add", logic.RoleMiddleware(logic.AddUser, 1))
	mux.HandleFunc("/admin/delete", logic.RoleMiddleware(logic.DeleteUser, 1))
	mux.HandleFunc("/admin/update", logic.RoleMiddleware(logic.UpdateUser, 1, 2))
	mux.HandleFunc("/tujuan/add", logic.RoleMiddleware(logic.CreateTujuan, 1))
	mux.HandleFunc("/tujuan/read", logic.RoleMiddleware(logic.GetTujuan, 0, 1, 2, 3))
	mux.HandleFunc("/tujuan/update", logic.RoleMiddleware(logic.UpdateTujuan, 1, 2))
	mux.HandleFunc("/tujuan/delete", logic.RoleMiddleware(logic.DeleteTujuan, 1))
	mux.HandleFunc("/paket/add", logic.RoleMiddleware(logic.CreatePaket, 1))
	mux.HandleFunc("/paket/read", logic.RoleMiddleware(logic.GetPaket, 0, 1, 2, 3))
	mux.HandleFunc("/paket/update", logic.RoleMiddleware(logic.UpdatePaket, 0, 1, 2, 3))
	mux.HandleFunc("/paket/delete", logic.RoleMiddleware(logic.DeletePaket, 1))
	mux.HandleFunc("/sort_paket/create", logic.RoleMiddleware(logic.CreateSortPaket, 1))
	mux.HandleFunc("/sort_paket/read", logic.RoleMiddleware(logic.GetSortPaket, 1, 2))
	mux.HandleFunc("/sort_paket/update", logic.RoleMiddleware(logic.UpdateSortPaket, 1, 2))
	mux.HandleFunc("/sort_paket/delete", logic.RoleMiddleware(logic.DeleteSortPaket, 1))

	// Wrap the mux with EnableCors middleware
	http.Handle("/", EnableCors(mux))
}
