package examples

import (
	"context"

	"github.com/trycourier/courier-go/v3"
)

func postBrand() {
	client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)

	client.PostBrand(context.Background(), courier.PostBrandBody{
		Name: "my-brand",
	})
}
