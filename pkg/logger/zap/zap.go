package zap

import (
	"encoding/json"

	"github.com/AVieraUY/go-forum/config"
	"github.com/AVieraUY/go-forum/pkg/logger"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// RegisterLogger zap
func RegisterLogger(c config.LoggerConfig) error {
	zLogger, err := initConfig(c)
	if err != nil {
		return errors.Wrap(err, "Failed to initialize logger")
	}

	defer zLogger.Sync()
	sugar := zLogger.Sugar()
	sugar.Info()
	logger.SetLogger(sugar)
	return nil
}

func initConfig(c config.LoggerConfig) (zap.Logger, error) {
	rawJSON := []byte(`{
		"level": "info",
		"Development": true,
		 "DisableCaller": false,
		"encoding": "console",
		"outputPaths": ["stdout", "../../../app.log"],
		"errorOutputPaths": ["stderr"],
		"encoderConfig": {
		   "timeKey":        "ts",
		   "levelKey":       "level",
		   "messageKey":     "msg",
			"nameKey":        "name",
		   "stacktraceKey":  "stacktrace",
			"callerKey":      "caller",
		   "lineEnding":     "\n\t",
		   "timeEncoder":     "time",
		   "levelEncoder":    "lowercaseLevel",
		   "durationEncoder": "stringDuration",
			"callerEncoder":   "shortCaller"
		}
	   }`)
	var cfg zap.Config
	var zap *zap.Logger
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		return *zap, errors.Wrap(err, "Failed to unmarshal config JSON")
	}

	if err := customizeConfig(&cfg, c); err != nil {
		return *zap, errors.Wrap(err, "Failed to customize config")
	}

	zap, err := cfg.Build()
	if err != nil {
		return *zap, errors.Wrap(err, "Failed to build config")
	}

	zap.Debug("Logger constructed.")

	return *zap, nil
}

func customizeConfig(cfg *zap.Config, c config.LoggerConfig) error {
	cfg.DisableCaller = !c.EnableCaller
	l := zap.NewAtomicLevel().Level()
	if err := l.Set(c.Level); err != nil {
		return errors.Wrap(err, "Failed to set custom level")
	}

	cfg.Level.SetLevel(l)
	return nil
}
