package user

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

// Get obtain a user by ID
func (s *manager) Get(id entity.ID) (*User, error) {
	return s.repo.Get(id)
}

// Search obtain users by query
func (s *manager) Search(query string) ([]*User, error) {
	return s.repo.Search(strings.ToLower(query))
}

// List list all users
func (s *manager) List() ([]*User, error) {
	return s.repo.List()
}

// Create creates a new user
func (s *manager) Create(e *User) (entity.ID, error) {
	e.ID = entity.NewID()
	e.CreatedAt = time.Now()
	return s.repo.Create(e)
}

// Update update an existing user
func (s *manager) Update(e *User) error {
	e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}

// Delete delete an existing user
func (s *manager) Delete(id entity.ID) error {
	_, err := s.repo.Get(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}
