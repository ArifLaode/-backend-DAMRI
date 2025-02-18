// main.go
package main

import (
	router "damri/Router" // Import package Router
	"log"
	"net/http"
)

func main() {
	db, err := ConnectToDatabase() // Memanggil fungsi ConnectToDatabase dari connect.go
	if err != nil {
		log.Fatalln("Database tidak dapat terhubung:", err)
	}
	defer db.Close()

	router.SetupRoutes()

	// SetupRoutes() harus didefinisikan di tempat lain
	// router.SetupRoutes()

	log.Println("Server berjalan di port 1034")
	log.Fatal(http.ListenAndServe(":1034", nil))
}
