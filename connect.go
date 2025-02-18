package main

import (
	initiation "damri/Initiation"
	logic "damri/Logic"
	"database/sql"
)

func ConnectToDatabase() (*sql.DB, error) {
	db, err := initiation.NewDB()
	if err != nil {
		return nil, err
	}
	err = initiation.CreateTable(db)
	if err != nil {
		return nil, err
	}

	logic.SetDB(db)
	return db, nil
}
