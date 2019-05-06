package repositories

import (
	"../../domain"
)

type UserRepository interface {
	FindAll() (domain.Users, error)
	FindById(int) (domain.User, error)
}
