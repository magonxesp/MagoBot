package main

import (
	"github.com/MagonxESP/MagoBot/commands"
	"github.com/MagonxESP/MagoBot/conversations"
	"github.com/MagonxESP/MagoBot/internal/infraestructure/helpers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
	}

	helpers.ConnectMongodb()
	defer helpers.DisconnectMongodb()

	helpers.ConnectRedis()
	defer helpers.DisconnectRedis()

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
