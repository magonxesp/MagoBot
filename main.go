package main

import (
	"log/slog"
	"os"

	"github.com/MagonxESP/MagoBot/commands"
	"github.com/MagonxESP/MagoBot/conversations"
	"github.com/MagonxESP/MagoBot/internal/infraestructure/helpers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		slog.Warn("failed loading .env", "error", err)
	}

	helpers.ConnectMongodb()
	defer helpers.DisconnectMongodb()

	helpers.ConnectRedis()
	defer helpers.DisconnectRedis()

	bot, err := tgbotapi.NewBotAPI(os.Getenv("MAGOBOT_TOKEN"))

	if err != nil {
		slog.Error("failed creating bot instance", "error", err)
		os.Exit(1)
	}

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates, err := bot.GetUpdatesChan(updateConfig)

	if err != nil {
		slog.Error("failed creating updates channel", "error", err)
		os.Exit(1)
	}

	for update := range updates {
		if commands.HandleCommand(bot, &update) {
			continue
		}

		if conversations.HandleConversation(bot, &update) {
			continue
		}
	}
}
