package router

import (
	"net/http"

	"../../infrastructure/database"
	"../../interfaces/controllers"
)

func Dispatch(mux *http.ServeMux) {
	userController := controllers.NewUserController(database.NewSqlHandler())

	mux.HandleFunc("/users", userController.Index)
	mux.HandleFunc("/user", userController.Show)
}
