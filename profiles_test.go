package courier_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trycourier/courier-go"
)

func TestClient_GetProfile(t *testing.T) {
	recipientID := "recipient001"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/profiles/"+recipientID, req.URL.String())

			rw.Header().Add("Content-Type", "application/json")

			_, _ = rw.Write([]byte(`{"profile": {"email":"foo@bar.com","first_name":"John","last_name":"Doe","phone_number":"+11234567890"}}`))

		}))
	defer server.Close()

	t.Run("makes request to get a profile", func(t *testing.T) {
		client := courier.CourierClient("key", server.URL, "", "")
		rsp, err := client.GetProfile(recipientID)
		assert.Nil(t, err)
		assert.NotNil(t, rsp)
	})
}

func TestClient_MergeProfile(t *testing.T) {
	recipientID := "recipient001"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/profiles/"+recipientID, req.URL.String())

			rw.WriteHeader(http.StatusNoContent)

		}))
	defer server.Close()

	t.Run("makes request to merge a profile", func(t *testing.T) {
		client := courier.CourierClient("key", server.URL, "", "")
		client.MergeProfile(recipientID, []byte(`{"profile": {"email":"foo@bar.com","first_name":"John","last_name":"Doe","phone_number":"+11234567890"}}`))
	})
}

func TestClient_UpdateProfile(t *testing.T) {
	recipientID := "recipient001"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/profiles/"+recipientID, req.URL.String())

			rw.WriteHeader(http.StatusNoContent)

		}))
	defer server.Close()

	t.Run("makes request to update a profile", func(t *testing.T) {
		client := courier.CourierClient("key", server.URL, "", "")
		client.UpdateProfile(recipientID, []byte(`{"profile": {"email":"buzz@bar.com","first_name":"Jane","last_name":"Doe","phone_number":"+11234567890"}}`))
	})
}

func TestClient_PatchProfile(t *testing.T) {
	recipientID := "recipient001"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/profiles/"+recipientID, req.URL.String())

			rw.WriteHeader(http.StatusNoContent)

		}))
	defer server.Close()

	t.Run("makes request to update a profile", func(t *testing.T) {
		client := courier.CourierClient("key", server.URL, "", "")
		client.PatchProfile(recipientID, []courier.PatchOp{})
	})
}

func TestClient_GetProfileLists(t *testing.T) {
	recipientID := "recipient001"
	listID := "list001"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/profiles/"+recipientID+"/lists", req.URL.String())

			rw.Header().Add("Content-Type", "application/json")

			rsp := `
			{
				"Paging": {
						"Cursor": "",
						"More": false
				},
				"Results": [
						{
								"Id": "%s",
								"Name": "My List",
								"Created": "2020-11-22T21:30:23.879Z",
								"Updated": "2021-01-31T21:17:19.019Z"
						}
				]
			}`
			_, _ = rw.Write([]byte(fmt.Sprintf(rsp, listID)))

		}))
	defer server.Close()

	t.Run("makes request to get profile lists", func(t *testing.T) {
		client := courier.CourierClient("key", server.URL, "", "")
		rsp, err := client.GetProfileLists(recipientID, "")
		assert.Nil(t, err)
		assert.NotNil(t, listID, rsp.Results[0].Id)
	})
}
