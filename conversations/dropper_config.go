package conversations

import (
	"errors"
	"github.com/MagonxESP/MagoBot/internal/application"
	"github.com/MagonxESP/MagoBot/internal/domain"
	"github.com/MagonxESP/MagoBot/internal/infraestructure/persistence/mongodb/repository"
	"github.com/MagonxESP/MagoBot/pkg/dropper"
	"github.com/MagonxESP/MagoBot/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/google/uuid"
	"log"
	"regexp"
	"strings"
)

func testDropperConnection(config *domain.DropperConfig) bool {
	client := dropper.NewClient(config.Url, config.ClientId, config.ClientSecret)
	err := client.Authenticate()

	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func DropperConfigConversationHandler(conversation *telegram.Conversation, bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	steps := map[int]telegram.ConversationStepHandler{
		0: DropperConfigConversationStep0,
	}

	conversation.HandleConversationStep(steps, bot, update)
}

func DropperConfigConversationStep0(conversation *telegram.Conversation, bot *tgbotapi.BotAPI, update *tgbotapi.Update) error {
	configLines := strings.Split(strings.ReplaceAll(update.Message.Text, " ", ""), "\n")
	configProperties := map[string]string{}
	lineRegex := regexp.MustCompile("^([a-z_]+):(.+)$")

	for _, line := range configLines {
		matches := lineRegex.FindSubmatch([]byte(line))

		if matches == nil {
			if _, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "El formato no es valido")); err != nil {
				log.Println(err)
			}

			return errors.New("invalid config format")
		}

		configProperties[string(matches[1])] = string(matches[2])
	}

	config := &domain.DropperConfig{
		Id:           uuid.New().String(),
		UserId:       update.Message.From.ID,
		Url:          configProperties["url"],
		ClientId:     configProperties["client_id"],
		ClientSecret: configProperties["client_secret"],
	}

	if !testDropperConnection(config) {
		message := "Algo esta mal en la configuracion, no se puede conectar o autenticar con el dropper. " +
			"Revisa los credenciales e intentalo de nuevo"
		if _, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, message)); err != nil {
			log.Println(err)
			return nil
		}

		return errors.New("dropper wrong connection or invalid credentials")
	}

	configurer := application.NewDropperConfigurer(repository.NewMongoDbDropperConfigRepository())
	err := configurer.CreateNewConfig(config)

	if err != nil && errors.Is(err, application.ExistingConfigError) {
		if _, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Ya hay configurado un dropper")); err != nil {
			log.Println(err)
		}

		return nil
	}

	if err != nil {
		if _, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Ha ocurrido un error mientras se guardaba la configuracion, intentalo de nuevo mas tarde")); err != nil {
			log.Println(err)
		}

		return err
	}

	telegram.SendOkSticker(bot, update)

	if _, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "El dropper se ha configurado exitosamente")); err != nil {
		log.Println(err)
	}

	return nil
}
