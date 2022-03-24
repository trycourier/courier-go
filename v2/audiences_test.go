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

func TestPutAudience(t *testing.T) {
	// expectedResponseID := "software-engineers-from-sf"
	server := httptest.NewServer(http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			assert.Equal(t, "/audiences/software-engineers-from-sf", req.URL.String())
			rw.Header().Add("Content-Type", "application/json")
			rsp := `
			{
				"audience": {
					"id": "software-engineers-from-sf",
					"name": "Software Engineers From SF",
					"description": "Updated descriptionss",
					"created_at": "2022-03-22T19:07:32.066Z",
					"updated_at": "2022-03-24T00:32:11.122Z",
					"filter": {
						"path": "title",
						"value": "Software Engineer",
						"operator": "EQ"
					}
				}
			}`
			_, _ = rw.Write([]byte(rsp))

		}))
	defer server.Close()

	t.Run("Put Audience", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)

		audiencePutResponse, err := client.PutAudience(context.Background(), "software-engineers-from-sf", courier.Audience{
			Description: "Updated descriptionss",
			Name:        "Software Engineers From SF",
			Filter: courier.SingleFilter{
				Operator: "EQ",
				Path:     "title",
				Value:    "Software Engineer",
			},
		})
		assert.Nil(t, err)

		assert.Equal(t, "software-engineers-from-sf", audiencePutResponse.Audience.Id)
	})
}

func TestGetAudience(t *testing.T) {
	expectedResponseID := "123456789"
	server := httptest.NewServer(http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			assert.Equal(t, "/audiences/123456789", req.URL.String())
			rw.Header().Add("Content-Type", "application/json")
			_, writeErr := rw.Write([]byte(fmt.Sprintf("{ \"id\" : \"%s\" }", expectedResponseID)))
			if writeErr != nil {
				t.Error(writeErr)
			}
		}))
	defer server.Close()

	t.Run("Get Audience", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		response, err := client.GetAudience(context.Background(), "123456789")
		assert.Nil(t, err)
		assert.Equal(t, expectedResponseID, response.Id)
	})
}
