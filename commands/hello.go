package commands

import (
	"fmt"
	"log"

	"github.com/MagonxESP/MagoBot/internal/infraestructure/helpers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartCommandHandler(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	message := tgbotapi.NewMessage(
		update.Message.Chat.ID,
		fmt.Sprintf("%s Hola k ase", helpers.MentionUserMd(*update.Message.From)),
	)

	message.ParseMode = "markdown"

	_, err := bot.Send(message)

	if err != nil {
		log.Fatal(err)
	}
}
