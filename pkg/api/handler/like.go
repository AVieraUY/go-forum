package handler

import (
	"net/http"

	"github.com/AVieraUY/go-forum/internal/domain/entity/like"
	"github.com/go-chi/chi"
)

func createLike(manager like.Manager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func getLike(manager like.Manager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func getLikeByPostID(manager like.Manager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func deleteLike(manager like.Manager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func deleteLikesFromPost(manager like.Manager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func deleteLikesFromUser(manager like.Manager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

// MakeLikeHandlers create handlers to operate with like
func MakeLikeHandlers(r *chi.Mux, manager like.Manager) {
	r.Route("/likes", func(r chi.Router) {
		r.Post("/", createLike(manager))

		r.Route("/{likeId}", func(r chi.Router) {
			r.Get("/", getLike(manager))
			r.Delete("/", deleteLike(manager))
		})

		r.Route("/fromPost/{postId}", func(r chi.Router) {
			r.Get("/", getLikeByPostID(manager))
			r.Delete("/", deleteLikesFromPost(manager))
		})

		r.Route("/fromUser/{userId}", func(r chi.Router) {
			r.Delete("/", deleteLikesFromUser(manager))
		})
	})
}
