package initiation

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() (*sql.DB, error) {
	dsn := "root@tcp(127.0.0.1:3306)/bus"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	log.Println("Koneksi ke database berhasil")
	return db, nil
}
