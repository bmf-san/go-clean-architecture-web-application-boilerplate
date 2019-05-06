package main

import (
	"log"
	"net/http"
	"os"

	"./infrastructure/env"
	"./infrastructure/router"
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
