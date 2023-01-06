package app

import (
	"log"
	"os"

	"github.com/SantiagoBedoya/openai-chat/config"
	"github.com/urfave/cli/v2"
)

var cfg *config.Config

// Run initialize the OpenAI chat
func Run() {
	app := &cli.App{
		Name:     "talk",
		Usage:    "Talk with OpenAI about anything",
		Commands: initiliazeCommands(cfg),
		Flags:    flags,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
