package post

import (
	"time"

	"github.com/AVieraUY/go-forum/internal/domain/entity"
)

// Post data
type Post struct {
	ID        entity.ID
	Title     string
	Content   string
	UserID    entity.ID
	CreatedAt time.Time
	UpdatedAt time.Time
}
