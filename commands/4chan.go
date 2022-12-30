package commands

import (
	"fmt"
	"github.com/MagonxESP/MagoBot/internal/infraestructure/helpers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func sendRandomThreadOfBoard(board string, bot *tgbotapi.BotAPI, update *tgbotapi.Update) (tgbotapi.Message, error) {
	thread, err := helpers.RandomThreadFromBoard(board)

	if err != nil {
		return bot.Send(tgbotapi.NewMessage(
			update.Message.Chat.ID,
			fmt.Sprintf("Ha ocurrido un error al buscar un thread aleatorio del board %s", board),
		))
	}

	return bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, helpers.ThreadUrl(thread)))
}

func sendRandomFileThreadOfBoard(board string, bot *tgbotapi.BotAPI, update *tgbotapi.Update) (tgbotapi.Message, error) {
	thread, err := helpers.RandomThreadFromBoard(board)

	if err != nil {
		return bot.Send(tgbotapi.NewMessage(
			update.Message.Chat.ID,
			fmt.Sprintf("Ha ocurrido un error al buscar un thread aleatorio del board %s", board),
		))
	}

	post, err := helpers.RandomPostFromThread(thread)

	if err != nil {
		return bot.Send(tgbotapi.NewMessage(
			update.Message.Chat.ID,
			fmt.Sprintf("Ha ocurrido un error al buscar un post aleatorio del board %s en el thread con id %d", board, thread.Id()),
		))
	}

	message, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, helpers.PostUrl(post)))

	if post.ImageURL() != "" {
		message, err = bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, post.ImageURL()))
	}

	return message, err
}

func ChanRandomWThreadCommandHandler(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	if _, err := sendRandomFileThreadOfBoard("w", bot, update); err != nil {
		log.Println(err)
	}
}

func ChanRandomBThreadCommandHandler(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	if _, err := sendRandomThreadOfBoard("b", bot, update); err != nil {
		log.Println(err)
	}
}

func ChanRandomHentaiThreadCommandHandler(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	if _, err := sendRandomFileThreadOfBoard("h", bot, update); err != nil {
		log.Println(err)
	}
}

func ChanRandomEcchiThreadCommandHandler(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	if _, err := sendRandomFileThreadOfBoard("e", bot, update); err != nil {
		log.Println(err)
	}
}
