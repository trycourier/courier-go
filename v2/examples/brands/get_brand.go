package examples

import (
	"context"
	"log"

	"github.com/trycourier/courier-go/v2"
)

func getBrand() {
	client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)

	data, err := client.GetBrand(context.Background(), "my-brand")

	if err != nil {
		log.Fatalln(err)
	}
	log.Println(data)
}
