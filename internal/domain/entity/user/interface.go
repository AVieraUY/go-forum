package user

import (
	"github.com/AVieraUY/go-forum/internal/domain/entity"
)

// Reader user reader
type Reader interface {
	Get(id entity.ID) (*User, error)
	Search(query string) ([]*User, error)
	List() ([]*User, error)
}

//Writer user writer
type Writer interface {
	Create(e *User) (entity.ID, error)
	Update(e *User) error
	Delete(id entity.ID) error
}

type repository interface {
	Reader
	Writer
}

// Manager user manager
type Manager interface {
	repository
}
