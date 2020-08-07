package token

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/google/uuid"
)

// Token type
type Token struct{}

// NewService creates a new service token
func NewService() *Token {
	return &Token{}
}

// Hash a token
func (s *Token) Hash(t string) string {
	hasher := md5.New()
	hasher.Write([]byte(t))
	theHash := hex.EncodeToString(hasher.Sum(nil))

	u, err := uuid.NewUUID()
	if err != nil {
		return ""
	}
	token := theHash + u.String()
	return token
}
