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

func TestGetAuditEvent(t *testing.T) {
	expectedResponseID := "123456789"
	server := httptest.NewServer(http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			assert.Equal(t, "/audit-events/123456789", req.URL.String())
			rw.Header().Add("Content-Type", "application/json")
			_, writeErr := rw.Write([]byte(fmt.Sprintf("{ \"auditEventId\" : \"%s\" }", expectedResponseID)))
			if writeErr != nil {
				t.Error(writeErr)
			}
		}))
	defer server.Close()

	t.Run("Get Audit Event Happy Path", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		response, err := client.GetAuditEvent(context.Background(), "123456789")
		assert.Nil(t, err)
		assert.Equal(t, expectedResponseID, response.AuditEventId)
	})

	t.Run("should throw Audit Event ID required if empty", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		_, err := client.GetAuditEvent(context.Background(), "")
		assert.Equal(t, err.Error(), "audit Event ID is required")
	})
}

func TestListAuditEvents(t *testing.T) {
	expectedAuditEventId := "ZX3xXUMNKL4y2NkiKgstl"
	server := httptest.NewServer(http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			assert.Equal(t, "/audit-events", req.URL.String())
			rw.Header().Add("Content-Type", "application/json")
			rsp := `
			{
				"results": [
					{
						"actor": {
							"email": "foo@bar.com",
							"id": "foo-user"
						},
						"auditEventId": "ZX3xXUMNKL4y2NkiKgstl",
						"source": "courier.studio",
						"target": {},
						"timestamp": "2022-07-22T22:33:59.552Z",
						"type": "notification:published"
					}
				],
				"paging": {
					"cursor": null,
					"more": false
				}
			}`
			_, _ = rw.Write([]byte(rsp))
		}))
	defer server.Close()

	t.Run("List Audit Events Happy Path", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		response, err := client.ListAuditEvents(context.Background(), "")
		assert.Nil(t, err)
		assert.Equal(t, expectedAuditEventId, response.Results[0].AuditEventId)
		assert.Equal(t, "foo-user", response.Results[0].Actor.Id)
	})
}
