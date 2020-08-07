package token

type Service interface {
	Hash(t string) string
}
