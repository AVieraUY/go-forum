package handler

import (
	"net/http"

	"github.com/AVieraUY/go-forum/internal/domain/entity/post"
	"github.com/go-chi/chi"
)

func createPost(post.Manager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func getPost(post.Manager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func getPosts(post.Manager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func searchPost(post.Manager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func updatePost(post.Manager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func deletePost(post.Manager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

// MakePostHandlers create the handlers to operate with post
func MakePostHandlers(r *chi.Mux, manager post.Manager) {
	r.Route("/posts", func(r chi.Router) {
		r.Post("/", createPost(manager))
		r.Get("/", getPosts(manager))

		r.Route("/{postId}", func(r chi.Router) {
			r.Get("/", getPost(manager))
			r.Put("/", updatePost(manager))
			r.Delete("/", deletePost(manager))
		})
	})
}
