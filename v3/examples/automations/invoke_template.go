package examples

import (
	"context"
	"log"

	"github.com/trycourier/courier-go/v3"
)

func invokeAutomationTemplate() {
	client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)

	runID, err := client.InvokeAutomationTemplate(context.Background(), "<AUTOMATION-TEMPLATE-ID>", courier.AutomationTemplateInvokeBody{
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
