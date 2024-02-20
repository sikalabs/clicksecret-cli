package client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func Client() {
	CLICKSECRET_ORIGIN := os.Getenv("CLICKSECRET_ORIGIN")
	CLICKSECRET_API_TOKEN := os.Getenv("CLICKSECRET_API_TOKEN")
	CLICKSECRET_EMAIL := os.Getenv("CLICKSECRET_EMAIL")

	if CLICKSECRET_ORIGIN == "" {
		log.Fatalln("CLICKSECRET_ORIGIN is empty")
	}

	if CLICKSECRET_API_TOKEN == "" {
		log.Fatalln("CLICKSECRET_API_TOKEN is empty")
	}

	if CLICKSECRET_EMAIL == "" {
		log.Fatalln("CLICKSECRET_EMAIL is empty")
	}

	data_request, err := get(CLICKSECRET_ORIGIN + "/api/v1/request?api_token=" + CLICKSECRET_API_TOKEN + "&email=" + CLICKSECRET_EMAIL)
	if err != nil {
		log.Fatalln(err.Error())
	}

	for {
		data_get, err := get(CLICKSECRET_ORIGIN + "/api/v1/get?api_token=" + CLICKSECRET_API_TOKEN + "&session_id=" + data_request["session_id"])
		if err != nil {
			log.Fatalln(err.Error())
		}
		if data_get["secret"] != "" {
			fmt.Println(data_get["secret"])
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func get(url string) (map[string]string, error) {
	// URL of the Echo server endpoint returning JSON
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		if err != nil {
			return nil, err
		}
	}

	// Unmarshal the JSON data into the struct
	var data map[string]string
	err = json.Unmarshal(body, &data)
	if err != nil {
		if err != nil {
			return nil, err
		}
	}

	return data, nil
}
