package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type SqlHandler interface {
	Query(string, ...interface{}) (Rows, error)
}

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type Rows interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
	Err() error
}

type SqlHandlerImpl struct {
	Conn *sql.DB
}

type ResultImpl struct {
	Result sql.Result
}

type RowsImpl struct {
	Rows *sql.Rows
}

func NewSqlHandler() SqlHandler {
	sqlHandlerImpl := &SqlHandlerImpl{}
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))
	conn, err := sql.Open(os.Getenv("DB_DRIVER"), dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	err = conn.Ping()
	if err != nil {
		fmt.Println("Failed to connected DB:", err)
	}
	sqlHandlerImpl.Conn = conn

	return sqlHandlerImpl
}

func (sqlHandler *SqlHandlerImpl) Query(query string, args ...interface{}) (Rows, error) {
	rows, err := sqlHandler.Conn.Query(query, args...)

	if err != nil {
		return nil, err
	}

	rowsImpl := &RowsImpl{}
	rowsImpl.Rows = rows

	return rowsImpl, err
}

func (result ResultImpl) LastInsertId() (int64, error) {
	return result.Result.LastInsertId()
}

func (result ResultImpl) RowsAffected() (int64, error) {
	return result.Result.RowsAffected()
}

func (rows RowsImpl) Scan(dest ...interface{}) error {
	return rows.Rows.Scan(dest...)
}

func (rows RowsImpl) Next() bool {
	return rows.Rows.Next()
}

func (rows RowsImpl) Close() error {
	return rows.Rows.Close()
}

func (rows RowsImpl) Err() error {
	return rows.Rows.Err()
}
