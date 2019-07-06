package usecases

import (
	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/domain"
)

// A PostInteractor belong to the usecases layer.
type PostInteractor struct {
	PostRepository PostRepository
}

// Index is display a listing of the resource.
func (pi *PostInteractor) Index() (posts domain.Posts, err error) {
	posts, err = pi.PostRepository.FindAll()

	return
}

// Store is store a newly created resource in storage.
func (pi *PostInteractor) Store(p domain.Post) (id int64, err error) {
	id, err = pi.PostRepository.Save(p)

	return
}

// Destroy is remove the specified resource from storage.
func (pi *PostInteractor) Destroy(postID int) (err error) {
	err = pi.PostRepository.DeleteByID(postID)

	return
}
