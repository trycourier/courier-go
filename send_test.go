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
			Content: courier.SendMessageParamsMessageContent{
				Body:  "Thanks for signing up, {{name}}",
				Title: "Welcome!",
			},
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
			Context: courier.MessageContextParam{
				TenantID: courier.String("tenant_id"),
			},
			Data: map[string]any{
				"name": "bar",
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
				Channels: []courier.MessageRoutingChannelUnionParam{{
					OfString: courier.String("email"),
				}},
				Method: "single",
			},
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
				OfSendMessagesMessageToObject: &courier.SendMessageParamsMessageToObject{
					AccountID: courier.String("account_id"),
					Context: courier.MessageContextParam{
						TenantID: courier.String("tenant_id"),
					},
					Data: map[string]any{
						"foo": "bar",
					},
					Email:       courier.String("email@example.com"),
					Locale:      courier.String("locale"),
					PhoneNumber: courier.String("phone_number"),
					Preferences: courier.SendMessageParamsMessageToObjectPreferences{
						Notifications: map[string]courier.PreferenceParam{
							"foo": {
								Status: courier.PreferenceStatusOptedIn,
								ChannelPreferences: []courier.PreferenceChannelPreferenceParam{{
									Channel: "direct_message",
								}},
								Rules: []courier.PreferenceRuleParam{{
									Until: "until",
									Start: courier.String("start"),
								}},
								Source: courier.PreferenceSourceSubscription,
							},
						},
						Categories: map[string]courier.PreferenceParam{
							"foo": {
								Status: courier.PreferenceStatusOptedIn,
								ChannelPreferences: []courier.PreferenceChannelPreferenceParam{{
									Channel: "direct_message",
								}},
								Rules: []courier.PreferenceRuleParam{{
									Until: "until",
									Start: courier.String("start"),
								}},
								Source: courier.PreferenceSourceSubscription,
							},
						},
						TemplateID: courier.String("templateId"),
					},
					TenantID: courier.String("tenant_id"),
					UserID:   courier.String("user_id"),
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
