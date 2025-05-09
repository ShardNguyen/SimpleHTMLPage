package config

import (
	"SimpleHTMLPage/utilities"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v3"
)

const keyLength = 16

type Config struct {
	Postgresql struct {
		Dsn string `yaml:"dsn"`
	} `yaml:"postgresql"`
	secretKey string
}

var config *Config

func GetConfig() *Config {
	return config
}

func (c *Config) createSecretKey() {
	c.secretKey = utilities.StringRand(keyLength)
}

func (c *Config) GetPostgresqlDSN() string {
	return c.Postgresql.Dsn
}

func (c *Config) GetSecretKey() string {
	if c.secretKey == "" {
		c.createSecretKey()
	}
	return c.secretKey
}

func ParseConfig() error {
	config = &Config{}

	// Get current project path
	path, err := os.Getwd()

	if err != nil {
		return err
	}

	// Get config file based on path
	configFile, err := os.ReadFile(filepath.Join(path, "config/config.yaml"))

	if err != nil {
		return err
	}

	// Parse the config file
	err = yaml.Unmarshal(configFile, config)

	if err != nil {
		return err
	}

	return nil
}
