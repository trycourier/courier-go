package examples

import (
	"context"

	"github.com/trycourier/courier-go/v2"
)

func postNotificationDraftVariations() {
	client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)

	client.PostNotificationDraftVariations(context.Background(), "my-notification", courier.NotificationVariationsRequestBody{
		Blocks: []interface{}{
			map[string]string{"id": "block_00d0bcb0-aba1-443f-a8dd-daac505500fe"},
			map[string]string{"type": "text"},
			// ...locales
		},
		Channels: []interface{}{
			map[string]string{"id": "channel_79d25574-83be-4da1-a5c3-3d4e2ab5f154"},
			// ...locales
		},
	})
}
