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

func TestTenantTemplateGet(t *testing.T) {
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
	_, err := client.Tenants.Templates.Get(
		context.TODO(),
		"template_id",
		courier.TenantTemplateGetParams{
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

func TestTenantTemplateListWithOptionalParams(t *testing.T) {
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
	_, err := client.Tenants.Templates.List(
		context.TODO(),
		"tenant_id",
		courier.TenantTemplateListParams{
			Cursor: courier.String("cursor"),
			Limit:  courier.Int(0),
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

func TestTenantTemplatePublishWithOptionalParams(t *testing.T) {
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
	_, err := client.Tenants.Templates.Publish(
		context.TODO(),
		"template_id",
		courier.TenantTemplatePublishParams{
			TenantID: "tenant_id",
			PostTenantTemplatePublishRequest: courier.PostTenantTemplatePublishRequestParam{
				Version: courier.String("version"),
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

func TestTenantTemplateReplaceWithOptionalParams(t *testing.T) {
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
	_, err := client.Tenants.Templates.Replace(
		context.TODO(),
		"template_id",
		courier.TenantTemplateReplaceParams{
			TenantID: "tenant_id",
			PutTenantTemplateRequest: courier.PutTenantTemplateRequestParam{
				Template: courier.TenantTemplateInputParam{
					Content: shared.ElementalContentParam{
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
						Version: "version",
					},
					Channels: map[string]courier.TenantTemplateInputChannelParam{
						"foo": {
							BrandID: courier.String("brand_id"),
							If:      courier.String("if"),
							Metadata: courier.TenantTemplateInputChannelMetadataParam{
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
							Providers:     []string{"string"},
							RoutingMethod: "all",
							Timeouts: courier.TenantTemplateInputChannelTimeoutsParam{
								Channel:  courier.Int(0),
								Provider: courier.Int(0),
							},
						},
					},
					Providers: map[string]courier.TenantTemplateInputProviderParam{
						"foo": {
							If: courier.String("if"),
							Metadata: courier.TenantTemplateInputProviderMetadataParam{
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
							Timeouts: courier.Int(0),
						},
					},
					Routing: shared.MessageRoutingParam{
						Channels: []shared.MessageRoutingChannelUnionParam{{
							OfString: courier.String("string"),
						}},
						Method: shared.MessageRoutingMethodAll,
					},
				},
				Published: courier.Bool(true),
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
