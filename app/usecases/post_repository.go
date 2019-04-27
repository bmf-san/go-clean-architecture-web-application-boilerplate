package usecases

// A PostRepository belong to the usecases layer.
type PostRepository interface {
	AddPost() (int64, error)
}
