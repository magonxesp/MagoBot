package utils

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func MentionUserMd(user tgbotapi.User) string {
	return fmt.Sprintf("[@{}](tg://user?id=%d)", user.ID)
}
