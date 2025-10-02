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

func TestSendMessage(t *testing.T) {
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
		Message: courier.MessageUnionParam{
			OfContentMessage: &courier.MessageContentMessageParam{
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
						"name": "bar",
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
							OfString: courier.String("email"),
						}},
						Method: "single",
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
				BaseMessageSendToParam: courier.BaseMessageSendToParam{
					To: courier.BaseMessageSendToToUnionParam{
						OfBaseMessageSendToToObject: &courier.BaseMessageSendToToObjectParam{
							Data: map[string]any{
								"foo": "bar",
							},
							Filters: []courier.BaseMessageSendToToObjectFilterParam{{
								Operator: "MEMBER_OF",
								Path:     "account_id",
								Value:    "value",
							}},
							ListID: courier.String("list_id"),
						},
					},
				},
				Content: courier.ContentUnionParam{
					OfElementalContentSugar: &courier.ContentElementalContentSugarParam{
						Body:  "Thanks for signing up, {{name}}",
						Title: "Welcome!",
					},
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
