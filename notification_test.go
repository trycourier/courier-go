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
	"github.com/trycourier/courier-go/v4/shared"
)

func TestNotificationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Notifications.New(context.TODO(), courier.NotificationNewParams{
		NotificationTemplateCreateRequest: courier.NotificationTemplateCreateRequestParam{
			Notification: courier.NotificationTemplatePayloadParam{
				Brand: courier.NotificationTemplatePayloadBrandParam{
					ID: "brand_abc",
				},
				Content: shared.ElementalContentParam{
					Elements: []shared.ElementalNodeUnionParam{{
						OfElementalChannelNodeWithType: &shared.ElementalChannelNodeWithTypeParam{
							ElementalChannelNodeParam: shared.ElementalChannelNodeParam{
								ElementalBaseNodeParam: shared.ElementalBaseNodeParam{},
							},
							Type: "channel",
						},
					}},
					Version: "2022-01-01",
				},
				Name: "Welcome Email",
				Routing: courier.NotificationTemplatePayloadRoutingParam{
					StrategyID: "rs_123",
				},
				Subscription: courier.NotificationTemplatePayloadSubscriptionParam{
					TopicID: "marketing",
				},
				Tags: []string{"onboarding", "welcome"},
			},
			State: courier.NotificationTemplateCreateRequestStateDraft,
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

func TestNotificationGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Notifications.Get(
		context.TODO(),
		"id",
		courier.NotificationGetParams{
			Version: courier.String("version"),
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

func TestNotificationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Notifications.List(context.TODO(), courier.NotificationListParams{
		Cursor:  courier.String("cursor"),
		EventID: courier.String("event_id"),
		Notes:   courier.Bool(true),
	})
	if err != nil {
		var apierr *courier.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNotificationArchive(t *testing.T) {
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
	err := client.Notifications.Archive(context.TODO(), "id")
	if err != nil {
		var apierr *courier.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNotificationListVersionsWithOptionalParams(t *testing.T) {
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
	_, err := client.Notifications.ListVersions(
		context.TODO(),
		"id",
		courier.NotificationListVersionsParams{
			Cursor: courier.String("cursor"),
			Limit:  courier.Int(10),
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

func TestNotificationPublishWithOptionalParams(t *testing.T) {
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
	err := client.Notifications.Publish(
		context.TODO(),
		"id",
		courier.NotificationPublishParams{
			NotificationTemplatePublishRequest: courier.NotificationTemplatePublishRequestParam{
				Version: courier.String("v321669910225"),
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

func TestNotificationPutContentWithOptionalParams(t *testing.T) {
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
	_, err := client.Notifications.PutContent(
		context.TODO(),
		"id",
		courier.NotificationPutContentParams{
			NotificationContentPutRequest: courier.NotificationContentPutRequestParam{
				Content: courier.NotificationContentPutRequestContentParam{
					Elements: []shared.ElementalNodeUnionParam{{
						OfElementalChannelNodeWithType: &shared.ElementalChannelNodeWithTypeParam{
							ElementalChannelNodeParam: shared.ElementalChannelNodeParam{
								ElementalBaseNodeParam: shared.ElementalBaseNodeParam{},
							},
							Type: "channel",
						},
					}},
					Version: courier.String("2022-01-01"),
				},
				State: courier.NotificationTemplateStateDraft,
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

func TestNotificationPutElementWithOptionalParams(t *testing.T) {
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
	_, err := client.Notifications.PutElement(
		context.TODO(),
		"elementId",
		courier.NotificationPutElementParams{
			ID: "id",
			NotificationElementPutRequest: courier.NotificationElementPutRequestParam{
				Type:     "text",
				Channels: []string{"string"},
				Data: map[string]any{
					"content": "bar",
				},
				If:    courier.String("if"),
				Loop:  courier.String("loop"),
				Ref:   courier.String("ref"),
				State: courier.NotificationTemplateStateDraft,
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

func TestNotificationPutLocaleWithOptionalParams(t *testing.T) {
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
	_, err := client.Notifications.PutLocale(
		context.TODO(),
		"localeId",
		courier.NotificationPutLocaleParams{
			ID: "id",
			NotificationLocalePutRequest: courier.NotificationLocalePutRequestParam{
				Elements: []courier.NotificationLocalePutRequestElementParam{{
					ID: "elem_1",
				}, {
					ID: "elem_2",
				}},
				State: courier.NotificationTemplateStateDraft,
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

func TestNotificationReplaceWithOptionalParams(t *testing.T) {
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
	_, err := client.Notifications.Replace(
		context.TODO(),
		"id",
		courier.NotificationReplaceParams{
			NotificationTemplateUpdateRequest: courier.NotificationTemplateUpdateRequestParam{
				Notification: courier.NotificationTemplatePayloadParam{
					Brand: courier.NotificationTemplatePayloadBrandParam{
						ID: "id",
					},
					Content: shared.ElementalContentParam{
						Elements: []shared.ElementalNodeUnionParam{{
							OfElementalChannelNodeWithType: &shared.ElementalChannelNodeWithTypeParam{
								ElementalChannelNodeParam: shared.ElementalChannelNodeParam{
									ElementalBaseNodeParam: shared.ElementalBaseNodeParam{},
								},
								Type: "channel",
							},
						}},
						Version: "2022-01-01",
					},
					Name: "Updated Name",
					Routing: courier.NotificationTemplatePayloadRoutingParam{
						StrategyID: "strategy_id",
					},
					Subscription: courier.NotificationTemplatePayloadSubscriptionParam{
						TopicID: "topic_id",
					},
					Tags: []string{"updated"},
				},
				State: courier.NotificationTemplateUpdateRequestStatePublished,
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

func TestNotificationGetContentWithOptionalParams(t *testing.T) {
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
	_, err := client.Notifications.GetContent(
		context.TODO(),
		"id",
		courier.NotificationGetContentParams{
			Version: courier.String("version"),
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
