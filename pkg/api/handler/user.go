package handler

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/AVieraUY/go-forum/internal/domain/entity/user"
)

func createUser(manager user.Manager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func getUser(manager user.Manager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func getUsers(manager user.Manager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func searchUsers(manager user.Manager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func updateUser(manager user.Manager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func deleteUser(manager user.Manager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

// MakeUserHandlers create the handlers to operate with user
func MakeUserHandlers(r *chi.Mux, manager user.Manager) {
	r.Route("/users", func(r chi.Router) {
		r.Post("/", createUser(manager))
		r.Get("/", getUsers(manager))

		r.Route("/{userId}", func(r chi.Router) {
			r.Get("/", getUser(manager))
			r.Put("/", updateUser(manager))
			r.Delete("/", deleteUser(manager))
		})
	})
}
