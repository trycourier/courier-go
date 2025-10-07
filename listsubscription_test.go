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

func TestListSubscriptionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Lists.Subscriptions.List(
		context.TODO(),
		"list_id",
		courier.ListSubscriptionListParams{
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

func TestListSubscriptionAdd(t *testing.T) {
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
	err := client.Lists.Subscriptions.Add(
		context.TODO(),
		"list_id",
		courier.ListSubscriptionAddParams{
			Recipients: []courier.PutSubscriptionsRecipientParam{{
				RecipientID: "recipientId",
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

func TestListSubscriptionSubscribe(t *testing.T) {
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
	err := client.Lists.Subscriptions.Subscribe(
		context.TODO(),
		"list_id",
		courier.ListSubscriptionSubscribeParams{
			Recipients: []courier.PutSubscriptionsRecipientParam{{
				RecipientID: "recipientId",
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

func TestListSubscriptionSubscribeUserWithOptionalParams(t *testing.T) {
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
	err := client.Lists.Subscriptions.SubscribeUser(
		context.TODO(),
		"user_id",
		courier.ListSubscriptionSubscribeUserParams{
			ListID: "list_id",
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

func TestListSubscriptionUnsubscribeUser(t *testing.T) {
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
	err := client.Lists.Subscriptions.UnsubscribeUser(
		context.TODO(),
		"user_id",
		courier.ListSubscriptionUnsubscribeUserParams{
			ListID: "list_id",
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
