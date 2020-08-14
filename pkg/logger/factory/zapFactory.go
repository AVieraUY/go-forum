package factory

import (
	"github.com/AVieraUY/go-forum/config"
	"github.com/AVieraUY/go-forum/pkg/logger/zap"
	"github.com/pkg/errors"
)

// ZapFactory zap factory type
type ZapFactory struct{}

// Build an instance of zap logger
func (f *ZapFactory) Build(c *config.LoggerConfig) error {
	if err := zap.RegisterLogger(*c); err != nil {
		return errors.Wrap(err, "Failed to register logger")
	}
	return nil
}
