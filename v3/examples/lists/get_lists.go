package examples

import (
	"context"
	"log"

	"github.com/trycourier/courier-go/v3"
)

func getLists() {
	client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)

	data, err := client.GetLists(context.Background(),
		"", // optional cursor
		"", // optional pattern
	)

	if err != nil {
		log.Fatalln(err)
	}
	log.Println(data)
}
