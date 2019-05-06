package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"../../interfaces/database"
)

type SqlHandler struct {
	Conn *sql.DB
}

type Result struct {
	Result sql.Result
}

type Row struct {
	Rows *sql.Rows
}

func NewSqlHandler() database.SqlHandler {
	sqlHandler := &SqlHandler{}
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))
	conn, err := sql.Open(os.Getenv("DB_DRIVER"), dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	err = conn.Ping()
	if err != nil {
		fmt.Println("Failed to connected DB:", err)
	}
	sqlHandler.Conn = conn

	return sqlHandler
}

func (sqlHandler *SqlHandler) Query(query string, args ...interface{}) (database.Row, error) {
	rows, err := sqlHandler.Conn.Query(query, args...)

	if err != nil {
		return nil, err
	}

	row := &Row{}
	row.Rows = rows

	return row, err
}

func (result Result) LastInsertId() (int64, error) {
	return result.Result.LastInsertId()
}

func (result Result) RowsAffected() (int64, error) {
	return result.Result.RowsAffected()
}

func (row Row) Scan(dest ...interface{}) error {
	return row.Rows.Scan(dest...)
}

func (row Row) Next() bool {
	return row.Rows.Next()
}

func (row Row) Close() error {
	return row.Rows.Close()
}

func (row Row) Err() error {
	return row.Rows.Err()
}
