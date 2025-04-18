package commands

import (
	"errors"

	"github.com/MagonxESP/MagoBot/internal/application"
	"github.com/MagonxESP/MagoBot/internal/infraestructure/persistence/mongodb/repository"
	"github.com/MagonxESP/MagoBot/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"golang.org/x/exp/slog"
)

func DropCommandHandler(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	finder := application.NewDropperConfigFinder(repository.NewMongoDbDropperConfigRepository())
	config, err := finder.FindByUserId(update.Message.From.ID)

	if err != nil {
		telegram.SendTextMessage(bot, update.Message.Chat.ID, "Ha ocurrido un error mientras se cargaba la configuracion del dropper")
		return
	}

	if config == nil {
		telegram.SendTextMessage(bot, update.Message.Chat.ID, "No hay ningun dropper configurado, lanza el comando /configdropper para configurarlo")
		return
	}

	conversation := telegram.NewConversation("drop", update.Message.Chat.ID, update.Message.From.ID)
	err = conversation.Save()

	if err != nil {
		telegram.SendCommandErrorMessage(bot, update)
		return
	}

	telegram.SendTextMessage(bot, update.Message.Chat.ID, "Pasame el link que contenga el archivo que quieres guardar")
}

func DropperConfigCommandHandler(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	conversation := telegram.NewConversation("dropper_config", update.Message.Chat.ID, update.Message.From.ID)
	err := conversation.Save()

	if err != nil {
		slog.Warn("failed saving dropper config conversation", "error", err)
		telegram.SendCommandErrorMessage(bot, update)
		return
	}

	text := "Enviame los datos necesarios para poder conectarme a tu dropper con el siguiente formato:\n\n" +
		"url: https://dominio-del-dropper.example\n" +
		"client_id: clientid\n" +
		"client_secret: clientsecret"

	telegram.SendTextMessage(bot, update.Message.Chat.ID, text)
}

func DeleteDropperConfigCommandHandler(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	deleter := application.NewDropperConfigDeleter(repository.NewMongoDbDropperConfigRepository())
	err := deleter.DeleteUserConfig(update.Message.From.ID)

	if err != nil && errors.Is(err, application.NotConfigExistsError) {
		telegram.SendTextMessage(bot, update.Message.Chat.ID, "No tienes ninguna configuracion guardada")
		return
	}

	if err != nil {
		telegram.SendTextMessage(bot, update.Message.Chat.ID, "Ha ocurrido un error al eliminar la configuracion")
		return
	}

	telegram.SendOkSticker(bot, update)
	telegram.SendTextMessage(bot, update.Message.Chat.ID, "Se ha eliminado la configuracion del dropper con exito")
}
