package api

import (
	"log"
	"net/http"
	"fmt"
	"encoding/json"
	"errors"
)

type Event struct {
	ID              int    `json:"id"`
	Subject         string `json:"subject"`
	Topic           string `json:"topic"`
	Owner           string `json:"owner"`
	StartingAt      string `json:"starting_at"`
	PlaceName       string `json:"place_name"`
	PlaceAddress    string `json:"place_address"`
	MaxParticipants int    `json:"max_participants"`
	Participants    int    `json:"participants"`
}

func GetUserEvents(telegramID int) ([]*Event, error) {
	c := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf(Host + "/api/users/%v/events", telegramID), nil)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res, err := c.Do(req)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	if res.StatusCode == http.StatusOK {
		u := &User{}
		err = json.NewDecoder(res.Body).Decode(u)
		log.Println("user response", u.ID)
		return u.ID, nil
	}
	return 0, errors.New("Login was unsuccessful")
	return nil, nil
}
