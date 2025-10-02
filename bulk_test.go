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
					Categories: map[string]courier.RecipientPreferencesCategoryParam{
						"foo": {
							Status: courier.PreferenceStatusOptedIn,
							ChannelPreferences: []courier.RecipientPreferencesCategoryChannelPreferenceParam{{
								Channel: courier.ChannelClassificationDirectMessage,
							}},
							Rules: []courier.RecipientPreferencesCategoryRuleParam{{
								Until: "until",
								Start: courier.String("start"),
							}},
						},
					},
					Notifications: map[string]courier.RecipientPreferencesNotificationParam{
						"foo": {
							Status: courier.PreferenceStatusOptedIn,
							ChannelPreferences: []courier.RecipientPreferencesNotificationChannelPreferenceParam{{
								Channel: courier.ChannelClassificationDirectMessage,
							}},
							Rules: []courier.RecipientPreferencesNotificationRuleParam{{
								Until: "until",
								Start: courier.String("start"),
							}},
						},
					},
				},
				Profile:   map[string]interface{}{},
				Recipient: courier.String("recipient"),
				To: courier.UserRecipientParam{
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
					Preferences: courier.UserRecipientPreferencesParam{
						Notifications: map[string]courier.UserRecipientPreferencesNotificationParam{
							"foo": {
								Status: courier.PreferenceStatusOptedIn,
								ChannelPreferences: []courier.UserRecipientPreferencesNotificationChannelPreferenceParam{{
									Channel: courier.ChannelClassificationDirectMessage,
								}},
								Rules: []courier.UserRecipientPreferencesNotificationRuleParam{{
									Until: "until",
									Start: courier.String("start"),
								}},
								Source: "subscription",
							},
						},
						Categories: map[string]courier.UserRecipientPreferencesCategoryParam{
							"foo": {
								Status: courier.PreferenceStatusOptedIn,
								ChannelPreferences: []courier.UserRecipientPreferencesCategoryChannelPreferenceParam{{
									Channel: courier.ChannelClassificationDirectMessage,
								}},
								Rules: []courier.UserRecipientPreferencesCategoryRuleParam{{
									Until: "until",
									Start: courier.String("start"),
								}},
								Source: "subscription",
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
		Message: courier.InboundBulkMessageParam{
			Brand: courier.String("brand"),
			Data: map[string]any{
				"foo": "bar",
			},
			Event: courier.String("event"),
			Locale: map[string]any{
				"foo": "bar",
			},
			Message: courier.InboundBulkMessageMessageUnionParam{
				OfInboundBulkTemplateMessage: &courier.InboundBulkMessageMessageInboundBulkTemplateMessageParam{
					BaseMessageParam: courier.BaseMessageParam{
						BrandID: courier.String("brand_id"),
						Channels: map[string]courier.BaseMessageChannelParam{
							"foo": {
								BrandID: courier.String("brand_id"),
								If:      courier.String("if"),
								Metadata: courier.BaseMessageChannelMetadataParam{
									Utm: courier.BaseMessageChannelMetadataUtmParam{
										Campaign: courier.String("campaign"),
										Content:  courier.String("content"),
										Medium:   courier.String("medium"),
										Source:   courier.String("source"),
										Term:     courier.String("term"),
									},
								},
								Override: map[string]any{
									"foo": "bar",
								},
								Providers:     []string{"string"},
								RoutingMethod: "all",
								Timeouts: courier.BaseMessageChannelTimeoutsParam{
									Channel:  courier.Int(0),
									Provider: courier.Int(0),
								},
							},
						},
						Context: courier.MessageContextParam{
							TenantID: courier.String("tenant_id"),
						},
						Data: map[string]any{
							"foo": "bar",
						},
						Delay: courier.BaseMessageDelayParam{
							Duration: courier.Int(0),
							Until:    courier.String("until"),
						},
						Expiry: courier.BaseMessageExpiryParam{
							ExpiresIn: courier.BaseMessageExpiryExpiresInUnionParam{
								OfString: courier.String("string"),
							},
							ExpiresAt: courier.String("expires_at"),
						},
						Metadata: courier.BaseMessageMetadataParam{
							Event:   courier.String("event"),
							Tags:    []string{"string"},
							TraceID: courier.String("trace_id"),
							Utm: courier.BaseMessageMetadataUtmParam{
								Campaign: courier.String("campaign"),
								Content:  courier.String("content"),
								Medium:   courier.String("medium"),
								Source:   courier.String("source"),
								Term:     courier.String("term"),
							},
						},
						Preferences: courier.BaseMessagePreferencesParam{
							SubscriptionTopicID: "subscription_topic_id",
						},
						Providers: map[string]courier.BaseMessageProviderParam{
							"foo": {
								If: courier.String("if"),
								Metadata: courier.BaseMessageProviderMetadataParam{
									Utm: courier.BaseMessageProviderMetadataUtmParam{
										Campaign: courier.String("campaign"),
										Content:  courier.String("content"),
										Medium:   courier.String("medium"),
										Source:   courier.String("source"),
										Term:     courier.String("term"),
									},
								},
								Override: map[string]any{
									"foo": "bar",
								},
								Timeouts: courier.Int(0),
							},
						},
						Routing: courier.BaseMessageRoutingParam{
							Channels: []courier.MessageRoutingChannelUnionParam{{
								OfString: courier.String("string"),
							}},
							Method: "all",
						},
						Timeout: courier.BaseMessageTimeoutParam{
							Channel: map[string]int64{
								"foo": 0,
							},
							Criteria:   "no-escalation",
							Escalation: courier.Int(0),
							Message:    courier.Int(0),
							Provider: map[string]int64{
								"foo": 0,
							},
						},
					},
					Template: "template",
				},
			},
			Override: map[string]interface{}{},
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
