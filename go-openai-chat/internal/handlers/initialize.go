package handlers

import (
	"fmt"
	"strings"
	"time"

	"github.com/SantiagoBedoya/openai-chat/internal/models"
	"github.com/urfave/cli/v2"
)

const (
	situationalContext = `You are a therapist. You are intelligent and witty. You are giving your client advice.
	Your conversation thus far has been:`
	inquiryToRespond = `Please respond, following the same conversation format.`
)

// Initialize will handle the chat with OpenAI
func (h *Handler) Initialize(ctx *cli.Context) error {
	now := time.Now()
	fmt.Println("Here you can type your message secure. Type `/exit` to finish the session (or just ctr + c)")
	for true {
		fmt.Print("> ")
		message, _ := h.reader.ReadString('\n')
		if strings.TrimSpace(message) == "/exit" {
			fmt.Println("Take care...")
			fmt.Printf("Time used: %s\n", getDurationString(now))
			break
		}
		var resp models.Response
		if err := h.openAI.CreateCompletion(models.Message{
			Model:       "text-davinci-003",
			Prompt:      buildPrompt(message),
			MaxTokens:   500,
			Temperature: 0.5,
		}, &resp); err != nil {
			return err
		}
		for _, choice := range resp.Choices {
			fmt.Printf("Talk: %s\n", strings.TrimSpace(choice.Text))
		}
	}
	return nil
}

func buildPrompt(message string) string {
	return fmt.Sprintf("%s\n: %s\n%s", situationalContext, message, inquiryToRespond)
}

func getDurationString(now time.Time) string {
	duration := time.Since(now)
	return duration.String()
}
