// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/trycourier/courier-go"
	"github.com/trycourier/courier-go/internal/testutil"
	"github.com/trycourier/courier-go/option"
	"github.com/trycourier/courier-go/shared"
)

func TestBulkAddUsers(t *testing.T) {
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
	err := client.Bulk.AddUsers(
		context.TODO(),
		"job_id",
		courier.BulkAddUsersParams{
			Users: []courier.InboundBulkMessageUserParam{{
				Data: map[string]interface{}{},
				Preferences: courier.RecipientPreferencesParam{
					Categories: map[string]courier.NotificationPreferenceDetailsParam{
						"foo": {
							Status: courier.PreferenceStatusOptedIn,
							ChannelPreferences: []shared.ChannelPreferenceParam{{
								Channel: courier.ChannelClassificationDirectMessage,
							}},
							Rules: []shared.RuleParam{{
								Until: "until",
								Start: courier.String("start"),
							}},
						},
					},
					Notifications: map[string]courier.NotificationPreferenceDetailsParam{
						"foo": {
							Status: courier.PreferenceStatusOptedIn,
							ChannelPreferences: []shared.ChannelPreferenceParam{{
								Channel: courier.ChannelClassificationDirectMessage,
							}},
							Rules: []shared.RuleParam{{
								Until: "until",
								Start: courier.String("start"),
							}},
						},
					},
				},
				Profile:   map[string]interface{}{},
				Recipient: courier.String("recipient"),
				To: shared.UserRecipientParam{
					AccountID: courier.String("account_id"),
					Context: courier.MessageContextParam{
						TenantID: courier.String("tenant_id"),
					},
					Data: map[string]any{
						"foo": "bar",
					},
					Email:       courier.String("email"),
					Locale:      courier.String("locale"),
					PhoneNumber: courier.String("phone_number"),
					Preferences: shared.UserRecipientPreferencesParam{
						Notifications: map[string]shared.PreferenceParam{
							"foo": {
								Status: courier.PreferenceStatusOptedIn,
								ChannelPreferences: []shared.ChannelPreferenceParam{{
									Channel: courier.ChannelClassificationDirectMessage,
								}},
								Rules: []shared.RuleParam{{
									Until: "until",
									Start: courier.String("start"),
								}},
								Source: shared.PreferenceSourceSubscription,
							},
						},
						Categories: map[string]shared.PreferenceParam{
							"foo": {
								Status: courier.PreferenceStatusOptedIn,
								ChannelPreferences: []shared.ChannelPreferenceParam{{
									Channel: courier.ChannelClassificationDirectMessage,
								}},
								Rules: []shared.RuleParam{{
									Until: "until",
									Start: courier.String("start"),
								}},
								Source: shared.PreferenceSourceSubscription,
							},
						},
						TemplateID: courier.String("templateId"),
					},
					TenantID: courier.String("tenant_id"),
					UserID:   courier.String("user_id"),
				},
			}},
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

func TestBulkNewJobWithOptionalParams(t *testing.T) {
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
	_, err := client.Bulk.NewJob(context.TODO(), courier.BulkNewJobParams{
		Message: courier.InboundBulkMessageUnionParam{
			OfInboundBulkTemplateMessage: &courier.InboundBulkMessageInboundBulkTemplateMessageParam{
				Template: "template",
				Brand:    courier.String("brand"),
				Data: map[string]any{
					"foo": "bar",
				},
				Event: courier.String("event"),
				Locale: map[string]map[string]any{
					"foo": {
						"foo": "bar",
					},
				},
				Override: map[string]any{
					"foo": "bar",
				},
			},
		},
	})
	if err != nil {
		var apierr *courier.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBulkListUsersWithOptionalParams(t *testing.T) {
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
	_, err := client.Bulk.ListUsers(
		context.TODO(),
		"job_id",
		courier.BulkListUsersParams{
			Cursor: courier.String("cursor"),
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

func TestBulkGetJob(t *testing.T) {
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
	_, err := client.Bulk.GetJob(context.TODO(), "job_id")
	if err != nil {
		var apierr *courier.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBulkRunJob(t *testing.T) {
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
	err := client.Bulk.RunJob(context.TODO(), "job_id")
	if err != nil {
		var apierr *courier.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
