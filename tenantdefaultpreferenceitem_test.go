// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/courier-go"
	"github.com/stainless-sdks/courier-go/internal/testutil"
	"github.com/stainless-sdks/courier-go/option"
)

func TestTenantDefaultPreferenceItemUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Tenants.DefaultPreferences.Items.Update(
		context.TODO(),
		"topic_id",
		courier.TenantDefaultPreferenceItemUpdateParams{
			TenantID: "tenant_id",
			SubscriptionTopicNew: courier.SubscriptionTopicNewParam{
				Status:           courier.SubscriptionTopicNewStatusOptedIn,
				CustomRouting:    []courier.ChannelClassification{courier.ChannelClassificationInbox},
				HasCustomRouting: courier.Bool(true),
			},
		},
	)
	if err != nil {
		var apierr *courier.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTenantDefaultPreferenceItemDelete(t *testing.T) {
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
	err := client.Tenants.DefaultPreferences.Items.Delete(
		context.TODO(),
		"topic_id",
		courier.TenantDefaultPreferenceItemDeleteParams{
			TenantID: "tenant_id",
		},
	)
	if err != nil {
		var apierr *courier.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
