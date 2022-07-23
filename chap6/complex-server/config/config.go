package config

import (
	"io"
	"log"
)

// AppConfig define data struct
type AppConfig struct {
	Logger *log.Logger
}

// InitConfig define initializer for config
func InitConfig(w io.Writer) AppConfig {
	return AppConfig{
		Logger: log.New(w, "", log.Ldate|log.Ltime|log.Lshortfile),
	}
}
