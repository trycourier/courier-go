# `trycourier`

This Go package helps you send notifications through Courier, the smartest way to design & deliver notifications. Design your notifications once using our drag & drop editor, then deliver to any channel through one API. Email, mobile push, SMS, Slack — you name it!

API Documenation: https://docs.trycourier.com/reference

!["Golang Gopher"](https://blog.golang.org/gopher/gopher.png)

## Installation
Add this line to your go package:
```bash
import "github.com/trycourier/courier-go"
```

## Usage

```go
func send() {
        type profile struct {
                Email string `json:"email"`
        }
        type data struct {
                Foo string `json:"foo"`
        }

        var eventID = "example-event"
        var recipientID = "example-recipient"

        client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)
        messageID, err := client.Send(context.Background(), eventID, recipientID, courier.SendBody{
                profile{
                        Email: "foo@example.com",
                },
                data{
                        Foo: "bar",
                },
        })
        if err != nil {
                log.Fatalln(err)
	}
        log.Println(messageID)
}
```

If you need to use a base url other than the default https://api.trycourier.app, you can pass it in as the second paramter to the `CourierClient`:

```go
client := courier.CreateClient("<AUTH_TOKEN>", "<BASE_URL>")
```

## Advanced Usage

```go
func getProfile() {
        var recipientID = "<YOUR_RECIPIENT_ID>"

        client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)
        response, err := client.GetProfile(context.Background(), recipientID)
        if err != nil {
                log.Fatalln(err)
        }

        log.Println(response.Profile)
}

func mergeProfile() {
        var recipientID = "<YOUR_RECIPIENT_ID>"
        var profile = []byte(`{
                profile: {
                        email: "example@example.com",
                        phone_number: "555-228-3890"
                }
        }`)

        client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)
        err := client.MergeProfile(context.Background(), recipientID, profile)
        if err != nil {
                log.Fatalln(err)
        }
}

func updateProfile() {
        var recipientID = "<YOUR_RECIPIENT_ID>"
        var profile = []byte(`{
                profile: {
                        email: "example@example.com",
                        phone_number: "555-228-3890"
                }
        }`)

        client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)
        err := client.UpdateProfile(context.Background(), recipientID, profile)
        if err != nil {
                log.Fatalln(err)
        }
}
```

## Staying Updated
To update this SDK to the latest version, use `go get -u github.com/trycourier/courier-go`.

## License
The package is available as open source under the terms of the MIT License.
[MIT License](http://www.opensource.org/licenses/mit-license.php)
