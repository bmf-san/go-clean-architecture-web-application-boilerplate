package interactors

import (
	"../../domain"
	"../repositories"
)

type UserInteractor struct {
	UserRepository repositories.UserRepository
}

func (userInteractor *UserInteractor) GetUsers() (users domain.Users, err error) {
	users, err = userInteractor.UserRepository.FindAll()

	return
}

func (UserInteractor *UserInteractor) GetUserById(userId int) (user domain.User, err error) {
	user, err = UserInteractor.UserRepository.FindById(userId)

	return
}
