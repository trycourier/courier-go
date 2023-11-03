package examples

import (
	"context"

	"github.com/trycourier/courier-go/v3"
)

func runJob() {
	client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)

	client.RunJob(context.Background(), "job123")
}
