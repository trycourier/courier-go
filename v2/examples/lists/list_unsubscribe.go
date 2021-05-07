package examples

import (
	"context"

	"github.com/trycourier/courier-go/v2"
)

func listUnsubscribe() {
	client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)

	client.ListUnsubscribe(context.Background(), "my-list", "my-recipient")
}
