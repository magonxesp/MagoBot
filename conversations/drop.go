package conversations

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/MagonxESP/MagoBot/internal/application"
	"github.com/MagonxESP/MagoBot/internal/infraestructure/persistence/mongodb/repository"
	"github.com/MagonxESP/MagoBot/pkg/dropper"
	"github.com/MagonxESP/MagoBot/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"golang.org/x/exp/slices"
	"golang.org/x/exp/slog"
)

func DropConversationHandler(conversation *telegram.Conversation, bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	steps := map[int]telegram.ConversationStepHandler{
		0: DropConversationStep0,
		1: DropConversationStep1,
	}

	conversation.HandleConversationStep(steps, bot, update)
}

func getUserDropperClient(update *tgbotapi.Update) (*dropper.Client, error) {
	finder := application.NewDropperConfigFinder(repository.NewMongoDbDropperConfigRepository())
	config, err := finder.FindByUserId(update.Message.From.ID)

	if err != nil {
		return nil, err
	}

	if config == nil {
		return nil, errors.New("dropper configuration not found")
	}

	client := dropper.NewClient(config.Url, config.ClientId, config.ClientSecret)

	if err = client.Authenticate(); err != nil {
		return nil, err
	}

	return client, nil
}

func DropConversationStep0(conversation *telegram.Conversation, bot *tgbotapi.BotAPI, update *tgbotapi.Update) error {
	url := update.Message.Text
	regex, err := regexp.Compile("^(https?:\\/\\/)([a-zA-Z0-9.]+)(\\/.*)$")

	if !regex.Match([]byte(url)) {
		telegram.SendTextMessage(bot, update.Message.Chat.ID, fmt.Sprintf("%s no es una url valida", url))
		return fmt.Errorf("message %s is an invalid url", url)
	}

	if err = conversation.SetState("url", url); err != nil {
		return err
	}

	client, err := getUserDropperClient(update)

	if err != nil {
		return err
	}

	buckets, err := client.GetAllBuckets()

	if len(buckets) == 0 {
		telegram.SendTextMessage(bot, update.Message.Chat.ID, "el dropper no tiene buckets configurados")
		return errors.New("missing dropper buckets")
	}

	if err != nil {
		return err
	}

	var buttons []tgbotapi.KeyboardButton
	var bucketNames []string
	for _, bucket := range buckets {
		bucketNames = append(bucketNames, bucket.Name)
		buttons = append(buttons, tgbotapi.NewKeyboardButton(bucket.Name))
	}

	if err = conversation.SetState("available_buckets", strings.Join(bucketNames, ";")); err != nil {
		return err
	}

	telegram.SendKeyboard(bot, update.Message.Chat.ID, "En que bucket lo quieres guardar?", buttons)
	return nil
}

func DropConversationStep1(conversation *telegram.Conversation, bot *tgbotapi.BotAPI, update *tgbotapi.Update) error {
	var url string
	var availableBuckets string
	var err error
	bucketName := update.Message.Text

	if url, err = conversation.GetState("url"); err != nil {
		return err
	}

	if availableBuckets, err = conversation.GetState("available_buckets"); err != nil {
		return err
	}

	if !slices.Contains(strings.Split(availableBuckets, ";"), bucketName) {
		telegram.SendTextMessage(bot, update.Message.Chat.ID, fmt.Sprintf("el bucket %s no existe", bucketName))
		return errors.New("wrong bucket name")
	}

	client, err := getUserDropperClient(update)

	if err != nil {
		return err
	}

	telegram.SendRemoveKeyboard(
		bot,
		update.Message.Chat.ID,
		fmt.Sprintf("Enviando peticion para descargar el contenido de la url %s al dropper", url),
	)

	go func() {
		if err := client.Drop(url, &dropper.Bucket{Name: bucketName}); err != nil {
			telegram.SendTextMessage(
				bot,
				update.Message.Chat.ID,
				fmt.Sprintf("Ha ocurrido un error mientras se estaba descargando el contenido de la url %s", url),
			)
			slog.Warn("failed drop url on dropper bucket", "bucket", bucketName, "error", err)
			return
		}

		telegram.SendTextMessage(
			bot,
			update.Message.Chat.ID,
			fmt.Sprintf("El contenido de la url %s se ha descargado con exito en el bucket %s", url, bucketName),
		)
	}()

	return nil
}
