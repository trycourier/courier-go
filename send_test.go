package courier_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trycourier/courier-go"
)

func TestSendMap(t *testing.T) {
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
	url := "/send"
	messageID := "123456789"

	server := httptest.NewServer(http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			assert.Equal(t, url, req.URL.String())

			buf := new(bytes.Buffer)
			buf.ReadFrom(req.Body)
			bodyJSON := buf.String()
			log.Println("send_test.go: TestSendMap")
			log.Println(bodyJSON)

			err := json.Unmarshal(buf.Bytes(), &requestBody)
			if err != nil {
				t.Error(err)
			}

			rw.Header().Add("Content-Type", "application/json")
			_, writeErr := rw.Write([]byte(fmt.Sprintf("{ \"messageId\" : \"%s\" }", messageID)))
			if writeErr != nil {
				t.Error(writeErr)
			}
		}),
	)
	defer server.Close()

	t.Run("sends request", func(t *testing.T) {
		eventID := "event-id"
		recipientID := "recipient-id"

		profile := make(map[string]interface{})
		profile["email"] = "foo@bar.com"

		data := make(map[string]interface{})
		data["foo"] = "bar"

		body := make(map[string]interface{})
		body["profile"] = profile
		body["data"] = data

		client := courier.CreateClient("key", &server.URL)
		result, err := client.SendMap(context.Background(), eventID, recipientID, body)

		assert.Nil(t, err)
		assert.Equal(t, messageID, result)
		assert.Equal(t, eventID, requestBody.Event)
		assert.Equal(t, recipientID, requestBody.Recipient)
		assert.Equal(t, "foo@bar.com", requestBody.Profile.Email)
		assert.Equal(t, "bar", requestBody.Data.Foo)
	})
}

func TestSendStruct(t *testing.T) {
	type RequestData struct {
		Foo string
	}
	type RequestProfile struct {
		Email string
	}
	type RequestBody struct {
		Event     string
		Recipient string
		Profile   RequestProfile
		Data      RequestData
	}

	var requestBody RequestBody
	url := "/send"
	messageID := "123456789"

	server := httptest.NewServer(http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			assert.Equal(t, url, req.URL.String())

			buf := new(bytes.Buffer)
			buf.ReadFrom(req.Body)
			bodyJSON := buf.String()
			log.Println("send_test.go: TestSendStruct")
			log.Println(bodyJSON)

			err := json.Unmarshal(buf.Bytes(), &requestBody)
			if err != nil {
				t.Error(err)
			}

			rw.Header().Add("Content-Type", "application/json")
			_, writeErr := rw.Write([]byte(fmt.Sprintf("{ \"messageId\" : \"%s\" }", messageID)))
			if writeErr != nil {
				t.Error(writeErr)
			}
		}),
	)
	defer server.Close()

	t.Run("sends request", func(t *testing.T) {
		type data struct {
			Foo string `json:"foo"`
		}
		type profile struct {
			Email string `json:"email"`
		}

		eventID := "event-id"
		recipientID := "recipient-id"

		client := courier.CreateClient("key", &server.URL)
		result, err := client.Send(context.Background(), eventID, recipientID, courier.SendBody{
			profile{
				Email: "foo@bar.com",
			},
			data{
				Foo: "bar",
			},
		})

		assert.Nil(t, err)
		assert.Equal(t, messageID, result)
		assert.Equal(t, eventID, requestBody.Event)
		assert.Equal(t, recipientID, requestBody.Recipient)
		assert.Equal(t, "foo@bar.com", requestBody.Profile.Email)
		assert.Equal(t, "bar", requestBody.Data.Foo)
	})
}
