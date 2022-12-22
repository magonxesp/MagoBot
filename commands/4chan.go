package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/moshee/go-4chan-api/api"
)

func ChanRandomWThread(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	posts := api.GetThreads("w")
}

func ChanRandomBThread(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {

}

func ChanRandomHentaiThread(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {

}

func ChanRandomEcchiThread(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {

}
