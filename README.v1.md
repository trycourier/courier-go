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
package main

import (
        "context
        "log"
        "github.com/trycourier/courier-go"
)
type profile struct {
        Email string `json:"email"`
}
type data struct {
        Foo string `json:"foo"`
}
func send() {
        var authToken = "<YOUR_AUTH_TOKEN>"
        var eventID = "<YOUR_EVENT_ID>"
        var recipientID = "<YOUR_RECIPIENT_ID>"
        profile := &profile{
                email: "foo@example.com",
        }
        data := &data{
                foo: "bar",
        }
        client := courier.CourierClient(authToken, "https://api.trycourier.app")
        messageID, err := client.Send(context.Background(), eventID, recipientID, profile, data)
        if err != nil {
                log.Fatalln(err)
	}
        log.Println(messageID)
}

func main() {
        send()
}
```

If you need to use a base url other than the default https://api.trycourier.app, you can pass it in as the second paramter to the `CourierClient`:

```go
client := courier.CourierClient("<AUTH_TOKEN>", "<BASE_URL>")
```

## Advanced Usage

```go
func getProfile() {
        client := courier.CourierClient("<AUTH_TOKEN>")
        var recipientId = "<RECIPIENT_ID>"
        response, err := client.GetProfile(recipientId)
        if err != nil {
                log.Fatalln(err)
        }
        log.Println(response.Profile)
}

func mergeProfile() {
        client := courier.CourierClient("<AUTH_TOKEN>")
        var profile = []byte(`{
                profile: {
                        email: "example@example.com",
                        phone_number: "555-228-3890"
                },
                data: {} // optional variables for merging into templates
        }`)
        err := client.MergeProfile("<RECIPIENT_ID>", profile)
        if err != nil {
                log.Fatalln(err)
        }
}

func updateProfile() {
        client := courier.CourierClient("<AUTH_TOKEN>")
        var profile = []byte(`{
                profile: {
                        email: "example@example.com",
                        phone_number: "555-228-3890"
                },
                data: {} // optional variables for merging into templates
        }`)
        err := client.UpdateProfile("5957debf-5e16-499f-ab35-a614a87fded5", profile)
        if err != nil {
                log.Fatalln(err)
        }
}
```

## Contributing
Bug reports and pull requests are welcome on GitHub at https://github.com/trycourier/courier-go.

## License
The package is available as open source under the terms of the MIT License.
[MIT License](http://www.opensource.org/licenses/mit-license.php)