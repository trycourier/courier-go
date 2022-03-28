package examples

import (
	"context"
	"log"
	"time"

	"github.com/trycourier/courier-go/v2"
)

func sendIdempotentMessage() {
	client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)
	expiration := time.Now().Add(time.Hour * 24)
	requestID, err := client.SendMessageWithOptions(context.Background(),
		map[string]interface{}{
			"template": "my-template",
			"to": map[string]string{
				"email": "foo@bar.com",
			},
		},
		"POST",
		courier.WithIdempotencyKey("fake-key"),
		courier.WithIdempotencyKeyExpiration(expiration),
	)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(requestID)
}
