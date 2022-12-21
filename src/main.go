package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))

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
		log.Print(update.Message.Text)
	}
}
