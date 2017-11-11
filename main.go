package main

import (
	"log"
	tg "gopkg.in/telegram-bot-api.v4"
	"github.com/joho/godotenv"
	"os"
	"github.com/rentziass/ares/telegram_bot/api"
	"fmt"
)

type LoginProcess struct {
	Started bool
	Username string
	UsernameInserted bool
	Password string
	PasswordInserted bool
}

var loginProcesses = map[int]*LoginProcess{}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	api.Host = os.Getenv("API_HOST")

	tgToken := os.Getenv("TG_TOKEN")
	bot, err := tg.NewBotAPI(tgToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tg.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Text == "/start" {
			Welcome(update, bot)
			continue
		}

		if update.Message.Text == "/login" {
			loginProcesses[update.Message.From.ID] = &LoginProcess{Started: true}
			msg := tg.NewMessage(update.Message.Chat.ID, "Please insert your username")
			bot.Send(msg)
			continue
		}

		// If a login process is ongoing for this user see where he's at
		if p, ok := loginProcesses[update.Message.From.ID]; ok {
			if p.UsernameInserted {
				// This is a password
				p.Password = update.Message.Text
				msg := tg.NewMessage(update.Message.Chat.ID, "Attempting to login 💫")
				bot.Send(msg)
				fmt.Println("something")
				LoginUser(p.Username, p.Password, update, bot)

				// do some actual login stuff
				continue
			}
			// Otherwise it's a username
			p.Username = update.Message.Text
			p.UsernameInserted = true
			msg := tg.NewMessage(update.Message.Chat.ID, "Username: " + loginProcesses[update.Message.From.ID].Username + ". Please insert your password. Don't forget to delete that message from this chat 🙌🏼")
			bot.Send(msg)
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tg.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
