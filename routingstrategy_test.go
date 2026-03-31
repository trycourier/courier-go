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

func TestRoutingStrategyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.RoutingStrategies.New(context.TODO(), courier.RoutingStrategyNewParams{
		RoutingStrategyCreateRequest: courier.RoutingStrategyCreateRequestParam{
			Name: "Email via SendGrid",
			Routing: shared.MessageRoutingParam{
				Channels: []shared.MessageRoutingChannelUnionParam{{
					OfString: courier.String("email"),
				}},
				Method: shared.MessageRoutingMethodSingle,
			},
			Channels: shared.MessageChannelsParam{
				"email": shared.ChannelParam{
					BrandID: courier.String("brand_id"),
					If:      courier.String("if"),
					Metadata: shared.ChannelMetadataParam{
						Utm: shared.UtmParam{
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
					Providers:     []string{"sendgrid", "ses"},
					RoutingMethod: shared.ChannelRoutingMethodAll,
					Timeouts: shared.TimeoutsParam{
						Channel:  courier.Int(0),
						Provider: courier.Int(0),
					},
				},
			},
			Description: courier.String("Routes email through sendgrid with SES failover"),
			Providers: shared.MessageProvidersParam{
				"sendgrid": shared.MessageProvidersTypeParam{
					If: courier.String("if"),
					Metadata: shared.MetadataParam{
						Utm: shared.UtmParam{
							Campaign: courier.String("campaign"),
							Content:  courier.String("content"),
							Medium:   courier.String("medium"),
							Source:   courier.String("source"),
							Term:     courier.String("term"),
						},
					},
					Override: map[string]any{},
					Timeouts: courier.Int(0),
				},
			},
			Tags: []string{"production", "email"},
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

func TestRoutingStrategyGet(t *testing.T) {
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
	_, err := client.RoutingStrategies.Get(context.TODO(), "id")
	if err != nil {
		var apierr *courier.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRoutingStrategyListWithOptionalParams(t *testing.T) {
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
	_, err := client.RoutingStrategies.List(context.TODO(), courier.RoutingStrategyListParams{
		Cursor: courier.String("cursor"),
		Limit:  courier.Int(1),
	})
	if err != nil {
		var apierr *courier.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRoutingStrategyArchive(t *testing.T) {
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
	err := client.RoutingStrategies.Archive(context.TODO(), "id")
	if err != nil {
		var apierr *courier.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRoutingStrategyReplaceWithOptionalParams(t *testing.T) {
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
	_, err := client.RoutingStrategies.Replace(
		context.TODO(),
		"id",
		courier.RoutingStrategyReplaceParams{
			RoutingStrategyReplaceRequest: courier.RoutingStrategyReplaceRequestParam{
				Name: "Email via SendGrid v2",
				Routing: shared.MessageRoutingParam{
					Channels: []shared.MessageRoutingChannelUnionParam{{
						OfString: courier.String("email"),
					}},
					Method: shared.MessageRoutingMethodSingle,
				},
				Channels: shared.MessageChannelsParam{
					"email": shared.ChannelParam{
						BrandID: courier.String("brand_id"),
						If:      courier.String("if"),
						Metadata: shared.ChannelMetadataParam{
							Utm: shared.UtmParam{
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
						Providers:     []string{"ses", "sendgrid"},
						RoutingMethod: shared.ChannelRoutingMethodAll,
						Timeouts: shared.TimeoutsParam{
							Channel:  courier.Int(0),
							Provider: courier.Int(0),
						},
					},
				},
				Description: courier.String("Updated routing with SES primary"),
				Providers: shared.MessageProvidersParam{
					"ses": shared.MessageProvidersTypeParam{
						If: courier.String("if"),
						Metadata: shared.MetadataParam{
							Utm: shared.UtmParam{
								Campaign: courier.String("campaign"),
								Content:  courier.String("content"),
								Medium:   courier.String("medium"),
								Source:   courier.String("source"),
								Term:     courier.String("term"),
							},
						},
						Override: map[string]any{},
						Timeouts: courier.Int(0),
					},
				},
				Tags: []string{"production", "email", "v2"},
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
