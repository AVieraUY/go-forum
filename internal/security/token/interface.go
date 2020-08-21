package token

// Service token interface
type Service interface {
	Hash(t string) string
}
