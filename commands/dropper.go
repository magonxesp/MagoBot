package commands

import (
	"github.com/MagonxESP/MagoBot/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func DropCommandHandler(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	conversation := telegram.NewConversation("drop", update.Message.Chat.ID, update.Message.From.ID)
	err := conversation.Save()

	if err != nil {
		telegram.SendCommandErrorMessage(bot, update)
		return
	}

	_, err = bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Pasame el link que contenga el archivo que quieres guardar"))

	if err != nil {
		log.Println(err)
	}
}
