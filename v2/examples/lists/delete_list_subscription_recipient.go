package examples

import (
	"context"

	"github.com/trycourier/courier-go/v2"
)

func deleteListSubscriptionRecipient() {
	client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)

	client.DeleteListSubscriptionsRecipient(context.Background(), "my-list", "my-recipient")
}
