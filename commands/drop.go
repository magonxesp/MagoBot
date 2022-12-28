package commands

import (
	"github.com/MagonxESP/MagoBot/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func DropCommandHandler(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	conversation := utils.NewConversation("drop", update.Message.Chat.ID, update.Message.From.ID)
	err := conversation.Save()

	if err != nil {
		utils.SendCommandErrorMessage(bot, update)
		return
	}

	_, err = bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Pasame el link que contenga el archivo que quieres guardar"))

	if err != nil {
		log.Println(err)
	}
}
