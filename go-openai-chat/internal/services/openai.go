package services

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/SantiagoBedoya/openai-chat/config"
	"github.com/SantiagoBedoya/openai-chat/internal/models"
)

type OpenAIService struct {
	cfg     *config.Config
	baseURL string
	client  *http.Client
}

func NewOpenAIService(cfg *config.Config) *OpenAIService {
	client := http.DefaultClient
	return &OpenAIService{
		cfg:     cfg,
		baseURL: cfg.OpenAIAPIURL,
		client:  client,
	}
}

func (s OpenAIService) CreateCompletion(message models.Message, resp *models.Response) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, s.baseURL+"/completions", bytes.NewBuffer(data))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+s.cfg.OpenAIAPISecret)

	response, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(resp); err != nil {
		return err
	}
	return nil
}

func (s OpenAIService) Custom(method, path string, resp interface{}) error {
	req, err := http.NewRequest(method, s.baseURL+path, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+s.cfg.OpenAIAPISecret)
	response, err := s.client.Do(req)

	if err != nil {
		return err
	}
	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(resp); err != nil {
		return err
	}
	return nil
}
