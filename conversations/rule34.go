package conversations

import (
	"fmt"
	"log/slog"

	"github.com/MagonxESP/MagoBot/internal/domain"
	"github.com/MagonxESP/MagoBot/internal/infraestructure/helpers"
	"github.com/MagonxESP/MagoBot/pkg/booru"
	"github.com/MagonxESP/MagoBot/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Rule34ConversationHandler(conversation *telegram.Conversation, bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	steps := map[int]telegram.ConversationStepHandler{
		0: Rule34ConversationStep0,
	}

	conversation.HandleConversationStep(steps, bot, update)
}

func Rule34ConversationStep0(conversation *telegram.Conversation, bot *tgbotapi.BotAPI, update *tgbotapi.Update) error {
	tag := domain.Slugify(update.Message.Text)
	slog.Debug("searhing rule34 posts", "tag", tag)

	apiKey := helpers.GetRule34ApiKey()
	userId := helpers.GetRule34UserId()
	request := booru.NewRule34PostListRequest(apiKey, userId, []string{tag})
	posts, err := booru.GetPostList(request)

	if err != nil {
		return err
	}

	if posts.Offset > 0 {
		request.Page = helpers.RandomInt(0, posts.Count/posts.Offset)
	}

	posts, err = booru.GetPostList(request)
	if err != nil {
		telegram.SendTextMessage(bot, update.Message.Chat.ID, fmt.Sprintf(
			"Ha ocurrido un error al buscar el rule34 de %s [%s]",
			update.Message.Text,
			tag,
		))

		return err
	}

	if len(posts.Items) == 0 {
		telegram.SendTextMessage(bot, update.Message.Chat.ID, fmt.Sprintf(
			"No se han encontrado resultados para %s [%s]",
			update.Message.Text,
			tag,
		))

		return nil
	}

	post := posts.Items[helpers.RandomInt(0, len(posts.Items)-1)]
	telegram.SendTextMessage(bot, update.Message.Chat.ID, post.FileURL)
	return nil
}
