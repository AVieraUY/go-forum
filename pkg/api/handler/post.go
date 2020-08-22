package handler

import (
	"encoding/json"
	"net/http"

	"github.com/AVieraUY/go-forum/internal/domain/entity"
	"github.com/AVieraUY/go-forum/pkg/api/presenter"
	"github.com/AVieraUY/go-forum/pkg/logger"

	"github.com/AVieraUY/go-forum/internal/domain/entity/post"
	"github.com/go-chi/chi"
)

func createPost(manager post.Manager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func getPost(manager post.Manager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error getting post"
		notFoundMessage := "Post not found"
		id, err := entity.StringToID(chi.URLParam(r, "id"))
		if err != nil {
			logger.Log.Errorf("%s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		data, err := manager.Get(id)
		if err != nil {
			logger.Log.Errorf("%s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(notFoundMessage))
			return
		}

		j := &presenter.Post{
			ID:        data.ID,
			Title:     data.Title,
			Content:   data.Content,
			UserID:    data.UserID,
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,
		}

		if err := json.NewEncoder(w).Encode(j); err != nil {
			logger.Log.Errorf("%s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func getPosts(manager post.Manager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error getting posts"
		notFoundMessage := "Posts not found"
		data, err := manager.List()
		if err != nil {
			logger.Log.Errorf("%s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if len(data) <= 0 || data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(notFoundMessage))
			return
		}

		var j []*presenter.Post
		for _, d := range data {
			j = append(j, &presenter.Post{
				ID:        d.ID,
				Title:     d.Title,
				Content:   d.Content,
				UserID:    d.UserID,
				CreatedAt: d.CreatedAt,
				UpdatedAt: d.UpdatedAt,
			})
		}

		if err := json.NewEncoder(w).Encode(j); err != nil {
			logger.Log.Errorf("%s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func searchPost(manager post.Manager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}

func updatePost(manager post.Manager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error updating post"
		var input struct {
			ID      string `json:"id"`
			Title   string `json:"title"`
			Content string `json:"content"`
			UserID  string `json:"userId"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			logger.Log.Errorf("%s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		id, err := entity.StringToID(input.ID)
		if err != nil {
			logger.Log.Errorf("%s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		userID, err := entity.StringToID(input.UserID)
		if err != nil {
			logger.Log.Errorf("%s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		p := &post.Post{
			ID:      id,
			Title:   input.Title,
			Content: input.Content,
			UserID:  userID,
		}
		err = manager.Update(p)
		if err != nil {
			logger.Log.Errorf("%s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}

func deletePost(manager post.Manager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error deleting post"
		id, err := entity.StringToID(chi.URLParam(r, "id"))
		if err != nil {
			logger.Log.Errorf("%s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		if err = manager.Delete(id); err != nil {
			logger.Log.Errorf("%s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

// MakePostHandlers create the handlers to operate with post
func MakePostHandlers(r *chi.Mux, manager post.Manager) {
	r.Route("/posts", func(r chi.Router) {
		r.Post("/", createPost(manager))
		r.Get("/", getPosts(manager))

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", getPost(manager))
			r.Put("/", updatePost(manager))
			r.Delete("/", deletePost(manager))
		})
	})
}
