package courier_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trycourier/courier-go"
)

func TestClient_GetEvent(t *testing.T) {
	eventID := "event-001"
	notificationID := "notification-001"
	eventType := "notification"

	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/events/"+eventID, req.URL.String())

			rw.Header().Add("Content-Type", "application/json")
			rsp := `
			{
				"Event": "",
				"Id": "%s",
				"Type": "%s"
			}`
			_, _ = rw.Write([]byte(fmt.Sprintf(rsp, notificationID, eventType)))

		}))
	defer server.Close()

	t.Run("makes request to get an event", func(t *testing.T) {
		client := courier.CourierClient("key", server.URL)
		rsp, err := client.GetEvent(context.Background(), eventID)
		assert.Nil(t, err)
		assert.Equal(t, notificationID, rsp.Id)
		assert.Equal(t, eventType, rsp.Type)
	})
}

func TestClient_PutEvent(t *testing.T) {
	type UpsertEventRequest struct {
		Id   string
		Type string
	}
	var requestBody UpsertEventRequest

	eventID := "event-001"
	notificationID := "notification-001"
	eventType := "notification"

	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/events/"+eventID, req.URL.String())

			decoder := json.NewDecoder(req.Body)
			err := decoder.Decode(&requestBody)
			if err != nil {
				t.Error(err)
			}

			rw.Header().Add("Content-Type", "application/json")
			_, writeErr := rw.Write([]byte("{ \"Status\" : \"whatever\" }"))
			if writeErr != nil {
				t.Error(writeErr)
			}

		}))
	defer server.Close()

	t.Run("makes request to upsert an event", func(t *testing.T) {
		client := courier.CourierClient("key", server.URL)
		_, err := client.PutEvent(context.Background(), eventID, notificationID, eventType)
		assert.Nil(t, err)
		assert.Equal(t, notificationID, requestBody.Id)
		assert.Equal(t, eventType, requestBody.Type)
	})
}
