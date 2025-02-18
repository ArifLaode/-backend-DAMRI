package initiation

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func CreateTable(db *sql.DB) error {
	queries := []string{
		`
		CREATE TABLE IF NOT EXISTS admin (
			id_admin CHAR(36) PRIMARY KEY,
			nama VARCHAR(100),
			password VARCHAR(100),
			role INT
		);`,

		`
		CREATE TABLE IF NOT EXISTS tujuan (
			id_tujuan CHAR(36) PRIMARY KEY,
			nama VARCHAR(100),
			sekitar JSON,
			jarak FLOAT,
			harga FLOAT,
			koordinat JSON
		);`,

		`
		CREATE TABLE IF NOT EXISTS paket (
			id_paket CHAR(20) PRIMARY KEY,
			penerima VARCHAR(100),
			pengirim VARCHAR(100),
			telp_penerima VARCHAR(20),
			telp_pengirim VARCHAR(20),
			nama_barang VARCHAR(100),
			harga FLOAT,
			berat FLOAT,
			status INT,
			id_tujuan CHAR(36),
			FOREIGN KEY (id_tujuan) REFERENCES tujuan(id_tujuan)
		);`,
	}

	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			return err
		}
	}

	log.Println("Tabel telah berhasil dibuat")
	return nil
}
