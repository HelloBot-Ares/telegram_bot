package api

import (
	"log"
	"net/http"
	"encoding/json"
)

func GetTopics() ([]Topic, error) {
	c := &http.Client{}
	req, err := http.NewRequest("GET", Host+"/api/topics", nil)
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
		topics := make([]Topic, 0)
		err = json.NewDecoder(res.Body).Decode(&topics)
		return topics, nil
	}
	return nil, nil
}
