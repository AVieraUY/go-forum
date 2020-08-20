package handler

import (
	"net/http"

	"github.com/AVieraUY/go-forum/internal/domain/entity/comment"
	"github.com/go-chi/chi"
)

func createComment(manager comment.Manager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func getComment(manager comment.Manager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func getComments(manager comment.Manager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func searchComments(manager comment.Manager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func updateComment(manager comment.Manager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func deleteComment(manager comment.Manager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

// MakeCommentHandlers create the handlers to operate with comment
func MakeCommentHandlers(r *chi.Mux, manager comment.Manager) {
	r.Route("/comment", func(r chi.Router) {
		r.Post("/", createComment(manager))
		r.Get("/", getComments(manager))

		r.Route("/{commentId}", func(r chi.Router) {
			r.Get("/", getComment(manager))
			r.Put("/", updateComment(manager))
			r.Delete("/", deleteComment(manager))
		})
	})
}
