package courier_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trycourier/courier-go"
)

// type TestResponseProfile struct {
// 	Foo string `json:"foo"`
// }

func TestProfiles_GetProfile(t *testing.T) {
	profileID := "example"
	foo := "bar"
	url := "/profiles/" + profileID

	server := httptest.NewServer(http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			assert.Equal(t, url, req.URL.String())
			rw.Header().Add("Content-Type", "application/json")
			rsp := `
				{
					"profile": {
						"foo": "%s"
					}
				}`
			_, _ = rw.Write([]byte(fmt.Sprintf(rsp, foo)))
		}),
	)
	defer server.Close()

	t.Run("makes requests for message ID", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		response, err := client.GetProfile(context.Background(), profileID)
		assert.Nil(t, err)

		var profile map[string]string
		err = json.Unmarshal(response["profile"], &profile)
		assert.Nil(t, err)

		assert.Equal(t, "bar", profile["foo"])
	})
}
