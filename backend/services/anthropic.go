package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type AnthropicProvider struct {
	APIKey  string
	BaseURL string
}

func NewAnthropicProvider(apiKey, baseURL string) *AnthropicProvider {
	return &AnthropicProvider{
		APIKey:  apiKey,
		BaseURL: baseURL,
	}
}

func (p *AnthropicProvider) ChatStream(message string, writer io.Writer) error {
	url := fmt.Sprintf("%s/v1/messages", p.BaseURL)

	payload := map[string]interface{}{
		"model":     "claude-3-5-sonnet-20241022",
		"max_tokens": 4096,
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": message,
			},
		},
		"stream": true,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", p.APIKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("anthropic api error: %s", string(body))
	}

	decoder := json.NewDecoder(resp.Body)
	for {
		var event struct {
			Type string `json:"type"`
			Delta struct {
				Text string `json:"text"`
			} `json:"delta,omitempty"`
		}

		if err := decoder.Decode(&event); err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		if event.Type == "content_block_delta" && event.Delta.Text != "" {
			_, _ = writer.Write([]byte(event.Delta.Text))
		}

		if event.Type == "message_stop" {
			break
		}
	}

	return nil
}

