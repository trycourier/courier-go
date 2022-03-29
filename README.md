# Courier Go SDK

This Go package helps you send notifications through Courier, the smartest way to design & deliver notifications. Design your notifications once using our drag & drop editor, then deliver to any channel through one API. Email, mobile push, SMS, Slack — you name it!

!["Golang Gopher"](https://blog.golang.org/gopher/gopher.png)

APIs supported:

- Audiences API
- Messages API
- Profiles API
- Send API

## Official Courier API docs

For a full description of request and response payloads and properties, please see the [official Courier API docs](https://docs.courier.com/reference).

## Installation

Add this line to your go package:

```bash
import "github.com/trycourier/courier-go"
```

## Usage

```go
func send() {
        var eventID = "example-event"
        var recipientID = "example-recipient"

        client := courier.CourierClient("<YOUR_AUTH_TOKEN>", nil)
        requestID, err := client.SendMessage(
                context.Background(),
                courier.SendMessageRequestBody{
                        "template": "<COURIER_TEMPLATE>",
                        "to": map[string]string{
                                "email": "test@email.com"
                        }
                }
        )
        if err != nil {
                log.Fatalln(err)
	}
        log.Println(requestID)
}
```

If you would like to send an idempotent message then be sure to use the SendMessageWithOptions method as shown below:

```go
func sendIdempotentMessage() {
        var eventID = "example-event"
        var recipientID = "example-recipient"
        expiration := time.Now().Add(time.Hour * 30)
        client := courier.CourierClient("<YOUR_AUTH_TOKEN>", nil)
        requestID, err := client.SendMessageWithOptions(
                context.Background(),
                map[string]interface{}{
                        "template": "<COURIER_TEMPLATE>",
                        "to": map[string]string{
                                "email": "test@email.com"
                        }
                },
                courier.WithIdempotencyKey("fake-key"),
		courier.WithIdempotencyKeyExpiration(expiration)
        )
        if err != nil {
                log.Fatalln(err)
	}
        log.Println(requestID)
}
```

If you need to use a base url other than the default https://api.courier.com, you can pass it in as the second paramter to the `CourierClient`:

```go
client := courier.CourierClient("<AUTH_TOKEN>", "<BASE_URL>")
```

## APIs

For a full description of request and response payloads and properties, please see the [official Courier API docs](https://docs.courier.com/reference).

### Send API
* ```SendMessage(context, message interface{}): object``` [[?]](https://www.courier.com/docs/reference/send/message/)
* ```Send(context, eventID string, recipientID string, profile interface{}, data interface{}, brand string, override interface{}, preferences interface{}): object``` [[?]](https://www.courier.com/docs/reference/send/message/)
* ```SendToList(context, eventID string, listID string, pattern string, data interface{}, brand string, override interface{}): object``` [[?]](https://www.courier.com/docs/reference/send/list/)
* ```SendMessageWithOptions(ctx context.Context, body map[string]interface{}, method string, opts ...Option) (string, error):``` [[?]](https://www.courier.com/docs/reference/idempotent-requests/)

### Messages API

- `GetMessage(context, messageID string): object` [[?]](https://www.courier.com/docs/reference/messages/by-id/)
- `GetMessages(context, cursor string, event string, list string, messageId string, notification string, recipient string): object` [[?]](https://www.courier.com/docs/reference/messages/list/)
- `GetMessageHistory(context, messageID string, _type string): object` [[?]](https://www.courier.com/docs/reference/messages/history-by-id/)

### Profiles API

- `GetProfile(id string): object` [[?]](https://docs.courier.com/reference/profiles-api#getprofilebyrecipientid)
- `MergeProfile(id string, profile byte[]): object` [[?]](https://docs.courier.com/reference/profiles-api#mergeprofilebyrecipientid)
- `UpdateProfile(id string, profile byte[]): object` [[?]](https://docs.courier.com/reference/profiles-api#patchprofilebyrecipientid)

## Staying Updated

To update this SDK to the latest version, use `go get -u github.com/trycourier/courier-go`.

## License

The package is available as open source under the terms of the MIT License.
[MIT License](http://www.opensource.org/licenses/mit-license.php)
