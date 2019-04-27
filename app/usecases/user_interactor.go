package usecases

import (
	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/domain"
)

// A UserInteractor belong to the usecases layer.
type UserInteractor struct {
	UserRepository UserRepository
}

// ShowUsers represents a service which is related to get all users.
func (ui *UserInteractor) ShowUsers() (users domain.Users, err error) {
	users, err = ui.UserRepository.FindAll()

	return
}

// ShowUserByID represents a service which is related to get a user.
func (ui *UserInteractor) ShowUserByID(userID int) (user domain.User, err error) {
	user, err = ui.UserRepository.FindByID(userID)

	return
}
