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
	setupLogger()
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
		slog.Debug("handling message received", "message", update.Message.Text)

		if commands.HandleCommand(bot, &update) {
			slog.Debug("command handled", "message", update.Message.Text)
			continue
		}

		if conversations.HandleConversation(bot, &update) {
			slog.Debug("conversation handled", "message", update.Message.Text)
			continue
		}
	}
}

func setupLogger() {
	logJson := os.Getenv("MAGOBOT_LOG_JSON")
	logLevel := os.Getenv("MAGOBOT_LOG_LEVEL")

	var handler slog.Handler
	var level slog.Level

	if logLevel == "info" {
		level = slog.LevelInfo
	} else if logLevel == "warning" {
		level = slog.LevelWarn
	} else if logLevel == "error" {
		level = slog.LevelError
	} else {
		level = slog.LevelDebug
	}

	if logJson == "true" {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		})

	} else {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		})
	}

	logger := slog.New(handler)
	slog.SetDefault(logger)
}
