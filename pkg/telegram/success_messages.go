package telegram

import (
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendOkSticker(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	file := tgbotapi.FileID("CAACAgQAAxkBAAEbh91jr2nYghzwG3k8CqEwq8wbT8T44wAC7wIAAtkjZCEHw4gOZ10OCi0E")
	if _, err := bot.Send(tgbotapi.NewSticker(update.Message.Chat.ID, file)); err != nil {
		slog.Warn("failed sending ok sticker", "error", err)
	}
}
