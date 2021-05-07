package examples

import (
	"context"

	"github.com/trycourier/courier-go/v2"
)

func deleteBrand() {
	client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)

	client.DeleteBrand(context.Background(), "my-brand")
}
