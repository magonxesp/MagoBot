package main

import (
	"github.com/MagonxESP/MagoBot/commands"
	"github.com/MagonxESP/MagoBot/conversations"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("MAGOBOT_TOKEN"))

	if err != nil {
		log.Fatal(err)
		return
	}

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates, err := bot.GetUpdatesChan(updateConfig)

	if err != nil {
		log.Fatal(err)
		return
	}

	for update := range updates {
		if commands.HandleCommand(bot, &update) {
			continue
		}

		if conversations.HandleConversation(bot, &update) {
			continue
		}
	}
}
