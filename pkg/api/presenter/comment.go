package presenter

import (
	"time"

	"github.com/AVieraUY/go-forum/internal/domain/entity"
)

// Comment data
type Comment struct {
	ID        entity.ID `json:"id"`
	User      User      `json:"user"`
	Post      Post      `json:"post"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
