package utils

import (
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strings"
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

func sendErrorSticker(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	panicStickers := []string{
		"CAACAgUAAxkBAAEbc9BjrFtNNmatYllsBTgzxkpc6uzWZQAC4gYAAgy_QFYYyoo8nxJpyCwE",
		"CAACAgEAAxkBAAEbc9ZjrFt-VHiPfaBWc-vmGxeF8q3fcAACMAADqv30J5qjo-YOA02iLAQ",
		"CAACAgIAAxkBAAEbc9hjrFuZcQi0FdHxooTSPI0cu9em-QACWxAAApBi8UgSt3M3QMXnoSwE",
		"CAACAgQAAxkBAAEbc9pjrFu1bBDvK9b0N8TQIW2X0veN0wACxgIAAkXeuwVOmYfpKSuFWSwE",
		"CAACAgQAAxkBAAEbc9xjrFvMvAtxiyK2iE8bBG0FxV19aQAC9QIAAtkjZCHhBEoAAfVdab8sBA",
		"CAACAgQAAxkBAAEbc95jrFvUrMcdgd7v1M1fFMvQCOAiggACBQMAAtkjZCElAfMvpD7j6iwE",
	}

	sticker := panicStickers[RandomInt(0, len(panicStickers)-1)]
	_, err := bot.Send(tgbotapi.NewStickerShare(update.Message.Chat.ID, sticker))

	if err != nil {
		log.Println(err)
	}
}

func SendCommandErrorMessage(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	sendErrorSticker(bot, update)
	_, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Ha ocurrido un error durante la ejecución del comando, vuelve a intentarlo mas tarde."))

	if err != nil {
		log.Println(err)
	}
}
