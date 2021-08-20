package examples

import (
	"context"
	"log"

	"github.com/trycourier/courier-go/v2"
)

func getNotificationDraftContent() {
	client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)

	data, err := client.GetNotificationDraftContent(context.Background(), "my-notification")

	if err != nil {
		log.Fatalln(err)
	}
	log.Println(data)
}
