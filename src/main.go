package main

import (
	"log"
	"net/http"
	"os"

	"./config"
	"./route"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config.Load()

	mux := http.NewServeMux()
	route.Dispatch(mux)

	if err := http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), mux); err != nil {
		log.Fatal(err)
	}
}
