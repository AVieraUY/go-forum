package comment

import (
	"time"

	"github.com/AVieraUY/go-forum/internal/domain/entity"
)

// Comment data
type Comment struct {
	ID        entity.ID
	UserID    entity.ID
	PostID    entity.ID
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
