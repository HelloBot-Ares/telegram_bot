package main

import (
	"fmt"

	"github.com/HelloBot-Ares/telegram_bot/api"
	tg "gopkg.in/telegram-bot-api.v4"
)

func FetchUserEvents(u tg.Update, bot *tg.BotAPI) {
	events, err := api.GetUserEvents(u.Message.From.ID)
	if err != nil {
		msg := tg.NewMessage(u.Message.Chat.ID, "Si è verificato un problema. 😭")
		bot.Send(msg)
		return
	}

	msgText := ""
	if len(events) == 0 {
		msgText = "Non hai gruppi studio in programma al momento, vuoi crearne uno?"
	}

	for _, e := range events {
		msgText = msgText + fmt.Sprintf(eventTemplate, e.Subject, e.Topic.Name, e.Owner.Username, e.PlaceName, e.StartingAt, e.PlaceAddress, len(e.Participants), e.MaxParticipants)
	}

	fmt.Println(msgText)

	msg := tg.NewMessage(u.Message.Chat.ID, msgText)
	msg.ParseMode = tg.ModeHTML
	//msg.Text = msgText
	bot.Send(msg)
}

var eventTemplate = `<b>%v</b>
<i>%v</i>
• Organizzatore: %v
• Dove: %v
• Quando: %v
• Indirizzo: %v
• Partecipanti: %v
• Max Partecipanti: %v

`
