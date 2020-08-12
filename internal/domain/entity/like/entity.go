package like

import (
	"time"

	"github.com/AVieraUY/go-forum/internal/domain/entity"
)

// Like data
type Like struct {
	ID        entity.ID
	UserID    entity.ID
	PostID    entity.ID
	CreatedAt time.Time
}
