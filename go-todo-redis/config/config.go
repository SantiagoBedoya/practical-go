package config

import (
	"log"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

// Config define data structure for config variables
type Config struct {
	Port      int    `toml:"port"`
	RedisHost string `toml:"redis_host"`
	RedisPort int    `toml:"redis_port"`
}

// LoadConfig load config from file
func LoadConfig(fileName string) *Config {
	cfgFile, _ := filepath.Abs(fileName)
	conf := Config{}
	if _, err := toml.DecodeFile(cfgFile, &conf); err != nil {
		logrus.Error(err)
		log.Fatal("error reading config file")
	}
	return &conf
}
