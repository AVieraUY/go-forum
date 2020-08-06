package user

import (
	"time"

	"github.com/AVieraUY/go-forum/internal/domain/entity"
)

// User data
type User struct {
	ID         entity.ID
	Username   string
	Email      string
	Password   string
	AvatarPath string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
