package examples

import (
	"context"

	"github.com/trycourier/courier-go/v2"
)

func deleteList() {
	client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)

	client.DeleteList(context.Background(), "my-list")
}
