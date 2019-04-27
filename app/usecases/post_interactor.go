package usecases

// A PostInteractor belong to the usecases layer.
type PostInteractor struct {
	PostRepository PostRepository
}

// StorePost represents a service which is related to store a Post.
func (ui *PostInteractor) StorePost() (id int64, err error) {
	id, err = ui.PostRepository.AddPost()

	return
}
