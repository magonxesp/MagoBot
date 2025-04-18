package telegram

import (
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func SendTextMessage(bot *tgbotapi.BotAPI, chatID int64, text string) {
	_, err := bot.Send(tgbotapi.NewMessage(chatID, text))

	if err != nil {
		slog.Warn("error sending message", "chat_id", chatID, "text", text, "error", err)
	}
}

func SendKeyboard(bot *tgbotapi.BotAPI, chatID int64, text string, buttons []tgbotapi.KeyboardButton) {
	message := tgbotapi.NewMessage(chatID, text)
	message.ReplyMarkup = tgbotapi.NewReplyKeyboard(tgbotapi.NewKeyboardButtonRow(buttons...))

	if _, err := bot.Send(message); err != nil {
		slog.Warn("failed sending keyboard", "chat_id", chatID, "text", text, "error", err)
	}
}

func SendRemoveKeyboard(bot *tgbotapi.BotAPI, chatID int64, text string) {
	message := tgbotapi.NewMessage(chatID, text)
	message.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)

	if _, err := bot.Send(message); err != nil {
		slog.Warn("failed sending remove keyboard", "chat_id", chatID, "text", text, "error", err)
	}
}
