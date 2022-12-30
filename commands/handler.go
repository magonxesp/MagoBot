package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

type CommandHandler func(bot *tgbotapi.BotAPI, update *tgbotapi.Update)

var handlers = map[string]CommandHandler{
	"start":                   StartCommandHandler,
	"roll":                    RollCommandHandler,
	"rule34":                  Rule34CommandHandler,
	"4chanrandomwthread":      ChanRandomWThreadCommandHandler,
	"4chanrandombthread":      ChanRandomBThreadCommandHandler,
	"4chanrandomhentaithread": ChanRandomHentaiThreadCommandHandler,
	"4chanrandomecchithread":  ChanRandomEcchiThreadCommandHandler,
	"drop":                    DropCommandHandler,
	"configdropper":           DropperConfigCommandHandler,
}

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
