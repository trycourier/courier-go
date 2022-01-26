package examples

import (
	"context"
	"log"

	"github.com/trycourier/courier-go/v2"
)

func createJob() {
	client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)

	jobID, err := client.CreateJob(context.Background(), courier.CreateJobBody{
		Message: map[string]string{
			"event": "foo",
			// more message properties
		},
	})

	if err != nil {
		log.Fatalln(err)
	}
	log.Println(jobID)
}
