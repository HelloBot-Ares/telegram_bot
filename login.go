package main

import (
	tg "gopkg.in/telegram-bot-api.v4"
	"github.com/rentziass/telegram_bot/api"
	"log"
	"github.com/rentziass/telegram_bot/actions"
)

func LoginUser(username, password string, u tg.Update, bot *tg.BotAPI) {
	// If successfull, login returns here the database ID for current user
	databaseID, err := api.LoginUser(username, password)

	// However it goes from here, remove the current user login process, we won't need it again
	delete(loginProcesses, u.Message.From.ID)

	// Reset the process if login was unsuccessful
	if err != nil {
		log.Println("Login was unsuccessful")
		failLogin(u.Message.Chat.ID, bot)
		return
	}


	// If login is successfull we set Telegram ID for current user via API
	// and we won't need to do it again
	err = api.SetTelegramIDForUser(u.Message.From.ID, databaseID)
	if err != nil {
		log.Println("Telegram ID was unsuccessful")
		failLogin(u.Message.Chat.ID, bot)
		return
	}

	msg := tg.NewMessage(u.Message.Chat.ID, "Bentornato " + u.Message.From.UserName + "!")
	msg.ReplyMarkup = tg.NewRemoveKeyboard(true)
	bot.Send(msg)

	msg = tg.NewMessage(u.Message.Chat.ID, "Come posso aiutarti?")
	msg.ReplyMarkup = actions.MainMenu()
	bot.Send(msg)
}

func failLogin(chatID int64, bot *tg.BotAPI) {
	msg := tg.NewMessage(chatID, "Non siamo riusciti a trovarti 😭 Riproviamo?")
	msg.ReplyMarkup = actions.UserSessionsActions()
	bot.Send(msg)
}
