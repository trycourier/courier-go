package examples

import (
	"context"
	"log"
	"time"

	"github.com/trycourier/courier-go/v3"
)

func sendIdempotentMessage() {
	client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)
	expiration := time.Now().Add(time.Hour * 24)
	requestID, err := client.SendMessageWithOptions(context.Background(),
		map[string]interface{}{
			"message": map[string]interface{}{
				"to": map[string]string{
					"email": "foo@example.com",
				},
				"content": map[string]string{
					"title": "Welcome!",
					"body":  "Thanks for signing up Harry Potter",
				},
				"routing": map[string]interface{}{
					"method":   "single",
					"channels": []string{"email"},
				},
			},
		},
		"POST",
		courier.WithIdempotencyKey("fake-key-6"),
		courier.WithIdempotencyKeyExpiration(expiration),
	)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(requestID)
}
