package presenter

import (
	"time"

	"github.com/AVieraUY/go-forum/internal/domain/entity"
)

// User data
type User struct {
	ID         entity.ID `json:"id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	AvatarPath string    `json:"avatar_path"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
