package examples

import (
	"context"
	"log"

	"github.com/trycourier/courier-go/v2"
)

func getListSubscriptions() {
	client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)

	data, err := client.GetListSubscriptions(context.Background(), "my-list", "")

	if err != nil {
		log.Fatalln(err)
	}
	log.Println(data)
}
