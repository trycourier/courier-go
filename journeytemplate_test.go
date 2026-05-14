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

func TestJourneyTemplateNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Journeys.Templates.New(
		context.TODO(),
		"x",
		courier.JourneyTemplateNewParams{
			JourneyTemplateCreateRequest: courier.JourneyTemplateCreateRequestParam{
				Channel: "email",
				Notification: courier.JourneyTemplateCreateRequestNotificationParam{
					Brand: courier.JourneyTemplateCreateRequestNotificationBrandParam{
						ID: "id",
					},
					Content: courier.JourneyTemplateCreateRequestNotificationContentParam{
						Elements: []shared.ElementalNodeUnionParam{{
							OfElementalTextNodeWithType: &shared.ElementalTextNodeWithTypeParam{
								ElementalBaseNodeParam: shared.ElementalBaseNodeParam{
									Channels: []string{"string"},
									If:       courier.String("if"),
									Loop:     courier.String("loop"),
									Ref:      courier.String("ref"),
								},
								Type: "text",
							},
						}},
						Version: "2022-01-01",
						Scope:   "default",
					},
					Name: "Welcome email",
					Subscription: courier.JourneyTemplateCreateRequestNotificationSubscriptionParam{
						TopicID: "topic_id",
					},
					Tags: []string{"string"},
				},
				ProviderKey: courier.String("x"),
				State:       courier.String("state"),
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

func TestJourneyTemplateGet(t *testing.T) {
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
	_, err := client.Journeys.Templates.Get(
		context.TODO(),
		"x",
		courier.JourneyTemplateGetParams{
			TemplateID: "x",
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

func TestJourneyTemplateListWithOptionalParams(t *testing.T) {
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
	_, err := client.Journeys.Templates.List(
		context.TODO(),
		"x",
		courier.JourneyTemplateListParams{
			Cursor: courier.String("cursor"),
			Limit:  courier.Int(1),
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

func TestJourneyTemplateArchive(t *testing.T) {
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
	err := client.Journeys.Templates.Archive(
		context.TODO(),
		"x",
		courier.JourneyTemplateArchiveParams{
			TemplateID: "x",
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

func TestJourneyTemplateListVersions(t *testing.T) {
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
	_, err := client.Journeys.Templates.ListVersions(
		context.TODO(),
		"x",
		courier.JourneyTemplateListVersionsParams{
			TemplateID: "x",
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

func TestJourneyTemplatePublishWithOptionalParams(t *testing.T) {
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
	err := client.Journeys.Templates.Publish(
		context.TODO(),
		"x",
		courier.JourneyTemplatePublishParams{
			TemplateID: "x",
			JourneyTemplatePublishRequest: courier.JourneyTemplatePublishRequestParam{
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

func TestJourneyTemplateReplaceWithOptionalParams(t *testing.T) {
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
	_, err := client.Journeys.Templates.Replace(
		context.TODO(),
		"x",
		courier.JourneyTemplateReplaceParams{
			TemplateID: "x",
			JourneyTemplateReplaceRequest: courier.JourneyTemplateReplaceRequestParam{
				Notification: courier.JourneyTemplateReplaceRequestNotificationParam{
					Brand: courier.JourneyTemplateReplaceRequestNotificationBrandParam{
						ID: "id",
					},
					Content: courier.JourneyTemplateReplaceRequestNotificationContentParam{
						Elements: []shared.ElementalNodeUnionParam{{
							OfElementalTextNodeWithType: &shared.ElementalTextNodeWithTypeParam{
								ElementalBaseNodeParam: shared.ElementalBaseNodeParam{
									Channels: []string{"string"},
									If:       courier.String("if"),
									Loop:     courier.String("loop"),
									Ref:      courier.String("ref"),
								},
								Type: "text",
							},
						}},
						Version: "2022-01-01",
						Scope:   "default",
					},
					Name: "name",
					Subscription: courier.JourneyTemplateReplaceRequestNotificationSubscriptionParam{
						TopicID: "topic_id",
					},
					Tags: []string{"string"},
				},
				State: courier.String("state"),
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
