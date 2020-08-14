package factory

// ZapFactory zap factory type
type ZapFactory struct{}

// Build an instance of zap logger
func (f *ZapFactory) Build() error {
	return nil
}
