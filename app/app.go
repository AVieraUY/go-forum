package app

import (
	"github.com/AVieraUY/go-forum/config"
	"github.com/AVieraUY/go-forum/pkg/logger/factory"
	"github.com/pkg/errors"
)

// InitApp initialize configuration of the app
func InitApp(filename string) (*config.Config, error) {
	c, err := config.ReadConfig(filename)
	if err != nil {
		return nil, errors.Wrap(err, "ReadConfig")
	}
	err = initLogger(&c.LoggerConfig)
	if err != nil {
		return nil, errors.Wrap(err, "initLogger")
	}

	return c, nil
}

func initLogger(c *config.LoggerConfig) error {
	err := factory.GetLoggerFactoryBuilder(config.Zap).Build(c)
	if err != nil {
		return errors.Wrap(err, "initLogger")
	}
	return nil
}
