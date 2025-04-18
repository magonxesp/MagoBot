package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"golang.org/x/exp/slog"
)

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
	"deletedropperconfig":     DeleteDropperConfigCommandHandler,
}

// HandleCommand Handles an incoming command it returns true if a command is handled or false if not
func HandleCommand(bot *tgbotapi.BotAPI, update *tgbotapi.Update) bool {
	message := update.Message.Text
	slog.Debug("handling command", "message", message)

	if !update.Message.IsCommand() {
		slog.Debug("message is not a command", "message", message)
		return false
	}

	handler, exists := handlers[update.Message.Command()]

	if exists {
		slog.Debug("executing command", "message", message)
		handler(bot, update)
		return true
	}

	slog.Debug("command not exists", "message", message)
	return false
}
