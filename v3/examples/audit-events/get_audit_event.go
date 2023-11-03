package examples

import (
	"context"
	"log"

	"github.com/trycourier/courier-go/v3"
)

func getAuditEvent() {
	client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)

	data, err := client.GetAuditEvent(context.Background(), "123")

	if err != nil {
		log.Fatalln(err)
	}
	log.Println(data)
}
