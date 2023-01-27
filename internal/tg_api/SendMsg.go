package tg_api

import (
	"TESTBOT/api"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SendMessage(text string, ChatId string) (bool, error) {
	// Global variables
	var err error

	var response *http.Response

	// Send the message
	url := fmt.Sprintf("%ssendMessage", api.UrlToken)
	body, _ := json.Marshal(map[string]string{
		"chat_id": ChatId,
		"text":    text,
	})
	response, err = http.Post(
		url,
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return false, err
	}

	// Close the request at the end
	defer response.Body.Close()

	// Body
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return false, err
	}

	// Log
	println("Message '%s' was sent", text)
	println("Response JSON: %s", string(body))

	// Return
	return true, nil
}
