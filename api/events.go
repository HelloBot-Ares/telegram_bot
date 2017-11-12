package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Event struct {
	ID              int     `json:"id"`
	Subject         string  `json:"subject"`
	Topic           *Topic  `json:"topic"`
	Owner           *User   `json:"owner"`
	StartingAt      string  `json:"starting_at_formatted"`
	PlaceName       string  `json:"place_name"`
	PlaceAddress    string  `json:"place_address"`
	MaxParticipants int     `json:"max_participants"`
	Participants    []*User `json:"participants"`
}

type Topic struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

func GetUserEvents(telegramID int) ([]Event, error) {
	c := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf(Host+"/api/users/%v/events", telegramID), nil)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	q := req.URL.Query()
	q.Add("from_telegram", "wazzabanna")
	req.URL.RawQuery = q.Encode()

	res, err := c.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if res.StatusCode == http.StatusOK {
		events := make([]Event, 0)
		err = json.NewDecoder(res.Body).Decode(&events)
		fmt.Println(events)
		return events, nil
	}
	return nil, errors.New("Events fetch was unsuccessful")
}

func CreateEvent(subject, placeName, topicID string, telegramID int) error {
	c := &http.Client{}
	jsonString := fmt.Sprintf(`
		{
			"event": {
				"subject": "%v",
				"place_name": "%v",
				"topic_id": "%v",
				"telegram_id": "%v"
			}
		}
	`, subject, placeName, topicID, telegramID)
	req, err := http.NewRequest("POST", Host+"/api/telegram_events", strings.NewReader(jsonString))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Println(err)
		return err
	}

	res, err := c.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}
	if res.StatusCode == http.StatusOK {
		events := make([]Event, 0)
		//bodyString, _ := ioutil.ReadAll(res.Body)
		//fmt.Println(string(bodyString))
		err = json.NewDecoder(res.Body).Decode(&events)
		fmt.Println(events)
		return nil
	}
	return errors.New("Problema con creazione evento")
}

func JoinEvent(eventID string, telegramID int) (*Event, error) {
	c := &http.Client{}
	jsonString := fmt.Sprintf(`
		{
			"telegram_id": "%v"
		}
	`, telegramID)
	req, err := http.NewRequest("POST", Host+"/api/events/" + eventID + "/telegram_participants", strings.NewReader(jsonString))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res, err := c.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if res.StatusCode == http.StatusOK {
		event := &Event{}
		//bodyString, _ := ioutil.ReadAll(res.Body)
		//fmt.Println(string(bodyString))
		err = json.NewDecoder(res.Body).Decode(event)
		fmt.Println(event)
		return event, nil
	}
	return nil, errors.New("Problema con join evento")
}
