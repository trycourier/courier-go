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

func TestFeedbackNewWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Feedback.New(context.TODO(), courier.FeedbackNewParams{
		Category:         courier.FeedbackNewParamsCategoryPossibleBug,
		Description:      "GET /lists/{list_id}/subscriptions returned items: [] with paging.more set to true, and following the cursor returned the remaining subscribers.",
		Surface:          courier.FeedbackNewParamsSurfaceAPI,
		Title:            "List subscriptions returns an empty page with a next cursor",
		ExpectedBehavior: courier.String("Every page before the last one contains at least one item."),
		Harness:          courier.String("Claude Code"),
		Location:         courier.String("GET /lists/{list_id}/subscriptions"),
		Model:            courier.String("claude-opus-5-5"),
		RequestID:        courier.String("1-67f340b1-58b7b231d9485ef0cda0a38b"),
	})
	if err != nil {
		var apierr *courier.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
