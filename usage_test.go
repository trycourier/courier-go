// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier_test

import (
	"context"
	"os"
	"testing"

	"github.com/trycourier/courier-go"
	"github.com/trycourier/courier-go/internal/testutil"
	"github.com/trycourier/courier-go/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := courier.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	response, err := client.Send.Message(context.TODO(), courier.SendMessageParams{
		Message: courier.SendMessageParamsMessage{
			To: courier.SendMessageParamsMessageToUnion{
				OfUserRecipient: &courier.UserRecipientParam{
					UserID: courier.String("your_user_id"),
				},
			},
			Template: courier.String("your_template"),
			Data: map[string]any{
				"foo": "bar",
			},
		},
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", response.RequestID)
}
