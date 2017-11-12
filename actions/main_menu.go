package actions

import (
	tg "gopkg.in/telegram-bot-api.v4"
)

func MainMenu() interface{} {
	createButton := tg.NewKeyboardButton("/crea gruppo")
	myGroupsButton := tg.NewKeyboardButton("/i miei gruppi")
	searchButton := tg.NewKeyboardButton("/gruppi vicino a me")
	markup := tg.NewReplyKeyboard(
		tg.NewKeyboardButtonRow(createButton),
		tg.NewKeyboardButtonRow(myGroupsButton),
		tg.NewKeyboardButtonRow(searchButton),
	)
	markup.Selective = true
	return markup
}