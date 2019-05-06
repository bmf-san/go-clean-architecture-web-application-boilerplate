package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"../../usecases/interactors"
	"../database"
	repository "../repositories"
)

type UserController struct {
	UserInteractor interactors.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		UserInteractor: interactors.UserInteractor{
			UserRepository: &repository.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (userController *UserController) Index(w http.ResponseWriter, r *http.Request) {
	users, err := userController.UserInteractor.GetUsers()
	if err != nil {
		log.Fatalf("show users list error: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (userController *UserController) Show(w http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.Atoi(r.URL.Query().Get("id"))

	user, err := userController.UserInteractor.GetUserById(userId)
	if err != nil {
		log.Fatalf("show a user error: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
