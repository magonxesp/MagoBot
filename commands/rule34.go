package commands

import (
	"fmt"
	"github.com/MagonxESP/MagoBot/lib/booru"
	"github.com/MagonxESP/MagoBot/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strings"
)

func Rule34CommandHandler(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	request := booru.NewPostListRequest(booru.Rule34, utils.GetCommandArguments(update))
	request.Page = utils.RandomInt(1, 100)
	posts, err := booru.GetPostList(request)

	if err != nil {
		log.Println(err)
		return
	}

	if len(posts) == 0 {
		request.Page = 1
		posts, err = booru.GetPostList(request)

		if err != nil {
			log.Println(err)
			return
		}
	}

	if len(posts) == 0 {
		_, err = bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf(
			"No se han encontrado resultados para el tag(s) %s",
			strings.Join(request.Tags, ", "),
		)))

		if err != nil {
			log.Println(err)
		}

		return
	}

	post := posts[utils.RandomInt(0, len(posts)-1)]
	bytes, err := utils.GetFileContentFromUrl(post.FileUrl)

	if err != nil {
		_, err = bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Ha ocurrido un error al recuperar la imagen del post aleatorio"))
		log.Println(err)
		return
	}

	_, err = bot.Send(tgbotapi.NewPhotoUpload(update.Message.Chat.ID, tgbotapi.FileBytes{Bytes: bytes}))

	if err != nil {
		log.Println(err)
	}
}
