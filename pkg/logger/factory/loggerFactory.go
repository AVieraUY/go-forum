package factory

import (
	"github.com/AVieraUY/go-forum/config"
)

// LoggerFactory logger factory
type LoggerFactory interface {
	Build(c *config.LoggerConfig) error
}

var loggerFactoryBuilderMap = map[string]LoggerFactory{
	config.Zap: &ZapFactory{},
}

// GetLoggerFactoryBuilder of concret implementation
func GetLoggerFactoryBuilder(key string) LoggerFactory {
	return loggerFactoryBuilderMap[key]
}
