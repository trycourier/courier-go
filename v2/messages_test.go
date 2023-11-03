package courier_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trycourier/courier-go/v2"
)

func TestMessages_GetMessage(t *testing.T) {
	requestMessageID := "1-23456789"
	status := "CLICKED"
	sent := int64(1589563865697)
	var expectedURL string
	var rsp courier.MessageResponse

	server := httptest.NewServer(http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			assert.Equal(t, expectedURL, req.URL.String())

			rw.WriteHeader(http.StatusOK)
			rw.Header().Add("Content-Type", "application/json")
			bytes, err := json.Marshal(rsp)
			assert.Nil(t, err)
			_, _ = rw.Write(bytes)
		}),
	)
	defer server.Close()

	t.Run("makes requests for message ID", func(t *testing.T) {
		rsp = createResponse(requestMessageID, sent, status)
		defer func() { rsp = courier.MessageResponse{} }() // reset response between tests

		expectedURL = "/messages/" + requestMessageID

		client := courier.CreateClient("key", &server.URL)
		response, err := client.GetMessage(context.Background(), requestMessageID)
		assert.Nil(t, err)

		assert.Equal(t, requestMessageID, response.ID)
		assert.Equal(t, sent, response.Sent)
		assert.Equal(t, status, response.Status)
	})

	t.Run("makes requests to cancel a message", func(t *testing.T) {
		rsp = createResponse(requestMessageID, sent, status)
		defer func() { rsp = courier.MessageResponse{} }() // reset response between tests

		canceledAt := int64(4567890)
		rsp.CanceledAt = canceledAt

		expectedURL = fmt.Sprintf("/messages/%s/cancel", requestMessageID)

		client := courier.CreateClient("key", &server.URL)
		response, err := client.CancelMessage(context.Background(), requestMessageID)
		assert.Nil(t, err)

		assert.Equal(t, requestMessageID, response.ID)
		assert.Equal(t, sent, response.Sent)
		assert.Equal(t, status, response.Status)
		assert.Equal(t, canceledAt, response.CanceledAt)
	})
}

func createResponse(requestMessageID string, sent int64, status string) courier.MessageResponse {
	return courier.MessageResponse{
		ID:           requestMessageID,
		Clicked:      1589563890843,
		Delivered:    1589563972000,
		Enqueued:     1589563863773,
		Event:        "GEFGNB2GNQ4MZVHW4WV4R8Q8ZVN5",
		Notification: "GEFGNB2GNQ4MZVHW4WV4R8Q8ZVN5",
		Providers: []*courier.ProvidersResponse{
			{
				Channel: &courier.ProvidersChannelResponse{
					Key:      "sendgrid",
					Template: "5e95b992-3505-4f66-8808-f91d5d0fe8c9",
				},
				Clicked:   1589563890843,
				Delivered: 1589563972000,
				Provider:  "sendgrid",
				Reference: map[string]string{
					"message_id":   "Xx14o44jToS8StSTPLGsTw.filterdrecv-p3iad2-8ddf98858-7qsdm-19-5EBED1D9-D3.0",
					"x-message-id": "Xx14o44jToS8StSTPLGsTw",
				},
				Sent:   1589563865697,
				Status: "DELIVERED",
			},
		},
		Recipient: "tony@trycourier.com",
		Sent:      sent,
		Status:    status,
	}
}
