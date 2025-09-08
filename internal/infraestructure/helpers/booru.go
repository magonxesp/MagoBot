package helpers

import (
	"log/slog"
	"os"
)

func GetRule34ApiKey() string {
	apiKey := os.Getenv("MAGOBOT_RULE34_API_KEY")

	if apiKey == "" {
		slog.Warn("MAGOBOT_RULE34_API_KEY env variable is missing")
	}

	return apiKey
}
