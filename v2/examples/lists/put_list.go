package examples

import (
	"context"

	"github.com/trycourier/courier-go/v2"
)

func putList() {
	client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)

	client.PutList(context.Background(),
		"my-list",
		courier.PutListBody{
			Name: "Besties",
		},
	)
}
