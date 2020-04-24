package courier_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/expel-io/courier-go"

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
			// nolint
			rw.Write([]byte(fmt.Sprintf("{ \"MessageId\" : \"%s\" }", expectedResponseID)))

		}))
	defer server.Close()

	t.Run("sends request", func(t *testing.T) {

		client := courier.CourierClient("key", server.URL)

		myData := &Data{
			Foo: "bar",
		}
		myProfile := &Profile{
			Email: "foo@bar.com",
		}
		eventID := "event-id"
		recipientID := "recpient-id"
		messageID, err := client.Send(context.Background(), eventID, recipientID, myProfile, myData)

		assert.Nil(t, err)
		assert.Equal(t, expectedResponseID, messageID)
		assert.Equal(t, eventID, requestBody.Event)
		assert.Equal(t, recipientID, requestBody.Recipient)
		assert.Equal(t, myData.Foo, requestBody.Data.Foo)
		assert.Equal(t, myProfile.Email, requestBody.Profile.Email)
	})

}
