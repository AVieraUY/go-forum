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

// Get likes from a post
func (s *manager) GetByPostID(id entity.ID) ([]*Like, error) {
	return s.repo.GetByPostID(id)
}

// Creates a new like
func (s *manager) Create(e *Like) (entity.ID, error) {
	e.ID = entity.NewID()
	e.CreatedAt = time.Now()
	return s.repo.Create(e)
}

// Delete an existing like
func (s *manager) Delete(id entity.ID) error {
	_, err := s.repo.Get(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}

// DeletePostLikes deletes all likes from a post
func (s *manager) DeletePostLikes(id entity.ID) error {
	return s.repo.DeletePostLikes(id)
}

// DeleteUserLikes deletes all likes from a user
func (s *manager) DeleteUserLikes(id entity.ID) error {
	return s.repo.DeleteUserLikes(id)
}
