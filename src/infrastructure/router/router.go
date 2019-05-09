package router

import (
	"net/http"

	"github.com/bmf-san/go-api-boilerplate/infrastructure/database"
	"github.com/bmf-san/go-api-boilerplate/interfaces/controllers"
)

func Dispatch(mux *http.ServeMux) {
	userController := controllers.NewUserController(database.NewSqlHandler())

	mux.HandleFunc("/users", userController.Index)
	mux.HandleFunc("/user", userController.Show)
}
