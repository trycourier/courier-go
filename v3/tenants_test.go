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

func TestGetTenant(t *testing.T) {
	expectedResponseID := "123456789"
	server := httptest.NewServer(http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			assert.Equal(t, "/tenants/123456789", req.URL.String())
			rw.Header().Add("Content-Type", "application/json")
			_, writeErr := rw.Write([]byte(fmt.Sprintf("{ \"id\" : \"%s\" }", expectedResponseID)))
			if writeErr != nil {
				t.Error(writeErr)
			}
		}))
	defer server.Close()

	t.Run("Get Tenant", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		response, err := client.GetTenant(context.Background(), "123456789")
		assert.Nil(t, err)
		assert.Equal(t, expectedResponseID, response.Id)
	})
}

func TestGetTenants(t *testing.T) {
	expectedAccountId := "ZX3xXUMNKL4y2NkiKgstl"
	server := httptest.NewServer(http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			assert.Equal(t, "/tenants", req.URL.String())
			rw.Header().Add("Content-Type", "application/json")
			rsp := `
			{
				"items": [
					{
						"id": "ZX3xXUMNKL4y2NkiKgstl",
						"name": "my-tenant"
					}
				],
				"has_more": false,
				"url": "/tenants"
			}`
			_, _ = rw.Write([]byte(rsp))
		}))
	defer server.Close()

	t.Run("Get Accounts", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		response, err := client.GetTenants(context.Background(), "", "")
		assert.Nil(t, err)
		assert.Equal(t, expectedAccountId, response.Items[0].Id)
		assert.Equal(t, "my-tenant", response.Items[0].Name)
		assert.Equal(t, "/tenants", response.Url)
	})
}

func TestPutTenant(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			assert.Equal(t, "/tenants/tenant-1", req.URL.String())
			rw.Header().Add("Content-Type", "application/json")
			rsp := `
			{
				"id": "tenant-1",
				"name": "My Tenant"
			}`
			_, _ = rw.Write([]byte(rsp))

		}))
	defer server.Close()

	t.Run("Put Account", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)

		accountPutResponse, err := client.PutTenant(context.Background(), "tenant-1", map[string]string{
			"name": "My Tenant",
		})
		assert.Nil(t, err)

		assert.Equal(t, "tenant-1", accountPutResponse.Id)
	})
}

func TestDeleteAccount(t *testing.T) {
	tenantId := "my-tenant"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/tenants/"+tenantId, req.URL.String())

			rw.WriteHeader(http.StatusNoContent)

		}))
	defer server.Close()

	t.Run("Delete Account", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		client.DeleteTenant(context.Background(), tenantId)
	})
}
