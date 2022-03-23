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
	"github.com/trycourier/courier-go/v2"
)

func TestSend_SendMap(t *testing.T) {
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

			rw.WriteHeader(http.StatusOK)
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

func TestSend_Send(t *testing.T) {
	type RequestData struct {
		Foo string
	}
	type RequestProfile struct {
		Email string
	}
	type RequestBodyOverride struct {
		ReplyBroadcast bool `json:"reply_broadcast"`
	}
	type RequestSlackOverride struct {
		Body RequestBodyOverride
	}
	type RequestOverride struct {
		Slack RequestSlackOverride
	}
	type RequestBody struct {
		Event     string
		Recipient string
		Profile   RequestProfile
		Data      RequestData
		Override  RequestOverride
		Brand     string
	}
	type data struct {
		Foo string `json:"foo"`
	}
	type profile struct {
		Email string `json:"email"`
	}

	var requestBody RequestBody
	url := "/send"
	messageID := "123456789"
	eventID := "event-id"
	recipientID := "recipient-id"

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

			rw.WriteHeader(http.StatusOK)
			rw.Header().Add("Content-Type", "application/json")
			_, writeErr := rw.Write([]byte(fmt.Sprintf("{ \"messageId\" : \"%s\" }", messageID)))
			if writeErr != nil {
				t.Error(writeErr)
			}
		}),
	)
	defer server.Close()

	t.Run("sends request", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		result, err := client.Send(context.Background(), eventID, recipientID, courier.SendBody{
			Profile: profile{
				Email: "foo@bar.com",
			},
			Data: data{
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

	t.Run("sends request with override", func(t *testing.T) {
		type bodyOverride struct {
			ReplyBroadcast bool `json:"reply_broadcast"`
		}
		type slackOverride struct {
			Body bodyOverride `json:"body"`
		}
		type override struct {
			Slack slackOverride `json:"slack"`
		}

		client := courier.CreateClient("key", &server.URL)
		result, err := client.Send(context.Background(), eventID, recipientID, courier.SendBody{
			Profile: profile{
				Email: "foo@bar.com",
			},
			Data: data{
				Foo: "bar",
			},
			Override: override{
				Slack: slackOverride{
					Body: bodyOverride{
						ReplyBroadcast: true,
					},
				},
			},
		})

		assert.Nil(t, err)
		assert.Equal(t, messageID, result)
		assert.Equal(t, eventID, requestBody.Event)
		assert.Equal(t, recipientID, requestBody.Recipient)
		assert.Equal(t, "foo@bar.com", requestBody.Profile.Email)
		assert.Equal(t, "bar", requestBody.Data.Foo)
		assert.Equal(t, true, requestBody.Override.Slack.Body.ReplyBroadcast)
	})

	t.Run("sends request with brand", func(t *testing.T) {
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
			Profile: profile{
				Email: "foo@bar.com",
			},
			Data: data{
				Foo: "bar",
			},
			Brand: "dispatcher",
		})

		assert.Nil(t, err)
		assert.Equal(t, messageID, result)
		assert.Equal(t, eventID, requestBody.Event)
		assert.Equal(t, recipientID, requestBody.Recipient)
		assert.Equal(t, "foo@bar.com", requestBody.Profile.Email)
		assert.Equal(t, "bar", requestBody.Data.Foo)
		assert.Equal(t, "dispatcher", requestBody.Brand)
	})
}

func TestSend_Send_400Error(t *testing.T) {
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

			rw.WriteHeader(http.StatusBadRequest)
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
		_, err := client.Send(context.Background(), eventID, recipientID, courier.SendBody{
			Profile: profile{
				Email: "foo@bar.com",
			},
			Data: data{
				Foo: "bar",
			},
		})

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusBadRequest, err.(*courier.HTTPError).StatusCode)
	})
}

func TestSend_SendMessage(t *testing.T) {
	var requestBody courier.SendMessageRequestBody
	url := "/send"
	requestID := "123456789"

	server := httptest.NewServer(http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			assert.Equal(t, url, req.URL.String())

			buf := new(bytes.Buffer)
			buf.ReadFrom(req.Body)

			err := json.Unmarshal(buf.Bytes(), &requestBody)
			if err != nil {
				t.Error(err)
			}

			rw.WriteHeader(http.StatusOK)
			rw.Header().Add("Content-Type", "application/json")
			_, writeErr := rw.Write([]byte(fmt.Sprintf("{ \"requestId\" : \"%s\" }", requestID)))
			if writeErr != nil {
				t.Error(writeErr)
			}
		}),
	)
	defer server.Close()

	t.Run("sends request", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		result, err := client.SendMessage(context.Background(), courier.SendMessageRequestBody{
			Message: map[string]interface{}{
				"template": "my-template",
				"to": map[string]string{
					"email": "foo@bar.com",
				},
			},
		})

		assert.Nil(t, err)
		assert.Equal(t, requestID, result)
	})
}

func TestSend_SendMessage_Timeout(t *testing.T) {
	var requestBody courier.SendMessageRequestBody
	url := "/send"
	requestID := "123456789"

	server := httptest.NewServer(http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			assert.Equal(t, url, req.URL.String())

			buf := new(bytes.Buffer)
			buf.ReadFrom(req.Body)

			err := json.Unmarshal(buf.Bytes(), &requestBody)
			if err != nil {
				t.Error(err)
			}

			rw.WriteHeader(http.StatusOK)
			rw.Header().Add("Content-Type", "application/json")
			_, writeErr := rw.Write([]byte(fmt.Sprintf("{ \"requestId\" : \"%s\" }", requestID)))
			if writeErr != nil {
				t.Error(writeErr)
			}
		}),
	)
	defer server.Close()

	t.Run("sends request with timeout", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		result, err := client.SendMessage(context.Background(), courier.SendMessageRequestBody{
			Message: map[string]interface{}{
				"template": "my-template",
				"to": map[string]string{
					"email": "foo@bar.com",
				},
				"timeout": int64(3600000), // 1 hour in milliseconds
			},
		})

		assert.Nil(t, err)
		assert.Equal(t, requestID, result)
	})
}
