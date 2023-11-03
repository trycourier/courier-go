package examples

import (
	"context"
	"log"

	"github.com/trycourier/courier-go/v3"
)

func getList() {
	client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)

	data, err := client.GetList(context.Background(), "my-list")

	if err != nil {
		log.Fatalln(err)
	}
	log.Println(data)
}
