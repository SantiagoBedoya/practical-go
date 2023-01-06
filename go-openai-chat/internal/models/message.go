package models

type Message struct {
	Model       string  `json:"model"`
	Prompt      string  `json:"prompt"`
	MaxTokens   uint    `json:"max_tokens"`
	Temperature float64 `json:"temperature"`
}
