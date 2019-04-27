package main

import (
	"net/http"

	"github.com/bmf-san/go-api-boilerplate/app/infrastructure"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	logger := infrastructure.NewLogger()

	infrastructure.Load(logger)

	sqlHandler, err := infrastructure.NewSQLHandler()
	if err != nil {
		logger.LogError("%s", err)
	}

	mux := http.NewServeMux()
	infrastructure.Dispatch(logger, sqlHandler, mux)
}
