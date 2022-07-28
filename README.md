# Courier Go SDK

This Go package helps you send notifications through Courier, the smartest way to design & deliver notifications. Design your notifications once using our drag & drop editor, then deliver to any channel through one API. Email, mobile push, SMS, Slack — you name it!

!["Golang Gopher"](https://blog.golang.org/gopher/gopher.png)

APIs supported:

- Audiences API
- Audit Events API
- Automations API
- Brand API
- Messages API
- Profiles API
- Send API
- List API

## Official Courier API docs

For a full description of request and response payloads and properties, please see the [official Courier API docs](https://docs.courier.com/reference).

## Installation

Download the package using `go get`:

```bash
go get -u github.com/trycourier/courier-go/v2
```

## Usage

```go
package main

import (
	"context"
	"fmt"

	"github.com/trycourier/courier-go/v2"
)

func main() {
	client := courier.CourierClient("<YOUR_AUTH_TOKEN>", nil)
	message := courier.SendMessageRequestBody{
		"template": "<COURIER_TEMPLATE>",
			"to": map[string]string{
				"email": "test@email.com"
		}
	}

	reqID, err := client.SendMessage(context.Background(), message)
	if err != nil {
		panic(err)
	}

	fmt.Println(reqID)
}
```

If you would like to send an idempotent message then be sure to use the SendMessageWithOptions method as shown below:

```go
package main

import (
	"fmt"
	"time"
	"context"

	"github.com/trycourier/courier-go/v2"
)

func main() {
	client := courier.CourierClient("<YOUR_AUTH_TOKEN>", nil)
	message := map[string]interface{}{
		"template": "<COURIER_TEMPLATE>",
			"to": map[string]string{
				"email": "test@email.com"
			}
	}

	reqID, err := client.SendMessageWithOptions(
		context.Background(),
		message,
		courier.WithIdempotencyKey("fake-key"),
		courier.WithIdempotencyKeyExpiration(time.Now().Add(time.Hour * 30))
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(reqID)
}
```

If you need to use a base url other than the default https://api.courier.com, you can pass it in as the second paramter to the `CourierClient`:

```go
client := courier.CourierClient("<AUTH_TOKEN>", "<BASE_URL>")
```

## APIs

For a full description of request and response payloads and properties, please see the [official Courier API docs](https://docs.courier.com/reference).

### Send API

- `SendMessage(ctx context.Context, body SendMessageRequestBody) (string, error)` [[?]](https://www.courier.com/docs/reference/send/message/)
- `Send(ctx context.Context, eventID, recipientID string, body interface{}) (string, error): object` [[?]](https://www.courier.com/docs/reference/send/message/)
- `SendMessageWithOptions(ctx context.Context, body map[string]interface{}, method string, opts ...Option) (string, error)` [[?]](https://www.courier.com/docs/reference/idempotent-requests/)

### Messages API

- `GetMessage(ctx context.Context, messageID string) (*MessageResponse, error)` [[?]](https://www.courier.com/docs/reference/messages/by-id/)

### Audiences API

- `PutAudience(ctx context.Context, audienceId string, audience Audience) (*AudienceResponse, error)` [[?]](https://www.courier.com/docs/reference/audiences/put/)
- `GetAudience(ctx context.Context, audienceId string) (*AudienceResponseBody, error)` [[?]](https://www.courier.com/docs/reference/audiences/by-id/)
- `GetAudienceMembers(ctx context.Context, audienceId string, cursor string) (*GetAudienceMembersResponse, error)` [[?]](https://www.courier.com/docs/reference/audiences/list-audience-members/)
- `GetAudiences(ctx context.Context, cursor string) (*GetAudiencesResponse, error)` [[?]](https://www.courier.com/docs/reference/audiences/list-audiences/)
- `DeleteAudience(ctx context.Context, audienceId string) error ` [[?]](https://www.courier.com/docs/reference/audiences/delete/)

### Brand API

- `GetBrands(ctx context.Context, cursor string) (*BrandsResponse, error)` [[?]](https://www.courier.com/docs/reference/brands/list/)
- `GetBrand(ctx context.Context, brandID string) (*BrandResponse, error)` [[?]](https://www.courier.com/docs/reference/brands/by-id/)
- `PostBrand(ctx context.Context, body PostBrandBody) (*BrandResponse, error)` [[?]](https://www.courier.com/docs/reference/brands/create/)
- `PutBrand(ctx context.Context, brandID string, body PutBrandBody) (*BrandResponse, error)` [[?]](https://www.courier.com/docs/reference/brands/replace/)
- `DeleteBrand(ctx context.Context, brandID string) error` [[?]](https://www.courier.com/docs/reference/brands/delete/)

### Profiles API

- `GetProfile(ctx context.Context, id string) (map[string]json.RawMessage, error): object` [[?]](https://docs.courier.com/reference/profiles-api#getprofilebyrecipientid)
- `MergeProfileBytes(ctx context.Context, id string, profile []byte) error` [[?]](https://docs.courier.com/reference/profiles-api#mergeprofilebyrecipientid)
- `UpdateProfileBytes(ctx context.Context, id string, profile []byte) error` [[?]](https://docs.courier.com/reference/profiles-api#patchprofilebyrecipientid)

### Lists API

- `GetLists(ctx context.Context, cursor string, pattern string) (*ListsResponse, error)` [[?]](https://www.courier.com/docs/reference/lists/list/)
- `GetList(ctx context.Context, listID string) (*ListResponse, error)` [[?]](https://www.courier.com/docs/reference/lists/by-id/)
- `PutList(ctx context.Context, listID string, body interface{}) error` [[?]](https://www.courier.com/docs/reference/lists/replace/)
- `DeleteList(ctx context.Context, listID string) error` [[?]](https://www.courier.com/docs/reference/lists/delete/)
- `RestoreList(ctx context.Context, listID string) error` [[?]](https://www.courier.com/docs/reference/lists/restore/)
- `GetListSubscriptions(ctx context.Context, listID string, cursor string) (*ListSubscriptionsResponse, error)` [[?]](https://www.courier.com/docs/reference/lists/subscriptions/)
- `PutListSubscriptions(ctx context.Context, listID string, body interface{}) error` [[?]](https://www.courier.com/docs/reference/lists/put-subscribe/)
- `PostListSubscriptions(ctx context.Context, listID string, body interface{}) error` [[?]](https://www.courier.com/docs/reference/lists/post-subscribe/)
- `ListSubscribe(ctx context.Context, listID string, recipientID string, body interface{}) error` [[?]](https://www.courier.com/docs/reference/lists/recipient-subscribe/)
- `ListUnsubscribe(ctx context.Context, listID string, recipientID string) error` [[?]](https://www.courier.com/docs/reference/lists/delete-subscription/)

### Automations API

- `InvokeAutomation(ctx context.Context, body interface{}) (string, error)`[[?]](https://www.courier.com/docs/reference/automation/invoke/)
- `InvokeAutomationTemplate(ctx context.Context, templateId string, body interface{}) (string, error)` [[?]](https://www.courier.com/docs/reference/automation/invoke-template/)

### Audit Events API

- `GetAuditEvent(ctx context.Context, auditEventID string) (*AuditEvent, error)` [[?]](https://www.courier.com/docs/reference/audit-events/by-id/)
- `ListAuditEvents(ctx context.Context, cursor string) (*ListAuditEventsResponse, error)` [[?]](https://www.courier.com/docs/reference/audit-events/list/)

## Staying Updated

To update this SDK to the latest version, use `go get -u github.com/trycourier/courier-go`.

## License

The package is available as open source under the terms of the MIT License.
[MIT License](http://www.opensource.org/licenses/mit-license.php)
