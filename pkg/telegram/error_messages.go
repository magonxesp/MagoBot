package telegram

import (
	"log/slog"

	"github.com/MagonxESP/MagoBot/internal/infraestructure/helpers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func sendErrorSticker(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	panicStickers := []string{
		"CAACAgUAAxkBAAEbc9BjrFtNNmatYllsBTgzxkpc6uzWZQAC4gYAAgy_QFYYyoo8nxJpyCwE",
		"CAACAgEAAxkBAAEbc9ZjrFt-VHiPfaBWc-vmGxeF8q3fcAACMAADqv30J5qjo-YOA02iLAQ",
		"CAACAgIAAxkBAAEbc9hjrFuZcQi0FdHxooTSPI0cu9em-QACWxAAApBi8UgSt3M3QMXnoSwE",
		"CAACAgQAAxkBAAEbc9pjrFu1bBDvK9b0N8TQIW2X0veN0wACxgIAAkXeuwVOmYfpKSuFWSwE",
		"CAACAgQAAxkBAAEbc9xjrFvMvAtxiyK2iE8bBG0FxV19aQAC9QIAAtkjZCHhBEoAAfVdab8sBA",
		"CAACAgQAAxkBAAEbc95jrFvUrMcdgd7v1M1fFMvQCOAiggACBQMAAtkjZCElAfMvpD7j6iwE",
	}

	sticker := panicStickers[helpers.RandomInt(0, len(panicStickers)-1)]
	file := tgbotapi.FileID(sticker)
	_, err := bot.Send(tgbotapi.NewSticker(update.Message.Chat.ID, file))

	if err != nil {
		slog.Warn("failed sending error sticker", "error", err)
	}
}

func SendCommandErrorMessage(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	sendErrorSticker(bot, update)
	_, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Ha ocurrido un error durante la ejecución del comando, vuelve a intentarlo mas tarde."))

	if err != nil {
		slog.Warn("failed sending command error message", "error", err)
	}
}

func SendConversationNextStepErrorMessage(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	sendErrorSticker(bot, update)
	_, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Ha ocurrido un error al guardar el estado de la conversacion, vuelve a intentarlo mas tarde."))

	if err != nil {
		slog.Warn("failed sending conversation state save error message", "error", err)
	}
}
