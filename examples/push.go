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

	// Make simple notification with text and title
	request.Text = "hi!"
	request.Title = "New title awesome!"

	// If you use param `IOS` `Text` and `Title` will be ignored
	// Custom configuration for apns
	// Form more details visit apple doc
	// https://developer.apple.com/library/archive/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/CreatingtheNotificationPayload.html#//apple_ref/doc/uid/TP40008194-CH10-SW1
	// https://developer.apple.com/library/archive/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/PayloadKeyReference.html#//apple_ref/doc/uid/TP40008194-CH17-SW1
	request.IOS = map[string]interface{}{
		"aps": map[string]interface{}{
			"alert": map[string]interface{}{
				"title":          "Game Request",
				"body":           "Bob wants to play poker",
				"action-loc-key": "PLAY",
			},
			"badge": 5,
			"sound": "default",
		},
		"acme1": "bar",
		"acme2": []string{"bang", "whiz"},
	}

	// If you use param `Android` `Text` and `Title` will be ignored
	// Custom message for FCM
	request.Android = &boomware.MessagingPushAndroid{
		// Custom data provided to client
		Data: map[string]interface{}{
			"acme1": "bar",
		},
		Notification: &boomware.AndroidNotification{
			Title: "Game Request",
			Body:  "Bob wants to play poker",
			Sound: "sound",
		},
	}

	// Send push
	response := client.MessagingPush(request)
	if response.Err() != nil {
		log.Println("sending push error", response.Err())
	} else {
		log.Println("successfully sent push, id:", response.ID)
	}
}
