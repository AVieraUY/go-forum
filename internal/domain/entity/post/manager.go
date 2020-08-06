package post

import (
	"strings"
	"time"

	"github.com/AVieraUY/go-forum/internal/domain/entity"
)

type manager struct {
	repo repository
}

// NewManager creates a new manager of given repository
func NewManager(r repository) *manager {
	return &manager{repo: r}
}

// Get obtain a post by ID
func (s *manager) Get(id entity.ID) (*Post, error) {
	return s.repo.Get(id)
}

// Search obtain posts by query
func (s *manager) Search(query string) ([]*Post, error) {
	return s.repo.Search(strings.ToLower(query))
}

// List list all posts
func (s *manager) List() ([]*Post, error) {
	return s.repo.List()
}

// Create creates a new post
func (s *manager) Create(e *Post) (entity.ID, error) {
	e.ID = entity.NewID()
	e.CreatedAt = time.Now()
	return s.repo.Create(e)
}

// Update update an existing post
func (s *manager) Update(e *Post) error {
	e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}

// Delete delete and existing post
func (s *manager) Delete(id entity.ID) error {
	_, err := s.repo.Get(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}
