// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/trycourier/courier-go/v3"
	"github.com/trycourier/courier-go/v3/internal/testutil"
	"github.com/trycourier/courier-go/v3/option"
	"github.com/trycourier/courier-go/v3/shared"
)

func TestTenantTenantDefaultPreferenceItemUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Tenants.TenantDefaultPreferences.Items.Update(
		context.TODO(),
		"topic_id",
		courier.TenantTenantDefaultPreferenceItemUpdateParams{
			TenantID: "tenant_id",
			SubscriptionTopicNew: shared.SubscriptionTopicNewParam{
				Status:           shared.SubscriptionTopicNewStatusOptedIn,
				CustomRouting:    []shared.ChannelClassification{shared.ChannelClassificationInbox},
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

func TestTenantTenantDefaultPreferenceItemDelete(t *testing.T) {
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
	err := client.Tenants.TenantDefaultPreferences.Items.Delete(
		context.TODO(),
		"topic_id",
		courier.TenantTenantDefaultPreferenceItemDeleteParams{
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
