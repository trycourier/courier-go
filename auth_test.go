// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/trycourier/courier-go/v4"
	"github.com/trycourier/courier-go/v4/internal/testutil"
	"github.com/trycourier/courier-go/v4/option"
)

func TestAuthIssueToken(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Auth.IssueToken(context.TODO(), courier.AuthIssueTokenParams{
		ExpiresIn: "expires_in",
		Scope:     "scope",
	})
	if err != nil {
		var apierr *courier.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
