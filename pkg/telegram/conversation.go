package telegram

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	"github.com/MagonxESP/MagoBot/internal/infraestructure/helpers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
)

type Conversation struct {
	Id     string `json:"id"`
	Key    string `json:"key"`
	ChatId int64  `json:"chat_id"`
	UserId int64  `json:"user_id"`
	Step   int    `json:"step"`
}

type ConversationKey string
type ConversationStepHandler func(conversation *Conversation, bot *tgbotapi.BotAPI, update *tgbotapi.Update) error

func NewConversation(key string, chatId int64, userId int64) *Conversation {
	return &Conversation{
		Id:     uuid.New().String(),
		Key:    key,
		ChatId: chatId,
		UserId: userId,
		Step:   0,
	}
}

func GetExistingConversation(key ConversationKey) (*Conversation, error) {
	var conversation Conversation

	client := helpers.GetRedisClient()
	encoded, err := client.Get(*helpers.GetRedisContext(), string(key)).Result()
	slog.Debug("get conversation", "content", encoded, "key", key)

	if err != nil {
		slog.Debug("failed getting conversation", "key", key, "error", err)
		return nil, err
	}

	if encoded == "" {
		slog.Debug("conversation has empty content", "key", key)
		return nil, nil
	}

	err = json.Unmarshal([]byte(encoded), &conversation)

	if err != nil {
		slog.Debug("failed deserialize json conversation", "key", key, "error", err)
		return nil, err
	}

	return &conversation, nil
}

func GetConversationKey(chatId int64, userId int64) ConversationKey {
	return ConversationKey(fmt.Sprintf("conversation_%d_%d", chatId, userId))
}

func (c *Conversation) NextStep() {
	c.Step += 1
}

func (c *Conversation) Save() error {
	key := string(GetConversationKey(c.ChatId, c.UserId))
	encoded, err := json.Marshal(c)
	slog.Debug("saving conversation", "content", encoded, "key", key)

	if err != nil {
		return err
	}

	client := helpers.GetRedisClient()
	err = client.Set(*helpers.GetRedisContext(), key, encoded, 7*24*time.Hour).Err()

	if err != nil {
		return err
	}

	return nil
}

func (c *Conversation) conversationValueKey(key string) string {
	return fmt.Sprintf("conversation_%s_%s", c.Id, key)
}

func (c *Conversation) SetState(key string, value string) error {
	client := helpers.GetRedisClient()
	return client.Set(*helpers.GetRedisContext(), c.conversationValueKey(key), value, 7*24*time.Hour).Err()
}

func (c *Conversation) GetState(key string) (string, error) {
	client := helpers.GetRedisClient()
	return client.Get(*helpers.GetRedisContext(), c.conversationValueKey(key)).Result()
}

func (c *Conversation) HandleConversationStep(stepHandler map[int]ConversationStepHandler, bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	handler, ok := stepHandler[c.Step]

	if ok {
		err := handler(c, bot, update)

		if err != nil {
			slog.Warn("conversation step finished with error", "conversation_key", c.Key, "error", err)
			return
		}

		if len(stepHandler) > c.Step {
			c.NextStep()
		}

		if len(stepHandler) == c.Step {
			client := helpers.GetRedisClient()
			if err := client.Del(*helpers.GetRedisContext(), string(GetConversationKey(c.ChatId, c.UserId))).Err(); err != nil {
				slog.Warn("failed removing conversation", "conversation_key", c.Key, "error", err)
			}

			return
		}

		if err = c.Save(); err != nil {
			SendConversationNextStepErrorMessage(bot, update)
		}
	}
}
