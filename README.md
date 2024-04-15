# Courier Go Library

[![Courier: Your Complete Communication Stack](https://marketing-assets-public.s3.us-west-1.amazonaws.com/github_nodejs.png)](https://courier.com)

[![go shield](https://img.shields.io/badge/go-docs-blue)](https://pkg.go.dev/github.com/merge-api/merge-go-client)
[![fern shield](https://img.shields.io/badge/%F0%9F%8C%BF-SDK%20generated%20by%20Fern-brightgreen)](https://buildwithfern.com/?utm_source=trycourier/courier-go/readme)

This is the official Go library for sending notifications with the [Courier](https://courier.com) REST API.

## Documentation

For a full description of request and response payloads and properties, please see the
[official Courier API docs](https://www.courier.com/docs/reference/).

## Requirements

This module requires Go version >= 1.13.

## Installation

Run the following command to use the Go library in your Go module:

```sh
go get github.com/trycourier/courier-go/v3
```

## Instantiation

```go
import (
  "context"
  "fmt"

  courierclient "github.com/trycourier/courier-go/v3/client"
  option "github.com/trycourier/courier-go/v3/option"
)

client := courierclient.NewClient(
  option.WithAuthorizationToken("<YOUR_TOKEN>"),
)
```

## Usage

```go
import (
  "context"
  "fmt"

  courier "github.com/trycourier/courier-go/v3"
  courierclient "github.com/trycourier/courier-go/v3/client"
  option "github.com/trycourier/courier-go/v3/option"

)

client := courierclient.NewClient(
  option.WithAuthorizationToken("<YOUR_TOKEN>"),
)
sendResponse, err := client.Send(
  context.TODO(),
  &courier.SendMessageRequest{
    Message: courier.NewMessageFromTemplateMessage({
      Template: "<COURIER_TEMPLATE>",
    }),
  },
)
if err != nil {
  return err
}
fmt.Printf("Sent message %s\n", sendResponse.RequestId)
```

## Unions

Our API, particularly the send method, uses several unions. Our Go SDK defines structs
to construct these unions, such as `courier.Message`, shown below:

```go
import (
  courier "github.com/trycourier/courier-go/v3"
)

request := &courier.SendMessageRequest{
  // Construct a content message.
  Message: &courier.Message{
    ContentMessage: &courier.ContentMessage{
      // Construct a single recepient that is a user recepient.
      To: &courier.MessageRecipient{
        Recipient: &courier.Recipient{
          UserRecipient: &courier.UserRecipient{
            Email: courier.String("marty_mcfly@email.com"),
            Data: &courier.MessageData{
              "name": "Marty",
            },
          },
        },
      },
      // Construct content from elemental content sugar.
      Content: &courier.Content{
        ElementalContentSugar: &courier.ElementalContentSugar{
          Title: "Back to the Future",
          Body:  "Oh my {{name}}, we need 1.21 Gigawatts!",
        },
      },
    },
  },
}
```

### Breaking change (3.0.7 -> 3.0.8)

We introduced a better union construction experience in [3.0.8](https://github.com/trycourier/courier-go/releases/tag/v3.0.8).
For example, the `courier.Message` type was previously constructed with the following:

```go
import (
  courier "github.com/trycourier/courier-go/v3"
)

request := courier.SendMessageRequest{
  // Construct a content message.
  Message: courier.NewMessageFromContentMessage(
    &courier.ContentMessage{
      // Construct a single recepient that is a user recepient.
      To: courier.NewMessageRecipientFromRecipient(
        courier.NewRecipientFromUserRecipient(
          &courier.UserRecipient{
            Email: courier.String("marty_mcfly@email.com"),
            Data: &courier.MessageData{
              "name": "Marty",
            },
          },
        ),
      ),
      // Construct content from elemental content sugar.
      Content: courier.NewContentFromElementalContentSugar(
        &courier.ElementalContentSugar{
          Title: "Back to the Future",
          Body:  "Oh my {{name}}, we need 1.21 Gigawatts!",
        },
      ),
    },
  ),
}
```

Although the construction looks fairly similar, the old approach required navigating a
variety of cumebersome constructor function names (e.g. `courier.NewContentFromElementalContentSugar`).

The new approach drops these constructors entirely, which simplifies the experience significantly.
Migrating from the old approach is as simple as setting the concrete type to the appropriate
field like so:

**Before**

```go
  ...
  Content: courier.NewContentFromElementalContentSugar(
    &courier.ElementalContentSugar{
      Title: "Back to the Future",
      Body:  "Oh my {{name}}, we need 1.21 Gigawatts!",
    },
  ),
  ...
```

**After**

```go
  ...
  Content: &courier.Content{
    ElementalContentSugar: &courier.ElementalContentSugar{
      Title: "Back to the Future",
      Body:  "Oh my {{name}}, we need 1.21 Gigawatts!",
    },
  },
  ...
```

## Timeouts

Setting a timeout for each individual request is as simple as using the standard
`context` library. Setting a one second timeout for an individual API call looks
like the following:

```go
ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
defer cancel()

response, err := client.Send(
  ctx,
  &courier.SendMessageRequest{
    Message: ...
  },
)
```

## Client Options

A variety of client options are included to adapt the behavior of the library, which
includes configuring authorization tokens to be sent on every request, or providing your
own instrumented `*http.Client`. Both of these options are shown below:

```go
client := courierclient.NewClient(
  option.WithAuthorizationToken("<YOUR_TOKEN>"),
  option.WithHTTPClient(
    &http.Client{
      Timeout: 5 * time.Second,
    },
  ),
)
```

> Providing your own `*http.Client` is recommended. Otherwise, the `http.DefaultClient` will be used,
> and your client will wait indefinitely for a response (unless the per-request, context-based timeout
> is used).

## Errors

Structured error types are returned from API calls that return non-success status codes. For example,
you can check if the error was due to a bad request (i.e. status code 400) with the following:

```go
response, err := client.Send(
  context.TODO(),
  &courier.SendMessageRequest{},
)
if err != nil {
  if apiErr, ok := err.(*core.APIError); ok && apiErr.StatusCode == http.StatusBadRequest {
    // Do something with the bad request ...
  }
  return err
}
```

These errors are also compatible with the `errors.Is` and `errors.As` APIs, so you can access the error
like so:

```go
response, err := client.Send(
  context.TODO(),
  &courier.SendMessageRequest{},
)
if err != nil {
  var apiErr *core.APIError
  if errors.As(err, apiError); ok {
    switch apiErr.StatusCode {
      case http.StatusBadRequest:
        // Do something with the bad request ...
    }
  }
  return err
}
```

If you'd like to wrap the errors with additional information and still retain the ability to access the type
with `errors.Is` and `errors.As`, you can use the `%w` directive:

```go
response, err := client.Send(
  context.TODO(),
  &courier.SendMessageRequest{},
)
if err != nil {
  return fmt.Errorf("failed to list employees: %w", err)
}
```

## Contributing

While we value open-source contributions to this SDK, this library is generated programmatically. Additions
made directly to this library would have to be moved over to our generation code, otherwise they would be
overwritten upon the next generated release. Feel free to open a PR as a proof of concept, but know that we
will not be able to merge it as-is. We suggest opening an issue first to discuss with us!

On the other hand, contributions to the README are always very welcome!
