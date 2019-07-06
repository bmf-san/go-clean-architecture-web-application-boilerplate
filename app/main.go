package main

import (
	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/infrastructure"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	logger := infrastructure.NewLogger()

	infrastructure.Load(logger)

	sqlHandler, err := infrastructure.NewSQLHandler()
	if err != nil {
		logger.LogError("%s", err)
	}

	infrastructure.Dispatch(logger, sqlHandler)
}
