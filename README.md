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
                "my-brand-id",
                nil,
                nil,
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
* ```Send(context, eventID string, recipientID string, profile interface{}, data interface{}, brand string, override interface{}, preferences interface{}): object``` [[?]](https://docs.courier.com/reference/send-api#sendmessage)
* ```SendToList(context, eventID string, listID string, pattern string, data interface{}, brand string, override interface{}): object``` [[?]](https://docs.courier.com/reference/send-api#sendlist)

### Messages API
* ```GetMessage(context, messageID string): object``` [[?]](https://docs.courier.com/reference/messages-api#getmessagebyid)
* ```GetMessages(context, cursor string, event string, list string, messageId string, notification string, recipient string): object``` [[?]](https://docs.courier.com/reference/messages-api#getmessages)
* ```GetMessageHistory(context, messageID string, _type string): object``` [[?]](https://docs.courier.com/reference/messages-api#getmessagehistorybyid)

### Events API
* ```GetEvents(context): object``` [[?]](https://docs.courier.com/reference/events-api#getevents)
* ```GetEvent(context, eventID string): object``` [[?]](https://docs.courier.com/reference/events-api#geteventbyid)
* ```PutEvent(context, eventID string, notificationID string, eventType string): object``` [[?]](https://docs.courier.com/reference/events-api#replaceeventbyid)

### Brands API
* ```GetBrands(context, cursor string): object``` [[?]](https://docs.courier.com/reference/brands-api#getbrands)
* ```PostBrand(context, brandID string, brandName string, brandSettings BrandSettings, brandSnippets BrandSnippets): object``` [[?]](https://docs.courier.com/reference/brands-api#createbrand)
* ```GetBrand(context, brandID string): object``` [[?]](https://docs.courier.com/reference/brands-api#getbrand)
* ```PutBrand(context, brandID string, brandName string, brandSettings BrandSettings, brandSnippets BrandSnippets): object``` [[?]](https://docs.courier.com/reference/brands-api#replacebrand)
* ```DeleteBrand(context, brandID string): error string``` [[?]](https://docs.courier.com/reference/brands-api#deletebrand)

### Lists API
* ```GetLists(context, cursor string, pattern string): object``` [[?]](https://docs.courier.com/reference/lists-api#getlists)
* ```GetList(context, listID string): object``` [[?]](https://docs.courier.com/reference/lists-api#getlist)
* ```PutList(context, listID string, listName string): object``` [[?]](https://docs.courier.com/reference/lists-api#putlist)
* ```DeleteList(context, listID string): error string``` [[?]](https://docs.courier.com/reference/lists-api#deletelist)
* ```RestoreList(context, listID string): error string``` [[?]](https://docs.courier.com/reference/lists-api#putlistrestore)
* ```GetListSubscriptions(context, listID string, cursor string): object``` [[?]](https://docs.courier.com/reference/lists-api#getlistsubscriptions)
* ```SubscribeMultipleRecipientsToList(context, listID string, recipients []Recipient): error string``` [[?]](https://docs.courier.com/reference/lists-api#createlistsubscriptions)
* ```SubscribeRecipientToList(context, listID string, recipientID string): error string``` [[?]](https://docs.courier.com/reference/lists-api#putlistsubscription)
* ```DeleteListSubscription(context, listID string, recipientID string): error string``` [[?]](https://docs.courier.com/reference/lists-api#deletelistsubscription)

### Profiles API
* ```GetProfile(recipientID string): object``` [[?]](https://docs.courier.com/reference/profiles-api#getprofilebyrecipientid)
* ```MergeProfile(recipientID string, profile byte[]): error string``` [[?]](https://docs.courier.com/reference/profiles-api#mergeprofilebyrecipientid)
* ```UpdateProfile(recipientID string, profile byte[]): error string``` [[?]](https://docs.courier.com/reference/profiles-api#patchprofilebyrecipientid)
* ```PatchProfile(recipientID string, patch []PatchOp): error string``` [[?]](https://docs.courier.com/reference/profiles-api#patchprofilebyrecipientid)
* ```GetProfileLists(recipientID string, cursor string): Object``` [[?]](https://docs.courier.com/reference/profiles-api#getlistsforprofilebyrecipientid)

## Staying Updated
To update this SDK to the latest version, use `go get -u github.com/trycourier/courier-go`.

## License
The package is available as open source under the terms of the MIT License.
[MIT License](http://www.opensource.org/licenses/mit-license.php)
