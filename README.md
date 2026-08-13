# Courier Go SDK

The Courier Go SDK provides typed access to the Courier REST API from Go applications. Use it to send notifications, manage user profiles, check message status, issue JWT tokens for client-side SDKs, and more.

## Installation

```bash
go get github.com/trycourier/courier-go/v4
```

Then import it:

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
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(response.RequestID)
}
```

The client reads `COURIER_API_KEY` from your environment automatically.

## Documentation

Full documentation: **[courier.com/docs/sdk-libraries/go](https://www.courier.com/docs/sdk-libraries/go/)**

- [Quickstart](https://www.courier.com/docs/getting-started/quickstart/)
- [Send API](https://www.courier.com/docs/platform/sending/send-message/)
- [API Reference](https://www.courier.com/docs/reference/get-started/)
