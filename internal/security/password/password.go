package password

import (
	"golang.org/x/crypto/bcrypt"
)

// Password type
type Password struct{}

// NewService creates a new password service
func NewService() *Password {
	return &Password{}
}

// Hash a raw password
func (s *Password) Hash(p string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
}

// Verify a hashed password against a raw
func (s *Password) Verify(hashed, p string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(p))
}
