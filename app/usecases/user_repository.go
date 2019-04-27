package usecases

import (
	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/domain"
)

// A UserRepository belong to the usecases layer.
type UserRepository interface {
	FindAll() (domain.Users, error)
	FindByID(int) (domain.User, error)
}
