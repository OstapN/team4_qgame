package actions

import (
	//"team4_qgame/betypes"
	"team4_qgame/db"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func ShowRankCommand(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {

	msg := tgbotapi.NewMessage(
		update.Message.Chat.ID, "Your rank is "+update.Message.Chat.FirstName)
	db.Test()
	bot.Send(msg)
}
