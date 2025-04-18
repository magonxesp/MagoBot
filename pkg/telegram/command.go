package telegram

import (
	"errors"
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetCommandArguments(update *tgbotapi.Update) []string {
	return strings.Split(update.Message.CommandArguments(), " ")
}

func GetCommandArgument(update *tgbotapi.Update, position int) (string, error) {
	args := GetCommandArguments(update)

	if len(args) < position {
		return "", errors.New(fmt.Sprintf("The index %d is greater than the arguments count", position))
	}

	return args[position], nil
}
