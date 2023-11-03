package examples

import (
	"context"

	"github.com/trycourier/courier-go/v3"
)

func listSubscribe() {
	client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)

	client.ListSubscribe(context.Background(), "my-list", "my-recipient", courier.ListSubscriptionRecipientBody{})
}
