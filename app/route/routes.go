package route

import (
	"database/sql"
	"net/http"

	"../handler"
)

func Dispatch(mux *http.ServeMux, db *sql.DB) {
	userHandler := handler.NewUserHandler(db)

	mux.HandleFunc("/users", userHandler.Index)
}
