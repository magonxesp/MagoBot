package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

type CommandHandler func(bot *tgbotapi.BotAPI, update *tgbotapi.Update)

var (
	handlers = map[string]CommandHandler{
		"start":                   StartCommandHandler,
		"roll":                    StartCommandHandler,
		"rule34":                  StartCommandHandler,
		"4chanrandomwthread":      ChanRandomWThread,
		"4chanrandombthread":      ChanRandomBThread,
		"4chanrandomhentaithread": ChanRandomHentaiThread,
		"4chanrandomecchithread":  ChanRandomEcchiThread,
	}
)

// HandleCommand Handles an incoming command it returns true if a command is handled or false if not
func HandleCommand(bot *tgbotapi.BotAPI, update *tgbotapi.Update) bool {
	if !update.Message.IsCommand() {
		return false
	}

	handler, exists := handlers[update.Message.Command()]

	if exists {
		handler(bot, update)
		return true
	}

	return false
}
