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

	// Make a sms request
	request := new(boomware.SMSRequest)
	// Required param number in e164 format https://en.wikipedia.org/wiki/E.164
	request.Number = "79052273940"
	// Required param text
	request.Text = "hi!"

	// Sent sms
	response := client.SMS(request)
	if response.Error != nil {
		log.Println("sending sms error", response.Error)
	} else {
		log.Println("successfully send sms, id:", response.ID)
	}
}
