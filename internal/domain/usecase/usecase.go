package usecase

import (
	"strings"

	validators "github.com/AVieraUY/go-validators/email"

	"github.com/AVieraUY/go-forum/internal/domain"
	"github.com/AVieraUY/go-forum/internal/domain/entity/comment"
	"github.com/AVieraUY/go-forum/internal/domain/entity/post"
	"github.com/AVieraUY/go-forum/internal/domain/entity/user"
)

type usecase struct {
	uManager user.Manager
	pManager post.Manager
	cManager comment.Manager
}

const (
	LOGIN           = "login"
	UPDATE          = "update"
	FORGOT_PASSWORD = "forgot_password"
)

// NewUseCase create new use case
func NewUseCase(u user.Manager, p post.Manager, c comment.Manager) *usecase {
	return &usecase{uManager: u, pManager: p, cManager: c}
}

// Register an user
func (s *usecase) Register(u *user.User) error {
	return nil
}

// Login an user
func (s *usecase) Login(u *user.User) error {
	return nil
}

// ForgotPassword send a token to update password
func (s *usecase) ForgotPassword(u *user.User) error {
	return nil
}

// UpdateDetails update details of an user
func (s *usecase) UpdateDetails(u *user.User) error {
	return nil
}

func validateUser(u *user.User, action string) map[string]string {
	var errorMessages = make(map[string]string)

	switch strings.ToLower(action) {
	case LOGIN:
		if u.Username == "" {
			errorMessages["invalid_username"] = domain.ErrInvalidUsername.Error()
		}
		if u.Password == "" || (u.Password != "" && len(u.Password) < 6) {
			errorMessages["invalid_password"] = domain.ErrInvalidPassword.Error()
		}
	case UPDATE, FORGOT_PASSWORD:
		if u.Email == "" || !validators.IsValid(u.Email) {
			errorMessages["invalid_email"] = domain.ErrInvalidEmail.Error()
		}
	default:
		if u.Username == "" {
			errorMessages["invalid_username"] = domain.ErrInvalidUsername.Error()
		}
		if u.Password == "" || (u.Password != "" && len(u.Password) < 6) {
			errorMessages["invalid_password"] = domain.ErrInvalidPassword.Error()
		}
		if u.Email == "" || !validators.IsValid(u.Email) {
			errorMessages["invalid_email"] = domain.ErrInvalidEmail.Error()
		}
	}

	return errorMessages
}
