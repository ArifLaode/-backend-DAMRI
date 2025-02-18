package logic

import (
	"net/http"
	"strings"
)

func RoleMiddleware(next http.HandlerFunc, requiredRoles ...int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userRole := getUserRole(r) // Fungsi untuk mengambil role pengguna

		if userRole == 0 {
			// Guest access only allowed for /paket/read and /paket/update
			if r.URL.Path == "/paket/read" || r.URL.Path == "/paket/update" || r.URL.Path == "/tujuan/read" {
				next.ServeHTTP(w, r)
				return
			}
		} else {
			for _, role := range requiredRoles {
				if role == userRole {
					next.ServeHTTP(w, r)
					return
				}
			}
		}

		http.Error(w, "Anda tidak memiliki akses ke halaman ini", http.StatusForbidden)
	}
}

func getUserRole(r *http.Request) int {
	roleStr := r.Header.Get("Role")
	switch strings.TrimSpace(roleStr) {
	case "1":
		return 1
	case "2":
		return 2
	case "3":
		return 3
	default:
		return 0 // Default role sebagai guest
	}
}
