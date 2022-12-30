package conversations

import (
	"github.com/MagonxESP/MagoBot/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type ConversationHandler func(conversation *telegram.Conversation, bot *tgbotapi.BotAPI, update *tgbotapi.Update)

var conversationHandler = map[string]ConversationHandler{
	"drop":           DropConversationHandler,
	"dropper_config": DropperConfigConversationHandler,
}

func HandleConversation(bot *tgbotapi.BotAPI, update *tgbotapi.Update) bool {
	conversation, err := telegram.GetExistingConversation(telegram.GetConversationKey(update.Message.Chat.ID, update.Message.From.ID))

	if err != nil {
		log.Println(err)
		return false
	}

	if conversation == nil {
		return false
	}

	handler, ok := conversationHandler[conversation.Key]

	if ok {
		handler(conversation, bot, update)
		return true
	}

	return false
}
