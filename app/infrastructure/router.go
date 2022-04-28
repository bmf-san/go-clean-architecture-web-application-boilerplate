package infrastructure

import (
	"net/http"
	"os"

	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/interfaces"
	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/usecases"
	"github.com/go-chi/chi"
)

// Dispatch is handle routing
func Dispatch(logger usecases.Logger, sqlHandler interfaces.SQLHandler) {
	userController := interfaces.NewUserController(sqlHandler, logger)
	postController := interfaces.NewPostController(sqlHandler, logger)

	r := chi.NewRouter()
	r.Get("/users", userController.Index)
	r.Get("/user", userController.Show)
	r.Get("/posts", postController.Index)
	r.Post("/post", postController.Store)
	r.Delete("/post", postController.Destroy)

	if err := http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), r); err != nil {
		logger.LogError("%s", err)
	}
}
