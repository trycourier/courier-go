package examples

import (
	"context"
	"log"

	"github.com/trycourier/courier-go/v3"
)

func getNotificationSubmissionChecks() {
	client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)

	data, err := client.GetNotificationSubmissionChecks(context.Background(), "my-notification", "my-submission")

	if err != nil {
		log.Fatalln(err)
	}
	log.Println(data)
}
