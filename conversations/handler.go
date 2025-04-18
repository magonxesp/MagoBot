package conversations

import (
	"github.com/MagonxESP/MagoBot/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"golang.org/x/exp/slog"
)

type ConversationHandler func(conversation *telegram.Conversation, bot *tgbotapi.BotAPI, update *tgbotapi.Update)

var conversationHandler = map[string]ConversationHandler{
	"drop":           DropConversationHandler,
	"dropper_config": DropperConfigConversationHandler,
	"rule34":         Rule34ConversationHandler,
}

func HandleConversation(bot *tgbotapi.BotAPI, update *tgbotapi.Update) bool {
	key := telegram.GetConversationKey(update.Message.Chat.ID, update.Message.From.ID)
	slog.Debug("handling conversation reply", "message", update.Message.Text, "key", key)
	conversation, err := telegram.GetExistingConversation(key)

	if err != nil {
		slog.Warn("failed retrieving conversation", "key", key, "error", err)
		return false
	}

	if conversation == nil {
		slog.Debug("conversation not exists", "message", update.Message.Text, "key", key)
		return false
	}

	handler, ok := conversationHandler[conversation.Key]

	if ok {
		slog.Debug("executing conversation handler", "message", update.Message.Text, "key", key)
		handler(conversation, bot, update)
		return true
	}

	return false
}
