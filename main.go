package main

import (
	"log"
	tg "gopkg.in/telegram-bot-api.v4"
	"github.com/joho/godotenv"
	"os"
	"github.com/rentziass/ares/telegram_bot/api"
	"github.com/rentziass/ares/telegram_bot/actions"
	"strings"
)

type LoginProcess struct {
	Started bool
	Username string
	UsernameInserted bool
	Password string
	PasswordInserted bool
}

type EventCreationProcess struct {
	TopicID string
	SeenTopics int
	TopicSet bool
	Subject string
	SubjectSet bool
	PlaceName string
}

var loginProcesses = map[int]*LoginProcess{}
var eventCreationProcesses = map[int]*EventCreationProcess{}
var eventJoiningProcesses = map[int]bool{}

var Topics []api.Topic

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	api.Host = os.Getenv("API_HOST")

	Topics, err = api.GetTopics()
	if err != nil {
		log.Fatal(err)
	}

	tgToken := os.Getenv("TG_TOKEN")
	bot, err := tg.NewBotAPI(tgToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

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
			msg := tg.NewMessage(update.Message.Chat.ID, "Per favore, inserisci il tuo username")
			bot.Send(msg)
			continue
		}

		if update.Message.Text == "I Miei Gruppi" {
			FetchUserEvents(update, bot)
			continue
		}

		if update.Message.Text == "Gruppi Vicino A Me" {
			SearchEvents(update, bot)
			eventJoiningProcesses[update.Message.From.ID] = true
			continue
		}

		if update.Message.Text == "Crea Gruppo" {
			msg := tg.NewMessage(update.Message.Chat.ID, "Seleziona la materia del tuo nuovo gruppo:")
			eventCreationProcesses[update.Message.From.ID] = &EventCreationProcess{}
			msg.ReplyMarkup = actions.TopicKeyboard(0, Topics)
			bot.Send(msg)
			continue
		}

		// If a login process is ongoing for this user see where he's at
		if p, ok := loginProcesses[update.Message.From.ID]; ok {
			if p.UsernameInserted {
				// This is a password
				p.Password = update.Message.Text
				msg := tg.NewMessage(update.Message.Chat.ID, "Accedo 💫")
				bot.Send(msg)
				LoginUser(p.Username, p.Password, update, bot)

				// do some actual login stuff
				continue
			}
			// Otherwise it's a username
			p.Username = update.Message.Text
			p.UsernameInserted = true
			msg := tg.NewMessage(update.Message.Chat.ID, "Username: " + loginProcesses[update.Message.From.ID].Username + ". Per favore, inserisci ora la tua password. Non dimenticarti di cancellare quel messaggio dalla chat 🙌🏼")
			bot.Send(msg)
			continue
		}

		// Check if the user is in the process of creating a group
		if p, ok := eventCreationProcesses[update.Message.From.ID]; ok {
			if p.TopicSet {
				if p.SubjectSet {
					p.PlaceName = update.Message.Text

					err := api.CreateEvent(p.Subject, p.PlaceName, p.TopicID, update.Message.From.ID)
					delete(eventCreationProcesses, update.Message.From.ID)
					if err != nil {
						msg := tg.NewMessage(update.Message.Chat.ID, "C'è stato un problema nella creazione del gruppo 😅")
						msg.ReplyMarkup = actions.MainMenu()
						bot.Send(msg)
						continue
					}
					msg := tg.NewMessage(update.Message.Chat.ID, "Gruppo creato! 💫")
					msg.ReplyMarkup = actions.MainMenu()
					bot.Send(msg)
					continue
				} else {
					p.Subject = update.Message.Text
					p.SubjectSet = true
					msg := tg.NewMessage(update.Message.Chat.ID, "Dove vuoi organizzarlo?")
					bot.Send(msg)
				}
			} else {
				if update.Message.Text == "/next" {
					p.SeenTopics = p.SeenTopics + 6
					msg := tg.NewMessage(update.Message.Chat.ID, "Seleziona la materia del tuo nuovo gruppo:")
					msg.ReplyMarkup = actions.TopicKeyboard(p.SeenTopics, Topics)
					bot.Send(msg)
				} else {
					parts := strings.Split(update.Message.Text, " ")
					id := strings.Replace(parts[0], "/", "", -1)
					p.TopicID = id
					p.TopicSet = true
					msg := tg.NewMessage(update.Message.Chat.ID, "Specifica un argomento per il tuo gruppo:")
					msg.ReplyMarkup = tg.NewRemoveKeyboard(true)
					bot.Send(msg)
				}
			}
			continue
		}

		// Check if the user may be choosing an event to join
		if _, ok := eventJoiningProcesses[update.Message.From.ID]; ok && strings.Contains(update.Message.Text, "/") {
			id := strings.Replace(update.Message.Text, "/", "", -1)
			event, err := api.JoinEvent(id, update.Message.From.ID)
			delete(eventJoiningProcesses, update.Message.From.ID)
			if err != nil {
				msg := tg.NewMessage(update.Message.Chat.ID, "Si è verificato un problema 😅")
				bot.Send(msg)
				continue
			}

			msg := tg.NewMessage(update.Message.Chat.ID, "Congratulazioni! Ti sei unito a " + event.Subject + "!")
			bot.Send(msg)
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		//msg := tg.NewMessage(update.Message.Chat.ID, update.Message.Text)
		//msg.ReplyToMessageID = update.Message.MessageID
		//
		//bot.Send(msg)
	}
}
