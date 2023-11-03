package examples

import (
	"context"

	"github.com/trycourier/courier-go/v3"
)

func putNotificationSubmissionChecks() {
	client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)

	client.PutNotificationSubmissionChecks(context.Background(), "my-notification", "my-submission", courier.NotificationSubmissionChecksRequest{
		Checks: []courier.SubmissionCheck{
			courier.SubmissionCheck{
				ID:     "my-check",
				Status: "RESOLVED",
				Type:   "custom",
			},
		},
	})
}
