package commands

import (
	"strconv"

	"github.com/MagonxESP/MagoBot/internal/infraestructure/helpers"
	"github.com/MagonxESP/MagoBot/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func RollCommandHandler(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	max := 100

	if arg, err := telegram.GetCommandArgument(update, 0); err == nil {
		if maxNumber, err := strconv.Atoi(arg); err == nil {
			max = maxNumber
		}
	}

	telegram.SendTextMessage(bot, update.Message.Chat.ID, strconv.Itoa(helpers.RandomInt(0, max)))
}
