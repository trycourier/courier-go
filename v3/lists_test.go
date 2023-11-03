package courier_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trycourier/courier-go/v3"
)

func TestGetLists(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/lists", req.URL.String())

			rw.Header().Add("Content-Type", "application/json")
			rsp := `
			{
				"Paging": {
						"Cursor": null,
						"More": false
				},
				"Items": [
						{
								"ID": "my-list",
								"Name": "My Awesome List",
								"Created": "2020-11-22T21:30:23.879Z",
								"Updated": "2020-11-24T20:06:12.332Z"
						}
				]
			}`
			_, _ = rw.Write([]byte(rsp))

		}))
	defer server.Close()

	t.Run("Get Lists", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		rsp, err := client.GetLists(context.Background(), "", "")
		assert.Nil(t, err)
		assert.Equal(t, "my-list", rsp.Items[0].ID)
	})
}

func TestGetList(t *testing.T) {
	listID := "my-list"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/lists/"+listID, req.URL.String())

			rw.Header().Add("Content-Type", "application/json")
			rsp := `
			{
				"ID": "%s",
				"Name": "",
				"Created": "2020-11-22T21:30:23.879Z",
				"Updated": "2020-11-24T20:06:12.332Z"
			}`
			_, _ = rw.Write([]byte(fmt.Sprintf(rsp, listID)))

		}))
	defer server.Close()

	t.Run("Get List", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		rsp, err := client.GetList(context.Background(), listID)
		assert.Nil(t, err)
		assert.Equal(t, listID, rsp.ID)
	})
}

func TestPutList(t *testing.T) {
	listID := "my-list"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/lists/"+listID, req.URL.String())

			rw.WriteHeader(http.StatusNoContent)

		}))
	defer server.Close()

	t.Run("Put List", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		client.PutList(context.Background(), listID, courier.PutListBody{
			Name: "Besties",
		})
	})
}

func TestDeleteList(t *testing.T) {
	listID := "my-list"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/lists/"+listID, req.URL.String())

			rw.WriteHeader(http.StatusNoContent)

		}))
	defer server.Close()

	t.Run("Delete List", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		client.DeleteList(context.Background(), listID)
	})
}

func TestRestoreList(t *testing.T) {
	listID := "my-list"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/lists/"+listID+"/restore", req.URL.String())

			rw.WriteHeader(http.StatusNoContent)

		}))
	defer server.Close()

	t.Run("Restore List", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		client.RestoreList(context.Background(), listID)
	})
}

func TestGetListSubscriptions(t *testing.T) {
	listID := "my-list"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/lists/"+listID+"/subscriptions", req.URL.String())

			rw.Header().Add("Content-Type", "application/json")
			rsp := `
			{
				"Paging": {
						"Cursor": null,
						"More": false
				},
				"Items": [
						{
								"RecipientID": "5ed558d4-d2eb-4e0f-984a-81a0f04054b1",
								"Created": "2020-11-24T20:06:12.352Z"
						}
				]
			}`
			_, _ = rw.Write([]byte(rsp))

		}))
	defer server.Close()

	t.Run("Get list subscriptions", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		rsp, err := client.GetListSubscriptions(context.Background(), listID, "")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(rsp.Items))
		assert.Equal(t, "5ed558d4-d2eb-4e0f-984a-81a0f04054b1", rsp.Items[0].RecipientID)
	})
}

func TestPutListSubscriptions(t *testing.T) {
	listID := "my-list"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/lists/"+listID+"/subscriptions", req.URL.String())

			rw.WriteHeader(http.StatusNoContent)

		}))
	defer server.Close()

	t.Run("Put list subscriptions", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		client.PutListSubscriptions(context.Background(), listID, courier.ListSubscriptionBody{
			Recipients: []courier.ListRecipient{
				courier.ListRecipient{
					RecipientID: "0460766e-8463-4905-ae98-b72c7aef41d6",
				},
			},
		})
	})
}

func TestPostListSubscriptions(t *testing.T) {
	listID := "my-list"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/lists/"+listID+"/subscriptions", req.URL.String())

			rw.WriteHeader(http.StatusNoContent)

		}))
	defer server.Close()

	t.Run("Post list subscriptions", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		client.PostListSubscriptions(context.Background(), listID, courier.ListSubscriptionBody{
			Recipients: []courier.ListRecipient{
				courier.ListRecipient{
					RecipientID: "0460766e-8463-4905-ae98-b72c7aef41d6",
				},
			},
		})
	})
}

func TestListSubscribe(t *testing.T) {
	listID := "my-list"
	recipientID := "foo"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/lists/"+listID+"/subscriptions/"+recipientID, req.URL.String())

			rw.WriteHeader(http.StatusNoContent)

		}))
	defer server.Close()

	t.Run("subscribe to a list", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		client.ListSubscribe(context.Background(), listID, recipientID, courier.ListSubscriptionRecipientBody{})
	})
}

func TestListUnsubscribe(t *testing.T) {
	listID := "my-list"
	recipientID := "foo"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/lists/"+listID+"/subscriptions/"+recipientID, req.URL.String())

			rw.WriteHeader(http.StatusNoContent)

		}))
	defer server.Close()

	t.Run("unsubscribe from a list", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		client.ListUnsubscribe(context.Background(), listID, recipientID)
	})
}
