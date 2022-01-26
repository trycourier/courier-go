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

func TestCreateJob(t *testing.T) {
	expectedResponseID := "123456789"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/bulk", req.URL.String())

			rw.Header().Add("Content-Type", "application/json")
			_, writeErr := rw.Write([]byte(fmt.Sprintf("{ \"jobId\" : \"%s\" }", expectedResponseID)))
			if writeErr != nil {
				t.Error(writeErr)
			}

		}))
	defer server.Close()

	t.Run("create job", func(t *testing.T) {

		client := courier.CreateClient("key", &server.URL)

		jobID, err := client.CreateJob(context.Background(), courier.CreateJobBody{
			Message: map[string]string{
				"event": "foo",
			},
		})

		assert.Nil(t, err)
		assert.Equal(t, expectedResponseID, jobID)
	})
}

func TestIngestJob(t *testing.T) {
	jobId := "job001"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/bulk/"+jobId, req.URL.String())

			rw.Header().Add("Content-Type", "application/json")

			rsp := `
			{
				"errors": [],
				"total": 1
			}`
			_, _ = rw.Write([]byte(rsp))
		}))
	defer server.Close()

	t.Run("ingest job", func(t *testing.T) {

		client := courier.CreateClient("key", &server.URL)

		response, err := client.IngestJob(context.Background(), "job001", courier.IngestJobBody{
			Users: []interface{}{
				map[string]string{"recipient": "johndoe"},
			},
		})

		assert.Nil(t, err)
		assert.Equal(t, 0, len(response.Errors))
		assert.Equal(t, 1, response.Total)
	})
}

func TestRunJob(t *testing.T) {
	jobId := "job001"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/bulk/"+jobId+"/run", req.URL.String())

			rw.WriteHeader(http.StatusAccepted)
		}))
	defer server.Close()

	t.Run("run job", func(t *testing.T) {

		client := courier.CreateClient("key", &server.URL)

		err := client.RunJob(context.Background(), "job001")

		assert.Nil(t, err)
	})
}

func TestGetJob(t *testing.T) {
	jobId := "job001"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/bulk/"+jobId, req.URL.String())

			rw.Header().Add("Content-Type", "application/json")

			rsp := `
			{
				"job": {
					"definition": {
						"event": "foo"
					},
					"enqueued": 1,
					"failures": 0,
					"received": 1,
					"status": "COMPLETED"
				}
			}`
			_, _ = rw.Write([]byte(rsp))
		}))
	defer server.Close()

	t.Run("get job", func(t *testing.T) {

		client := courier.CreateClient("key", &server.URL)

		response, err := client.GetJob(context.Background(), "job001")

		assert.Nil(t, err)
		assert.NotNil(t, response.Job)
	})
}

func TestGetJobUsers(t *testing.T) {
	jobId := "job001"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/bulk/"+jobId+"/users", req.URL.String())

			rw.Header().Add("Content-Type", "application/json")

			rsp := `
			{
				"items": [
					{
						"recipient": "johndoe"
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

	t.Run("get job users", func(t *testing.T) {

		client := courier.CreateClient("key", &server.URL)

		response, err := client.GetJobUsers(context.Background(), "job001", "")

		assert.Nil(t, err)
		assert.Equal(t, 1, len(response.Items))
		assert.Nil(t, response.Paging.Cursor)
		assert.Equal(t, false, response.Paging.More)
	})
}
