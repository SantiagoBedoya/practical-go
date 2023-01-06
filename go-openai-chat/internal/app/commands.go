package app

import (
	"github.com/SantiagoBedoya/openai-chat/config"
	"github.com/SantiagoBedoya/openai-chat/internal/handlers"
	"github.com/urfave/cli/v2"
)

func initiliazeCommands(cfg *config.Config) []*cli.Command {
	handler := handlers.NewHandler(cfg)
	commands := []*cli.Command{
		{
			Name:   "start",
			Usage:  "Initialize the chat",
			Action: handler.Initialize,
		},
	}
	return commands
}
