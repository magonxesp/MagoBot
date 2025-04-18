package commands

import (
	"github.com/MagonxESP/MagoBot/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func Rule34CommandHandler(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	conversation := telegram.NewConversation("rule34", update.Message.Chat.ID, update.Message.From.ID)
	err := conversation.Save()

	if err != nil {
		telegram.SendCommandErrorMessage(bot, update)
		return
	}

	telegram.SendTextMessage(bot, update.Message.Chat.ID, "De que personaje quieres que busque?")
}
