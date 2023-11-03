package examples

import (
	"context"

	"github.com/trycourier/courier-go/v3"
)

func cancelNotificationSubmission() {
	client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)

	client.CancelNotificationSubmission(context.Background(), "my-notification", "my-submission")
}
