package api

import (
	"log"
	"net/http"
	"strings"
	"fmt"
	"github.com/pkg/errors"
	"encoding/json"
)

func LoginUser(username, password string) (int, error){
	c := &http.Client{}
	jsonString := fmt.Sprintf(`{"user": {"username": "%v", "password": "%v"}}`, username, password)
	req, err := http.NewRequest("POST", Host + "/api/signin", strings.NewReader(jsonString))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Println(err)
		return 0, err
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
}

func SetTelegramIDForUser(telegramID, databaseID int) error {
	c := &http.Client{}
	jsonString := fmt.Sprintf(`{"telegram_id": "%v"}`, telegramID)
	req, err := http.NewRequest("POST", fmt.Sprintf("%v/api/users/%v/set_telegram_id", Host, databaseID), strings.NewReader(jsonString))
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
		return nil
	}
	return errors.New("TelegramID update was unsuccessful")
}

type User struct {
	ID int `json:"id"`
}

type UserResponse struct {
	User *User `json:"user"`
}
