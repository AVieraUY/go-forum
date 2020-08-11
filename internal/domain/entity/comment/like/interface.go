package like

import "github.com/AVieraUY/go-forum/internal/domain/entity"

// Reader like reader
type Reader interface {
	Get(id entity.ID) (*Like, error)
	Search(query string) (*[]Like, error)
	List() (*[]Like, error)
}

// Writer like writer
type Writer interface {
	Create(e *Like) (entity.ID, error)
	Update(e *Like) error
	Delete(id entity.ID) error
}

type repository interface {
	Reader
	Writer
}

// Manager like manager
type Manager interface {
	repository
}
