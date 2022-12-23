package commands

import (
	"fmt"
	"github.com/MagonxESP/MagoBot/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func sendRandomThreadOfBoard(board string, bot *tgbotapi.BotAPI, update *tgbotapi.Update) (tgbotapi.Message, error) {
	thread, err := utils.RandomThreadFromBoard(board)

	if err != nil {
		return bot.Send(tgbotapi.NewMessage(
			update.Message.Chat.ID,
			fmt.Sprintf("Ha ocurrido un error al buscar un thread aleatorio del board %s", board),
		))
	}

	return bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, utils.ThreadUrl(thread)))
}

func sendRandomFileThreadOfBoard(board string, bot *tgbotapi.BotAPI, update *tgbotapi.Update) (tgbotapi.Message, error) {
	thread, err := utils.RandomThreadFromBoard(board)

	if err != nil {
		return bot.Send(tgbotapi.NewMessage(
			update.Message.Chat.ID,
			fmt.Sprintf("Ha ocurrido un error al buscar un thread aleatorio del board %s", board),
		))
	}

	post, err := utils.RandomPostFromThread(thread)

	if err != nil {
		return bot.Send(tgbotapi.NewMessage(
			update.Message.Chat.ID,
			fmt.Sprintf("Ha ocurrido un error al buscar un post aleatorio del board %s en el thread con id %d", board, thread.Id()),
		))
	}

	return bot.Send(tgbotapi.NewPhotoUpload(update.Message.Chat.ID, post.ImageURL()))
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
