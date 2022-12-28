package conversations

import (
	"github.com/MagonxESP/MagoBot/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type ConversationHandler func(conversation *utils.Conversation, bot *tgbotapi.BotAPI, update *tgbotapi.Update)

var conversationHandler = map[string]ConversationHandler{
	"drop": DropConversationHandler,
}

func HandleConversation(bot *tgbotapi.BotAPI, update *tgbotapi.Update) bool {
	conversation, err := utils.GetExistingConversation(utils.GetConversationKey(update.Message.Chat.ID, update.Message.From.ID))

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
