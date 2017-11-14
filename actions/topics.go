package actions

import (
	"fmt"

	"github.com/HelloBot-Ares/telegram_bot/api"
	"gopkg.in/telegram-bot-api.v4"
)

func TopicKeyboard(position int, topics []api.Topic) interface{} {
	visibleTopics := topics[position : position+5]
	buttons := []tgbotapi.KeyboardButton{}
	for _, t := range visibleTopics {
		buttons = append(buttons, tgbotapi.NewKeyboardButton(fmt.Sprintf(`/%v %v`, t.ID, t.Name)))
	}

	buttons = append(buttons, tgbotapi.NewKeyboardButton("/next"))
	row1 := tgbotapi.NewKeyboardButtonRow(buttons[0:2]...)
	row2 := tgbotapi.NewKeyboardButtonRow(buttons[2:4]...)
	row3 := tgbotapi.NewKeyboardButtonRow(buttons[4:6]...)

	return tgbotapi.NewReplyKeyboard(row1, row2, row3)
}
