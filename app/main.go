package main

import (
	"log"
	"net/http"
	"os"

	"./config"
	"./database"
	"./route"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config.LoadEnv()
	db := database.Connect()

	mux := http.NewServeMux()
	route.Dispatch(mux, db)

	if err := http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), mux); err != nil {
		log.Fatal(err)
	}
}
