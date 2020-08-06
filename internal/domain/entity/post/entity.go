package post

import (
	"time"

	"github.com/AVieraUY/go-forum/internal/domain/entity"
	"github.com/AVieraUY/go-forum/internal/domain/entity/user"
)

// Post data
type Post struct {
	ID        entity.ID
	Title     string
	Content   string
	Author    user.User
	CreatedAt time.Time
	UpdatedAt time.Time
}
