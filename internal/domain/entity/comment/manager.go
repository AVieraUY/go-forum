package comment

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

// Get obtain a comment by id
func (s *manager) Get(id entity.ID) (*Comment, error) {
	return s.repo.Get(id)
}

// Search obtain comments by query
func (s *manager) Search(query string) ([]*Comment, error) {
	return s.repo.Search(strings.ToLower(query))
}

// List lists all comments
func (s *manager) List() ([]*Comment, error) {
	return s.repo.List()
}

// Create creates a new comment
func (s *manager) Create(e *Comment) (entity.ID, error) {
	e.ID = entity.NewID()
	e.CreatedAt = time.Now()
	return s.repo.Create(e)
}

// Update update an existing comment
func (s *manager) Update(e *Comment) error {
	e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}

// Delete delete an existing comment
func (s *manager) Delete(id entity.ID) error {
	_, err := s.repo.Get(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}
