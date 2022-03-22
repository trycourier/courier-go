package examples

import (
	"context"
	"log"

	"github.com/trycourier/courier-go/v2"
)

func sendMessage() {
	client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)

	requestID, err := client.SendMessage(context.Background(), courier.SendMessageRequestBody{
		Message: map[string]interface{}{
			"template": "my-template",
			"to": map[string]string{
				"email": "foo@bar.com",
			},
			"timeout": int64(3600000), // 1 hour in milliseconds
		},
	})

	if err != nil {
		log.Fatalln(err)
	}
	log.Println(requestID)
}