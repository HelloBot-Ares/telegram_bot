package actions

import (
	tg "gopkg.in/telegram-bot-api.v4"
)

func MainMenu() interface{} {
	createButton := tg.NewKeyboardButton("Crea Gruppo")
	myGroupsButton := tg.NewKeyboardButton("I Miei Gruppi")
	searchButton := tg.NewKeyboardButton("Gruppi Vicino A Me")
	markup := tg.NewReplyKeyboard(
		tg.NewKeyboardButtonRow(createButton),
		tg.NewKeyboardButtonRow(myGroupsButton),
		tg.NewKeyboardButtonRow(searchButton),
	)
	markup.Selective = true
	return markup
}