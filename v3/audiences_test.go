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

func TestPutAudience(t *testing.T) {
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

	t.Run("Put Audience Happy Path", func(t *testing.T) {
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

	t.Run("Get Audience Happy Path", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		response, err := client.GetAudience(context.Background(), "123456789")
		assert.Nil(t, err)
		assert.Equal(t, expectedResponseID, response.Id)
	})

	t.Run("should throw Audience ID required if empty", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		_, err := client.GetAudience(context.Background(), "")
		assert.Equal(t, err.Error(), "Audience ID is required")
	})
}

func TestGetAudienceMembers(t *testing.T) {
	expectedMemberId := "test-member-id"
	server := httptest.NewServer(http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			assert.Equal(t, "/audiences/123456789/members", req.URL.String())
			rw.Header().Add("Content-Type", "application/json")
			rsp := `
			{
				"items": [
					{
						"added_at": "2022-03-22T19:13:13.137Z",
						"audience_id": "software-engineers-from-sf",
						"audience_version": 3,
						"member_id": "test-member-id",
						"reason": "EQ('title', 'Software Engineer') => true"
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

	t.Run("Get Audience Members Happy Path", func(t *testing.T) {
		expectedAudienceVersion := 3
		client := courier.CreateClient("key", &server.URL)
		response, err := client.GetAudienceMembers(context.Background(), "123456789", "")
		assert.Nil(t, err)
		assert.Equal(t, expectedMemberId, response.Items[0].MemberId)
		assert.Equal(t, expectedAudienceVersion, response.Items[0].AudienceVersion)
	})
}

func TestAudienceMembersError(t *testing.T) {
	expectedResponseID := "123456789"
	server := httptest.NewServer(http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			rw.Header().Add("Content-Type", "application/json")
			_, writeErr := rw.Write([]byte(fmt.Sprintf("{ \"id\" : \"%s\" }", expectedResponseID)))
			if writeErr != nil {
				t.Error(writeErr)
			}
		}))
	defer server.Close()

	t.Run("should throw Audience ID required if empty", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		_, err := client.GetAudience(context.Background(), "")
		assert.Equal(t, err.Error(), "Audience ID is required")
	})
}

func TestGetAudiences(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			assert.Equal(t, "/audiences", req.URL.String())
			rw.Header().Add("Content-Type", "application/json")
			rsp := `
			{
				"items": [
					{
						"created_at": "2022-03-22T19:07:32.066Z",
						"description": "Updated descriptionss",
						"id": "software-engineers-from-sf",
						"name": "Software Engineers From SF",
						"filter": {
							"path": "title",
							"value": "Software Engineer",
							"operator": "EQ"
						},
						"updated_at": "2022-03-24T05:30:58.659Z"
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

	t.Run("Get Audiences Happy path", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		response, err := client.GetAudiences(context.Background(), "")
		assert.Nil(t, err)
		assert.Equal(t, "software-engineers-from-sf", response.Items[0].Id)
		assert.Equal(t, "Software Engineers From SF", response.Items[0].Name)
	})
}

func TestDeleteAudience(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			assert.Equal(t, "/audiences/123456789", req.URL.String())
			rw.Header().Add("Content-Type", "application/json")
			rw.WriteHeader(http.StatusNoContent)
		}))
	defer server.Close()

	t.Run("Delete Audience Happy Path", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		err := client.DeleteAudience(context.Background(), "123456789")
		assert.NoError(t, err)
	})

	t.Run("should throw Audience ID required if empty", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		err := client.DeleteAudience(context.Background(), "")
		assert.EqualError(t, err, "Audience ID is required")
	})
}
