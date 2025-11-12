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
)

func TestBrandNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Brands.New(context.TODO(), courier.BrandNewParams{
		Name: "name",
		ID:   courier.String("id"),
		Settings: courier.BrandSettingsParam{
			Colors: courier.BrandColorsParam{
				Primary:   courier.String("primary"),
				Secondary: courier.String("secondary"),
			},
			Email: courier.BrandSettingsEmailParam{
				Footer: courier.EmailFooterParam{
					Content:        courier.String("content"),
					InheritDefault: courier.Bool(true),
				},
				Head: courier.EmailHeadParam{
					InheritDefault: true,
					Content:        courier.String("content"),
				},
				Header: courier.EmailHeaderParam{
					Logo: courier.LogoParam{
						Href:  courier.String("href"),
						Image: courier.String("image"),
					},
					BarColor:       courier.String("barColor"),
					InheritDefault: courier.Bool(true),
				},
				TemplateOverride: courier.BrandSettingsEmailTemplateOverrideParam{
					BrandTemplateParam: courier.BrandTemplateParam{
						Enabled:               true,
						BackgroundColor:       courier.String("backgroundColor"),
						BlocksBackgroundColor: courier.String("blocksBackgroundColor"),
						Footer:                courier.String("footer"),
						Head:                  courier.String("head"),
						Header:                courier.String("header"),
						Width:                 courier.String("width"),
					},
					Mjml: courier.BrandTemplateParam{
						Enabled:               true,
						BackgroundColor:       courier.String("backgroundColor"),
						BlocksBackgroundColor: courier.String("blocksBackgroundColor"),
						Footer:                courier.String("footer"),
						Head:                  courier.String("head"),
						Header:                courier.String("header"),
						Width:                 courier.String("width"),
					},
					FooterBackgroundColor: courier.String("footerBackgroundColor"),
					FooterFullWidth:       courier.Bool(true),
				},
			},
			Inapp: courier.BrandSettingsInAppParam{
				Colors: courier.BrandColorsParam{
					Primary:   courier.String("primary"),
					Secondary: courier.String("secondary"),
				},
				Icons: courier.IconsParam{
					Bell:    courier.String("bell"),
					Message: courier.String("message"),
				},
				WidgetBackground: courier.WidgetBackgroundParam{
					BottomColor: courier.String("bottomColor"),
					TopColor:    courier.String("topColor"),
				},
				BorderRadius:       courier.String("borderRadius"),
				DisableMessageIcon: courier.Bool(true),
				FontFamily:         courier.String("fontFamily"),
				Placement:          courier.BrandSettingsInAppPlacementTop,
			},
		},
		Snippets: courier.BrandSnippetsParam{
			Items: []courier.BrandSnippetParam{{
				Name:  "name",
				Value: "value",
			}},
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

func TestBrandGet(t *testing.T) {
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
	_, err := client.Brands.Get(context.TODO(), "brand_id")
	if err != nil {
		var apierr *courier.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBrandUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Brands.Update(
		context.TODO(),
		"brand_id",
		courier.BrandUpdateParams{
			Name: "name",
			Settings: courier.BrandSettingsParam{
				Colors: courier.BrandColorsParam{
					Primary:   courier.String("primary"),
					Secondary: courier.String("secondary"),
				},
				Email: courier.BrandSettingsEmailParam{
					Footer: courier.EmailFooterParam{
						Content:        courier.String("content"),
						InheritDefault: courier.Bool(true),
					},
					Head: courier.EmailHeadParam{
						InheritDefault: true,
						Content:        courier.String("content"),
					},
					Header: courier.EmailHeaderParam{
						Logo: courier.LogoParam{
							Href:  courier.String("href"),
							Image: courier.String("image"),
						},
						BarColor:       courier.String("barColor"),
						InheritDefault: courier.Bool(true),
					},
					TemplateOverride: courier.BrandSettingsEmailTemplateOverrideParam{
						BrandTemplateParam: courier.BrandTemplateParam{
							Enabled:               true,
							BackgroundColor:       courier.String("backgroundColor"),
							BlocksBackgroundColor: courier.String("blocksBackgroundColor"),
							Footer:                courier.String("footer"),
							Head:                  courier.String("head"),
							Header:                courier.String("header"),
							Width:                 courier.String("width"),
						},
						Mjml: courier.BrandTemplateParam{
							Enabled:               true,
							BackgroundColor:       courier.String("backgroundColor"),
							BlocksBackgroundColor: courier.String("blocksBackgroundColor"),
							Footer:                courier.String("footer"),
							Head:                  courier.String("head"),
							Header:                courier.String("header"),
							Width:                 courier.String("width"),
						},
						FooterBackgroundColor: courier.String("footerBackgroundColor"),
						FooterFullWidth:       courier.Bool(true),
					},
				},
				Inapp: courier.BrandSettingsInAppParam{
					Colors: courier.BrandColorsParam{
						Primary:   courier.String("primary"),
						Secondary: courier.String("secondary"),
					},
					Icons: courier.IconsParam{
						Bell:    courier.String("bell"),
						Message: courier.String("message"),
					},
					WidgetBackground: courier.WidgetBackgroundParam{
						BottomColor: courier.String("bottomColor"),
						TopColor:    courier.String("topColor"),
					},
					BorderRadius:       courier.String("borderRadius"),
					DisableMessageIcon: courier.Bool(true),
					FontFamily:         courier.String("fontFamily"),
					Placement:          courier.BrandSettingsInAppPlacementTop,
				},
			},
			Snippets: courier.BrandSnippetsParam{
				Items: []courier.BrandSnippetParam{{
					Name:  "name",
					Value: "value",
				}},
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

func TestBrandListWithOptionalParams(t *testing.T) {
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
	_, err := client.Brands.List(context.TODO(), courier.BrandListParams{
		Cursor: courier.String("cursor"),
	})
	if err != nil {
		var apierr *courier.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBrandDelete(t *testing.T) {
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
	err := client.Brands.Delete(context.TODO(), "brand_id")
	if err != nil {
		var apierr *courier.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
