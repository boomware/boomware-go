package main

import (
	"github.com/boomware/boomware-go"
	"log"
	"os"
)

func main() {

	// Get your token here
	// To generate and get token here
	// Visit https://console.boomware.com
	token := os.Getenv("BW_EXAMPLE_TOKEN")

	// Create a default client
	client := boomware.New(token)

	// Make a push request
	request := new(boomware.MessagingPushRequest)
	// Required param number in e164 format https://en.wikipedia.org/wiki/E.164
	request.Number = "79052273940"

	// Optional params
	request.Text = "hi!"
	request.Title = "Game request"

	// Custom configuration for push notification ios optional
	request.IOS = map[string]interface{}{
		"apn": map[string]interface{}{
			"alert": map[string]interface{}{
				"title":          "Game Request",
				"body":           "Bob wants to play poker",
				"action-loc-key": "PLAY",
			},
			"badge": 5,
		},
		"acme1": "bar",
		"acme2": []string{"bang", "whiz"},
	}

	// Sent push
	response := client.MessagingPush(request)
	if response.Error != nil {
		log.Println("sending push error", response.Error)
	} else {
		log.Println("successfully send push, id:", response.ID)
	}
}
