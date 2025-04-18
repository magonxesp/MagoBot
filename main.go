package main

import (
	"log/slog"
	"os"

	"github.com/MagonxESP/MagoBot/commands"
	"github.com/MagonxESP/MagoBot/conversations"
	"github.com/MagonxESP/MagoBot/internal/infraestructure/helpers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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

	if os.Getenv("MAGOBOT_LOG_LEVEL") == "debug" || os.Getenv("MAGOBOT_LOG_LEVEL") == "" {
		bot.Debug = true
	}

	if err != nil {
		slog.Error("failed creating bot instance", "error", err)
		os.Exit(1)
	}

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		slog.Debug(
			"handling message received",
			"id", update.UpdateID,
			"message", update.Message.Text,
			"chat", update.Message.Chat.ID,
			"type", update.Message.Chat.Type,
		)
		HandleMessage(bot, &update)
	}
}

func HandleMessage(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	if commands.HandleCommand(bot, update) {
		slog.Debug(
			"command handled",
			"id", update.UpdateID,
			"message", update.Message.Text,
			"chat", update.Message.Chat.ID,
			"type", update.Message.Chat.Type,
		)
		return
	}

	if conversations.HandleConversation(bot, update) {
		slog.Debug(
			"conversation handled",
			"id", update.UpdateID,
			"message", update.Message.Text,
			"chat", update.Message.Chat.ID,
			"type", update.Message.Chat.Type,
		)
		return
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
