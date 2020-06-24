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

func TestProfiles_GetProfile(t *testing.T) {
	profileID := "example"
	foo := "bar"
	url := "/profiles/" + profileID

	server := httptest.NewServer(http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			assert.Equal(t, url, req.URL.String())
			rw.WriteHeader(http.StatusOK)
			rw.Header().Add("Content-Type", "application/json")
			rsp := `
				{
					"profile": {
						"foo": "%s"
					}
				}`
			_, _ = rw.Write([]byte(fmt.Sprintf(rsp, foo)))
		}),
	)
	defer server.Close()

	t.Run("makes requests for message ID", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		response, err := client.GetProfile(context.Background(), profileID)
		assert.Nil(t, err)

		var profile map[string]string
		err = json.Unmarshal(response["profile"], &profile)
		assert.Nil(t, err)

		assert.Equal(t, "bar", profile["foo"])
	})
}

func TestProfiles_ProfilePOST(t *testing.T) {
	type RequestBody struct {
		Profile interface{}
	}

	var requestBody RequestBody
	recipientID := "123456789"
	url := "/profiles/" + recipientID

	server := httptest.NewServer(http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			assert.Equal(t, url, req.URL.String())

			buf := new(bytes.Buffer)
			buf.ReadFrom(req.Body)
			bodyJSON := buf.String()
			log.Println("profiles_test.go: TestProfiles_MergeProfile")
			log.Println(bodyJSON)

			err := json.Unmarshal(buf.Bytes(), &requestBody)
			if err != nil {
				t.Error(err)
			}

			rw.WriteHeader(http.StatusOK)
			rw.Header().Add("Content-Type", "application/json")
			_, writeErr := rw.Write([]byte(fmt.Sprintf("{ \"status\" : \"%s\" }", "SUCCESS")))
			if writeErr != nil {
				t.Error(writeErr)
			}
		}),
	)
	defer server.Close()

	t.Run("sends request", func(t *testing.T) {
		profile := make(map[string]interface{})
		profile["email"] = "foo@bar.com"

		body := make(map[string]interface{})
		body["profile"] = profile

		bytes, err := json.Marshal(body)
		if err != nil {
			t.Error(err)
		}

		client := courier.CreateClient("key", &server.URL)
		_, err = client.API.SendRequestWithBytes(context.Background(), "POST", "/profiles/"+recipientID, bytes)

		assert.Nil(t, err)
	})
}
