package commands

import (
	"fmt"
	"github.com/MagonxESP/MagoBot/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func StartCommandHandler(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	message := tgbotapi.NewMessage(
		update.Message.Chat.ID,
		fmt.Sprintf("%s Hola k ase", utils.MentionUserMd(*update.Message.From)),
	)

	message.ParseMode = "markdown"

	_, err := bot.Send(message)

	if err != nil {
		log.Fatal(err)
	}
}
