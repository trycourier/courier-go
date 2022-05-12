package examples

import (
	"context"
	"log"

	"github.com/trycourier/courier-go/v2"
)

func sendMessageWithGranularMetadata() {
	client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)

	requestID, err := client.SendMessage(context.Background(), courier.SendMessageRequestBody{
		Message: map[string]interface{}{
			"template": "my-template",
			"to": map[string]string{
				"email": "foo@bar.com",
			},
			"metadata": map[string]interface{}{
				"utm": map[string]string{
					"source": "go",
				},
			},
			"channels": map[string]interface{}{
				"email": map[string]interface{}{
					"metadata": map[string]interface{}{
						"utm": map[string]string{
							"medium": "email",
						},
					},
				},
			},
			"providers": map[string]interface{}{
				"sendgrid": map[string]interface{}{
					"metadata": map[string]interface{}{
						"utm": map[string]string{
							"campaign": "sendgrid",
						},
					},
				},
			},
		},
	})

	if err != nil {
		log.Fatalln(err)
	}
	log.Println(requestID)
}
