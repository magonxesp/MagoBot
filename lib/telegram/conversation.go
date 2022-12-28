package telegram

import (
	"encoding/json"
	"fmt"
	"github.com/MagonxESP/MagoBot/utils"
	"github.com/google/uuid"
	"time"
)

type Conversation struct {
	Id     string `json:"id"`
	Key    string `json:"key"`
	ChatId int64  `json:"chat_id"`
	UserId int    `json:"user_id"`
	Step   int    `json:"step"`
}

type ConversationKey string

func NewConversation(key string, chatId int64, userId int) *Conversation {
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

	client := utils.CreateRedisClient()
	encoded, err := client.Get(utils.RedisContext, string(key)).Result()

	if err != nil {
		return nil, err
	}

	if encoded == "" {
		return nil, nil
	}

	err = json.Unmarshal([]byte(encoded), &conversation)

	if err != nil {
		return nil, err
	}

	return &conversation, nil
}

func GetConversationKey(chatId int64, userId int) ConversationKey {
	return ConversationKey(fmt.Sprintf("conversation_%d_%d", chatId, userId))
}

func (c *Conversation) NextStep() {
	c.Step += 1
}

func (c *Conversation) Save() error {
	encoded, err := json.Marshal(c)

	if err != nil {
		return err
	}

	client := utils.CreateRedisClient()
	err = client.Set(utils.RedisContext, string(GetConversationKey(c.ChatId, c.UserId)), encoded, 7*24*time.Hour).Err()

	if err != nil {
		return err
	}

	return nil
}
