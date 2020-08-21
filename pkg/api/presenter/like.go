package presenter

import (
	"time"

	"github.com/AVieraUY/go-forum/internal/domain/entity"
)

// Like data
type Like struct {
	ID        entity.ID `json:"id"`
	UserID    entity.ID `json:"userId"`
	PostID    entity.ID `json:"postId"`
	CreatedAt time.Time `json:"createdAt"`
}
