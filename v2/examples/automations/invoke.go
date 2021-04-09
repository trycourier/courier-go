package examples

import (
	"context"
	"log"

	"github.com/trycourier/courier-go/v2"
)

func invokeAutomation() {
	client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)

	runID, err := client.InvokeAutomation(context.Background(), courier.AutomationInvokeBody{
		Automation: courier.Automation{
			Steps: []courier.AutomationStep{
				courier.AutomationStep{
					Action: "send",
				},
			},
		},
		Brand:     "",  // optional brand ID
		Data:      nil, // optional data map
		Profile:   nil, // optional profile map
		Recipient: "",  // optional recipient ID
		Template:  "",  // optional notification template ID
	})

	if err != nil {
		log.Fatalln(err)
	}
	log.Println(runID)
}
