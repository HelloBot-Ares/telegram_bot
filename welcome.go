package main

import (
	"fmt"

	"github.com/HelloBot-Ares/telegram_bot/actions"
	"github.com/HelloBot-Ares/telegram_bot/api"
	tg "gopkg.in/telegram-bot-api.v4"
)

func Welcome(u tg.Update, bot *tg.BotAPI) {
	if api.CheckTelegramUserPresence(u.Message.From.ID) {
		msg := tg.NewMessage(u.Message.Chat.ID, fmt.Sprintf("Bentornato %v!", u.Message.From.UserName))
		bot.Send(msg)
		msg = tg.NewMessage(u.Message.Chat.ID, "Come posso aiutarti?")
		msg.ReplyMarkup = actions.MainMenu()
		bot.Send(msg)
		return
	}

	msg := tg.NewMessage(u.Message.Chat.ID, "Benvenuto! Effettua il login se sei già un utente oppure registrati 😏")
	msg.ReplyMarkup = actions.UserSessionsActions()
	bot.Send(msg)
}
