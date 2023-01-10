package config

import (
	"log"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Host       string `toml:"host"`
	Port       int    `toml:"port"`
	DbURI      string `toml:"db_uri"`
	PrivateKey string `toml:"private_key"`
	PublicKey  string `toml:"public_key"`
}

func LoadConfig(filename string) *Config {
	path, err := filepath.Abs(filename)
	if err != nil {
		log.Fatal(err)
	}
	var cfg Config
	if _, err := toml.DecodeFile(path, &cfg); err != nil {
		log.Fatal(err)
	}
	return &cfg
}
