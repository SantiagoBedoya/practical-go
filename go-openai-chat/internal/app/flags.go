package app

import (
	"fmt"
	"strings"

	"github.com/SantiagoBedoya/openai-chat/config"
	"github.com/urfave/cli/v2"
)

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:    "config",
		Aliases: []string{"c"},
		Usage:   "Load configuration from `FILEPATH`.cfg",
		Value:   "config.cfg",
		Action: func(ctx *cli.Context, s string) error {
			filePath := strings.TrimSpace(s)
			if len(filePath) == 0 {
				return fmt.Errorf("invalid file path")
			}
			cfg = config.LoadConfig(s)
			return nil
		},
	},
}
