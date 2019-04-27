package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"../database"
	"../repository"
)

type userHandler struct {
	UserRepository repository.UserRepositoryImpl
}

func NewUserHandler(sqlHandler database.SqlHandler) *userHandler {
	return &userHandler{
		UserRepository: repository.UserRepositoryImpl{
			SqlHandler: sqlHandler,
		},
	}
}

func (handler *userHandler) List(w http.ResponseWriter, r *http.Request) {
	users, err := handler.UserRepository.GetUsers()
	if err != nil {
		log.Fatalf("show users list error: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
