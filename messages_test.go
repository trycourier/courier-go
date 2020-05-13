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

func TestClient_GetMessage(t *testing.T) {

	requestMessageID := "1-23456789"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/messages/"+requestMessageID, req.URL.String())

			rw.Header().Add("Content-Type", "application/json")
			rsp := `
			{ "id" : "%s" }
			`
			_, _ = rw.Write([]byte(fmt.Sprintf(rsp, requestMessageID)))

		}))
	defer server.Close()

	t.Run("makes requests for message ID", func(t *testing.T) {
		client := courier.CourierClient("key", server.URL)
		rsp, err := client.GetMessage(context.Background(), requestMessageID)
		assert.Nil(t, err)
		assert.Equal(t, requestMessageID, rsp.Id)
	})
}
