package examples

import (
	"context"
	"log"

	"github.com/trycourier/courier-go/v2"
)

func sendMessageWithMetadata() {
	client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)

	requestID, err := client.SendMessage(context.Background(), courier.SendMessageRequestBody{
		Message: map[string]interface{}{
			"template": "my-template",
			"to": map[string]string{
				"email": "foo@bar.com",
			},
			"metadata": map[string]string{
				"trace_id": "Wassup?",
			},
		},
	})

	if err != nil {
		log.Fatalln(err)
	}
	log.Println(requestID)
}
