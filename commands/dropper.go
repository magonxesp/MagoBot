package commands

import (
	"github.com/MagonxESP/MagoBot/internal/application"
	"github.com/MagonxESP/MagoBot/internal/infraestructure/persistence/mongodb/repository"
	"github.com/MagonxESP/MagoBot/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func DropCommandHandler(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	finder := application.NewDropperConfigFinder(repository.NewMongoDbDropperConfigRepository())
	config, err := finder.FindByUserId(update.Message.From.ID)

	if err != nil {
		if _, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Ha ocurrido un error mientras se cargaba la configuracion del dropper")); err != nil {
			log.Println(err)
		}
		return
	}

	if config == nil {
		if _, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "No hay ningun dropper configurado, lanza el comando /configdropper para configurarlo")); err != nil {
			log.Println(err)
		}
		return
	}

	conversation := telegram.NewConversation("drop", update.Message.Chat.ID, update.Message.From.ID)
	err = conversation.Save()

	if err != nil {
		telegram.SendCommandErrorMessage(bot, update)
		return
	}

	_, err = bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Pasame el link que contenga el archivo que quieres guardar"))

	if err != nil {
		log.Println(err)
	}
}

func DropperConfigCommandHandler(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	conversation := telegram.NewConversation("dropper_config", update.Message.Chat.ID, update.Message.From.ID)
	err := conversation.Save()

	if err != nil {
		telegram.SendCommandErrorMessage(bot, update)
		return
	}

	text := "Enviame los datos necesarios para poder conectarme a tu dropper con el siguiente formato:\n\n" +
		"url: https://dominio-del-dropper.example\n" +
		"client_id: clientid\n" +
		"client_secret: clientsecret"

	_, err = bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, text))

	if err != nil {
		log.Println(err)
	}
}
