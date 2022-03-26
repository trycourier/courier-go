package examples

import (
	"context"
	"log"

	"github.com/trycourier/courier-go/v2"
)

func sendMessageIdempotent() {
	client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)

	requestID, err := client.SendIdempotentMessage(context.Background(),
		map[string]interface{}{
			"template": "my-template",
			"to": map[string]string{
				"email": "foo@bar.com",
			},
		},
		"POST",
		"fake-idempotency-key",
	)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(requestID)
}
