package comment

import (
	"time"

	"github.com/AVieraUY/go-forum/internal/domain/entity"

	"github.com/AVieraUY/go-forum/internal/domain/entity/post"
	"github.com/AVieraUY/go-forum/internal/domain/entity/user"
)

// Comment data
type Comment struct {
	ID        entity.ID
	User      user.User
	Post      post.Post
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
