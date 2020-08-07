package usecase

import "github.com/AVieraUY/go-forum/internal/domain/entity/user"

// UseCase use case interface
type UseCase interface {
	Register(u *user.User) error
	Login(u *user.User) error
	ForgotPassword(u *user.User) error
	UpdateDetails(u *user.User) error
}
