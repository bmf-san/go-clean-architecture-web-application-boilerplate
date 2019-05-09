package repositories

import (
	"github.com/bmf-san/go-api-boilerplate/domain"
)

type UserRepository interface {
	FindAll() (domain.Users, error)
	FindById(int) (domain.User, error)
}
