package config

import (
	"crypto/rsa"
	"encoding/pem"
	"errors"
	"os"
	"path/filepath"

	utilkey "SimpleHTMLPage/utilities/keys"
	utilstr "SimpleHTMLPage/utilities/string"

	yaml "gopkg.in/yaml.v3"
)

const keyLength = 16

type Config struct {
	Postgresql struct {
		Dsn string `yaml:"dsn"`
	} `yaml:"postgresql"`

	Redis struct {
		Addr     string `yaml:"addr"`
		DB       int    `yaml:"db"`
		Password string `yaml:"password"`
	} `yaml:"redis"`

	Jwt struct {
		PublicKeyPath  string `yaml:"publicKeyPath"`
		PrivateKeyPath string `yaml:"privateKeyPath"`
		ExpireDuration int    `yaml:"expireDuration"`
	} `yaml:"jwt"`

	secretKey string

	publicKey  *rsa.PublicKey  // Write "openssl rsa -in <path to private key> -pubout > <path to output public key> to generate a public key based on private key"
	privateKey *rsa.PrivateKey // Write "openssl genrsa -out <path to output private key> <keysize>" to generate a private key?
}

var config *Config

func GetConfig() *Config {
	return config
}

func (c *Config) createSecretKey() {
	c.secretKey = utilstr.StringRand(keyLength)
}

func (c *Config) GetPostgresqlDSN() string {
	return c.Postgresql.Dsn
}

func (c *Config) GetRedisAddress() string {
	return c.Redis.Addr
}

func (c *Config) GetRedisDB() int {
	return c.Redis.DB
}

func (c *Config) GetRedisPassword() string {
	return c.Redis.Password
}

func (c *Config) GetJWTExpireDuration() int {
	if c.Jwt.ExpireDuration < 0 {
		return 1800
	}
	return c.Jwt.ExpireDuration
}

func (c *Config) GetSecretKey() string {
	if c.secretKey == "" {
		c.createSecretKey()
	}
	return c.secretKey
}

func (c *Config) GetPrivateKey() *rsa.PrivateKey {
	return c.privateKey
}

func (c *Config) GetPublicKey() *rsa.PublicKey {
	return c.publicKey
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

	// Get key from files
	privateKey, err := getPrivateKeyFromFile(config.Jwt.PrivateKeyPath)
	if err != nil {
		return err
	}

	publicKey, err := getPublicKeyFromFile(config.Jwt.PublicKeyPath)
	if err != nil {
		return err
	}

	config.privateKey = privateKey
	config.publicKey = publicKey

	return nil
}

func getPrivateKeyFromFile(filename string) (*rsa.PrivateKey, error) {
	keyBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil, errors.New("failed to decode private PEM block")
	}

	key, err := utilkey.ParsePrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func getPublicKeyFromFile(filename string) (*rsa.PublicKey, error) {
	keyBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil, errors.New("failed to decode private PEM block")
	}

	key, err := utilkey.ParsePublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return key, nil
}
