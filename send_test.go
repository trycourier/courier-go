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

func TestSendMessageWithOptionalParams(t *testing.T) {
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
	_, err := client.Send.Message(context.TODO(), courier.SendMessageParams{
		Message: courier.SendMessageParamsMessage{
			BrandID: courier.String("brand_id"),
			Channels: map[string]courier.SendMessageParamsMessageChannel{
				"foo": {
					BrandID: courier.String("brand_id"),
					If:      courier.String("if"),
					Metadata: courier.SendMessageParamsMessageChannelMetadata{
						Utm: courier.UtmParam{
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
					Timeouts: courier.SendMessageParamsMessageChannelTimeouts{
						Channel:  courier.Int(0),
						Provider: courier.Int(0),
					},
				},
			},
			Content: courier.SendMessageParamsMessageContentUnion{
				OfElementalContentSugar: &shared.ElementalContentSugarParam{
					Body:  "body",
					Title: "title",
				},
			},
			Context: courier.MessageContextParam{
				TenantID: courier.String("tenant_id"),
			},
			Data: map[string]any{
				"foo": "bar",
			},
			Delay: courier.SendMessageParamsMessageDelay{
				Duration: courier.Int(0),
				Until:    courier.String("until"),
			},
			Expiry: courier.SendMessageParamsMessageExpiry{
				ExpiresIn: courier.SendMessageParamsMessageExpiryExpiresInUnion{
					OfString: courier.String("string"),
				},
				ExpiresAt: courier.String("expires_at"),
			},
			Metadata: courier.SendMessageParamsMessageMetadata{
				Event:   courier.String("event"),
				Tags:    []string{"string"},
				TraceID: courier.String("trace_id"),
				Utm: courier.UtmParam{
					Campaign: courier.String("campaign"),
					Content:  courier.String("content"),
					Medium:   courier.String("medium"),
					Source:   courier.String("source"),
					Term:     courier.String("term"),
				},
			},
			Preferences: courier.SendMessageParamsMessagePreferences{
				SubscriptionTopicID: "subscription_topic_id",
			},
			Providers: map[string]courier.SendMessageParamsMessageProvider{
				"foo": {
					If: courier.String("if"),
					Metadata: courier.SendMessageParamsMessageProviderMetadata{
						Utm: courier.UtmParam{
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
			Routing: courier.SendMessageParamsMessageRouting{
				Channels: []shared.MessageRoutingChannelUnionParam{{
					OfString: courier.String("string"),
				}},
				Method: "all",
			},
			Template: courier.String("template_id"),
			Timeout: courier.SendMessageParamsMessageTimeout{
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
			To: courier.SendMessageParamsMessageToUnion{
				OfUserRecipient: &courier.UserRecipientParam{
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
					UserID:   courier.String("example_user"),
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
