package telegram

import (
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func SendOkSticker(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	if _, err := bot.Send(tgbotapi.NewStickerShare(update.Message.Chat.ID, "CAACAgQAAxkBAAEbh91jr2nYghzwG3k8CqEwq8wbT8T44wAC7wIAAtkjZCEHw4gOZ10OCi0E")); err != nil {
		slog.Warn("failed sending ok sticker", "error", err)
	}
}
