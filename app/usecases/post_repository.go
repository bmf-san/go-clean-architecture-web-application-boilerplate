package usecases

import (
	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/domain"
)

// A PostRepository belong to the usecases layer.
type PostRepository interface {
	FindAll() (domain.Posts, error)
	Save(domain.Post) (int64, error)
	DeleteByID(int) error
}
