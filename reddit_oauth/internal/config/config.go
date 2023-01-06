package config

import (
	"log"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Port            int    `toml:"port"`
	ReadTimeout     int    `toml:"read_timeout"`
	WriteTimeout    int    `toml:"write_timeout"`
	IdleTimeout     int    `toml:"idle_timoeut"`
	MongoURI        string `toml:"mongo_uri"`
	MongoDB         string `toml:"mongo_db"`
	MongoCollection string `toml:"mongo_collection"`
	NotificatorHost string `toml:"notificator_host"`
	NotificatorPort int    `toml:"notificator_port"`
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
