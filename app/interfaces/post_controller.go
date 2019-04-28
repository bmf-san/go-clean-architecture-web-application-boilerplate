package interfaces

import (
	"encoding/json"
	"net/http"

	"github.com/bmf-san/go-api-boilerplate/app/usecases"
)

// A PostController belong to the interface layer.
type PostController struct {
	PostInteractor usecases.PostInteractor
	Logger         usecases.Logger
}

// NewPostController returns the resource of Posts.
func NewPostController(sqlHandler SQLHandler, logger usecases.Logger) *PostController {
	return &PostController{
		PostInteractor: usecases.PostInteractor{
			PostRepository: &PostRepository{
				SQLHandler: sqlHandler,
			},
		},
		Logger: logger,
	}
}

// Store a newly created resource of a Post.
func (uc *PostController) Store(w http.ResponseWriter, r *http.Request) {
	uc.Logger.LogAccess("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

	Post, err := uc.PostInteractor.StorePost()
	if err != nil {
		uc.Logger.LogError("%s", err)
		// TODO: To return a error response message
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Post)
}
