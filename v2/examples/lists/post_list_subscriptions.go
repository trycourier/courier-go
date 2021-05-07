package examples

import (
	"context"

	"github.com/trycourier/courier-go/v2"
)

func postListSubscriptions() {
	client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)

	client.PostListSubscriptions(context.Background(), "my-list", courier.ListSubscriptionBody{
		Recipients: []courier.ListRecipient{
			courier.ListRecipient{
				RecipientID: "0460766e-8463-4905-ae98-b72c7aef41d6",
			},
		},
	})
}
