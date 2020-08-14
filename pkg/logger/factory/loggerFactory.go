package factory

import (
	"github.com/AVieraUY/go-forum/config"
	"github.com/AVieraUY/go-forum/pkg/logger"
)

// LoggerFactory logger factory
type LoggerFactory interface {
	Build(c *config.LoggerConfig) (logger.Logger, error)
}
