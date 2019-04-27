package infrastructure

import (
	"net/http"
	"os"

	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/interfaces"
	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/usecases"
)

// Dispatch is handle routing
func Dispatch(logger usecases.Logger, sqlHandler interfaces.SQLHandler, mux *http.ServeMux) {
	userController := interfaces.NewUserController(sqlHandler, logger)
	postController := interfaces.NewPostController(sqlHandler, logger)

	// TODO: Maybe I'll implement a routing package ...
	mux.HandleFunc("/users", userController.Index)
	mux.HandleFunc("/user", userController.Show)
	mux.HandleFunc("/post", postController.Store)

	if err := http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), mux); err != nil {
		logger.LogError("%s", err)
	}
}
