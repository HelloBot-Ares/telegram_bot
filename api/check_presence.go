package api

import (
	"net/http"
	"fmt"
	"strings"
	"log"
)

func CheckTelegramUserPresence(id int) bool {
	c := &http.Client{}
	jsonString := fmt.Sprintf(`{"telegram_id": "%v"}`, id)
	req, err := http.NewRequest("POST", Host + "/api/users/by_telegram_id", strings.NewReader(jsonString))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Println(err)
		return false
	}

	res, err := c.Do(req)
	if err != nil {
		log.Println(err)
		return false
	}
	if res.StatusCode == http.StatusOK {
		return true
	}
	return false
}
