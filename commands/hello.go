package commands

import (
	"fmt"
	"github.com/MagonxESP/MagoBot/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func StartCommandHandler(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	_, err := bot.Send(tgbotapi.NewMessage(
		update.Message.Chat.ID,
		fmt.Sprintf("%s Hola k ase", utils.MentionUserMd(*update.Message.From)),
	))

	if err != nil {
		log.Fatal(err)
	}
}
