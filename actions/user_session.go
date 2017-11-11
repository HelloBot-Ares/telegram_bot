package actions
import (
	tg "gopkg.in/telegram-bot-api.v4"
)

func UserSessionsActions() interface{} {
	loginButton := tg.NewKeyboardButton("/login")
	signupButton := tg.NewKeyboardButton("/registrati")
	buttonsRow := tg.NewKeyboardButtonRow(loginButton, signupButton)
	markup := tg.NewReplyKeyboard(buttonsRow)
	markup.Selective = true
	return markup
}
