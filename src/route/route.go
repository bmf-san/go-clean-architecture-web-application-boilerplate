package route

import (
	"net/http"

	"../database"
	"../handler"
)

func Dispatch(mux *http.ServeMux) {
	userHandler := handler.NewUserHandler(database.NewSqlHandler())

	mux.HandleFunc("/users", userHandler.List)
	mux.HandleFunc("/user", userHandler.Show)
}
