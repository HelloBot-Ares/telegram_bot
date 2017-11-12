package main

import (
	tg "gopkg.in/telegram-bot-api.v4"
	"github.com/rentziass/ares/telegram_bot/api"
	"fmt"
)

func SearchEvents(u tg.Update, bot *tg.BotAPI) {
	events, err := api.SearchEventsAroundMe()
	if err != nil {
		msg := tg.NewMessage(u.Message.Chat.ID, "Si è verificato un problema. 😭")
		bot.Send(msg)
		return
	}

	msgText := ""
	if len(events) == 0 {
		msgText = "Non ci sono gruppi intorno a te al momento, vuoi crearne uno?"
	} else {
		msgText = msgText + `<b>Risultati per: Torino</b>
		`
		for _, e := range events {
			msgText = msgText + fmt.Sprintf(searchTemplate, e.ID, e.Subject, e.Topic.Name, e.Owner.Username, e.PlaceName, e.StartingAt, e.PlaceAddress, len(e.Participants), e.MaxParticipants)
		}
	}


	fmt.Println(msgText)

	msg := tg.NewMessage(u.Message.Chat.ID, msgText)
	msg.ParseMode = tg.ModeHTML
	//msg.Text = msgText
	bot.Send(msg)
}

var searchTemplate = `/%v <b>%v</b>
<i>%v</i>
• Organizzatore: %v
• Dove: %v
• Quando: %v
• Indirizzo: %v
• Partecipanti: %v
• Max Partecipanti: %v

`
