package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"

	"github.com/pkg/errors"
)

// Filename of each enviroment
const (
	DevFilename     = "./config_dev.yaml"
	TestingFilename = "./config_testing.yaml"
	StagingFilename = "./config_staging.yaml"
	ProdFilename    = "./config_prod.yaml"
)

// LibName of loggerConfig
const (
	Zap = "zap"
)

// Config represents the entire application config
type Config struct {
	dbConfig     DBConfig     `yaml:"dbConfig"`
	apiConfig    APIConfig    `yaml:"apiConfig"`
	loggerConfig LoggerConfig `yaml:"loggerConfig"`
	zapConfig    LoggerConfig `yaml:"zapConfig"`
}

// DBConfig for db configuration
type DBConfig struct {
	DbName     string `yaml:"dbName"`
	DbUser     string `yaml:"dbUser"`
	DbPassword string `yaml:"dbPassword"`
	DbHost     string `yaml:"dbHost"`
	DbPort     int16  `yaml:"dbPort"`
}

// APIConfig for api configuration
type APIConfig struct {
	APIPort   int16  `yaml:"apiPort"`
	APISecret string `yaml:"apiSecret"`
}

// LoggerConfig for log configuration
type LoggerConfig struct {
	LibName      string `yaml:"libName"`
	Level        string `yaml:"level"`
	EnableCaller bool   `yaml:"enableCaller"`
}

// ReadConfig of the filename given
func ReadConfig(filename string) (*Config, error) {
	var c Config
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrap(err, "error reading config file")
	}

	err = yaml.Unmarshal(file, &c)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling file")
	}

	return &c, nil
}
