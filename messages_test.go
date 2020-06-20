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

	t.Run("makes requests for message ID", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		rsp, err := client.GetMessage(context.Background(), requestMessageID)
		assert.Nil(t, err)
		assert.Equal(t, requestMessageID, rsp.ID)
		assert.Equal(t, sent, rsp.Sent)
		assert.Equal(t, status, rsp.Status)
	})
}
