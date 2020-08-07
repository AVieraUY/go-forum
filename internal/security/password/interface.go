package password

type Service interface {
	Hash(p string) ([]byte, error)
	Verify(hashed, p string) error
}
