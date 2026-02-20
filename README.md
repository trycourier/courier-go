<!-- AUTO-GENERATED-OVERVIEW:START — Do not edit this section. It is synced from mintlify-docs. -->
# Courier Go SDK

The Courier Go SDK provides typed access to the Courier REST API from applications written in Go. It uses strongly typed request structs, automatic retries, and returns parsed response types.

## Installation

```go
import (
	"github.com/trycourier/courier-go/v4" // imported as courier
)
```

Requires Go 1.22+.

## Quick Start

```go
package main

import (
	"context"
	"fmt"

	"github.com/trycourier/courier-go/v4"
	"github.com/trycourier/courier-go/v4/option"
	"github.com/trycourier/courier-go/v4/shared"
)

func main() {
	client := courier.NewClient(
		option.WithAPIKey("My API Key"), // defaults to os.LookupEnv("COURIER_API_KEY")
	)
	response, err := client.Send.Message(context.TODO(), courier.SendMessageParams{
		Message: courier.SendMessageParamsMessage{
			To: courier.SendMessageParamsMessageToUnion{
				OfUserRecipient: &shared.UserRecipientParam{
					UserID: courier.String("your_user_id"),
				},
			},
			Template: courier.String("your_template_id"),
			Data: map[string]any{"foo": "bar"},
		},
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%+v
", response.RequestID)
}
```

The client reads `COURIER_API_KEY` from your environment automatically.

## Documentation

Full documentation: **[courier.com/docs/sdk-libraries/go](https://www.courier.com/docs/sdk-libraries/go/)**

- [Quickstart](https://www.courier.com/docs/getting-started/quickstart/)
- [Send API](https://www.courier.com/docs/platform/sending/send-message/)
- [API Reference](https://www.courier.com/docs/reference/get-started/)
<!-- AUTO-GENERATED-OVERVIEW:END -->
