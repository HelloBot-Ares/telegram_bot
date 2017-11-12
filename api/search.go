package api

import (
	"log"
	"net/http"
	"fmt"
	"encoding/json"
	"strings"
	"os"
)

func SearchEventsAroundMe() ([]Event, error) {
	c := &http.Client{}
	jsonString := fmt.Sprintf(`{"search": {"from_date": "11/11/2017 15:30", "to_date": "15/11/2017 15:30", "topic_id": "%v", "location": "Torino"}}`, os.Getenv("TOPIC_ID"))
	req, err := http.NewRequest("POST", Host+"/api/search", strings.NewReader(jsonString))
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
		events := make([]Event, 0)
		//bodyString, _ := ioutil.ReadAll(res.Body)
		//fmt.Println(string(bodyString))
		err = json.NewDecoder(res.Body).Decode(&events)
		fmt.Println(events)
		return events, nil
	}
	return nil, nil
}
