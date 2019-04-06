package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func Connect() (db *sql.DB) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))
	db, err := sql.Open(os.Getenv("DB_DRIVER"), dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Failed to connected DB:", err)
	}

	return db
}
