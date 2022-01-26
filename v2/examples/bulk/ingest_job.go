package examples

import (
	"context"
	"log"

	"github.com/trycourier/courier-go/v2"
)

func ingestUsers() {
	client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)

	data, err := client.IngestJob(context.Background(), "job123", courier.IngestJobBody{
		Users: []interface{}{
			map[string]string{"recipient": "johndoe"},
			// more user properties
		},
	})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(data)
}
