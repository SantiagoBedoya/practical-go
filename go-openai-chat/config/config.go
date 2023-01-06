package config

import (
	"log"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	OpenAIAPISecret string `toml:"OPENAI_API_SECRET"`
	OpenAIAPIURL    string `toml:"OPENAI_API_URL"`
}

func LoadConfig(fileName string) *Config {
	path, err := filepath.Abs(fileName)
	if err != nil {
		log.Fatal(err)
	}
	cfg := Config{}
	if _, err := toml.DecodeFile(path, &cfg); err != nil {
		log.Fatal(err)
	}
	return &cfg
}
