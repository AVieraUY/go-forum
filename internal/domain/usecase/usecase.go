package usecase

import (
	"strings"

	"github.com/AVieraUY/go-forum/internal/security/password"

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
	REGISTER        = "register"
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
	err := validateUser(u, REGISTER)
	if err != nil {
		return err
	}

	usrs, err := s.uManager.Search(u.Username)
	if err != nil {
		return err
	}
	if len(usrs) >= 1 {
		return domain.ErrDuplicatedUsername
	}

	_, err = s.uManager.Create(u)
	if err != nil {
		return err
	}

	return nil
}

// Return an access token
func (s *usecase) Login(u *user.User) error {
	err := validateUser(u, LOGIN)
	if err != nil {
		return err
	}

	usr, err := s.uManager.Get(u.ID)
	if err != nil {
		return err
	}

	err = password.NewService().Verify(usr.Password, u.Password)
	if err != nil {
		return err
	}

	return nil
}

// ForgotPassword send a token to update password
func (s *usecase) ForgotPassword(u *user.User) error {
	return nil
}

// UpdateDetails update details of an user
func (s *usecase) UpdateDetails(u *user.User) error {
	err := validateUser(u, UPDATE)
	if err != nil {
		return err
	}

	_, err = s.uManager.Get(u.ID)
	if err != nil {
		return err
	}

	err = s.uManager.Update(u)
	if err != nil {
		return err
	}

	return nil
}

func validateUser(u *user.User, action string) error {
	var err error
	switch strings.ToLower(action) {
	case LOGIN:
		if u.Username == "" {
			err = domain.ErrInvalidUsername
			return err
		}
		if u.Password == "" || (u.Password != "" && len(u.Password) < 6) {
			err = domain.ErrInvalidPassword
			return err
		}
	case UPDATE, FORGOT_PASSWORD:
		if u.Email == "" || !validators.IsValid(u.Email) {
			err = domain.ErrInvalidEmail
			return err
		}
	case REGISTER:
		if u.Username == "" {
			err = domain.ErrInvalidUsername
			return err
		}
		if u.Password == "" || (u.Password != "" && len(u.Password) < 6) {
			err = domain.ErrInvalidPassword
			return err
		}
		if u.Email == "" || !validators.IsValid(u.Email) {
			err = domain.ErrInvalidEmail
			return err
		}
	}

	if err != nil {
		return err
	}
	return nil
}
