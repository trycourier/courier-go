package courier_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trycourier/courier-go"
)

func TestClient_GetMessage(t *testing.T) {

	requestMessageID := "1-23456789"
	status := "CLICKED"
	sent := int64(1589563865697)

	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/messages/"+requestMessageID, req.URL.String())

			rw.Header().Add("Content-Type", "application/json")
			rsp := `
			{
  				"id": "%s",
  				"clicked": 1589563890843,
  				"delivered": 1589563972000,
  				"enqueued": 1589563863773,
  				"event": "GEFGNB2GNQ4MZVHW4WV4R8Q8ZVN5",
  				"notification": "GEFGNB2GNQ4MZVHW4WV4R8Q8ZVN5",
  				"providers": [
				{
      				"channel": {
        				"key": "sendgrid",
        				"template": "5e95b992-3505-4f66-8808-f91d5d0fe8c9"
      				},
      				"clicked": 1589563890843,
      				"delivered": 1589563972000,
      				"provider": "sendgrid",
      				"reference": {
        			"message_id": "Xx14o44jToS8StSTPLGsTw.filterdrecv-p3iad2-8ddf98858-7qsdm-19-5EBED1D9-D3.0",
        			"x-message-id": "Xx14o44jToS8StSTPLGsTw"
					},
      				"sent": 1589563865697,
      				"status": "DELIVERED"
    			}
  				],
  				"recipient": "tony@trycourier.com",
  				"sent": %d,
  				"status": "%s"
			}`
			_, _ = rw.Write([]byte(fmt.Sprintf(rsp, requestMessageID, sent, status)))

		}))
	defer server.Close()

	t.Run("makes request for message ID", func(t *testing.T) {
		client := courier.CourierClient("key", server.URL)
		rsp, err := client.GetMessage(context.Background(), requestMessageID)
		assert.Nil(t, err)
		assert.Equal(t, requestMessageID, rsp.Id)
		assert.Equal(t, sent, rsp.Sent)
		assert.Equal(t, status, rsp.Status)
	})
}

func TestClient_GetMessages(t *testing.T) {

	event := "courier-quickstart"
	recipient := "Github_6154318"

	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/messages?event="+event+"&recipient="+recipient, req.URL.String())

			rw.Header().Add("Content-Type", "application/json")
			rsp := `
			{
				"Paging":{
					 "Cursor":"",
					 "More":false
				},
				"Results":[
					 {
							"Id":"1-60122b54-7e5d45d35b820ace1c0ffc49",
							"Event":"%s",
							"Notification":"",
							"Status":"UNMAPPED",
							"Error":"",
							"Reason":"",
							"Recipient":"%s",
							"Enqueued":1611803477021,
							"Delivered":0,
							"Sent":0,
							"Clicked":0,
							"Opened":0,
							"Providers":null
					 }
				]
		 }`
			_, _ = rw.Write([]byte(fmt.Sprintf(rsp, event, recipient)))

		}))
	defer server.Close()

	t.Run("makes request for messages", func(t *testing.T) {
		client := courier.CourierClient("key", server.URL)
		rsp, err := client.GetMessages(context.Background(), "", event, "", "", "", recipient)

		assert.Nil(t, err)
		assert.Equal(t, event, rsp.Results[0].Event)
		assert.Equal(t, recipient, rsp.Results[0].Recipient)
	})
}

func TestClient_GetMessageHistory(t *testing.T) {
	messageId := "1-60136482-4f3e07390677c06e2e248de6"
	_type := "ENQUEUED"

	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/messages/"+messageId+"/history?type="+_type, req.URL.String())

			rw.Header().Add("Content-Type", "application/json")
			rsp := `
			{
				"Results": [
						{
								"data": {
										"favoriteAdjective": "awesomeness"
								},
								"event": "courier-quickstart",
								"profile": {
										"email": "foo@example.com"
								},
								"recipient": "Github_6154318",
								"ts": 1611883650743,
								"type": "%s"
						}
				]
		}`
			_, _ = rw.Write([]byte(fmt.Sprintf(rsp, _type)))

		}))
	defer server.Close()

	t.Run("makes request for message history", func(t *testing.T) {
		client := courier.CourierClient("key", server.URL)
		rsp, err := client.GetMessageHistory(context.Background(), "1-60136482-4f3e07390677c06e2e248de6", _type)

		assert.Nil(t, err)
		assert.Equal(t, 1, len(rsp.Results))
	})
}
