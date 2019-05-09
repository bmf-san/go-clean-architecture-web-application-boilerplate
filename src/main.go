package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bmf-san/go-api-boilerplate/infrastructure/env"
	"github.com/bmf-san/go-api-boilerplate/infrastructure/router"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	env.Load()

	mux := http.NewServeMux()
	router.Dispatch(mux)

	if err := http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), mux); err != nil {
		log.Fatal(err)
	}
}
