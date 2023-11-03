package examples

import (
	"context"
	"log"

	"github.com/trycourier/courier-go/v3"
)

func getNotificationContent() {
	client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)

	data, err := client.GetNotificationContent(context.Background(), "my-notification")

	if err != nil {
		log.Fatalln(err)
	}
	log.Println(data)
}
