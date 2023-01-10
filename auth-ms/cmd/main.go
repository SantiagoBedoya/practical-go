package main

import (
	"flag"

	"github.com/SantiagoBedoya/auth-ms/internal/app"
)

func main() {
	var config string
	flag.StringVar(&config, "config", ".config.cfg", "config file path")
	flag.Parse()

	app.Initialize(config)
}
