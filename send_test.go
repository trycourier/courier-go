package courier_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/trycourier/courier-go"

	"github.com/stretchr/testify/assert"
)

func TestSend(t *testing.T) {

	type Data struct {
		Foo string
	}
	type Profile struct {
		Email string
	}
	type RequestBody struct {
		Event     string
		Recipient string
		Brand     string
		Profile   Profile
		Data      Data
	}

	var requestBody RequestBody

	expectedResponseID := "123456789"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/send", req.URL.String())

			decoder := json.NewDecoder(req.Body)
			err := decoder.Decode(&requestBody)
			if err != nil {
				t.Error(err)
			}

			rw.Header().Add("Content-Type", "application/json")
			_, writeErr := rw.Write([]byte(fmt.Sprintf("{ \"MessageId\" : \"%s\" }", expectedResponseID)))
			if writeErr != nil {
				t.Error(writeErr)
			}

		}))
	defer server.Close()

	t.Run("sends request", func(t *testing.T) {

		client := courier.CourierClient("key", server.URL)

		data := &Data{
			Foo: "bar",
		}
		profile := &Profile{
			Email: "foo@bar.com",
		}
		eventID := "event-id"
		recipientID := "recipient-id"
		brandID := "brand-id"
		messageID, err := client.Send(context.Background(), eventID, recipientID, brandID, profile, data)

		assert.Nil(t, err)
		assert.Equal(t, expectedResponseID, messageID)
		assert.Equal(t, eventID, requestBody.Event)
		assert.Equal(t, recipientID, requestBody.Recipient)
		assert.Equal(t, brandID, requestBody.Brand)
		assert.Equal(t, data.Foo, requestBody.Data.Foo)
		assert.Equal(t, profile.Email, requestBody.Profile.Email)
	})

}
