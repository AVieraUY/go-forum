package comment

import "github.com/AVieraUY/go-forum/internal/domain/entity"

// Reader comment reader
type Reader interface {
	Get(id entity.ID) (*Comment, error)
	Search(query string) ([]*Comment, error)
	List() ([]*Comment, error)
}

// Writer comment writer
type Writer interface {
	Create(e *Comment) (entity.ID, error)
	Update(e *Comment) error
	Delete(id entity.ID) error
}

type repository interface {
	Reader
	Writer
}

// Manager comment manager
type Manager interface {
	repository
}
