package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type OpenAIProvider struct {
	APIKey  string
	BaseURL string
}

func NewOpenAIProvider(apiKey, baseURL string) *OpenAIProvider {
	return &OpenAIProvider{
		APIKey:  apiKey,
		BaseURL: baseURL,
	}
}

func (p *OpenAIProvider) ChatStream(message string, writer io.Writer) error {
	url := fmt.Sprintf("%s/chat/completions", p.BaseURL)

	payload := map[string]interface{}{
		"model": "gpt-3.5-turbo",
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
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", p.APIKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("openai api error: %s", string(body))
	}

	buffer := make([]byte, 4096)
	for {
		n, err := resp.Body.Read(buffer)
		if n > 0 {
			data := string(buffer[:n])
			lines := parseSSE(data)
			for _, line := range lines {
				if line == "" || line == "[DONE]" {
					continue
				}
				
				var streamResp struct {
					Choices []struct {
						Delta struct {
							Content string `json:"content"`
						} `json:"delta"`
					} `json:"choices"`
				}

				if err := json.Unmarshal([]byte(line), &streamResp); err != nil {
					continue
				}

				if len(streamResp.Choices) > 0 && streamResp.Choices[0].Delta.Content != "" {
					_, _ = writer.Write([]byte(streamResp.Choices[0].Delta.Content))
				}
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func parseSSE(data string) []string {
	var lines []string
	current := ""
	
	for _, char := range data {
		if char == '\n' {
			if current != "" {
				if len(current) > 6 && current[:6] == "data: " {
					lines = append(lines, current[6:])
				}
				current = ""
			}
		} else {
			current += string(char)
		}
	}
	if current != "" && len(current) > 6 && current[:6] == "data: " {
		lines = append(lines, current[6:])
	}
	
	return lines
}

