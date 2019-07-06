package interfaces

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/domain"
	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/usecases"
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

// Index is display a listing of the resource.
func (pc *PostController) Index(w http.ResponseWriter, r *http.Request) {
	pc.Logger.LogAccess("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

	posts, err := pc.PostInteractor.Index()
	if err != nil {
		pc.Logger.LogError("%s", err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

// Store is stora a newly created resource in storage.
func (pc *PostController) Store(w http.ResponseWriter, r *http.Request) {
	pc.Logger.LogAccess("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

	p := domain.Post{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		pc.Logger.LogError("%s", err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
	}

	post, err := pc.PostInteractor.Store(p)
	if err != nil {
		pc.Logger.LogError("%s", err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

// Destroy is remove the specified resource from storage.
func (pc *PostController) Destroy(w http.ResponseWriter, r *http.Request) {
	pc.Logger.LogAccess("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

	postID, _ := strconv.Atoi(r.URL.Query().Get("id"))

	err := pc.PostInteractor.Destroy(postID)
	if err != nil {
		pc.Logger.LogError("%s", err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nil)
}
