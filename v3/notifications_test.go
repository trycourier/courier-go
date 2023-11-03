package courier_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trycourier/courier-go/v3"
)

func TestGetNotifications(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/notifications", req.URL.String())

			rw.Header().Add("Content-Type", "application/json")

			rsp := `
			{
				"Paging": {
						"Cursor": null,
						"More": false
				},
				"Results": [
						{
								"ID": "my-notification",
								"Title": "My Awesome Notification"
						}
				]
			}`
			_, _ = rw.Write([]byte(rsp))

		}))
	defer server.Close()

	t.Run("Get Notifications", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		rsp, err := client.GetNotifications(context.Background(), "")
		assert.Nil(t, err)
		assert.Equal(t, "my-notification", rsp.Results[0].ID)
		assert.Equal(t, "My Awesome Notification", rsp.Results[0].Title)
	})
}

func TestGetNotificationContent(t *testing.T) {
	notificationID := "my-notification"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/notifications/"+notificationID+"/content", req.URL.String())

			rw.Header().Add("Content-Type", "application/json")

			rsp := `
			{
				"blocks": [{
					"id": "block_00d0bcb0-aba1-443f-a8dd-daac505500fe",
					"type": "text",
					"content": "text block",
					"checksum": "block-checksum",
					"locales": {
						"fr_FR": "french text block"
					}
				}],
				"channels": [{
					"id": "channel_79d25574-83be-4da1-a5c3-3d4e2ab5f154",
					"type": "email",
					"content": {
						"subject": "Hey hey!"
					},
					"checksum": "channel-checksum",
        	"locales": {
          	"fr_FR": {
							"subject": "French hey!"
          	}
        	}
				}]
			}`
			_, _ = rw.Write([]byte(rsp))

		}))
	defer server.Close()

	t.Run("Get Notification Content", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		rsp, err := client.GetNotificationContent(context.Background(), notificationID)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(rsp.Blocks))
		assert.Equal(t, 1, len(rsp.Channels))
	})
}

func TestGetNotificationDraftContent(t *testing.T) {
	notificationID := "my-notification"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/notifications/"+notificationID+"/draft/content", req.URL.String())

			rw.Header().Add("Content-Type", "application/json")

			rsp := `
			{
				"blocks": [{
					"id": "block_00d0bcb0-aba1-443f-a8dd-daac505500fe",
					"type": "text",
					"content": "text block",
					"checksum": "block-checksum",
					"locales": {
						"fr_FR": "french text block"
					}
				}],
				"channels": [{
					"id": "channel_79d25574-83be-4da1-a5c3-3d4e2ab5f154",
					"type": "email",
					"content": {
						"subject": "Hey hey!"
					},
					"checksum": "channel-checksum",
        	"locales": {
          	"fr_FR": {
							"subject": "French hey!"
          	}
        	}
				}]
			}`
			_, _ = rw.Write([]byte(rsp))

		}))
	defer server.Close()

	t.Run("Get Notification Draft Content", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		rsp, err := client.GetNotificationDraftContent(context.Background(), notificationID)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(rsp.Blocks))
		assert.Equal(t, 1, len(rsp.Channels))
	})
}

func TestPostNotificationVariations(t *testing.T) {
	notificationID := "my-notification"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/notifications/"+notificationID+"/variations", req.URL.String())

			rw.WriteHeader(http.StatusNoContent)

		}))
	defer server.Close()

	t.Run("Post Notification Variations", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		client.PostNotificationVariations(context.Background(), notificationID, courier.NotificationVariationsRequestBody{
			Blocks: []interface{}{
				map[string]string{"id": "block_00d0bcb0-aba1-443f-a8dd-daac505500fe"},
				map[string]string{"type": "text"},
				// ...locales
			},
			Channels: []interface{}{
				map[string]string{"id": "channel_79d25574-83be-4da1-a5c3-3d4e2ab5f154"},
				// ...locales
			},
		})
	})
}

func TestPostNotificationDraftVariations(t *testing.T) {
	notificationID := "my-notification"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/notifications/"+notificationID+"/draft/variations", req.URL.String())

			rw.WriteHeader(http.StatusNoContent)

		}))
	defer server.Close()

	t.Run("Post Notification Draft Variations", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		client.PostNotificationDraftVariations(context.Background(), notificationID, courier.NotificationVariationsRequestBody{
			Blocks: []interface{}{
				map[string]string{"id": "block_00d0bcb0-aba1-443f-a8dd-daac505500fe"},
				map[string]string{"type": "text"},
				// ...locales
			},
			Channels: []interface{}{
				map[string]string{"id": "channel_79d25574-83be-4da1-a5c3-3d4e2ab5f154"},
				// ...locales
			},
		})
	})
}

func TestGetNotificationSubmissionChecks(t *testing.T) {
	notificationID := "my-notification"
	submissionID := "my-submission"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/notifications/"+notificationID+"/"+submissionID+"/checks", req.URL.String())

			rw.Header().Add("Content-Type", "application/json")

			rsp := `
			{
				"Checks": [
						{
								"ID": "my-check",
								"Status": "PENDING",
								"Type": "custom",
								"Updated": 1629169195778
						}
				]
			}`
			_, _ = rw.Write([]byte(rsp))

		}))
	defer server.Close()

	t.Run("Get Notification Submission Checks", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		rsp, err := client.GetNotificationSubmissionChecks(context.Background(), notificationID, submissionID)
		assert.Nil(t, err)
		assert.Equal(t, "my-check", rsp.Checks[0].ID)
		assert.Equal(t, "PENDING", rsp.Checks[0].Status)
		assert.Equal(t, "custom", rsp.Checks[0].Type)
		assert.Equal(t, 1629169195778, rsp.Checks[0].Updated)
	})
}

func TestPutNotificationSubmissionChecks(t *testing.T) {
	notificationID := "my-notification"
	submissionID := "my-submission"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/notifications/"+notificationID+"/"+submissionID+"/checks", req.URL.String())

			rw.Header().Add("Content-Type", "application/json")

			rsp := `
			{
				"Checks": [
						{
								"ID": "my-check",
								"Status": "RESOLVED",
								"Type": "custom",
								"Updated": 1629169195779
						}
				]
			}`
			_, _ = rw.Write([]byte(rsp))

		}))
	defer server.Close()

	t.Run("Put Notification Submission Checks", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		client.PutNotificationSubmissionChecks(context.Background(), notificationID, submissionID, courier.NotificationSubmissionChecksRequest{
			Checks: []courier.SubmissionCheck{
				courier.SubmissionCheck{
					ID:     "my-check",
					Status: "RESOLVED",
					Type:   "custom",
				},
			},
		})
	})
}

func TestCancelNotificationSubmission(t *testing.T) {
	notificationID := "my-notification"
	submissionID := "my-submission"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/notifications/"+notificationID+"/"+submissionID+"/checks", req.URL.String())

			rw.WriteHeader(http.StatusNoContent)

		}))
	defer server.Close()

	t.Run("Delete Notification Submission", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		client.CancelNotificationSubmission(context.Background(), notificationID, submissionID)
	})
}
