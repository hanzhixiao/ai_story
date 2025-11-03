package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port               string
	CorsAllowedOrigins string
	OpenAIAPIKey       string
	OpenAIBaseURL      string
	AnthropicAPIKey    string
	AnthropicBaseURL   string
	DatabasePath       string
}

func LoadConfig() (*Config, error) {
	// 尝试加载.env文件，如果不存在也不报错
	_ = godotenv.Load()

	return &Config{
		Port:               getEnv("PORT", "8080"),
		CorsAllowedOrigins: getEnv("CORS_ALLOWED_ORIGINS", "http://localhost:3000,http://localhost:5173"),
		OpenAIAPIKey:       getEnv("OPENAI_API_KEY", ""),
		OpenAIBaseURL:      getEnv("OPENAI_BASE_URL", "https://api.openai.com/v1"),
		AnthropicAPIKey:    getEnv("ANTHROPIC_API_KEY", ""),
		AnthropicBaseURL:   getEnv("ANTHROPIC_BASE_URL", "https://api.anthropic.com"),
		DatabasePath:       getEnv("DATABASE_PATH", "grandma.db"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
