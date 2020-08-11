package like

import (
	"time"

	"github.com/AVieraUY/go-forum/internal/domain/entity"
)

type manager struct {
	repo repository
}

// NewManager creates a new like manager
func NewManager(r repository) *manager {
	return &manager{repo: r}
}

// Get a like
func (s *manager) Get(id entity.ID) (*Like, error) {
	return s.repo.Get(id)
}

// Search likes
func (s *manager) Search(query string) (*[]Like, error) {
	return s.repo.Search(query)
}

// List likes
func (s *manager) List() (*[]Like, error) {
	return s.repo.List()
}

// Creates a new like
func (s *manager) Create(e *Like) (entity.ID, error) {
	e.ID = entity.NewID()
	e.CreatedAt = time.Now()
	return s.repo.Create(e)
}

// Update an existing like
func (s *manager) Update(e *Like) error {
	e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}

// Delete an existing like
func (s *manager) Delete(id entity.ID) error {
	_, err := s.repo.Get(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}
