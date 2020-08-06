package post

import (
	"github.com/AVieraUY/go-forum/internal/domain/entity"
)

// Reader post reader
type Reader interface {
	Get(id entity.ID) (*Post, error)
	Search(query string) ([]*Post, error)
	List() ([]*Post, error)
}

// Writer post writer
type Writer interface {
	Create(e *Post) (entity.ID, error)
	Update(e *Post) error
	Delete(id entity.ID) error
}

type repository interface {
	Reader
	Writer
}

// Manager post manager
type Manager interface {
	repository
}
