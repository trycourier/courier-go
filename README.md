# courier-go

A go package for communicating with the Courier REST API.

API Documenation: https://docs.trycourier.com/reference

## Installation (via [npm](https://www.npmjs.com/package/@trycourier/courier))

```bash
import "github.com/trycourier/courier-go"
```

## Usage

```go
package main

import (
	"log"
	"github.com/trycourier/courier-go"
)

func send() {
	client := courier.CourierClient("<AUTH_TOKEN>")  // get from the Courier UI
	var message = []byte(`{
		eventId: "<EVENT_ID>", // get from the Courier UI
    recipientId: "<RECIPIENT_ID>"
		profile: {
      email: "example@example.com",
      phone_number: "555-228-3890"
    },
    data: {} // optional variables for merging into templates
	}`)
	data, err := client.Send(message)
	if err != nil {
		log.Fatalln(err)
  }
  log.Println(data)
}

func main() {
	send()
}
```

## Environment Variables

`courier-go` supports credential storage in environment variables. If no `authorizationToken` is provided when instantiating the Courier client (e.g., `client := courier.CourierClient("<AUTH_TOKEN>")`), the value in the `COURIER_AUTH_TOKEN` env var will be used.

If you need to use a base url other than the default https://api.trycourier.app, you can set it using the `COURIER_BASE_URL` env var.

## Advanced Usage

```go
func getProfile() {
	client := courier.CourierClient("<AUTH_TOKEN>")
	var recipientId = "<RECIPIENT_ID>"
	profile, err := client.GetProfile(recipientId)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(profile.Profile)
}

func mergeProfile() {
	client := courier.CourierClient("<AUTH_TOKEN>")
	var newProfile = []byte(`{
		profile: {
      email: "example@example.com",
      phone_number: "555-228-3890"
    },
    data: {} // optional variables for merging into templates
	}`)
	err := client.MergeProfile("<RECIPIENT_ID>", newProfile)
	if err != nil {
		log.Fatalln(err)
	}
}

func updateProfile() {
	client := courier.CourierClient("<AUTH_TOKEN>")
	var newProfile = []byte(`{
		profile: {
      email: "example@example.com",
      phone_number: "555-228-3890"
    },
    data: {} // optional variables for merging into templates
	}`)
	err := client.UpdateProfile("5957debf-5e16-499f-ab35-a614a87fded5", newProfile)
	if err != nil {
		log.Fatalln(err)
	}
}
```

## License

[MIT License](http://www.opensource.org/licenses/mit-license.php)

## Author

[Courier](https://github.com/trycourier) ([support@trycourier.com](mailto:support@trycourier.com))
