package examples

import (
	"context"
	"log"

	"github.com/trycourier/courier-go/v2"
)

func getJobUsers() {
	client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)

	data, err := client.GetJobUsers(context.Background(),
		"job123", // required job ID
		"",       // optional cursor
	)

	if err != nil {
		log.Fatalln(err)
	}
	log.Println(data)
}
