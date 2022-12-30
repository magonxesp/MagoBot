package conversations

import (
	"fmt"
	"github.com/MagonxESP/MagoBot/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"regexp"
)

func DropConversationHandler(conversation *telegram.Conversation, bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	steps := map[int]telegram.ConversationStepHandler{
		0: DropConversationStep0,
	}

	conversation.HandleConversationStep(steps, bot, update)
}

func DropConversationStep0(conversation *telegram.Conversation, bot *tgbotapi.BotAPI, update *tgbotapi.Update) error {
	url := update.Message.Text
	regex, err := regexp.Compile("^(https?:\\/\\/)([a-zA-Z0-9.]+)(\\/.*)$")

	if !regex.Match([]byte(url)) {
		message := fmt.Sprintf("%s no es una url valida", url)
		if _, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, message)); err != nil {
			log.Println(err)
		}

		return fmt.Errorf("message %s is an invalid url", url)
	}

	if err = conversation.SetState("url", url); err != nil {
		return err
	}

	if _, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "En que bucket lo quieres guardar?")); err != nil {
		log.Println(err)
	}

	return nil
}

func DropConversationStep1(conversation *telegram.Conversation, bot *tgbotapi.BotAPI, update *tgbotapi.Update) error {
	return nil
}
