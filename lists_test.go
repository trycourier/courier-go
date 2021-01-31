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

func TestClient_GetLists(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/lists", req.URL.String())

			rw.Header().Add("Content-Type", "application/json")
			rsp := `
			{
				"Paging": {
						"Cursor": "",
						"More": false
				},
				"Items": [
						{
								"Id": "my-list",
								"Name": "",
								"Created": "2020-11-22T21:30:23.879Z",
								"Updated": "2020-11-24T20:06:12.332Z"
						}
				]
			}`
			_, _ = rw.Write([]byte(rsp))

		}))
	defer server.Close()

	t.Run("makes request to get all lists", func(t *testing.T) {
		client := courier.CourierClient("key", server.URL)
		rsp, err := client.GetLists(context.Background(), "", "")
		assert.Nil(t, err)
		assert.Equal(t, "my-list", rsp.Items[0].Id)
	})
}

func TestClient_GetList(t *testing.T) {
	listID := "my-list"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/lists/"+listID, req.URL.String())

			rw.Header().Add("Content-Type", "application/json")
			rsp := `
			{
				"Id": "%s",
				"Name": "",
				"Created": "2020-11-22T21:30:23.879Z",
				"Updated": "2020-11-24T20:06:12.332Z"
			}`
			_, _ = rw.Write([]byte(fmt.Sprintf(rsp, listID)))

		}))
	defer server.Close()

	t.Run("makes request to get a list", func(t *testing.T) {
		client := courier.CourierClient("key", server.URL)
		rsp, err := client.GetList(context.Background(), listID)
		assert.Nil(t, err)
		assert.Equal(t, listID, rsp.Id)
	})
}

func TestClient_PutList(t *testing.T) {
	listID := "my-list"
	listUpdatedName := "Besties"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/lists/"+listID, req.URL.String())

			rw.WriteHeader(http.StatusNoContent)

		}))
	defer server.Close()

	t.Run("makes request to update a list", func(t *testing.T) {
		client := courier.CourierClient("key", server.URL)
		client.PutList(context.Background(), listID, listUpdatedName)
	})
}

func TestClient_DeleteList(t *testing.T) {
	listID := "my-list"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/lists/"+listID, req.URL.String())

			rw.WriteHeader(http.StatusNoContent)

		}))
	defer server.Close()

	t.Run("makes request to delete a list", func(t *testing.T) {
		client := courier.CourierClient("key", server.URL)
		client.DeleteList(context.Background(), listID)
	})
}

func TestClient_RestoreList(t *testing.T) {
	listID := "my-list"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/lists/"+listID+"/restore", req.URL.String())

			rw.WriteHeader(http.StatusNoContent)

		}))
	defer server.Close()

	t.Run("makes request to restore a list", func(t *testing.T) {
		client := courier.CourierClient("key", server.URL)
		client.RestoreList(context.Background(), listID)
	})
}

func TestClient_GetListSubscriptions(t *testing.T) {
	listID := "my-list"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/lists/"+listID+"/subscriptions", req.URL.String())

			rw.Header().Add("Content-Type", "application/json")
			rsp := `
			{
				"Paging": {
						"Cursor": "",
						"More": false
				},
				"Items": [
						{
								"RecipientId": "5ed558d4-d2eb-4e0f-984a-81a0f04054b1",
								"Created": "2020-11-24T20:06:12.352Z"
						}
				]
			}`
			_, _ = rw.Write([]byte(rsp))

		}))
	defer server.Close()

	t.Run("makes request to get a list's all subscriptions", func(t *testing.T) {
		client := courier.CourierClient("key", server.URL)
		rsp, err := client.GetListSubscriptions(context.Background(), listID, "")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(rsp.Items))
		assert.Equal(t, "5ed558d4-d2eb-4e0f-984a-81a0f04054b1", rsp.Items[0].RecipientId)
	})
}

func TestClient_SubscribeMultipleRecipientsToList(t *testing.T) {
	listID := "my-list"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/lists/"+listID+"/subscriptions", req.URL.String())

			rw.WriteHeader(http.StatusNoContent)

		}))
	defer server.Close()

	t.Run("makes request to update a list", func(t *testing.T) {
		client := courier.CourierClient("key", server.URL)
		client.SubscribeMultipleRecipientsToList(context.Background(), listID, []courier.Recipient{{"recipient001"}})
	})
}

func TestClient_SubscribeRecipientToList(t *testing.T) {
	listID := "my-list"
	recipientID := "recipient001"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/lists/"+listID+"/subscriptions/"+recipientID, req.URL.String())

			rw.WriteHeader(http.StatusNoContent)

		}))
	defer server.Close()

	t.Run("makes request to update a list", func(t *testing.T) {
		client := courier.CourierClient("key", server.URL)
		client.SubscribeRecipientToList(context.Background(), listID, recipientID)
	})
}

func TestClient_DeleteListSubscription(t *testing.T) {
	listID := "my-list"
	recipientID := "recipient001"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/lists/"+listID+"/subscriptions/"+recipientID, req.URL.String())

			rw.WriteHeader(http.StatusNoContent)

		}))
	defer server.Close()

	t.Run("makes request to update a list", func(t *testing.T) {
		client := courier.CourierClient("key", server.URL)
		client.DeleteListSubscription(context.Background(), listID, recipientID)
	})
}
