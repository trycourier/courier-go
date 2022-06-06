package examples

import (
	"context"
	"log"

	"github.com/trycourier/courier-go/v2"
)

func sendMessageTimeout() {
	client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)

	requestID, err := client.SendMessage(context.Background(), courier.SendMessageRequestBody{
		Message: map[string]interface{}{
			"template": "my-template",
			"to": map[string]string{
				"email": "foo@bar.com",
			},
			"timeout": map[string]interface{}{
				"message": int64(3600000), // 1 hour in milliseconds
				"channel": map[string]int64{
					"email": int64(300000),
				},
				"provider": map[string]int64{
					"sendgrid": int64(0), // disable retry
				},
			},
		},
	})

	if err != nil {
		log.Fatalln(err)
	}
	log.Println(requestID)
}
