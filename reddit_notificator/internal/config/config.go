package config

import (
	"log"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Port int    `toml:"port"`
	Host string `toml:"host"`
}

func LoadConfig(fileName string) *Config {
	path, err := filepath.Abs(fileName)
	if err != nil {
		log.Fatal(err)
	}
	var cfg Config
	if _, err = toml.DecodeFile(path, &cfg); err != nil {
		log.Fatal(err)
	}
	return &cfg
}
