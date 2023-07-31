package courier_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trycourier/courier-go/v2"
)

func TestGetAccount(t *testing.T) {
	expectedResponseID := "123456789"
	server := httptest.NewServer(http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			assert.Equal(t, "/accounts/123456789", req.URL.String())
			rw.Header().Add("Content-Type", "application/json")
			_, writeErr := rw.Write([]byte(fmt.Sprintf("{ \"id\" : \"%s\" }", expectedResponseID)))
			if writeErr != nil {
				t.Error(writeErr)
			}
		}))
	defer server.Close()

	t.Run("Get Account", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		response, err := client.GetAccount(context.Background(), "123456789")
		assert.Nil(t, err)
		assert.Equal(t, expectedResponseID, response.Id)
	})
}

func TestGetAccounts(t *testing.T) {
	expectedAccountId := "ZX3xXUMNKL4y2NkiKgstl"
	server := httptest.NewServer(http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			assert.Equal(t, "/accounts", req.URL.String())
			rw.Header().Add("Content-Type", "application/json")
			rsp := `
			{
				"items": [
					{
						"id": "ZX3xXUMNKL4y2NkiKgstl",
						"name": "my-account"
					}
				],
				"has_more": false,
				"url": "/accounts"
			}`
			_, _ = rw.Write([]byte(rsp))
		}))
	defer server.Close()

	t.Run("Get Accounts", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		response, err := client.GetAccounts(context.Background(), "", "")
		assert.Nil(t, err)
		assert.Equal(t, expectedAccountId, response.Items[0].Id)
		assert.Equal(t, "my-account", response.Items[0].Name)
		assert.Equal(t, "/accounts", response.Url)
	})
}

func TestPutAccount(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			assert.Equal(t, "/accounts/account-1", req.URL.String())
			rw.Header().Add("Content-Type", "application/json")
			rsp := `
			{
				"id": "account-1",
				"name": "My Account"
			}`
			_, _ = rw.Write([]byte(rsp))

		}))
	defer server.Close()

	t.Run("Put Account", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)

		accountPutResponse, err := client.PutAccount(context.Background(), "account-1", map[string]string{
			"name": "My Account",
		})
		assert.Nil(t, err)

		assert.Equal(t, "account-1", accountPutResponse.Id)
	})
}

func TestDeleteAccount(t *testing.T) {
	accountID := "my-account"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/accounts/"+accountID, req.URL.String())

			rw.WriteHeader(http.StatusNoContent)

		}))
	defer server.Close()

	t.Run("Delete Account", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		client.DeleteAccount(context.Background(), accountID)
	})
}
