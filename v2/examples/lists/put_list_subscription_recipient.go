package examples

import (
	"context"

	"github.com/trycourier/courier-go/v2"
)

func putListSubscriptionRecipient() {
	client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)

	client.PutListSubscriptionsRecipient(context.Background(), "my-list", "my-recipient", courier.ListSubscriptionRecipientBody{})
}
