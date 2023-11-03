package examples

import (
	"context"
	"log"

	"github.com/trycourier/courier-go/v3"
)

func getJob() {
	client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)

	data, err := client.GetJob(context.Background(), "job123")

	if err != nil {
		log.Fatalln(err)
	}
	log.Println(data)
}
