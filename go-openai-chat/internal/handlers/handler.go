package handlers

import (
	"bufio"
	"os"

	"github.com/SantiagoBedoya/openai-chat/config"
	"github.com/SantiagoBedoya/openai-chat/internal/services"
)

// Handler save the configurations, stdin reader, and OpenAI service
type Handler struct {
	cfg    *config.Config
	reader *bufio.Reader
	openAI *services.OpenAIService
}

// NewHandler initialize the handler for commands
func NewHandler(cfg *config.Config) *Handler {
	if cfg == nil {
		cfg = config.LoadConfig("config.cfg")
	}
	client := services.NewOpenAIService(cfg)
	return &Handler{
		cfg:    cfg,
		reader: bufio.NewReader(os.Stdin),
		openAI: client,
	}
}
