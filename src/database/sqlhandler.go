package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type SqlHandler interface {
	Connect() (conn *sql.DB)
}

type SqlHandlerImpl struct {
	conn *sql.DB
}

func NewSqlHandler() SqlHandler {
	sqlHandlerImpl := &SqlHandlerImpl{}
	conn := sqlHandlerImpl.Connect()
	sqlHandlerImpl.conn = conn

	return sqlHandlerImpl
}

func (sqlHandler *SqlHandlerImpl) Connect() (conn *sql.DB) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))
	conn, err := sql.Open(os.Getenv("DB_DRIVER"), dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	err = conn.Ping()
	if err != nil {
		fmt.Println("Failed to connected DB:", err)
	}

	return conn
}
