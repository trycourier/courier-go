# Courier Go SDK

This Go package helps you send notifications through Courier, the smartest way to design & deliver notifications. Design your notifications once using our drag & drop editor, then deliver to any channel through one API. Email, mobile push, SMS, Slack — you name it!

!["Golang Gopher"](https://blog.golang.org/gopher/gopher.png)

APIs supported:
* Send API
* Messages API
* Profiles API

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
        messageID, err := client.Send(
                context.Background(), 
                eventID,
                recipientID,
                map[string]string{"email": "foo@example.com"},
		map[string]string{"favoriteAdjective": "awesomeness"},
        })
        if err != nil {
                log.Fatalln(err)
	}
        log.Println(messageID)
}
```

If you need to use a base url other than the default https://api.courier.com, you can pass it in as the second paramter to the `CourierClient`:

```go
client := courier.CourierClient("<AUTH_TOKEN>", "<BASE_URL>")
```

## APIs

For a full description of request and response payloads and properties, please see the [official Courier API docs](https://docs.courier.com/reference).

### Send API
* ```Send(context, eventID string, recipientID string, profile, data): object``` [[?]](https://docs.courier.com/reference/send-api#sendmessage)

### Messages API
* ```GetMessage(context, messageID string): object``` [[?]](https://docs.courier.com/reference/messages-api#getmessagebyid)

### Profiles API
* ```GetProfile(id string): object``` [[?]](https://docs.courier.com/reference/profiles-api#getprofilebyrecipientid)
* ```MergeProfile(id string, profile byte[]): object``` [[?]](https://docs.courier.com/reference/profiles-api#mergeprofilebyrecipientid)
* ```UpdateProfile(id string, profile byte[]): object``` [[?]](https://docs.courier.com/reference/profiles-api#patchprofilebyrecipientid)


## Staying Updated
To update this SDK to the latest version, use `go get -u github.com/trycourier/courier-go`.

## License
The package is available as open source under the terms of the MIT License.
[MIT License](http://www.opensource.org/licenses/mit-license.php)
