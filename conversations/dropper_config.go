package conversations

import (
	"errors"
	"regexp"
	"strings"

	"github.com/MagonxESP/MagoBot/internal/application"
	"github.com/MagonxESP/MagoBot/internal/domain"
	"github.com/MagonxESP/MagoBot/internal/infraestructure/persistence/mongodb/repository"
	"github.com/MagonxESP/MagoBot/pkg/dropper"
	"github.com/MagonxESP/MagoBot/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
	"golang.org/x/exp/slog"
)

func testDropperConnection(config *domain.DropperConfig) bool {
	client := dropper.NewClient(config.Url, config.ClientId, config.ClientSecret)
	err := client.Authenticate()

	if err != nil {
		slog.Warn("dropper connection failed", "client", config.ClientId, "url", config.Url, "error", err)
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
			telegram.SendTextMessage(bot, update.Message.Chat.ID, "El formato no es valido")
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
		telegram.SendTextMessage(bot, update.Message.Chat.ID, message)
		return errors.New("dropper wrong connection or invalid credentials")
	}

	configurer := application.NewDropperConfigurer(repository.NewMongoDbDropperConfigRepository())
	err := configurer.CreateNewConfig(config)

	if err != nil && errors.Is(err, application.ExistingConfigError) {
		telegram.SendTextMessage(bot, update.Message.Chat.ID, "Ya hay configurado un dropper")
		return nil
	}

	if err != nil {
		telegram.SendTextMessage(bot, update.Message.Chat.ID, "Ha ocurrido un error mientras se guardaba la configuracion, intentalo de nuevo mas tarde")
		slog.Warn("failed saving dropper config", "error", err)
		return err
	}

	telegram.SendOkSticker(bot, update)
	telegram.SendTextMessage(bot, update.Message.Chat.ID, "El dropper se ha configurado exitosamente")
	return nil
}
