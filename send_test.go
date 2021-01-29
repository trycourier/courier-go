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
	type Override struct {
		Bar string
	}
	type Preferences struct {
		Fizz string
	}
	type RequestBody struct {
		Event       string
		Recipient   string
		Profile     Profile
		Data        Data
		Brand       string
		Override    Override
		Preferences Preferences
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
		recipientID := "recpient-id"
		brand := "brand-id"
		override := &Override{
			Bar: "bar",
		}
		preferences := &Preferences{
			Fizz: "fizz",
		}
		messageID, err := client.Send(context.Background(), eventID, recipientID, profile, data, brand, override, preferences)

		assert.Nil(t, err)
		assert.Equal(t, expectedResponseID, messageID)
		assert.Equal(t, eventID, requestBody.Event)
		assert.Equal(t, recipientID, requestBody.Recipient)
		assert.Equal(t, data.Foo, requestBody.Data.Foo)
		assert.Equal(t, profile.Email, requestBody.Profile.Email)
		assert.Equal(t, brand, requestBody.Brand)
		assert.Equal(t, override.Bar, requestBody.Override.Bar)
		assert.Equal(t, preferences.Fizz, requestBody.Preferences.Fizz)
	})

}

func TestSendToList(t *testing.T) {

	type Data struct {
		Foo string
	}

	type Override struct {
		Bar string
	}

	type RequestBody struct {
		Event    string
		List     string
		Pattern  string
		Data     Data
		Brand    string
		Override Override
	}

	var requestBody RequestBody

	expectedResponseID := "123456789"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/send/list", req.URL.String())

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
		eventID := "event-id"
		listID := "list-id"
		brand := "brand-id"
		override := &Override{
			Bar: "bar",
		}
		messageID, err := client.SendToList(context.Background(), eventID, listID, "", data, brand, override)

		assert.Nil(t, err)
		assert.Equal(t, expectedResponseID, messageID)
		assert.Equal(t, eventID, requestBody.Event)
		assert.Equal(t, listID, requestBody.List)
		assert.Equal(t, data.Foo, requestBody.Data.Foo)
		assert.Equal(t, brand, requestBody.Brand)
		assert.Equal(t, override.Bar, requestBody.Override.Bar)
	})

}
