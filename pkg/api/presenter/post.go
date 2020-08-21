package presenter

import (
	"time"

	"github.com/AVieraUY/go-forum/internal/domain/entity"
)

// Post data
type Post struct {
	ID        entity.ID `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserID    entity.ID `json:"userId"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
